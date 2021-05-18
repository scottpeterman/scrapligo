package generic

import (
	"time"

	"github.com/scrapli/scrapligo/driver/base"
)

// FullSendCommand same as `SendCommand` but requiring explicit options.
func (d *Driver) FullSendCommand(
	c string,
	failedWhenContains []string,
	stripPrompt, eager bool,
	timeoutOps time.Duration,
) (*base.Response, error) {
	r := base.NewResponse(d.Host, d.Port, c, failedWhenContains)

	rawResult, err := d.Channel.SendInput(c, stripPrompt, eager, timeoutOps)

	r.Record(rawResult, string(rawResult))

	if err != nil {
		r.Failed = true
	}

	return r, err
}

// SendCommand send a command to a device, accepts a string command and variadic of `SendOption`s.
func (d *Driver) SendCommand(c string, o ...base.SendOption) (*base.Response, error) {
	finalOpts := d.ParseSendOptions(o)

	return d.FullSendCommand(
		c,
		finalOpts.FailedWhenContains,
		finalOpts.StripPrompt,
		finalOpts.Eager,
		finalOpts.TimeoutOps,
	)
}
