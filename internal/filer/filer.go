package filer

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type InventoryFile struct {
	Inventories []Inventory `yaml:"inventories"`
}

type Inventory struct {
	Name        string `yaml:"name"`
	Ip          string `yaml:"ip"`
	Domain      string `yaml:"domain"`
	Description string `yaml:"description"`
}

func ReadInventories() *InventoryFile {
	yamlFile, err := ioutil.ReadFile("inventories.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var inventory InventoryFile
	c := &inventory

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
