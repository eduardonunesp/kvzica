package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/eduardonunesp/kvzika/pkg/server"
	"golang.org/x/net/context"
)

var portFlag = flag.Int("port", server.DefaultPort, "")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	s, err := server.NewServer(*portFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	go func() {
		err := s.Run(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server started at port", *portFlag)
	<-interrupt
	log.Println("Server stopped")
}
