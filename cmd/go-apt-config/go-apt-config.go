package main

import (
	"log"
	"os"

	"github.com/julian-klode/go-apt"
)

func main() {
	/*session := goapt.NewSession()

	apt := session.Cache().Lookup("apt", "native")
	cand := session.Policy().CandidateVersion(apt)

	println("The current candidate for APT is:", cand.VerStr())*/
	session := &apt.Session{}
	cfg, err := session.Config()
	if err != nil {
		log.Fatal("Could not load configuration:", err)
	}
	for _, arg := range os.Args {
		println(arg, "=>", cfg.Find(arg))
	}
}