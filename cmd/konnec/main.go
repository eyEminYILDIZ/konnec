package main

import (
	"fmt"

	"github.com/eyEminYILDIZ/konnec/internal/filer"
)

func main() {
	inventory := filer.ReadInventories()

	fmt.Println(inventory.Resources)
}
