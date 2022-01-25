package main

import (
	"fmt"
	"os"

	"github.com/eyEminYILDIZ/konnec/internal/dns_resolver"
	"github.com/eyEminYILDIZ/konnec/internal/filer"
	"github.com/eyEminYILDIZ/konnec/internal/http_checker"
	"github.com/eyEminYILDIZ/konnec/internal/pinger"
)

func main() {
	// Get inventoryFilePath
	inventoryFilePath := "inventories.yaml"
	if len(os.Args) > 1 {
		inventoryFilePath = os.Args[1]
	}

	// Get checklistFilePath
	checklistFilePath := "checklist.yaml"
	if len(os.Args) > 2 {
		checklistFilePath = os.Args[2]
	}

	// Read inventory and checklist files
	inventory := filer.ReadInventoryFile(inventoryFilePath)
	checklist := filer.ReadChecklistFile(checklistFilePath)

	// Check conditions
	for _, checklistItem := range checklist.Items {
		// Determine resource from checklist item resource name
		matchedResource := inventory.GetResourceByName(checklistItem.ResourceName)
		if matchedResource == nil {
			fmt.Println("Resource Name did not matched with any inventory resource: ", checklistItem.ResourceName)
			continue
		}

		fmt.Println("\n#", matchedResource.Name)

		// Execute each condition
		for _, checklistCondition := range checklistItem.Conditions {
			flagSuccedded := false
			switch checklistCondition.Type {

			case "domain_ip_resolution":
				flagSuccedded = dns_resolver.ResolveDomainIpMatching(matchedResource.Domain, checklistCondition.Value)

			case "ping_to_domain":
				flagSuccedded = pinger.PingWithDomainName(checklistCondition.Value)

			case "ping_to_ip":
				flagSuccedded = pinger.PingWithIpAddress(checklistCondition.Value)

			case "http_checker":
				flagSuccedded = http_checker.CheckHttpEndpoint(checklistCondition.Value)

			default:
				fmt.Println("Condition did not matched with any Konnec feature name: ", checklistCondition.Type)
			}

			if flagSuccedded {
				fmt.Printf("[Success] [%s] %s\n", checklistCondition.Type, checklistCondition.Message)
			} else {
				fmt.Printf("[Failed ] [%s] %s\n", checklistCondition.Type, checklistCondition.Message)
			}
		}
	}

	fmt.Println("\nProgram finished...")
}
