package server

import (
	"fmt"

	"github.com/lthibault/log"
	"go.uber.org/fx"

	casm "github.com/wetware/casm/pkg"
	"github.com/wetware/casm/pkg/cluster"
	ww "github.com/wetware/ww/pkg"
)

type Joiner struct {
	fx.In

	Log log.Logger
	Vat casm.Vat
	// Metrics casm.Metrics          `optional:"true"`  // XXX - implement before merging
	Options []cluster.Option `group:"cluster"`
}

func (j Joiner) Join(ps cluster.PubSub) (*Node, error) {
	// join the cluster topic
	c, err := cluster.New(ps, j.options()...)
	if err != nil {
		return nil, fmt.Errorf("join cluster: %w", err)
	}

	// export the root Host capability
	j.Vat.Export(
		ww.HostCapability,
		ww.HostServer{ViewProvider: c})

	return &Node{
		Vat:  j.Vat,
		Node: c,
	}, nil
}

func (j Joiner) options() []cluster.Option {
	return append([]cluster.Option{
		cluster.WithLogger(j.Log),
		cluster.WithNamespace(j.Vat.NS),
		// cluster.WithMetrics(j.metrics()),  // TODO:  metrics should track view size
	}, j.Options...)
}
