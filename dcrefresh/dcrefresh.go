package main

import (
	"flag"
	"fmt"

	"go.scj.io/adds"
)

var (
	domainName string
	siteName   string
	force      bool
)

func init() {
	flag.StringVar(&domainName, "domain", "", "")
	flag.StringVar(&domainName, "d", "", "")
	flag.StringVar(&siteName, "site", "", "")
	flag.StringVar(&siteName, "s", "", "")
	flag.BoolVar(&force, "force", false, "")
	flag.BoolVar(&force, "f", false, "")
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		refresh("")
	} else {
		for _, computerName := range flag.Args() {
			refresh(computerName)
		}
	}
}

func refresh(computerName string) {
	if len(computerName) == 0 {
		fmt.Print("(local machine): ")
	} else {
		fmt.Printf("%v: ", computerName)
	}
	var info *adds.DomainControllerInfo
	var err error
	if force {
		info, err = adds.GetDcName(computerName, domainName, siteName, 0)
	} else {
		info, err = adds.GetDcName(computerName, domainName, siteName, adds.DS_FORCE_REDISCOVERY)
	}
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%+v\n", info)
	}
}
