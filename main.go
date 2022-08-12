package main

import (
	"fmt"
	"os"

	"github.com/rs/xid"
	"github.com/google/uuid"
)

func main() {
	a := ""
	if len(os.Args) > 1 {
		a = os.Args[1]
	} else {
		a = "xid"
	}
	switch a {
	case "xid":
		guid := xid.New()
		println(guid.String())
		if len(os.Args) > 2 && os.Args[2] == "-i" {
			fmt.Printf(
				"[machine: %x | pid: %d | time: %d | counter: %d]\n",
				guid.Machine(),
				guid.Pid(),
				guid.Time().Unix(),
				guid.Counter(),
			)
		}
	case "uuid":
		println(uuid.New().String())
	default:
		panic("Unknown command")
	}
}
