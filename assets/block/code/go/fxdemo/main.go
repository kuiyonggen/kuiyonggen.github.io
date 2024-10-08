package main

import (
    "fmt"
    "io/ioutil"
    "context"
    "net"
    "net/http"
    "go.uber.org/zap"
    "go.uber.org/fx"
    "go.uber.org/fx/fxevent"
)


// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group
func AsRoute(f any) any {
    return fx.Annotate(
        f,
        fx.As(new(Route)),
        fx.ResultTags(`group:"routes"`),
    )
}

// Route is an http.Handler that knows the mux pattern
// under which it will be registerd.
type Route interface {
    http.Handler

    // Pattern reports the path at which this is registered.
    Pattern() string
}

// HelloHandler is an HTTP handler that prints a greeting to the user.
type HelloHandler struct {
    log *zap.Logger
}

// EchoHandler is an http.Handler that copies its request body back to the response.
type EchoHandler struct{
    log *zap.Logger
}

func (*HelloHandler) Pattern() string {
    return "/hello"
}

func (*EchoHandler) Pattern() string {
    return "/echo"
}

// NewHelloHandler builds a new HelloHandler
func NewHelloHandler(log *zap.Logger) *HelloHandler {
    return &HelloHandler{log: log}
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
    return &EchoHandler{log: log}
}

// ServeHttp handles an HTTP request to the /hello endpoint
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        h.log.Warn("Error reading /hello request body")
        return
    }
    strBody := string(body[:])
    h.log.Info("Received a /hello request", zap.String("body", strBody))
    if _, err := fmt.Fprintf(w, "Hello, %s.\n", strBody); err != nil {
        h.log.Warn("Failed to handle /hello request", zap.Error(err))
    }
}

// ServeHttp handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        h.log.Warn("Error reading /echo request body", zap.Error(err))
        return
    }
    strBody := string(body[:])
    h.log.Info("Received a /echo request", zap.String("body", strBody))
    if _, err := fmt.Fprintf(w, "%s\n", strBody); err != nil {
        h.log.Warn("Failed to handle /echo request", zap.Error(err))
    }
}

// NewServeMux builds a ServeMux that will route requests to the given EchoHandler
func NewServeMux(routes []Route) *http.ServeMux {
    mux := http.NewServeMux()
    for _, route := range routes {
        mux.Handle(route.Pattern(), route)
    }
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
            fx.Annotate(
                NewServeMux,
                fx.ParamTags(`group:"routes"`),
            ),
            AsRoute(NewHelloHandler),
            AsRoute(NewEchoHandler),
            zap.NewExample,
        ),
        fx.Invoke(func(*http.Server) {}),
    ).Run()
    }

