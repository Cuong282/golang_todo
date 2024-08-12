package main

import (
	"flag"
	"fmt"
	"log"
	"main/api"
	"main/storage"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "the server address")
	flag.Parse()
	store := storage.NewMemoryStorage()
	server := api.NewServer(*listenAddr, store)
	fmt.Println("server runing 3000:", *listenAddr)

	log.Fatal(server.Start())
	fmt.Println("cuong")

}
