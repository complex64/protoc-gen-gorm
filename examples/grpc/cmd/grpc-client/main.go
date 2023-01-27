package main

import (
	"context"
	"fmt"

	"github.com/complex64/protoc-gen-gorm/examples/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	const addr = "localhost:1234"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	ctp := context.Background()
	client := pb.NewUsersServiceClient(conn)

	{
		fmt.Println("# Create User 'Alice'")
		u, err := client.CreateUser(ctx, &pb.CreateUserRequest{
			User: &pb.User{
				Name:       "users/alice",
				GivenName:  "Alice",
				FamilyName: "Trent",
				Email:      "alice@internet.org",
			},
		})
		if err != nil {
			panic(err)
		}
		Print(u)
		fmt.Println()
	}

	{
		fmt.Println("# Create User 'Bob'")
		u, err := client.CreateUser(ctx, &pb.CreateUserRequest{
			User: &pb.User{
				Name:       "users/bob",
				GivenName:  "Bob",
				FamilyName: "Walter",
				Email:      "bob@internet.org",
			},
		})
		if err != nil {
			panic(err)
		}
		Print(u)
		fmt.Println()
	}

	{
		fmt.Println("# Get User 'Alice'")
		u, err := client.GetUser(ctx, &pb.GetUserRequest{Name: "users/alice"})
		if err != nil {
			panic(err)
		}
		Print(u)
		fmt.Println()
	}

	{
		fmt.Println("# List All Users")
		res, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
		if err != nil {
			panic(err)
		}
		for _, u := range res.Users {
			Print(u)
		}
		fmt.Println()
	}

	{
		fmt.Println("# Update User 'Bob'")
		u, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
			User: &pb.User{
				Name:      "users/bob",
				GivenName: "Bobby",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"given_name",
				},
			},
		})
		if err != nil {
			panic(err)
		}
		Print(u)
		fmt.Println()
	}

	{
		fmt.Println("# Delete User 'Alice'")
		_, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Name: "users/alice"})
		if err != nil {
			panic(err)
		}
		fmt.Printf("=> done\n\n")
	}

	{
		fmt.Println("# List All Users")
		res, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
		if err != nil {
			panic(err)
		}
		for _, u := range res.Users {
			Print(u)
		}
	}
}

func Print(u *pb.User) {
	fmt.Printf(""+
		"=> Name:        %s\n"+
		"   Given Name:  %s\n"+
		"   Family Name: %s\n"+
		"   Email:       %s\n"+
		"   Created:     %s\n"+
		"   Updated:     %s\n",
		u.Name,
		u.GivenName,
		u.FamilyName,
		u.Email,
		u.CreateTime.AsTime().String(),
		u.UpdateTime.AsTime().String(),
	)
}
