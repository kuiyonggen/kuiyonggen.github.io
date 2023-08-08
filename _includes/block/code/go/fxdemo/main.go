package main

import (
    "io"
    "io/ioutil"
    "context"
    "net"
    "net/http"
    "go.uber.org/zap"
    "go.uber.org/fx"
    "go.uber.org/fx/fxevent"
)


// Route is an http.Handler that knows the mux pattern
// under which it will be registerd.
type Route interface {
    http.Handler

    // Pattern reports the path at which this is registered.
    Pattern() string
}

// EchoHandler is an http.Handler that copies its request body back to the response.
type EchoHandler struct{
    log *zap.Logger
}

func (*EchoHandler) Pattern() string {
    return "/echo"
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
    return &EchoHandler{log: log}
}

// ServerHttp handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        h.log.Warn("Error reading request body")
        return
    }
    h.log.Info("Received a request", zap.String("body", string(body[:])))
    if _, err := io.Copy(w, r.Body); err != nil {
        h.log.Warn("Failed to handle request", zap.Error(err))
    }
}

// NewServeMux builds a ServeMux that will route requests to the given EchoHandler
func NewServeMux(route Route) *http.ServeMux {
    mux := http.NewServeMux()
    mux.Handle(route.Pattern(), route)
    return mux
}

// NewHTTPServer builds an HTTP server that will begin serving requests
// when the Fx application starts.
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
    srv := &http.Server{Addr: ":8884", Handler: mux}
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            ln, err := net.Listen("tcp", srv.Addr)
            if err != nil {
                return err
            }
            log.Info("Starting HTTP server at", zap.String("addr", srv.Addr))
            go srv.Serve(ln)
            return nil
        },

        OnStop: func(ctx context.Context) error {
            log.Info("Stopping HTTP server.")
            return srv.Shutdown(ctx)
        },
    })
    return srv
}



func main() {
    fx.New(
        fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
            return &fxevent.ZapLogger{Logger: log}
        }),
        fx.Provide(
            NewHTTPServer,
            NewServeMux,
            fx.Annotate(
                NewEchoHandler,
                fx.As(new(Route)),
             ),
            zap.NewExample,
        ),
        fx.Invoke(func(*http.Server) {}),
    ).Run()
    }

