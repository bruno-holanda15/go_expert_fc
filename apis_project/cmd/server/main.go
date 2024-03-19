package main

import "github.com/bruno-holanda15/go_expert_fc/APIs_project/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}

