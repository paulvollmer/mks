package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var (
	version = "0.1.0"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  mks target/path/file.txt")
	fmt.Println("  mks target/path/file.txt -b 'file body here' -b 'multiple lines can be added too'")
	fmt.Println("\nFlags")
	flag.PrintDefaults()
	fmt.Println("\nPlease report issues at https://github.com/paulvollmer/mks/issues")
	fmt.Println("Copyright 2019, Paul Vollmer")
}

func main() {
	// TODO: -b flag to create the body
	flagMode := flag.Int("m", 0, "the file mode")
	flagVersion := flag.Bool("version", false, "print the version and exit")
	flag.Usage = usage
	flag.Parse()
	if *flagVersion {
		fmt.Printf("v%s\n", version)
		os.Exit(0)
	}

	nargs := flag.NArg()
	args := flag.Args()

	// fmt.Println(nargs, args)
	if nargs == 0 {
		usage()
		os.Exit(1)
	} else if nargs >= 1 {

		body := []byte("")
		// fmt.Println("NARGS", nargs)
		if nargs >= 2 {
			body = []byte(args[1])
			// fmt.Println("BODY", string(body))
		}
		fmt.Printf("==> create file: %q, body: %q\n", args[0], body)

		mode := os.ModePerm
		if *flagMode != 0 {
			mode = os.FileMode(*flagMode)
		}
		err := createSource(args[0], body, mode)
		if err != nil {
			fmt.Println("ERROR:> ", err)
			os.Exit(1)
		}

	} else {
		fmt.Printf("==> unknown command\n")
	}
}

func createSource(src string, body []byte, mode os.FileMode) error {
	dir, file := path.Split(src)
	// fmt.Println("==> DIR", dir, "FILE", file)
	if dir != "" {
		// check if the directory exist
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, mode)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(src, body, 0755)
		}
	}

	// check if a file exist
	if _, err := os.Stat(src); os.IsNotExist(err) {
		// path exists, creating the file...
		fmt.Println("==> create file", file, src)
		ioutil.WriteFile(src, body, 0755)
	}

	return nil
}
