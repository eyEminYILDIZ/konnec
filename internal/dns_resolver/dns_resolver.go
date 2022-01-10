package dns_resolver

import (
	"fmt"
	"net"

	"github.com/eyEminYILDIZ/konnec/internal/types"
)

func CheckDomainNameIpMatching(inventory *types.InventoryFile) (bool, []string) {
	resolveErrors := []string{}

	for _, resource := range inventory.Resources {
		flagMatched := false

		ips, _ := net.LookupIP(resource.Domain)

		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				if ipv4.String() == resource.Ip {
					flagMatched = true
					break
				}
			}
		}

		if !flagMatched {
			resolveErrors = append(resolveErrors, fmt.Sprint("- Not Matched: ", resource.Domain))
		}
	}

	hasError := len(resolveErrors) > 0

	return hasError, resolveErrors
}
