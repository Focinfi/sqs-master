package main

import "github.com/Focinfi/gosqs/master"

func main() {
	master.NewService(":8080").Start()
}
