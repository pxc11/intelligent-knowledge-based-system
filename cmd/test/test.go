package main

import (
	"ikbs/lib/basic"
	"ikbs/lib/config"
	"ikbs/lib/db"
	"ikbs/lib/logger"
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

}
