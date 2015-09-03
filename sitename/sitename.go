package main

import (
	"fmt"
	"os"

	"go.scj.io/adds"
)

func main() {
	if len(os.Args) < 2 {
		lookup("")
	} else {
		for _, arg := range os.Args[1:] {
			siteName, err := adds.GetSiteName(arg)
			if err != nil {
				panic(err)
			}
			fmt.Println(siteName)
		}
	}
}

func lookup(computerName string) {
	siteName, err := adds.GetSiteName(computerName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(siteName)
	}
}
