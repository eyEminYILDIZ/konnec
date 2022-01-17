package pinger

import (
	"os/exec"

	"github.com/eyEminYILDIZ/konnec/internal/dns_resolver"
)

func PingWithDomainName(domainName string) bool {
	ipV4addresses := dns_resolver.GetIpAddressesFromDomainName(domainName)

	if len(ipV4addresses) == 0 {
		return false
	}

	return PingWithIpAddress(ipV4addresses[0])
}

func PingWithIpAddress(ipAddress string) bool {
	cmd := exec.Command("ping", ipAddress, "-c", "4", "-W", "5")
	err := cmd.Run()

	return err == nil
}
