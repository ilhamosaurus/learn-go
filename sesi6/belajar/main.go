package main

import "belajar/router"

func main() {
	PORT := ":5003"

	router.StartServer().Run(PORT)
}
