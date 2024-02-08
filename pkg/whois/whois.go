package whois

import (
	"encoding/json"
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
		whois, err := whoisparser.Parse(result)
		isAddr := net.ParseIP(whoisVar)
		data := map[string]string{
			"domain":          whois.Domain.Domain,
			"registrar":       whois.Registrar.Name,
			"expiration_date": whois.Domain.ExpirationDate,
		}

		// Marshal the map into a JSON string
		jsonString, err := json.Marshal(data)
		if isAddr != nil {
			fmt.Println("Only valid domains are allowed.")
		} else {
			if err == nil {
				fmt.Println(string(jsonString))
			} else {
				fmt.Print("")
			}
		}
	}
}
