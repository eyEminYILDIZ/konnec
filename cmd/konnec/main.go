package main

import (
	"fmt"
	"os"

	"github.com/eyEminYILDIZ/konnec/internal/dns_resolver"
	"github.com/eyEminYILDIZ/konnec/internal/filer"
)

func main() {
	// Get inventoryFilePath
	inventoryFilePath := "inventories.yaml"
	if len(os.Args) > 1 {
		inventoryFilePath = os.Args[1]
	}

	// Read inventory file
	inventory := filer.ReadInventories(inventoryFilePath)

	// Check domainName-ip matching
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
