package driver

import (
	"github.com/oeoen/policy/driver/config"
	"github.com/oeoen/policy/pkg/handler"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"
)

type Registry interface {
	Tracer() *tracing.Tracer
	WithConfig(c config.Provider) Registry
	WithLogger(l *logrusx.Logger) Registry
	Init() error
	notification.Registry
	Handler() handler.Provider
}

func CallRegistry(r Registry) {
	r.Tracer()
}
