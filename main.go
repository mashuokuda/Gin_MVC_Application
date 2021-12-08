package main

import "Gin_MVC/router"

//go:generate  air

func main() {
	r := router.GetRouter()
	r.Run(":3000")
}
