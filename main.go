package roadiescript

import (
	"context"
	"net/http"

	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
)

const (
	typeName = "CustomScript"
)

// customScript middleware
type customScript struct {
	next http.Handler
	name string
}

// New creates a new handler.
func New(ctx context.Context, next http.Handler, config config.RoadieScript, name string) (http.Handler, error) {
	logger := log.FromContext(middlewares.GetLoggerCtx(ctx, name, typeName))
	logger.Debug("Creating middleware")

	m := &customScript{
		next: next,
		name: name,
	}

	return m, nil
}

func (l *customScript) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	logger := log.FromContext(middlewares.GetLoggerCtx(req.Context(), l.name, typeName))
	logger.Debug("Adjusting response")
	_, _ = rw.Write([]byte("hello from damn api"))
	l.next.ServeHTTP(rw, req)
}
