package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"github.com/lifeisgoodvega/otus-go-2019/hw06/src/copy_lib"
)

func main() {
	p := argparse.NewParser("Parser", "Ordinary parse")
	from := p.String("f", "from", &argparse.Options{Required: true, Help: "Source file"})
	to := p.String("t", "to", &argparse.Options{Required: true, Help: "Destination file"})
	limit := p.Int("l", "limit", &argparse.Options{Required: false, Help: "Maximum number of bytes to copy", Default: 0})
	offset := p.Int("o", "offset", &argparse.Options{Required: false, Help: "Offset from a begining of source file", Default: 0})

	err := p.Parse(os.Args)
	if err != nil {
		fmt.Print(p.Usage(err))
	}

	err = copy.Copy(*from, *to, *limit, *offset)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("File copied sucessfully")
	}
}
