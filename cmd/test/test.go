package main

import "ikbs/lib/logger"

type A struct {
	a int64
}

type B struct {
	A A
	b int64
}

func main() {

	logger.Info("dwdwdwd", "dwdwd")
	aa := B{
		A: A{
			a: 1,
		},
		b: 2,
	}
	logger.Error(aa, aa)
}
