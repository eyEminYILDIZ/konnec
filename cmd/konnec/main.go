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

	// Get checklistFilePath
	checklistFilePath := "checklist.yaml"
	if len(os.Args) > 2 {
		checklistFilePath = os.Args[2]
	}

	// Read inventory file
	inventory := filer.ReadInventory(inventoryFilePath)
	checklist := filer.ReadChecklist(checklistFilePath)

	for _, checklistItem := range checklist.Items {
		matchedResource := inventory.GetResourceByName(checklistItem.ResourceName)
		if matchedResource == nil {
			fmt.Println("Resource Name did not matched with any inventory resource: ", checklistItem.ResourceName)
			return
		}

		fmt.Println("\n#", matchedResource.Name)

		for _, checklistCondition := range checklistItem.Conditions {
			flagSuccedded := false
			switch checklistCondition.Type {

			case "domain_ip_resolution":
				flagSuccedded = dns_resolver.ResolveDomainIpMatching(matchedResource.Domain, checklistCondition.Value)

			default:
				fmt.Println("Condition did not matched any Konnec Feature: ", checklistCondition.Type)
			}

			if flagSuccedded {
				fmt.Println("Ok:", checklistCondition.Type)
			} else {
				fmt.Println("Failed:", checklistCondition.Type)
			}
		}

	}

	fmt.Println("\nProgram finished...")
}
