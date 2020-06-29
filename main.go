package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  mks target/path/file.txt")
	fmt.Println("  mks target/path/file.txt 'file body here'")
	fmt.Println("\nFlags")
	flag.PrintDefaults()
	fmt.Println("\nPlease report issues at https://github.com/paulvollmer/mks/issues")
	fmt.Println("Copyright 2019-2020, Paul Vollmer")
}

func main() {
	flagVersion := flag.Bool("version", false, "print the version and exit")
	flag.Usage = usage
	flag.Parse()
	if *flagVersion {
		fmt.Printf("version:  %s\n", version)
		fmt.Printf("commit:   %s\n", commit)
		fmt.Printf("build at: %s\n", date)
		os.Exit(0)
	}

	nargs := flag.NArg()
	args := flag.Args()

	if nargs == 0 {
		usage()
		os.Exit(1)
	} else if nargs >= 1 {

		body := []byte("")
		if nargs >= 2 {
			body = []byte(args[1])
		}

		err := createSource(args[0], body)
		if err != nil {
			fmt.Printf("ERROR %s\n", err)
			os.Exit(1)
		}

	} else {
		usage()
		os.Exit(1)
	}
}

func createSource(src string, body []byte) error {
	dir, _ := filepath.Split(src)
	if dir != "" {
		// check if the directory exist
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(src, body, 0644)
		}
	}

	// check if a file exist
	if _, err := os.Stat(src); os.IsNotExist(err) {
		// path exists, creating the file...
		ioutil.WriteFile(src, body, 0644)
	}
	return nil
}
