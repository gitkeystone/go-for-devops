// os.Stat vs os.LStat
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	symlinkPath := "example_symlink"

	statInfo, err := os.Stat(symlinkPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("os.Stat() - Target file info: %+v\n", statInfo)

	fmt.Println("---")

	lstatInfo, err := os.Lstat(symlinkPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("os.Lstat() - Symlink file info: %+v\n", lstatInfo)

	fmt.Println("---")
	fmt.Println(lstatInfo.Mode().String())
	fmt.Println(lstatInfo.Mode())
	if lstatInfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println("This is a symbolic link")
	}
}
