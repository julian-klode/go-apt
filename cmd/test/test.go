package main

import "github.com/julian-klode/goapt"
import "os"

func main() {
	/*session := goapt.NewSession()

	apt := session.Cache().Lookup("apt", "native")
	cand := session.Policy().CandidateVersion(apt)

	println("The current candidate for APT is:", cand.VerStr())*/

	cfg := goapt.GetConfig()
	for _, arg := range os.Args {
		println(arg, "=>", cfg.Find(arg))
	}
}
