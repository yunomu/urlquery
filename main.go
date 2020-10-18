package main

import (
	"flag"
	"io"
	"log"
	"net/url"
	"os"
)

func init() {
	flag.Parse()
	log.SetOutput(os.Stderr)
}

func main() {
	in, out := os.Stdin, os.Stdout

	for {
		buf := [0x1000]byte{}
		n, err := in.Read(buf[:])
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Read: %v", err)
		}

		s := url.QueryEscape(string(buf[:n]))

		if _, err = out.Write([]byte(s)); err != nil {
			log.Fatalf("Write: %v", err)
		}
	}
}
