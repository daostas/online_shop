package client_service

import (
	"context"
	"fmt"
	"log"

	"net"

	"online_shop/client-svc/config"
	"online_shop/client-svc/pb"
	"online_shop/repository"
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
func TestClientServices(t *testing.T) {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		fmt.Printf("Error loading loadConfig: %v", err)
		return
	}
	SqlDB, Db, err := repository.ConnectToDb()
	if err != nil {
		fmt.Printf("Cant connect to Database: %v", err)
	}

	clientsrv := NewClientsServiceServer(Db, &cfg)
	pb.RegisterClientsServer(s, clientsrv)
	clientgroupssrv := NewClientGroupsServer(Db, &cfg)
	pb.RegisterClientGroupsServer(s, clientgroupssrv)

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

	// userclient := pb.NewUsersClient(conn)

	// t.Run("UpdateUserInfoTest", func(t *testing.T) {

	// 	t.Run("UpdateUserInfoTest1", func(t *testing.T) {

	// 		user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
	// 		if err != nil {
	// 			t.Errorf("UpdateUserInfoTest1 failed: %v", err)
	// 		}

	// 		req := &pb.UpdateUserInfoReq{
	// 			Id:      user.(*models.Users).UserID,
	// 			Name:    "Stas",
	// 			Number:  "",
	// 			Email:   "daostas@gmail.com",
	// 			Dob:     "",
	// 			Address: "",
	// 		}

	// 		res, _ := userclient.UpdateUserInfo(ctx, req)

	// 		if res.Err != "success" {
	// 			t.Errorf("UpdateUserInfoTest1 failed: %v", res.Err)
	// 		}

	// 	})

	// })

	// t.Run("UpdateUserPassTest", func(t *testing.T) {

	// 	t.Run("UpdateUserPassTest1", func(t *testing.T) {

	// 		user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
	// 		if err != nil {
	// 			t.Errorf("UpdateUserPassTest1 failed: %v", err)
	// 		}

	// 		req := &pb.UpdateUserPassReq{
	// 			Id:    user.(*models.Users).UserID,
	// 			Pass:  "password123",
	// 			Pass1: "password123",
	// 			Pass2: "password123",
	// 		}

	// 		res, err := userclient.UpdateUserPass(ctx, req)
	// 		if err != nil {
	// 			t.Errorf("UpdateUserPassTest1 failed: %v", err)
	// 		}

	// 		if res.Err != "the entered old password and the current one did not match" {
	// 			t.Errorf("UpdateUserPassTest1 failed: %v", res.Err)
	// 		}

	// 	})

	// 	t.Run("UpdateUserPassTest2", func(t *testing.T) {

	// 		user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
	// 		if err != nil {
	// 			t.Errorf("UpdateUserPassTest2 failed: %v", err)
	// 		}

	// 		req := &pb.UpdateUserPassReq{
	// 			Id:    user.(*models.Users).UserID,
	// 			Pass:  "password",
	// 			Pass1: "password123",
	// 			Pass2: "password1234",
	// 		}

	// 		res, _ := userclient.UpdateUserPass(ctx, req)

	// 		if res.Err != "the new password and its repetition do not match" {
	// 			t.Errorf("UpdateUserPassTest2 failed: %v", res.Err)
	// 		}

	// 	})

	// 	t.Run("UpdateUserPassTest3", func(t *testing.T) {

	// 		user, err := Db.SelectOneFrom(models.UsersTable, "where email = $1", "daostas@gmail.com")
	// 		if err != nil {
	// 			t.Errorf("UpdateUserPassTest3 failed: %v", err)
	// 		}

	// 		req := &pb.UpdateUserPassReq{
	// 			Id:    user.(*models.Users).UserID,
	// 			Pass:  "password",
	// 			Pass1: "password123",
	// 			Pass2: "password123",
	// 		}

	// 		res, _ := userclient.UpdateUserPass(ctx, req)

	// 		if res.Err != "success" {
	// 			t.Errorf("UpdateUserPassTest3 failed: %v", res.Err)
	// 		}

	// 	})

	// })

	// t.Run("DeleteUserTest", func(t *testing.T) {

	// 	t.Run("DeleteUserTest1", func(t *testing.T) {

	// 		req := &pb.DeleteUserReq{
	// 			Login: "daostas@gmail.com",
	// 		}

	// 		res, _ := userclient.DeleteUser(ctx, req)

	// 		if res.Err != "success" {
	// 			t.Errorf("DeleteUserTest1 failed: %v", res.Err)
	// 		}

	// 	})

	// 	t.Run("DeleteUserTest2", func(t *testing.T) {

	// 		req := &pb.DeleteUserReq{
	// 			Login: "87019852218",
	// 		}

	// 		res, _ := userclient.DeleteUser(ctx, req)

	// 		if res.Err != "success" {
	// 			t.Errorf("DeleteUserTest2 failed: %v", res.Err)
	// 		}

	// 	})
	// })

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}

}
