// +build windows

package base

import (
	"errors"
	"time"

	"github.com/scrapli/scrapligo/transport"
)

// NewDriver create a new instance of `Driver`, accepts a host and variadic of options to modify
// the driver behavior.
func NewDriver(
	host string,
	options ...Option,
) (*Driver, error) {
	d := &Driver{
		Host:               host,
		Port:               22,
		AuthStrictKey:      true,
		TimeoutSocket:      30 * time.Second,
		TimeoutTransport:   45 * time.Second,
		TransportType:      transport.StandardTransportName,
		transportPtyHeight: 80,
		transportPtyWidth:  256,
		FailedWhenContains: []string{},
		PrivilegeLevels:    map[string]*PrivilegeLevel{},
		DefaultDesiredPriv: "",
	}

	for _, option := range options {
		err := option(d)

		if err != nil {
			if errors.Is(err, ErrIgnoredOption) {
				continue
			} else {
				return nil, err
			}
		}
	}

	baseTransportArgs := &transport.BaseTransportArgs{
		Host:             d.Host,
		Port:             d.Port,
		AuthUsername:     d.AuthUsername,
		TimeoutSocket:    d.TimeoutSocket,
		TimeoutTransport: d.TimeoutTransport,
		PtyHeight:        d.transportPtyHeight,
		PtyWidth:         d.transportPtyWidth,
	}

	if d.Transport == nil {
		switch d.TransportType {
		case transport.StandardTransportName:
			standardTransportArgs := &transport.StandardTransportArgs{
				AuthPassword:      d.AuthPassword,
				AuthPrivateKey:    d.AuthPrivateKey,
				AuthStrictKey:     d.AuthStrictKey,
				SSHConfigFile:     d.SSHConfigFile,
				SSHKnownHostsFile: d.SSHKnownHostsFile,
			}
			tImpl := &transport.Standard{
				StandardTransportArgs: standardTransportArgs,
			}
			t := &transport.Transport{
				Impl:              tImpl,
				BaseTransportArgs: baseTransportArgs,
			}
			d.Transport = t
		case transport.TelnetTransportName:
			telnetTransportArgs := &transport.TelnetTransportArgs{}
			tImpl := &transport.Telnet{
				TelnetTransportArgs: telnetTransportArgs,
			}
			t := &transport.Transport{
				Impl:              tImpl,
				BaseTransportArgs: baseTransportArgs,
			}
			d.Transport = t
		default:
			return nil, transport.ErrUnknownTransport
		}
	}

	c, err := NewChannel(d.Transport, options...)
	if err != nil {
		return nil, err
	}

	d.Channel = c

	return d, nil
}
