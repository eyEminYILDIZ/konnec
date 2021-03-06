package dns_resolver

import (
	"net"
)

func ResolveDomainIpMatching(domainName string, ipAddress string) bool {
	flagMatched := false
	ips, _ := net.LookupIP(domainName)

	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			if ipv4.String() == ipAddress {
				flagMatched = true
				break
			}
		}
	}

	return flagMatched
}

func GetIpAddressesFromDomainName(domainName string) []string {
	ipAddresses := []string{}
	ips, _ := net.LookupIP(domainName)

	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			ipAddresses = append(ipAddresses, ipv4.String())
		}
	}

	return ipAddresses
}
