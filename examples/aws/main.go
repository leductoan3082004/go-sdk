package main

import (
	"context"
	"fmt"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"log"
)

func main() {
	service := goservice.New(
		goservice.WithName("demo"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(aws.New("aws")),
	)
	service.OutEnv()
	_ = service.Init()
	logoFile := "logo.png" // put this file on project root to test

	s3 := service.MustGet("aws").(aws.S3)
	url, err := s3.Upload(context.Background(), logoFile, "image")
	if err != nil {
		log.Fatalln(err)
	}

	urls := []string{"image/1687082453071499600.png", "media/1572143325272547000.png"} // put fileKey need remove into array

	if err := s3.DeleteImages(context.Background(), urls); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(url)
}
