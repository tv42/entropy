package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/sys/unix"
)

var program = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", program)
	fmt.Fprintf(os.Stderr, "  %s BYTES >FILE\n", program)
	fmt.Fprintf(os.Stderr, "  %s BYTES | PROGRAM\n", program)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Outputs the requested amount of entropy.\n")
	fmt.Fprintf(os.Stderr, "\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(program + ": ")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}

	count, err := strconv.ParseInt(flag.Arg(0), 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	if count < 0 {
		count = 0
	}

	if _, err := unix.IoctlGetTermios(unix.Stdout, unix.TCGETS); err == nil {
		// is a terminal
		log.Fatal("stdout is a terminal, refusing to output binary")
	}

	if _, err := io.CopyN(os.Stdout, rand.Reader, count); err != nil {
		log.Fatal(err)
	}
}
