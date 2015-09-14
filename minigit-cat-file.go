package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	path := os.Args[1]

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
		return 1
	}

	br := bytes.NewReader(b)
	r, err := zlib.NewReader(br)
	if err != nil {
		log.Fatalln(err.Error())
		return 1
	}
	defer r.Close()

	var blob bytes.Buffer
	io.Copy(&blob, r)
	fmt.Printf("%s", blob.Bytes())

	return 0
}
