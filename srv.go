package veneur

import (
	"errors"
	"fmt"
	"net/url"
	"net"
)

// Srv is a Discoverer that uses SRV records to find
// healthy instances.

type SrvDiscoverer struct {
	Nameserver string
	Port string
	Protocol string
}

// NewSrv creates a new instance of a Srv Discoverer
// note, it should load in an optional nameserver
func NewSrv() (*SrvDiscoverer, error) {
	return &SrvDiscoverer {
		Nameserver: nil,
		Port: "8126", //TODO: make this configurable
		Protocol: "tcp", //TODO: make this configurable
	}, nil
}

// GetDestinationsForService updates the list of destinations based on healthy nodes
// found via SRV.
func (c *SrvDiscoverer) GetDestinationsForService(serviceName string) ([]string, error) {
	// I'm not implementing the full SRV spec because kubernetes does not support it:
	// https://github.com/kubernetes/kubernetes/issues/29420
	// still works pretty well without it
	cname, addr, err := net.LookupSRV("", "", serviceName)
	if err != nil {
		return nil, err
	}

	numHosts := len(addr)
	if numHosts < 1 {
		return nil, errors.New("SRV record empty")
	}

	// Make a slice to hold our returned hosts
	hosts := make([]string, numHosts)
	for index, a := range addr {

		dest := url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%d", a.Target, Port),
		}

		hosts[index] = dest.String()
	}

	return hosts, nil
}
