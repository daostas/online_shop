package user_service

import (
	"context"
	"online_shop/auth-svc/config"
	"online_shop/auth-svc/pb"
	rep "online_shop/repository"
	"online_shop/repository/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
)

type UserClaims struct {
	jwt.StandardClaims `json:"jwt-standard-claims"`
	ID                 int32  `json:"id"`
	Role               string `json:"role"`
}

type AuthServer struct {
	pb.UnimplementedAuthServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewAuthServer(db *reform.DB, cfg *config.Config) *AuthServer {
	return &AuthServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *AuthServer) RegisterUser(_ context.Context, req *pb.RegReq) (*pb.AuthRes, error) {
	only_numbers, is_email := true, false // Определение - был введен номер или почта
	for i := 0; i < len(req.Login); i++ {
		if req.Login[i] >= '0' && req.Login[i] <= '9' {
			continue
		} else {
			only_numbers = false
			if req.Login[i] == '@' {
				is_email = true
				break
			}
		}
	}

	var user *models.Users
	if only_numbers { //код по созданию пользоватьеля по номеру

		_, err := s.Db.SelectOneFrom(models.UsersTable, "where number = $1", req.Login) //Проверка на не занятость номера
		if err == nil {                                                                 // Если занят
			return &pb.AuthRes{Err: "already exists"}, nil
		}

		hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), int(s.Cfg.BcryptCost)) // хэширование пароля
		if err != nil {
			return &pb.AuthRes{Err: "err"}, nil
		}

		user = rep.NewUsers(nil, &req.Login, nil, nil, nil, string(hashed_pass)) // сохранение нового пользователя
		if err := s.Db.Insert(user); err != nil {
			return &pb.AuthRes{Err: "user creation error"}, nil
		}

		// user2, err := s.Db.SelectOneFrom(models.UsersTable, "where number = $1", req.Login) // достаю айди нового пользователя
		// if err != nil {
		// 	return &pb.AuthRes{Err: "err"}, nil
		// }

	} else if is_email { //код по созданию пользователя по почте

		_, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) //Проверка на не занятость почты
		if err == nil {                                                                // Если занята
			return &pb.AuthRes{Err: "already exists"}, nil
		}

		hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), int(s.Cfg.BcryptCost)) // хэширования пароля
		if err != nil {
			return &pb.AuthRes{Err: "err"}, nil
		}

		user = rep.NewUsers(nil, nil, &req.Login, nil, nil, string(hashed_pass)) // сохраненение нового пользователя
		if err := s.Db.Insert(user); err != nil {
			return &pb.AuthRes{Err: "user creation error"}, nil
		}

	} else { // возврат ошибки о некоректных данных
		return &pb.AuthRes{Err: "invalid data"}, nil
	}

	basket := rep.NewBasket(user.UserID)
	if err := s.Db.Insert(basket); err != nil { // Создаю корзину с айди пользователя
		return &pb.AuthRes{Err: "basket creation error"}, nil
	}

	favourite := rep.NewFavourite(user.UserID)
	if err := s.Db.Insert(favourite); err != nil { // создаю избранное с айди пользователя
		return &pb.AuthRes{Err: "favourite creation error"}, nil
	}

	return &pb.AuthRes{Err: "success"}, nil
}

// only_numbers, is_email := true, false // Определение - был введен номер или почта
// 	for i := 0; i < len(req.Login); i++ {
// 		if req.Login[i] >= '0' && req.Login[i] <= '9' {
// 			continue
// 		} else {
// 			only_numbers = false
// 			if req.Login[i] == '@' {
// 				is_email = true
// 				break
// 			}
// 		}
// 	}

// func (s *AuthServer) SignInUserByNum(_ context.Context, req *pb.SignInReq) (*pb.SignInRes, error) {

// 	user, err := s.Db.SelectOneFrom(models.UsersTable, "where number = $1", req.Login) // достаю пароль
// 	if err != nil {
// 		return &pb.SignInRes{
// 			Status: int32(codes.Internal),
// 			Err: "User not found or error in search",
// 		}, status.Errorf(codes.Internal, "%v", err)
// 	}
// 	err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
// 	if err != nil {                                                                                      // если не совпали
// 		return &pb.SignInRes{
// 			Status: int32(codes.InvalidArgument),
// 			Err: "wrong password",
// 			}, status.Errorf(codes.InvalidArgument, "%v", err)
// 	}

// 	claims := UserClaims{
// 		ID:    user.(*models.Users).UserID,
// 		Role:  nil,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(s.Cfg.TokenDuration)).Unix(),
// 			Issuer:    "Online shop",
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString([]byte(s.Cfg.JWTSecretKey))
// 	if err != nil {
// 		return &pb.SignInRes{
// 			Status: int32(codes.Internal),
// 			Err:  err.Error(),
// 		}, status.Errorf(codes.Internal, "%v", err)
// 	}

// 	return &pb.SignInRes{
// 		Status: int32(codes.OK),
// 		Err:  "",
// 		Token:  signedToken,
// 	}, nil

// }

