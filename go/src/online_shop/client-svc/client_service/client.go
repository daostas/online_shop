package client_service

import (
	"context"
	"online_shop/client-svc/config"
	"online_shop/client-svc/pb"
	"online_shop/repository/models"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/reform.v1"
)

type ClientsServer struct {
	pb.UnimplementedClientsServer
	Db  *reform.DB
	Cfg *config.Config
}

func NewClientsServiceServer(db *reform.DB, cfg *config.Config) *ClientsServer {
	return &ClientsServer{
		Db:  db,
		Cfg: cfg,
	}
}

func (s *ClientsServer) UpdateClientInfo(_ context.Context, req *pb.UpdateClientInfoReq) (*pb.ClientRes, error) {

	// код пока задуман так, что будет использоваться айди пользователя из токена

	user, err := s.Db.SelectOneFrom(models.UsersTable, "where user_id = $1", req.Id) // беру старые данные по айди пользователя
	if err != nil {
		return &pb.ClientRes{Err: "searching data error"}, nil
	}
	user.(*models.Users).UserName = &req.Name // если где-то были изменения, они сохранятся
	user.(*models.Users).Number = &req.Number
	user.(*models.Users).Email = &req.Email
	//user.(*models.Users).Dob = &req.Dob
	user.(*models.Users).Address = &req.Address

	if err := s.Db.Save(user.(*models.Users)); err != nil { // сохраняется в базу данных
		return &pb.ClientRes{Err: "Update data error"}, nil
	}

	return &pb.ClientRes{Err: "success"}, nil
}

func (s *ClientsServer) UpdateClientPass(_ context.Context, req *pb.UpdateClientPassReq) (*pb.ClientRes, error) {

	// код пока задуман так, что будет использоваться айди пользователя из токена

	user, err := s.Db.SelectOneFrom(models.UsersTable, "where user_id = $1", req.Id) // достаю старый пароль по айди пользователя
	if err != nil {
		return &pb.ClientRes{Err: "User not found or error in search"}, nil
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.(*models.Users).UserPassword), []byte(req.Pass)); err == nil { // Сравнение введенного и текущего установленного пароля

		if req.Pass1 == req.Pass2 { // совпадяют ли новый пароль и его повторение

			hashed_pass, err := bcrypt.GenerateFromPassword([]byte(req.Pass1), int(s.Cfg.BcryptCost)) // хэширование нового пароля
			if err != nil {
				return &pb.ClientRes{Err: "err"}, nil
			}
			user.(*models.Users).UserPassword = string(hashed_pass) // замена старого пароля новым хэшированным

			if err := s.Db.Save(user.(*models.Users)); err != nil { // сохраняю в базу
				return &pb.ClientRes{Err: "Password savind error"}, nil
			}
		} else { // новый пароль и его повторение не совпадают
			return &pb.ClientRes{Err: "the new password and its repetition do not match"}, nil
		}
	} else { // введеный старый пароль и текущий не совпали
		return &pb.ClientRes{Err: "the entered old password and the current one did not match"}, nil
	}

	return &pb.ClientRes{Err: "success"}, nil
}

func (s *ClientsServer) DeleteClient(_ context.Context, req *pb.DeleteClientReq) (*pb.ClientRes, error) {

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
			return &pb.ClientRes{Err: "User not found or error in search"}, nil
		}

		basket, err := s.Db.SelectOneFrom(models.BasketsTable, "where basket_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.ClientRes{Err: "Basket not found or error in search"}, nil
		}

		err = s.Db.Delete(basket.(*models.Baskets))
		if err != nil {
			return &pb.ClientRes{Err: "Basket deleting error"}, nil
		}

		favourite, err := s.Db.SelectOneFrom(models.FavouritesTable, "where favourite_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.ClientRes{Err: "Favourite not found or error in search"}, nil
		}

		err = s.Db.Delete(favourite.(*models.Favourites))
		if err != nil {
			return &pb.ClientRes{Err: "Favourite deleting error"}, nil
		}

		err = s.Db.Delete(user.(*models.Users))
		if err != nil {
			return &pb.ClientRes{Err: "User deleting error"}, nil
		}

	} else if is_email { //код по созданию пользователя по почте

		user, err := s.Db.SelectOneFrom(models.UsersTable, "where email = $1", req.Login) //Проверка на не занятость почты
		if err != nil {
			return &pb.ClientRes{Err: "User not found or error in search"}, nil
		}

		basket, err := s.Db.SelectOneFrom(models.BasketsTable, "where basket_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.ClientRes{Err: "Basket not found or error in search"}, nil
		}

		err = s.Db.Delete(basket.(*models.Baskets))
		if err != nil {
			return &pb.ClientRes{Err: "Basket deleting error"}, nil
		}

		favourite, err := s.Db.SelectOneFrom(models.FavouritesTable, "where favourite_id = $1", user.(*models.Users).UserID)
		if err != nil {
			return &pb.ClientRes{Err: "Favourite not found or error in search"}, nil
		}

		err = s.Db.Delete(favourite.(*models.Favourites))
		if err != nil {
			return &pb.ClientRes{Err: "Favourite deleting error"}, nil
		}

		err = s.Db.Delete(user.(*models.Users))
		if err != nil {
			return &pb.ClientRes{Err: "User deleting error"}, nil
		}

	} else { // возврат ошибки о некоректных данных
		return &pb.ClientRes{Err: "unvalid data"}, nil
	}

	return &pb.ClientRes{Err: "success"}, nil
}
