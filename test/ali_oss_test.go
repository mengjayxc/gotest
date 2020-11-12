package test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"testing"
)

func TestCreateAliClient(t *testing.T) {
	endPoint := ""
	accessKey := ""
	accessSecret := ""
	client, err := oss.New(
		endPoint,
		accessKey,
		accessSecret,
	)

	if err != nil {
		fmt.Printf("create client err: %s\n", err.Error())
	}

	imageURL := fmt.Sprintf("https://%s.%s/%s",
		"abc",
		endPoint,
		"filename")

	fmt.Printf("client: %v\n", client)
	fmt.Printf("image: %s\n", imageURL)
}
