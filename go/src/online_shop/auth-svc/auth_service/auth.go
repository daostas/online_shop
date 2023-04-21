package user_service

import (
	"context"
	"fmt"
	"net/http"
	"online_shop/auth-svc/config"
	"online_shop/auth-svc/pb"
	rep "online_shop/repository"
	"online_shop/repository/models"
	st "online_shop/status"
	"time"

	"gopkg.in/reform.v1"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	tr, err := s.Db.Begin()
	if err != nil {
		return &pb.AuthRes{
			Status: st.StatusInternalServerError,
			Err:    "error in beginning of the transaction: " + fmt.Sprint(err),
		}, nil
	}

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

		_, err := tr.SelectOneFrom(models.UsersTable, "where number = $1", req.Login) //Проверка на не занятость номера
		if err == nil {                                                               // Если занят
			tr.Rollback()
			return &pb.AuthRes{
				Status: st.StatusAlreadyExist,
				Err:    "already exists"}, nil
		}

		hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), int(s.Cfg.BcryptCost)) // хэширование пароля
		if err != nil {
			tr.Rollback()
			return &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "error in creation of password: " + fmt.Sprint(err)}, nil
		}

		user = rep.NewUsers(nil, &req.Login, nil, nil, nil, string(hashed_pass)) // сохранение нового пользователя
		if err := tr.Insert(user); err != nil {
			tr.Rollback()
			return &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "user creation error: " + fmt.Sprint(err)}, nil
		}

	} else if is_email { //код по созданию пользователя по почте

		_, err := tr.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) //Проверка на не занятость почты
		if err == nil {                                                              // Если занята
			tr.Rollback()
			return &pb.AuthRes{
				Status: st.StatusAlreadyExist,
				Err:    "already exists"}, nil
		}

		hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), int(s.Cfg.BcryptCost)) // хэширования пароля
		if err != nil {
			tr.Rollback()
			return &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "error in creation of password: " + fmt.Sprint(err)}, nil
		}

		user = rep.NewUsers(nil, nil, &req.Login, nil, nil, string(hashed_pass)) // сохраненение нового пользователя
		if err := tr.Insert(user); err != nil {
			tr.Rollback()
			return &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "user creation error: " + fmt.Sprint(err)}, nil
		}

	} else { // возврат ошибки о некоректных данных
		tr.Rollback()
		return &pb.AuthRes{
			Status: st.StatusInvalidData,
			Err:    "invalid data"}, nil
	}

	basket := rep.NewBasket(user.UserID)
	if err := tr.Insert(basket); err != nil { // Создаю корзину с айди пользователя
		tr.Rollback()
		return &pb.AuthRes{
			Status: st.StatusInternalServerError,
			Err:    "basket creation error:: " + fmt.Sprint(err)}, nil
	}

	favourite := rep.NewFavourite(user.UserID)
	if err := tr.Insert(favourite); err != nil { // создаю избранное с айди пользователя
		tr.Rollback()
		return &pb.AuthRes{
			Status: st.StatusInternalServerError,
			Err:    "favourite creation error: " + fmt.Sprint(err)}, nil
	}

	tr.Commit()
	return &pb.AuthRes{
		Status: st.StatusOK,
		Err:    "success"}, nil
}

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
			if user == nil {
				return &pb.SignInRes{
					Status: st.StatusRecordNotFound,
					Err:    "user not found"}, nil
			}
			return &pb.SignInRes{
				Status: st.StatusInternalServerError,
				Err:    "error in search"}, nil
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
		if err != nil {                                                                                      // если не совпали
			return &pb.SignInRes{
				Status: st.StatusInvalidArgument,
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

	} else if is_email { // если почта

		user, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) // достаю пароль
		if err != nil {
			if user == nil {
				return &pb.SignInRes{
					Status: st.StatusRecordNotFound,
					Err:    "user not found"}, nil
			}
			return &pb.SignInRes{
				Status: st.StatusInternalServerError,
				Err:    "error in search"}, nil
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
		if err != nil {                                                                                      // если не совпали
			return &pb.SignInRes{
				Status: st.StatusInvalidArgument,
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
			Status: st.StatusInvalidData,
			Err:    "invalid data",
			Token:  "",
		}, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.Cfg.JWTSecretKey))
	if err != nil {
		return &pb.SignInRes{
			Status: st.StatusInternalServerError,
			Err:    "error is signing token",
			Token:  "",
		}, nil
	}

	return &pb.SignInRes{
		Status: http.StatusOK,
		Err:    "success",
		Token:  signedToken,
	}, nil
}

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

