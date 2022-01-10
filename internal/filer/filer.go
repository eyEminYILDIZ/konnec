package filer

import (
	"io/ioutil"
	"log"

	"github.com/eyEminYILDIZ/konnec/internal/types"
	"gopkg.in/yaml.v2"
)

func ReadInventories() *types.InventoryFile {
	yamlFile, err := ioutil.ReadFile("inventories.yaml")
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
