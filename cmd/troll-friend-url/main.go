package main

import (
	"fmt"

	"github.com/DevGiV/troll-friend-url/internal/config"
)

func main() {
	//init config

	cfg := config.MustLoad()

	fmt.Println(cfg)

	//init storage

	//init router

	//run server
}
