package main

//go:generate rm -f gorm.db
//go:generate find pb -type f -name *.go -delete
//go:generate protoc -Iprotos --go_out=./ --gorm_out=./ protos/models.proto

import (
	"fmt"

	"github.com/complex64/protoc-gen-gorm/examples/protoc/pb"
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

	first := new(pb.UserModel)
	if err := db.First(first).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			alice := &pb.UserModel{Name: "Alice"}
			if err := db.Create(alice).Error; err != nil {
				panic(err)
			}
			fmt.Printf("Created first user: %s\n", alice.Name)
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("The first user is: %s\n", first.Name)
	}
}
