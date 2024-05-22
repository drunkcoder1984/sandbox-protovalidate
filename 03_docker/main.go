package main

import (
	"fmt"
	"log"

	userpb "buf-confv2/gen/go/user"

	"github.com/bufbuild/protovalidate-go"
)

func main() {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("Error protovalidate", err)
	}

	// Nameのバリデーション5文字以上なのでerrorになる
	user := &userpb.User{Name: "hoge"}
	if err := v.Validate(user); err != nil {
		fmt.Printf("Name: %s error: %v\n\n", user.Name, err)
	} else {
		fmt.Printf("Name: %s ok\n\n", user.Name)
	}

	// Nameのバリデーション5文字以上なのでokになる
	user = &userpb.User{Name: "hogehoge"}
	if err := v.Validate(user); err != nil {
		fmt.Printf("Name: %s error: %v\n\n", user.Name, err)
	} else {
		fmt.Printf("Name: %s ok\n\n", user.Name)
	}
}
