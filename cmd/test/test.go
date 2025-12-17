package main

import (
	"ikbs/lib/basic"
	"ikbs/lib/logger"
	"time"
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
	logger.Init()

	logger.Info("dwdwdwd", "dwdwd")
	aa := B{
		A: A{
			a: 1,
		},
		B: 2,
	}
	logger.Error(aa, aa)

	time.Sleep(5 * time.Second)

}
