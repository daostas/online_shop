package user_service

import (
	"context"
	rep "online_shop/repository"
	"online_shop/repository/models"
	"online_shop/user-svc/config"
	"online_shop/user-svc/pb"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/reform.v1"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewUserServiceServer(db *reform.DB, cfg *config.Config) *UserServiceServer {
	return &UserServiceServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *UserServiceServer) RegisterUser(_ context.Context, req *pb.RegUserReq) (*pb.UserRes, error) {
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
			return &pb.UserRes{Err: "already exists"}, nil
		}

		hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), int(s.Cfg.BcryptCost)) // хэширование пароля
		if err != nil {
			return &pb.UserRes{Err: "err"}, nil
		}

		user = rep.NewUsers(nil, &req.Login, nil, nil, nil, string(hashed_pass)) // сохранение нового пользователя
		if err := s.Db.Insert(user); err != nil {
			return &pb.UserRes{Err: "user creation error"}, nil
		}

		// user2, err := s.Db.SelectOneFrom(models.UsersTable, "where number = $1", req.Login) // достаю айди нового пользователя
		// if err != nil {
		// 	return &pb.UserRes{Err: "err"}, nil
		// }

	} else if is_email { //код по созданию пользователя по почте

		_, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) //Проверка на не занятость почты
		if err == nil {                                                                // Если занята
			return &pb.UserRes{Err: "already exists"}, nil
		}

		hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), int(s.Cfg.BcryptCost)) // хэширования пароля
		if err != nil {
			return &pb.UserRes{Err: "err"}, nil
		}

		user = rep.NewUsers(nil, nil, &req.Login, nil, nil, string(hashed_pass)) // сохраненение нового пользователя
		if err := s.Db.Insert(user); err != nil {
			return &pb.UserRes{Err: "user creation error"}, nil
		}

	} else { // возврат ошибки о некоректных данных
		return &pb.UserRes{Err: "invalid data"}, nil
	}

	basket := rep.NewBasket(user.UserID)
	if err := s.Db.Insert(basket); err != nil { // Создаю корзину с айди пользователя
		return &pb.UserRes{Err: "basket creation error"}, nil
	}

	favourite := rep.NewFavourite(user.UserID)
	if err := s.Db.Insert(favourite); err != nil { // создаю избранное с айди пользователя
		return &pb.UserRes{Err: "favourite creation error"}, nil
	}

	return &pb.UserRes{Err: "success"}, nil
}

func (s *UserServiceServer) SignInUser(_ context.Context, req *pb.SignInReq) (*pb.UserRes, error) {

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
			return &pb.UserRes{Err: "User not found or error in search"}, nil
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
		if err != nil {                                                                                      // если не совпали
			return &pb.UserRes{Err: "wrong password"}, nil
		}

	} else if is_email { // если почта

		user, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) // достаю пароль
		if err != nil {
			return &pb.UserRes{Err: "User not found or error in search"}, nil
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Password)) // сравниваю пароли
		if err != nil {                                                                                      // если не совпали
			return &pb.UserRes{Err: "wrong password"}, nil
		}

	} else { //вернуть ошибку о некорректности ввода данных
		return &pb.UserRes{Err: "invalid data"}, nil
	}

	return &pb.UserRes{Err: "success"}, nil
}

func (s *UserServiceServer) UpdateUserInfo(_ context.Context, req *pb.UpdateUserInfoReq) (*pb.UserRes, error) {

	// код пока задуман так, что будет использоваться айди пользователя из токена

	user, err := s.Db.SelectOneFrom(models.UsersTable, "where user_id = $1", req.Id) // беру старые данные по айди пользователя
	if err != nil {
		return &pb.UserRes{Err: "searching data error"}, nil
	}
	user.(*models.Users).UserName = &req.Name // если где-то были изменения, они сохранятся
	user.(*models.Users).Number = &req.Number
	user.(*models.Users).Email = &req.Email
	//user.(*models.Users).Dob = &req.Dob
	user.(*models.Users).Address = &req.Address

	if err := s.Db.Save(user.(*models.Users)); err != nil { // сохраняется в базу данных
		return &pb.UserRes{Err: "Update data error"}, nil
	}

	return &pb.UserRes{Err: "success"}, nil
}

