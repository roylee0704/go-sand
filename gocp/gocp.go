package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	l := len(os.Args[1:])

	if l != 2 {
		fmt.Fprintf(os.Stderr, "gocp: missing argument(s): <file> <dest-file-path>\n")
		return
	}

	filename, destpath := os.Args[1], os.Args[2]

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gocp: err: %v\n", err)
		return
	}
	defer fi.Close()

	if destinfo, err := os.Stat(destpath); os.IsNotExist(err) || !destinfo.IsDir() {
		fmt.Fprintf(os.Stderr, "gocp: err: <directory:%s> not found \n", destpath)
		return
	}

	fo, err := os.Create(path.Join(destpath, filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "gocp: err: %v\n", err)
		return
	}
	defer fo.Close()

	copy(fi, fo)

}

func copy(fi *os.File, fo *os.File) {
	buffer := make([]byte, 1024)

	for {
		n := 0
		if n, _ = fi.Read(buffer); n == 0 {
			break
		}
		fo.Write(buffer[:n])
		fmt.Printf("%d bytes transfered.\n", n)
	}

}
