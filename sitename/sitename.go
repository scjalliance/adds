package main

import (
	"fmt"

	"go.scj.io/adds"
)

func main() {
	siteName, err := adds.GetSiteName("")
	if err != nil {
		panic(err)
	}
	fmt.Println(siteName)
}
