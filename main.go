package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var fatalErr error

	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()
	var (
		path = flag.String("path", ".", "patch to search dead symlinks")
	)
	flag.Parse()
	err := filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fatalErr = err
			return err
		}
		fmt.Println(path)
		if info.Mode()&os.ModeSymlink != 0 {
			in, err := os.Open(path)
			if err != nil {
				var answer string
				fmt.Println("Found dead symlink at :", path, " Remove it ? (Y/n")
				fmt.Scan(&answer)
				answer = strings.TrimSpace(answer)
				answer = strings.ToLower(answer)
				if answer == "y" {
					os.Remove(path)
				}
			}
			in.Close()
		}
		return nil
	})
	if err != nil {
		fatalErr = err
		return
	}
}
