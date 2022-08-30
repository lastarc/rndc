package main

import (
	"fmt"

	"github.com/rs/xid"
	"github.com/alecthomas/kong"
	"github.com/google/uuid"
)

var cli struct {
	Xid struct {
		Info bool `help:"Print details" short:"i"`
	} `cmd:"" name:"xid" help:"random xid"`

	Uuid struct {
	} `cmd:"" name:"uuid" help:"random uuid"`
}

func main() {
	ctx := kong.Parse(
		&cli,
		kong.Name("rndc"),
		kong.Description("cli tool for random strings (xid, uuid)"),
	)
	switch ctx.Command() {
	case "xid":
		guid := xid.New()
		println(guid.String())
		if cli.Xid.Info {
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
		panic(ctx.Command())
	}
}
