package whois

import (
	"fmt"
	"net"

	whois "github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func Whois(whoisVar string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("")
		}
	}()
	result, err := whois.Whois(whoisVar)
	if err != nil {
		fmt.Println(err)
	} else {
		data, err := whoisparser.Parse(result)
		isAddr := net.ParseIP(whoisVar)
		if isAddr != nil {
			fmt.Println("Only valid domains are allowed.")
		} else {
			if err == nil {
				fmt.Println("domain:", data.Domain.Domain, "|", "registrar:", data.Registrar.Name, "|", "expiration_date:", data.Domain.ExpirationDate)
			} else {
				fmt.Print("")
			}
		}
	}
}
