package user_service

import (
	"context"
	"fmt"
	"log"

	"net"

	"online_shop/repository"
	"online_shop/repository/models"
	"online_shop/user-svc/config"
	"online_shop/user-svc/pb"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
func TestUserServices(t *testing.T) {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		fmt.Printf("Error loading loadConfig: %v", err)
		return
	}
	SqlDB, Db, err := repository.Conect_to_DB()
	if err != nil {
		fmt.Printf("Cant connect to Database: %v", err)
	}

	srv := NewUserServiceServer(Db, &cfg)

	pb.RegisterUserServiceServer(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
	ctx := context.Background()

	// pb.RegisterUserServiceServer(s, srv)

	// lis, err := net.Listen("tcp", cfg.Port)
	// if err != nil {
	// 	logger.Fatalln("Failed to listing:", err)
	// }

	// go func() {
	// 	logger.Println("User Svc on", cfg.Port)
	// 	fmt.Println("User Svc on", cfg.Port)
	// 	err := s.Serve(lis)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	t.Run("RegisterTest", func(t *testing.T) {
		t.Run("RegisterTest1", func(t *testing.T) {
			req := &pb.RegUserReq{
				Login:    "daostas@gmail.com",
				Password: "password",
			}

			res, _ := client.RegisterUser(ctx, req)

			if res.Err != "success" {
				t.Errorf("RegisterTest1 failed: %v", res.Err)
			}

		})

		t.Run("RegisterTest2", func(t *testing.T) {
			req := &pb.RegUserReq{
				Login:    "daostas@gmail.com",
				Password: "password",
			}

			res, _ := srv.RegisterUser(ctx, req)

			if res.Err != "already exists" {
				t.Errorf("RegisterTest2 failed: %v", res.Err)
			}

		})

		t.Run("RegisterTest3", func(t *testing.T) {
			req := &pb.RegUserReq{
				Login:    "dao11stasgmail.com",
				Password: "password",
			}

			res, _ := client.RegisterUser(ctx, req)

			if res.Err != "invalid data" {
				t.Errorf("RegisterTest3 failed: %v", res.Err)
			}

		})

		t.Run("RegisterTest4", func(t *testing.T) {
			req := &pb.RegUserReq{
				Login:    "87019852218",
				Password: "password",
			}

			res, _ := client.RegisterUser(ctx, req)

			if res.Err != "success" {
				t.Errorf("RegisterTest4 failed: %v", res.Err)
			}

		})
	})

	t.Run("SignInTests", func(t *testing.T) {

		t.Run("SignInTest1", func(t *testing.T) {
			req := &pb.SignInReq{
				Login:    "daostas@gmail.com",
				Password: "password",
			}

			res, _ := client.SignInUser(ctx, req)

			if res.Err != "success" {
				t.Errorf("SignInTest1 failed: %v", res.Err)
			}

		})
		t.Run("SignInTest2", func(t *testing.T) {
			req := &pb.SignInReq{
				Login:    "daostas@gmail.com",
				Password: "password123",
			}

			res, _ := client.SignInUser(ctx, req)

			if res.Err != "wrong password" {
				t.Errorf("SignInTest2 failed: %v", res.Err)
			}

		})

		t.Run("SignInTest3", func(t *testing.T) {
			req := &pb.SignInReq{
				Login:    "87019852218",
				Password: "password",
			}

			res, _ := client.SignInUser(ctx, req)

			if res.Err != "success" {
				t.Errorf("SignInTest3 failed: %v", res.Err)
			}

		})

		t.Run("SignInTest4", func(t *testing.T) {
			req := &pb.SignInReq{
				Login:    "da12ostasgmail.com",
				Password: "password",
			}

			res, _ := client.SignInUser(ctx, req)

			if res.Err != "invalid data" {
				t.Errorf("SignInTest4 failed: %v", res.Err)
			}

		})

	})

	t.Run("UpdateUserInfoTest", func(t *testing.T) {

		t.Run("UpdateUserInfoTest1", func(t *testing.T) {

			user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
			if err != nil {
				t.Errorf("UpdateUserInfoTest1 failed: %v", err)
			}

			req := &pb.UpdateUserInfoReq{
				Id:      user.(*models.Users).UserID,
				Name:    "Stas",
				Number:  "",
				Email:   "daostas@gmail.com",
				Dob:     "",
				Address: "",
			}

			res, _ := client.UpdateUserInfo(ctx, req)

			if res.Err != "success" {
				t.Errorf("UpdateUserInfoTest1 failed: %v", res.Err)
			}

		})

	})

	t.Run("UpdateUserPassTest", func(t *testing.T) {

		t.Run("UpdateUserPassTest1", func(t *testing.T) {

			user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
			if err != nil {
				t.Errorf("UpdateUserPassTest1 failed: %v", err)
			}

			req := &pb.UpdateUserPassReq{
				Id:    user.(*models.Users).UserID,
				Pass:  "password123",
				Pass1: "password123",
				Pass2: "password123",
			}

			res, err := client.UpdateUserPass(ctx, req)
			if err != nil {
				t.Errorf("UpdateUserPassTest1 failed: %v", err)
			}

			if res.Err != "the entered old password and the current one did not match" {
				t.Errorf("UpdateUserPassTest1 failed: %v", res.Err)
			}

		})

		t.Run("UpdateUserPassTest2", func(t *testing.T) {

			user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
			if err != nil {
				t.Errorf("UpdateUserPassTest2 failed: %v", err)
			}

			req := &pb.UpdateUserPassReq{
				Id:    user.(*models.Users).UserID,
				Pass:  "password",
				Pass1: "password123",
				Pass2: "password1234",
			}

			res, _ := client.UpdateUserPass(ctx, req)

			if res.Err != "the new password and its repetition do not match" {
				t.Errorf("UpdateUserPassTest2 failed: %v", res.Err)
			}

		})

		t.Run("UpdateUserPassTest3", func(t *testing.T) {

			user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
			if err != nil {
				t.Errorf("UpdateUserPassTest3 failed: %v", err)
			}

			req := &pb.UpdateUserPassReq{
				Id:    user.(*models.Users).UserID,
				Pass:  "password",
				Pass1: "password123",
				Pass2: "password123",
			}

			res, _ := client.UpdateUserPass(ctx, req)

			if res.Err != "success" {
				t.Errorf("UpdateUserPassTest3 failed: %v", res.Err)
			}

		})

	})

	t.Run("DeleteUserTest", func(t *testing.T) {

		t.Run("DeleteUserTest1", func(t *testing.T) {

			req := &pb.DeleteUserReq{
				Login: "daostas@gmail.com",
			}

			res, _ := client.DeleteUser(ctx, req)

			if res.Err != "success" {
				t.Errorf("DeleteUserTest1 failed: %v", res.Err)
			}

		})

		t.Run("DeleteUserTest2", func(t *testing.T) {

			req := &pb.DeleteUserReq{
				Login: "87019852218",
			}

			res, _ := client.DeleteUser(ctx, req)

			if res.Err != "success" {
				t.Errorf("DeleteUserTest2 failed: %v", res.Err)
			}

		})
	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
