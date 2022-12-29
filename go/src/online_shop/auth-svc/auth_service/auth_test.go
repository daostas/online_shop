package user_service

import (
	"context"
	"fmt"
	"log"

	"net"

	"online_shop/auth-svc/config"
	"online_shop/auth-svc/pb"
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
	SqlDB, Db, err := repository.Conect_to_DB()
	if err != nil {
		fmt.Printf("Cant connect to Database: %v", err)
	}

	authsrv := NewAuthServer(Db, &cfg)
	pb.RegisterAuthServer(s, authsrv)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	authclient := pb.NewAuthClient(conn)

	t.Run("RegisterTest", func(t *testing.T) {
		t.Run("RegisterTest1", func(t *testing.T) {
			req := &pb.RegReq{
				Login:    "daostas@gmail.com",
				Password: "password",
			}

			res, _ := authclient.RegisterUser(ctx, req)

			if res.Err != "success" {
				t.Errorf("RegisterTest1 failed: %v", res.Err)
			}

		})

		t.Run("RegisterTest2", func(t *testing.T) {
			req := &pb.RegReq{
				Login:    "daostas@gmail.com",
				Password: "password",
			}

			res, _ := authclient.RegisterUser(ctx, req)

			if res.Err != "already exists" {
				t.Errorf("RegisterTest2 failed: %v", res.Err)
			}

		})

		t.Run("RegisterTest3", func(t *testing.T) {
			req := &pb.RegReq{
				Login:    "dao11stasgmail.com",
				Password: "password",
			}

			res, _ := authclient.RegisterUser(ctx, req)

			if res.Err != "invalid data" {
				t.Errorf("RegisterTest3 failed: %v", res.Err)
			}

		})

		t.Run("RegisterTest4", func(t *testing.T) {
			req := &pb.RegReq{
				Login:    "87019852218",
				Password: "password",
			}

			res, _ := authclient.RegisterUser(ctx, req)

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

			res, _ := authclient.SignInUser(ctx, req)

			if res.Err != "" {
				t.Errorf("SignInTest1 failed: %v", res.Err)
			}

		})
		t.Run("SignInTest2", func(t *testing.T) {
			req := &pb.SignInReq{
				Login:    "daostas@gmail.com",
				Password: "password123",
			}

			res, _ := authclient.SignInUser(ctx, req)

			if res.Err != "wrong password" {
				t.Errorf("SignInTest2 failed: %v", res.Err)
			}

		})

		t.Run("SignInTest3", func(t *testing.T) {
			req := &pb.SignInReq{
				Login:    "87019852218",
				Password: "password",
			}

			res, _ := authclient.SignInUser(ctx, req)

			if res.Err != "" {
				t.Errorf("SignInTest3 failed: %v", res.Err)
			}

		})

	})

	if err := SqlDB.Close(); err != nil {
		t.Errorf("Cant close Database: %v", err)
	}
}
