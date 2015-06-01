package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inFile  = flag.String("infile", "infile", "the input file")
	outFile = flag.String("outfile", "outfile.json", "the output file")
)

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile(*inFile)
	if err != nil {
		fmt.Printf("couldn't read file [%s]\n", err)
		os.Exit(1)
	}
	spl := strings.Split(string(b), "\n")
	b, err = json.MarshalIndent(spl, "\t", "")
	if err != nil {
		fmt.Printf("couldn't encode as json list [%s]\n", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(*outFile, b, os.ModePerm)
	if err != nil {
		fmt.Printf("couldn't write output file [%s]\n", err)
		os.Exit(1)
	}
	fmt.Println("success")
	os.Exit(0)
}
