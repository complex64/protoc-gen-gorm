package main

//go:generate rm -f gorm.db
//go:generate rm -f pb/*.pb.go
//go:generate buf generate protos
//go:generate rm -rf pb/gorm

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/complex64/protoc-gen-gorm/examples/buf/pb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&pb.FileModel{}); err != nil {
		panic(err)
	}

	// Truncate
	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&pb.FileModel{}).Error; err != nil {
		panic(err)
	}

	fmt.Printf("Inserting...\n")

	// Insert all files in the working directory into the database.
	if err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		f := &pb.FileModel{
			Uuid:  uuid.NewString(),
			Name:  path,
			Bytes: info.Size(),
		}
		if err := db.Create(f).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		panic(err)
	}

	fmt.Printf("Files in database:\n")

	var all []pb.FileModel
	if err := db.Order("bytes asc").Find(&all).Error; err != nil {
		panic(err)
	}

	for _, file := range all {
		fmt.Printf(
			"%s\t%d\t%s\t%s\t%s\n",
			file.Uuid,
			file.Bytes,
			file.Created.Format(time.Kitchen),
			file.Updated.Format(time.Kitchen),
			file.Name,
		)
	}
}