func (s *UserServiceServer) UpdateUserPass(_ context.Context, req *pb.UpdateUserPassReq) (*pb.UserRes, error) {

	// код пока задуман так, что будет использоваться айди пользователя из токена

	user, err := s.Db.SelectOneFrom(models.UsersTable, "where user_id = $1", req.Id) // достаю старый пароль по айди пользователя
	if err != nil {
		return &pb.UserRes{Err: "User not found or error in search"}, nil
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Pass)); err == nil { // Сравнение введенного и текущего установленного пароля

		if req.Pass1 == req.Pass2 { // совпадяют ли новый пароль и его повторение

			hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Pass1), int(s.Cfg.BcryptCost)) // хэширование нового пароля
			if err != nil {
				return &pb.UserRes{Err: "err"}, nil
			}
			user.(*models.Users).UserPassword = string(hashed_pass) // замена старого пароля новым хэшированным

			if err := s.Db.Save(user.(*models.Users)); err != nil { // сохраняю в базу
				return &pb.UserRes{Err: "Password savind error"}, nil
			}
		} else { // новый пароль и его повторение не совпадают
			return &pb.UserRes{Err: "the new password and its repetition do not match"}, nil
		}
	} else { // введеный старый пароль и текущий не совпали
		return &pb.UserRes{Err: "the entered old password and the current one did not match"}, nil
	}

	return &pb.UserRes{Err: "success"}, nil
}

func (s *UserServiceServer) DeleteUser(_ context.Context, req *pb.DeleteUserReq) (*pb.UserRes, error) {

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

	if only_numbers { //код по созданию пользоватьеля по номеру

		user, err := s.Db.SelectOneFrom(models.UsersTable, "where number = $1", req.Login) //Проверка на не занятость номера
		if err != nil {
			return &pb.UserRes{Err: "User not found or error in search"}, nil
		}

		basket, err := s.Db.SelectOneFrom(models.BasketsTable, "where basket_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.UserRes{Err: "Basket not found or error in search"}, nil
		}

		err = s.Db.Delete(basket.(*models.Baskets))
		if err != nil {
			return &pb.UserRes{Err: "Basket deleting error"}, nil
		}

		favourite, err := s.Db.SelectOneFrom(models.FavouritesTable, "where favourite_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.UserRes{Err: "Favourite not found or error in search"}, nil
		}

		err = s.Db.Delete(favourite.(*models.Favourites))
		if err != nil {
			return &pb.UserRes{Err: "Favourite deleting error"}, nil
		}

		err = s.Db.Delete(user.(*models.Users))
		if err != nil {
			return &pb.UserRes{Err: "User deleting error"}, nil
		}

	} else if is_email { //код по созданию пользователя по почте

		user, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) //Проверка на не занятость почты
		if err != nil {
			return &pb.UserRes{Err: "User not found or error in search"}, nil
		}

		basket, err := s.Db.SelectOneFrom(models.BasketsTable, "where basket_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.UserRes{Err: "Basket not found or error in search"}, nil
		}

		err = s.Db.Delete(basket.(*models.Baskets))
		if err != nil {
			return &pb.UserRes{Err: "Basket deleting error"}, nil
		}

		favourite, err := s.Db.SelectOneFrom(models.FavouritesTable, "where favourite_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.UserRes{Err: "Favourite not found or error in search"}, nil
		}

		err = s.Db.Delete(favourite.(*models.Favourites))
		if err != nil {
			return &pb.UserRes{Err: "Favourite deleting error"}, nil
		}

		err = s.Db.Delete(user.(*models.Users))
		if err != nil {
			return &pb.UserRes{Err: "User deleting error"}, nil
		}

	} else { // возврат ошибки о некоректных данных
		return &pb.UserRes{Err: "unvalid data"}, nil
	}

	return &pb.UserRes{Err: "success"}, nil
}

// func (s *UserServiceServer) AddToBasket(_ context.Context, req *pb.AddToBasketReq) (*pb.UserRes, error) {

// 	// код пока задуман так, что будет использоваться айди пользователя из токена
// 	basket, err := s.Db.SelectOneFrom(models.BasketsTable, "where user_id = $1", req.UserId)
// 	if err!= nil {
// 		return &pb.UserRes{Err: "Basket not found or error in search"}, nil
// 	}
// 	record := rep.NewBasketProduct(basket.(*models.Baskets).BasketID, req.ProdId)
// 	if err := s.Db.Insert(record); err != nil { // сохраняется в базу данных
// 		return &pb.UserRes{Err: "Save data error"}, nil
// 	}

// 	return &pb.UserRes{Err: "success"}, nil
// }