func (s *AuthServer) SignInUser(_ context.Context, req *pb.SignInReq) (*pb.SignInRes, error) {

	var claims UserClaims
	only_numbers, is_email := true, false // Определение - был введен номер или почта
	for i := 0; i < len(req.Login); i++ {
		if req.Login[i] >= '0' && req.Login[i] <= '9' {
			continue
		} else {
			only_numbers = false
			if req.Login[i] == '@' {
				is_email = true
				break
			}
		}
	}
	if only_numbers { // если номер

		user, err := s.Db.SelectOneFrom(models.UsersTable, "where number = $1", req.Login) // достаю пароль
		if err != nil {
			return &pb.SignInRes{Err: "User not found or error in search"}, nil
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
		if err != nil {                                                                                      // если не совпали
			return &pb.SignInRes{
				Status: int32(codes.InvalidArgument),
				Err:    "wrong password",
				Token:  "",
			}, status.Errorf(codes.InvalidArgument, "%v", err)
		}

		claims = UserClaims{
			ID:   user.(*models.Users).UserID,
			Role: "ROLE_CLIENT",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(s.Cfg.TokenDuration)).Unix(),
				Issuer:    "Online shop",
			},
		}

	} else if is_email { // если почта

		user, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) // достаю пароль
		if err != nil {
			return &pb.SignInRes{Err: "User not found or error in search"}, nil
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
		if err != nil {                                                                                      // если не совпали
			return &pb.SignInRes{
				Status: int32(codes.InvalidArgument),
				Err:    "wrong password",
				Token:  "",
			}, nil
		}

		claims = UserClaims{
			ID:   user.(*models.Users).UserID,
			Role: "ROLE_CLIENT",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(s.Cfg.TokenDuration)).Unix(),
				Issuer:    "Online shop",
			},
		}

	} else { //вернуть ошибку о некорректности ввода данных
		return &pb.SignInRes{
			Status: int32(codes.InvalidArgument),
			Err:    "invalid data",
			Token:  "",
		}, status.Errorf(codes.InvalidArgument, "")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.Cfg.JWTSecretKey))
	if err != nil {
		return &pb.SignInRes{
			Status: int32(codes.Internal),
			Err:    err.Error(),
		}, status.Errorf(codes.Internal, "%v", err)
	}

	return &pb.SignInRes{
		Status: int32(codes.OK),
		Err:    "",
		Token:  signedToken,
	}, nil
}

// func (s *AuthServer) SignInUserByEmail(_ context.Context, req *pb.SignInReq) (*pb.SignInRes, error) {

// 	user, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) // достаю пароль
// 	if err != nil {
// 		return &pb.SignInRes{
// 			Status: int32(codes.Internal),
// 			Err: "User not found or error in search",
// 		}, status.Errorf(codes.Internal, "%v", err)
// 	}
// 	err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
// 	if err != nil {                                                                                      // если не совпали
// 		return &pb.SignInRes{
// 			Status: int32(codes.InvalidArgument),
// 			Err: "wrong password",
// 			Token: "",
// 			}, status.Errorf(codes.InvalidArgument, "%v", err)
// 	}

// 	claims := UserClaims{
// 		ID:    user.(*models.Users).UserID,
// 		Role:  nil,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(s.Cfg.TokenDuration)).Unix(),
// 			Issuer:    "Online shop",
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString([]byte(s.Cfg.JWTSecretKey))
// 	if err != nil {
// 		return &pb.SignInRes{
// 			Status: int32(codes.Internal),
// 			Err:  err.Error(),
// 		}, status.Errorf(codes.Internal, "%v", err)
// 	}

// 	return &pb.SignInRes{
// 		Status: int32(codes.OK),
// 		Err:  "",
// 		Token:  signedToken,
// 	}, nil

// }

func (s *AuthServer) Validate(ctx context.Context, req *pb.ValidateReq) (*pb.ValidateRes, error) {
	token, err := jwt.ParseWithClaims(req.Token, &UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Cfg.JWTSecretKey), nil
		},
	)
	if err != nil {
		return &pb.ValidateRes{
			Status: int32(codes.PermissionDenied),
			Error:  err.Error(),
			UserId: 0,
		}, status.Errorf(codes.PermissionDenied, "%v", err)
	}
	claims, ok := token.Claims.(*UserClaims)

	if !ok {
		return &pb.ValidateRes{
			Status: int32(codes.Internal),
			Error:  err.Error(),
			UserId: 0,
		}, status.Errorf(codes.Internal, "%v", err)
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return &pb.ValidateRes{
			Status: int32(codes.PermissionDenied),
			Error:  "Token expired",
			UserId: 0,
		}, status.Errorf(codes.PermissionDenied, "%v", err)
	}
	return &pb.ValidateRes{
		Status: int32(codes.OK),
		Error:  "",
		UserId: claims.ID,
		Role:   claims.Role,
	}, err

}
func (s *AuthServer) RegisterMainAdmin(ctx context.Context, req *pb.RegReqAdmin) (*pb.AuthRes, error) {
	return &pb.AuthRes{Err: "success"}, nil
}

func (s *AuthServer) RegisterAdmin(ctx context.Context, req *pb.RegReqAdmin) (*pb.AuthRes, error) {
	return &pb.AuthRes{Err: "success"}, nil
}

func (s *AuthServer) GetReqUser(ctx context.Context, req *pb.GetReq) (*pb.AuthRes, error) {
	return &pb.AuthRes{Err: "success"}, nil
}

func (s *AuthServer) GetAdmin(ctx context.Context, req *pb.GetReq) (*pb.AuthRes, error) {
	return &pb.AuthRes{Err: "success"}, nil
}
