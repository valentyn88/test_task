package main

import (
	"fmt"
	"net"
	"os"
	"test_task/consumer/storage"
	"test_task/user"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type userServer struct {
	storage storage.Storager
}

func (s userServer) Add(ctx context.Context, u *user.User) (*user.User, error) {
	usr := s.storage.AddRow(*u)
	return &usr, nil
}

func main() {
	srv := grpc.NewServer()
	users := userServer{storage: storage.New()}
	user.RegisterUsersServer(srv, users)
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not listen to :9090 %s", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, srv.Serve(l))
	os.Exit(1)
}
