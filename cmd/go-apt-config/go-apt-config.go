package main

import (
	"log"
	"os"

	"github.com/julian-klode/go-apt"
)

func main() {
	session, err := apt.NewSession()
	if err != nil {
		log.Fatal("Could not load session:", err)
	}
	cfg, err := session.Config()
	if err != nil {
		log.Fatal("Could not load configuration:", err)
	}
	for _, arg := range os.Args {
		println(arg, "=>", cfg.Find(arg))
	}
}
