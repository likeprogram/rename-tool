package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var strFlag string

func main() {
	flag.StringVar(&strFlag, "dir", "", "The source directory")
	flag.Parse()

	curDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Getwd failed: %v\n", err)
	}
	dir := curDir
	fmt.Println("Current dir: ", dir)
	if strFlag != "" {
		dir = strFlag
		fmt.Println("Passed dir: ", dir)
	}
	readDir, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	re, err := regexp.Compile(".+_(bin.+\\.dat)")
	fmt.Printf("re: %v\n", re)
	if err != nil {
		log.Fatal(err)
		return
	}

	destDir := filepath.Join(dir, "dest")
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			fmt.Printf("Create dest folder failed: %v\n\n", err)
			return
		}
	}
	for _, e := range readDir {
		if !e.IsDir() {
			fmt.Printf("%s\n", e.Name())
			parts := re.FindStringSubmatch(e.Name())
			//fmt.Printf("parts: %v\n", parts)
			if len(parts) == 2 {
				newFileName := parts[1]
				if err := os.Rename(filepath.Join(dir, e.Name()), filepath.Join(destDir, newFileName)); err != nil {
					fmt.Printf("Rename failed: %v\n\n", err)
				}
			}
		}
	}
}
