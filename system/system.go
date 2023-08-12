package system

import (
	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	local "github.com/libp2p/go-libp2p/core/host"
	"github.com/urfave/cli/v2"
	"github.com/wetware/pkg/client"
	"golang.org/x/exp/slog"
)

func Boot[T ~capnp.ClientKind](c *cli.Context, h local.Host) (T, error) {
	conn, err := dial(c, h)
	if err != nil {
		return T{}, err
	}

	go func() {
		defer conn.Close()

		select {
		case <-c.Done():
		case <-conn.Done():
		}
	}()

	client := conn.Bootstrap(c.Context)
	return T(client), client.Resolve(c.Context)
}

func dial(c *cli.Context, h local.Host) (*rpc.Conn, error) {
	return client.Dialer{
		Logger:   slog.Default(),
		NS:       c.String("ns"),
		Peers:    c.StringSlice("peer"),
		Discover: c.String("discover"),
	}.Dial(c.Context, h)
}
