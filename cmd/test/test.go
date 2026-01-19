package main

import (
	"fmt"
	"ikbs/lib/basic"
	"ikbs/lib/config"
	"ikbs/lib/db"
	"ikbs/lib/logger"
	"reflect"
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

	type User struct {
		Name string `json:"name,omitempty" binding:"required" a:"bb"`
	}

	t := reflect.TypeOf(User{})

	field, _ := t.FieldByName("Name")

	jsonTag := field.Tag.Get("json")
	a := field.Tag.Get("a")

	fmt.Println(jsonTag)
	fmt.Println(a)

}