func (s *AuthServer) RegisterMainAdmin(_ context.Context, req *pb.RegReqAdmin) (*pb.AuthRes, error) {

	hashed_pass, err := bcrypt.GenerateFromPassword([]byte("password"), int(s.Cfg.BcryptCost)) // хэширование пароля
	if err != nil {
		return &pb.AuthRes{Err: "err"}, nil
	}

	admin := rep.NewAdmin("main", string(hashed_pass), "ROLE_MAIN_ADMIN")

	err = s.Db.Insert(admin)
	if err != nil {
		return &pb.AuthRes{Err: "error in creating of main admin"}, nil
	}
	return &pb.AuthRes{Err: "success"}, nil

}

func (s *AuthServer) RegisterAdmin(_ context.Context, req *pb.RegReqAdmin) (*pb.AuthRes, error) {

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

	var admin *models.Admins
	if only_numbers || is_email { //код по созданию пользоватьеля по номеру

		_, err := s.Db.SelectOneFrom(models.AdminsTable, "where login = $1", req.Login) //Проверка на не занятость номера
		if err == nil {                                                                 // Если занят
			return &pb.AuthRes{
				Status: st.StatusAlreadyExist,
				Err:    "already exists"}, nil
		}

		hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), int(s.Cfg.BcryptCost)) // хэширование пароля
		if err != nil {
			return &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "error in generating password"}, nil
		}

		admin = rep.NewAdmin(req.Login, string(hashed_pass), "ROLE_ADMIN")
		err = s.Db.Insert(admin)
		if err != nil {
			return &pb.AuthRes{
				Status: st.StatusInternalServerError,
				Err:    "error in inserting data in admins table"}, nil
		}
		return &pb.AuthRes{
			Status: st.StatusOK,
			Err:    "success"}, nil

	} else { // возврат ошибки о некоректных данных
		return &pb.AuthRes{
			Status: st.StatusInvalidData,
			Err:    "invalid data"}, nil
	}

}

func (s *AuthServer) SignInAdmin(_ context.Context, req *pb.SignInReq) (*pb.SignInRes, error) {

	var claims UserClaims
	admin, err := s.Db.SelectOneFrom(models.AdminsTable, "where login = $1", req.Login) // достаю пароль
	if err != nil {
		if admin == nil {
			return &pb.SignInRes{
				Status: st.StatusRecordNotFound,
				Err:    "admin not found"}, nil
		}
		return &pb.SignInRes{
			Status: st.StatusInternalServerError,
			Err:    "error in search"}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.(*models.Admins).Password), []byte(req.Password)) // сравниваю пароли
	if err != nil {                                                                                    // если не совпали
		return &pb.SignInRes{
			Status: st.StatusInvalidArgument,
			Err:    "wrong password",
			Token:  "",
		}, nil
	}

	claims = UserClaims{
		ID:   admin.(*models.Admins).AdminID,
		Role: admin.(*models.Admins).Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(s.Cfg.TokenDuration)).Unix(),
			Issuer:    "Online shop",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.Cfg.JWTSecretKey))
	if err != nil {
		return &pb.SignInRes{
			Status: st.StatusInternalServerError,
			Err:    "error in signing of token",
			Token:  "",
		}, nil
	}

	return &pb.SignInRes{
		Status: st.StatusOK,
		Err:    "success",
		Token:  signedToken,
	}, nil
}
