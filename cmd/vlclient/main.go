package main

import (
	"flag"
	"fmt"

	"github.com/superloach/vlclient"
)

var (
	base   = flag.String("base", "http://localhost:8080", "base url")
	passwd = flag.String("passwd", "", "password")
)

func main() {
	flag.Parse()

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
