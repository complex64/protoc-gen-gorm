package main

import (
	"context"
	"net"

	"github.com/complex64/protoc-gen-gorm/examples/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&pb.UserModel{}); err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	impl := &server{db: db}
	pb.RegisterUsersServiceServer(srv, impl)

	const addr = "localhost:1234"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

var _ pb.UsersServiceServer = (*server)(nil)

type server struct {
	db *gorm.DB
	pb.UnimplementedUsersServiceServer
}

func (s server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := (&pb.User{Name: req.Name}).WithDB(s.db).Get(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := (&pb.User{}).WithDB(s.db).List(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListUsersResponse{Users: users}, nil
}

func (s server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user, err := req.User.WithDB(s.db).Create(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	err := req.User.WithDB(s.db).Patch(ctx, req.UpdateMask)
	if err != nil {
		return nil, err
	}
	return s.GetUser(ctx, &pb.GetUserRequest{Name: req.User.Name})
}

func (s server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	user := &pb.User{Name: req.Name}
	if err := user.WithDB(s.db).Delete(ctx); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
