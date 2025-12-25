package main

import (
	"context"
	"ikbs/internal/model"
	"ikbs/lib/basic"
	"ikbs/lib/config"
	"ikbs/lib/db"
	"ikbs/lib/logger"

	"gorm.io/gorm"
)

type A struct {
	a int64
}

type B struct {
	A A
	B int64
}

func main() {
	basic.Init()
	config.Init()
	logger.Init()
	db.Init()
	context2 := context.Background()
	err := gorm.G[model.User](db.GetDb()).Create(context2, &model.User{
		Username: "admin",
		Password: "111",
	})
	if err != nil {
		return
	}
}
