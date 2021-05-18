package network

import (
	"github.com/scrapli/scrapligo/driver/base"
)

// SendCommand basically the same as the generic driver flavor, but acquires the
// `DefaultDesiredPriv` prior to sending the command.
func (d *Driver) SendCommand(c string, o ...base.SendOption) (*base.Response, error) {
	finalOpts := d.ParseSendOptions(o)

	if d.CurrentPriv != d.DefaultDesiredPriv {
		err := d.AcquirePriv(d.DefaultDesiredPriv)
		if err != nil {
			r := base.NewResponse(d.Host, d.Port, c, []string{})
			return r, err
		}
	}

	return d.Driver.FullSendCommand(
		c,
		finalOpts.FailedWhenContains,
		finalOpts.StripPrompt,
		finalOpts.Eager,
		finalOpts.TimeoutOps,
	)
}
