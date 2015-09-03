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
		for _, computerName := range os.Args[1:] {
			fmt.Printf("%v: ", computerName)
			siteName, err := adds.GetSiteName(computerName)
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("%v\n", siteName)
			}
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
