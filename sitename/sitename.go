package main

import (
	"flag"
	"fmt"

	"go.scj.io/adds"
)

func main() {
	if len(flag.Args()) < 2 {
		lookup("")
	} else {
		for _, computerName := range flag.Args()[1:] {
			lookup(computerName)
		}
	}
}

func lookup(computerName string) {
	if len(computerName) == 0 {
		fmt.Print("(local machine): ")
	} else {
		fmt.Printf("%v: ", computerName)
	}
	siteName, err := adds.GetSiteName(computerName)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", siteName)
	}
}
