package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hsmtkk/easy_crypt_go/pkg/easycrypt"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s in out", os.Args[0])
	}
	inPath := os.Args[1]
	outPath := os.Args[2]
	in, err := ioutil.ReadFile(inPath)
	if err != nil {
		log.Fatal(err)
	}
	out := easycrypt.Crypt(in)
	if err := ioutil.WriteFile(outPath, out, 0644); err != nil {
		log.Fatal(err)
	}
}
