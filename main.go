package main

import (
	"log"
	"os"

	"github.com/Focinfi/sqs/master"
)

func main() {
	addr := os.Getenv("SQS_MASTER_ADDR")
	log.Println("SQS_MASTER_ADDR:", addr)
	if addr == "" {
		panic("Need SQS_MASTER_ADDR environment parameter")
	}

	master.NewService(addr).Start()
}
