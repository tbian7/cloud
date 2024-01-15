package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// io pipe
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		if _, err := fmt.Fprintln(pw, "hello"); err != nil {
			panic(err)
		}
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		panic(err)
	}

	// string reader
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	// MultiReader
	header := strings.NewReader("<msg>")
	body := strings.NewReader("hello")
	footer := strings.NewReader("</msg>\n")

	mr := io.MultiReader(header, body, footer)
	if _, err := io.Copy(os.Stdout, mr); err != nil {
		panic(err)
	}

	// MultiWrtier
	buf := &bytes.Buffer{}
	mw := io.MultiWriter(os.Stdout, os.Stderr, buf)

	fmt.Fprintln(mw, "hello")
	fmt.Println("from buffer: ", buf)
}
