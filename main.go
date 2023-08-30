package main

import (
	"d_gita_be/config"
	"d_gita_be/route"
)

func main() {

	config.InitDB()
	route.InitRoute()
}
