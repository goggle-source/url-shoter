package main

import (
	"fmt"

	"github.com/url-shoter/iternal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	
	//TODO: init logger: slog

	//TODO: init storage: sqlite

	//TODO: init router: chi, chi renger

	//TODO: run server
}