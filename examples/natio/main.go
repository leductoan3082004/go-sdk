package main

import (
	"context"
	goservice "github.com/leductoan3082004/go-sdk"
	pb "github.com/leductoan3082004/go-sdk/plugin/pubsub"
	"github.com/leductoan3082004/go-sdk/plugin/pubsub/natspb"
	"log"
)

type TestData struct {
	Value int `json:"value"`
}

func main() {
	service := goservice.New(
		goservice.WithName("demo"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(natspb.NewNatsPubSub("nats", "")),
	)
	if err := service.Init(); err != nil {
		log.Fatalln(err)
	}

	natpb := service.MustGet("nats").(pb.Provider)

	log.Println(natpb)

	ch, _ := natpb.Subscribe(context.Background(), "test")
	ch1, _ := natpb.Subscribe(context.Background(), "test")

	for i := 1; i <= 10; i++ {
		natpb.Publish(context.Background(), "test", &pb.Event{
			Data: TestData{Value: i},
		})
	}

	go func() {
		for v := range ch {
			log.Println(string(v.RemoteData))
		}
	}()

	go func() {
		for v := range ch1 {
			log.Println(string(v.RemoteData))
		}
	}()

	//defer close()
	//defer close2()

	service.Start()

	//s3 := service.MustGet("nats").(natspb.NatsOpt)
}
