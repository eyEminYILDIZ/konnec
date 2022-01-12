package filer

import (
	"io/ioutil"
	"log"

	"github.com/eyEminYILDIZ/konnec/internal/types"
	"gopkg.in/yaml.v2"
)

func ReadInventoryFile(inventoryFilePath string) *types.InventoryFile {
	yamlFile, err := ioutil.ReadFile(inventoryFilePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var inventory types.InventoryFile

	err = yaml.Unmarshal(yamlFile, &inventory)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &inventory
}

func ReadChecklistFile(checklistFilePath string) *types.Checklist {
	yamlFile, err := ioutil.ReadFile(checklistFilePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var checklist types.Checklist

	err = yaml.Unmarshal(yamlFile, &checklist)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &checklist
}
