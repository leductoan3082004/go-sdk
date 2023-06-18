package main

import (
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/storage/sdkgorm"
	"gorm.io/gorm"
)

func main() {
	sc := goservice.New(
		goservice.WithName("ahihi"),
		goservice.WithVersion("1.1.1"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", "gorm")),
	)

	_ = sc.Init()
	_ = sc.Start()

	_ = sc.MustGet("gorm").(*gorm.DB)
}
