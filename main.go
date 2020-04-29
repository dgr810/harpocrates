package main

import (
	"github.com/BESTSELLER/harpocrates/cmd"
	"github.com/BESTSELLER/harpocrates/config"
)

var secretJSON string

func main() {
	config.LoadConfig()
	cmd.Execute()
}
