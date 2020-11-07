package main

import (
	"bufio"
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
	fmt.Println("  echo 'file body here' | mks target/path/file.txt")
	fmt.Println("\nFlags")
	fmt.Println("  -perm int    the file permission (default 0644)")
	fmt.Println("  -version     print the version and exit")
	fmt.Println("\n\nPlease report issues at https://github.com/paulvollmer/mks/issues")
	fmt.Println("Copyright 2019-2020, Paul Vollmer")
}

func main() {
	flagVersion := flag.Bool("version", false, "")
	flagFilePermission := flag.Int("perm", 0644, "")
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

	stdin, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if nargs == 0 && stdin.Size() == 0 {
		usage()
		os.Exit(0)
	}

	body := []byte("")

	if stdin.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		stdinText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		body = []byte(stdinText)
	} else if nargs >= 2 {
		body = []byte(args[1])
	}

	err = createSource(args[0], body, os.FileMode(*flagFilePermission))
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		os.Exit(1)
	}
}

func createSource(src string, body []byte, perm os.FileMode) error {
	// check if a file exist
	_, err := os.Stat(src)
	if err == nil {
		return nil
	}

	dir, _ := filepath.Split(src)
	if dir != "" {
		// check if the directory exist
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}
	}

	return ioutil.WriteFile(src, body, perm)
}
