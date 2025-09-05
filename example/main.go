package main

import (
	"fmt"
	"os"

	"github.com/nxtgo/env"
)

func init() {
	err := env.Load("")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(".env file is loaded :D")
	fmt.Println(os.Getenv("BASED"))
}
