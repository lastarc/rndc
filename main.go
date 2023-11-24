package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/google/uuid"
	"github.com/rs/xid"
	"github.com/thanhpk/randstr"
)

var cli struct {
	Xid struct {
		Info bool `help:"Print details" short:"i"`
	} `cmd:"" name:"xid" help:"random xid"`

	Uuid struct {
	} `cmd:"" name:"uuid" help:"random uuid"`

	Hex struct {
		Length int `help:"length" short:"l" default:"4"`
	} `cmd:"" name:"hex" help:"random hex string"`

	Ascii struct {
		Length int `help:"length" short:"l" default:"8"`
	} `cmd:"" name:"ascii" help:"random ascii string"`

	Base62 struct {
		Length int `help:"length" short:"l" default:"8"`
	} `cmd:"" name:"base62" help:"random base62 string"`

	Base64 struct {
		Length int `help:"length" short:"l" default:"8"`
	} `cmd:"" name:"base64" help:"random base64 string"`
}

func main() {
	ctx := kong.Parse(
		&cli,
		kong.Name("rndc"),
		kong.Description("cli tool for random strings (xid, uuid, hex, ascii, base62, base64)"),
	)
	switch ctx.Command() {
	case "xid":
		guid := xid.New()
		fmt.Println(guid.String())
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
		fmt.Println(uuid.New().String())
	case "hex":
		fmt.Println(randstr.Hex(cli.Hex.Length))
	case "ascii":
		fmt.Println(randstr.String(cli.Ascii.Length))
	case "base62":
		fmt.Println(randstr.Base62(cli.Base62.Length))
	case "base64":
		fmt.Println(randstr.Base64(cli.Base64.Length))
	default:
		panic(ctx.Command())
	}
}
