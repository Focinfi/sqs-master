package main

import "github.com/Focinfi/sqs/master"

func main() {
	master.NewService(":8080").Start()
}
