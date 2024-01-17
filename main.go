package main

import (
	"fmt"
	"os"

	"github.com/tr3ss/whodis/pkg/whois"

	"github.com/speedata/optionparser"
)

var (
	whoisVar string
	banner   = "Usage: ./whodis [OPTIONS] DOMAIN\n"
)

func main() {
	op := optionparser.NewOptionParser()
	op.Banner = banner
	op.On("-w", "--whois domain", "WHOIS information", &whoisVar)

	err := op.Parse()
	if err != nil {
		fmt.Println("Unknown option: %s\n", os.Args[1])
		op.Help()
		os.Exit(0)
	} else if len(os.Args) < 2 {
		op.Help()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "-w", "-whois", "--whois":
		whois.Whois(whoisVar)
	default:
		fmt.Println("Bad command: %s\n", os.Args[1])
		op.Help()
		os.Exit(0)
	}
}
