package docker

import (
	"github.com/docker/docker/client"
)

type Option func(*Client) error

func WithClient(cli *client.Client) Option {
	return func(d *Client) error {
		d.inner = cli
		return nil
	}
}

func WithClientOpts(opts ...client.Opt) Option {
	return func(d *Client) error {
		if opts == nil {
			return nil
		}
		if d.copts == nil {
			d.copts = make([]client.Opt, 0)
		}
		d.copts = append(d.copts, opts...)
		return nil
	}
}
