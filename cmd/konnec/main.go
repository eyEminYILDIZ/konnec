package main

import (
	"fmt"

	"github.com/eyEminYILDIZ/konnec/internal/dns_resolver"
	"github.com/eyEminYILDIZ/konnec/internal/filer"
)

func main() {
	inventory := filer.ReadInventories()

	// Check DomainName-IP Matching
	hasError, dnsResolveErrors := dns_resolver.CheckDomainNameIpMatching(inventory)

	if hasError {
		fmt.Println("Some resource(s) domain name results did not matched provided IP address !")

		for _, errString := range dnsResolveErrors {
			fmt.Println(errString)
		}
	} else {
		fmt.Println("All resources domain name DNS resolving matched provided IP address")
	}

	fmt.Println("Program finished...")
}
