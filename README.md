# rndc

```
$ rndc -h
Usage: rndc <command>

cli tool for random strings (xid, uuid, hex, ascii, base62, base64)

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  xid
    random xid

  uuid
    random uuid

  hex
    random hex string

  ascii
    random ascii string

  base62
    random base62 string

  base64
    random base64 string

Run "rndc <command> --help" for more information on a command.
```

## Building & Installation

```bash
# clone the source code
$ git clone https://github.com/lastarc/rndc.git
# change to the directory
$ cd rndc
# build the executable
$ go build
# run from the directory...
$ ./rndc uuid
# or install globally
$ go install
$ rndc xid
```

**Note**: for global installation to work $GOPATH/bin (or $HOME/go/bin if the GOPATH environment variable is not set) must be in $PATH
