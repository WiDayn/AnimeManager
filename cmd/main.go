package main

import (
	"AnimeManager/internal/config"
	"AnimeManager/internal/service"
)

func main() {
	config.DefaultLoader()

	service.Rename()
}
