package main

import (
	"fmt"
	"log"
	"net"
	"online_shop/repository"
	"online_shop/user-svc/config"
	"online_shop/user-svc/pb"
	"online_shop/user-svc/user_service"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {

	logf, err := os.OpenFile("./user-svc/log/qs.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Cant open log file: %v", err)
	}
	defer func(logf *os.File) {
		err := logf.Close()
		if err != nil {
			log.Fatalf("Cant close log file: %v", err)
		}
	}(logf)
	logger := log.New(logf, "User Service: ", log.Flags())

	SqlDB, Db, err := repository.Conect_to_DB()
	if err != nil {
		log.Fatalf("Cant connect to Database: %v", err)
	}

	cfg, err := config.LoadConfig("./user-svc/config")
	if err != nil {
		logger.Fatalf("Error loading loadConfig: %v", err)
	}

	s := grpc.NewServer()
	srv := user_service.NewUserServiceServer(Db, &cfg)
	pb.RegisterUserServiceServer(s, srv)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		logger.Fatalln("Failed to listing:", err)
	}

	go func() {
		logger.Println("User Svc on", cfg.Port)
		fmt.Println("User Svc on", cfg.Port)
		err := s.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	waitExitSignal(s)

	if err := SqlDB.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("shutting down")
	logger.Println("Shutting down! Bye!")
}

func waitExitSignal(s *grpc.Server) {
	sigCh := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		log.Printf("Signal %s\n", sig)
		s.Stop()
		done <- true
	}()
	<-done
}
