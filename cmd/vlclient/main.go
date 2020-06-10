package main

import (
	"flag"
	"fmt"

	"github.com/superloach/vlclient"
)

var (
	base   = flag.String("base", "", "base url")
	passwd = flag.String("passwd", "", "password")
)

func main() {
	flag.Parse()
	args := flag.Args()

	if *base == "" {
		if len(args) == 0 {
			panic("please provide a base url")
		} else {
			*base = args[0]
			args = args[1:]
		}
	}

	if *passwd == "" {
		if len(args) == 0 {
			panic("please provide a password")
		} else {
			*passwd = args[0]
			args = args[1:]
		}
	}

	v, err := vlclient.NewVLClient(*base, *passwd)
	if err != nil {
		panic(err)
	}

	status, err := v.InPlay("file:///home/superloach/med/aud/Music/Hot Flash Heat Wave - Gutter Girl.mp3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", status)
}
