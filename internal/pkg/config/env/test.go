package main

import "fmt"

type testStruct struct {
	name string `env:"NAME" default:"Thinh"`
	age  int    `env:"AGE"`
}

func main() {
	var conf testStruct
	Load(&conf)
	fmt.Println(conf)
}
