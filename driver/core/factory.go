package core

import (
	"errors"

	"github.com/scrapli/scrapligo/driver/base"

	"github.com/scrapli/scrapligo/driver/network"
)

// ErrUnknownPlatform raised when user provides an unknown platform... duh.
var ErrUnknownPlatform = errors.New("unknown platform provided")

// SupportedPlatforms pseudo constant providing slice of all core platform types.
func SupportedPlatforms() []string {
	return []string{"cisco_iosxe", "cisco_iosxr", "cisco_nxos", "arista_eos", "juniper_junos"}
}

// NewCoreDriver return a new core driver for a given platform.
func NewCoreDriver(
	host,
	platform string,
	options ...base.Option,
) (*network.Driver, error) {
	if platform == "cisco_iosxe" {
		return NewIOSXEDriver(host, options...)
	} else if platform == "cisco_iosxr" {
		return NewIOSXRDriver(host, options...)
	} else if platform == "cisco_nxos" {
		return NewNXOSDriver(host, options...)
	} else if platform == "arista_eos" {
		return NewEOSDriver(host, options...)
	} else if platform == "juniper_junos" {
		return NewJUNOSDriver(host, options...)
	}

	return nil, ErrUnknownPlatform
}
