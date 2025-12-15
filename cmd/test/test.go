package main

import (
	"fmt"

	"ikbs/lib/config"
)

func main() {
	cfg, _ := config.LoadConfig()
	fmt.Printf("%+v\n", cfg)
}
