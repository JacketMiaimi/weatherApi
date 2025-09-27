package main

import (
	"fmt"

	"go.mod/internal/config"
)

func main() {
	cfg := config.LoadConfig("C:\\IT\\Go\\petProject\\unfinished\\weatherApi\\configs\\local.yaml")
	fmt.Println(cfg.Adress)
}