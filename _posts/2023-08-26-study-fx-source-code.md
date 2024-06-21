---
layout: post
title: "Study FX Source Code"
categories: ["tech", "golang"]
---
I decided to study the Uber FX framework.

{% raw %}
## Structure

```bash

fx/
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── README.md
├── annotated.go
├── annotated_test.go
├── app.go
├── app_internal_test.go
├── app_test.go
├── app_unixes.go
├── app_wasm.go
├── app_windows.go
├── app_windows_test.go
├── checklicense.sh
├── decorate.go
├── decorate_test.go
├── doc.go
├── docs
│   ├── Makefile
│   ├── annotate.md
│   ├── bin
│   │   └── region
│   ├── changelog.md -> ../CHANGELOG.md
│   ├── contributing.md -> ../CONTRIBUTING.md
│   ├── ex
│   │   ├── annotate
│   │   │   ├── cast.go
│   │   │   ├── cast_bad.go
│   │   │   ├── cast_test.go
│   │   │   ├── github
│   │   │   │   └── stub.go
│   │   │   ├── sample.go
│   │   │   └── sample_test.go
│   │   ├── get-started
│   │   │   ├── 01-minimal
│   │   │   │   ├── main.go
│   │   │   │   └── main_test.go
│   │   │   ├── 02-http-server
│   │   │   │   ├── main.go
│   │   │   │   └── main_test.go
│   │   │   ├── 03-echo-handler
│   │   │   │   ├── main.go
│   │   │   │   └── main_test.go
│   │   │   ├── 04-logger
│   │   │   │   ├── main.go
│   │   │   │   └── main_test.go
│   │   │   ├── 05-registration
│   │   │   │   ├── main.go
│   │   │   │   └── main_test.go
│   │   │   ├── 06-another-handler
│   │   │   │   ├── main.go
│   │   │   │   └── main_test.go
│   │   │   └── 07-many-handlers
│   │   │       ├── main.go
│   │   │       └── main_test.go
│   │   ├── modules
│   │   │   ├── module.go
│   │   │   └── module_test.go
│   │   ├── parameter-objects
│   │   │   ├── define.go
│   │   │   ├── define_test.go
│   │   │   ├── extend.go
│   │   │   └── extend_test.go
│   │   ├── result-objects
│   │   │   ├── define.go
│   │   │   ├── define_test.go
│   │   │   ├── extend.go
│   │   │   └── extend_test.go
│   │   └── value-groups
│   │       ├── consume
│   │       │   ├── annotate.go
│   │       │   ├── consume_test.go
│   │       │   └── param.go
│   │       └── feed
│   │           ├── annotate.go
│   │           ├── feed_test.go
│   │           └── result.go
│   ├── get-started
│   │   ├── README.md
│   │   ├── another-handler.md
│   │   ├── conclusion.md
│   │   ├── echo-handler.md
│   │   ├── http-server.md
│   │   ├── logger.md
│   │   ├── many-handlers.md
│   │   ├── minimal.md
│   │   └── registration.md
│   ├── go.mod
│   ├── go.sum
│   ├── index.md
│   ├── internal
│   │   ├── apptest
│   │   │   ├── run.go
│   │   │   └── run_test.go
│   │   ├── exectest
│   │   │   ├── cmd.go
│   │   │   ├── cmd_test.go
│   │   │   ├── output.go
│   │   │   └── output_test.go
│   │   ├── httptest
│   │   │   ├── http.go
│   │   │   └── http_test.go
│   │   ├── iotest
│   │   │   ├── read.go
│   │   │   └── read_test.go
│   │   └── test
│   │       ├── fake.go
│   │       ├── fake_test.go
│   │       ├── t.go
│   │       └── t_test.go
│   ├── intro.md
│   ├── lifecycle.md
│   ├── modules.md
│   ├── package.json
│   ├── parameter-objects.md
│   ├── result-objects.md
│   ├── value-groups
│   │   ├── README.md
│   │   ├── consume.md
│   │   └── feed.md
│   ├── value-groups.md
│   └── yarn.lock
├── error_example_test.go
├── example_test.go
├── extract.go
├── extract_test.go
├── fxevent
│   ├── console.go
│   ├── console_test.go
│   ├── event.go
│   ├── event_test.go
│   ├── logger.go
│   ├── zap.go
│   └── zap_test.go
├── fxtest
│   ├── app.go
│   ├── app_test.go
│   ├── lifecycle.go
│   ├── lifecycle_test.go
│   ├── printer.go
│   ├── printer_test.go
│   ├── tb.go
│   └── tb_test.go
├── go.mod
├── go.sum
├── inout.go
├── inout_test.go
├── internal
│   ├── e2e
│   │   ├── README.md
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── shutdowner_run_exitcode
│   │   │   ├── main.go
│   │   │   └── main_test.go
│   │   └── shutdowner_wait_exitcode
│   │       ├── main.go
│   │       └── main_test.go
│   ├── fxclock
│   │   ├── clock.go
│   │   └── clock_test.go
│   ├── fxlog
│   │   ├── default.go
│   │   ├── default_test.go
│   │   ├── foovendor
│   │   │   └── foovendor.go
│   │   ├── sample.git
│   │   │   └── sample.go
│   │   ├── spy.go
│   │   └── spy_test.go
│   ├── fxreflect
│   │   ├── fxreflect.go
│   │   ├── fxreflect_test.go
│   │   ├── stack.go
│   │   ├── stack_120_test.go
│   │   ├── stack_121_test.go
│   │   └── stack_test.go
│   ├── leaky_test
│   │   └── leaky_test.go
│   ├── lifecycle
│   │   ├── lifecycle.go
│   │   └── lifecycle_test.go
│   └── testutil
│       ├── writer.go
│       └── writer_test.go
├── invoke.go
├── lifecycle.go
├── log.go
├── log_test.go
├── module.go
├── module_test.go
├── populate.go
├── populate_example_test.go
├── populate_test.go
├── printer_writer.go
├── provide.go
├── replace.go
├── replace_test.go
├── shutdown.go
├── shutdown_test.go
├── signal.go
├── signal_test.go
├── supply.go
├── supply_test.go
├── tools
│   ├── analysis
│   │   └── passes
│   │       └── allfxevents
│   │           ├── analysis.go
│   │           ├── analysis_test.go
│   │           └── testdata
│   │               └── src
│   │                   ├── a
│   │                   │   ├── full.go
│   │                   │   ├── nop.go
│   │                   │   ├── not_a_logger.go
│   │                   │   ├── partial_test.go
│   │                   │   ├── ptr.go
│   │                   │   └── value.go
│   │                   ├── b
│   │                   │   ├── fxevent
│   │                   │   │   └── logger.go
│   │                   │   └── not_real_fxevent.go
│   │                   └── go.uber.org
│   │                       └── fx
│   │                           └── fxevent
│   │                               ├── fxevent.go
│   │                               └── partial.go
│   ├── cmd
│   │   └── fxlint
│   │       └── main.go
│   ├── go.mod
│   ├── go.sum
│   └── tools.go
└── version.go
```


## app.go

```go

// Copyright (c) 2020-2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"go.uber.org/dig"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/internal/fxclock"
	"go.uber.org/fx/internal/fxlog"
	"go.uber.org/fx/internal/fxreflect"
	"go.uber.org/fx/internal/lifecycle"
	"go.uber.org/multierr"
)

// DefaultTimeout is the default timeout for starting or stopping an
// application. It can be configured with the StartTimeout and StopTimeout
// options.
const DefaultTimeout = 15 * time.Second

// An Option configures an App using the functional options paradigm
// popularized by Rob Pike. If you're unfamiliar with this style, see
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html.
type Option interface {
	fmt.Stringer

	apply(*module)
}

// Error registers any number of errors with the application to short-circuit
// startup. If more than one error is given, the errors are combined into a
// single error.
//
// Similar to invocations, errors are applied in order. All Provide and Invoke
// options registered before or after an Error option will not be applied.
func Error(errs ...error) Option {
	return errorOption(errs)
}

type errorOption []error

func (errs errorOption) apply(mod *module) {
	mod.app.err = multierr.Append(mod.app.err, multierr.Combine(errs...))
}

func (errs errorOption) String() string {
	return fmt.Sprintf("fx.Error(%v)", multierr.Combine(errs...))
}

// Options converts a collection of Options into a single Option. This allows
// packages to bundle sophisticated functionality into easy-to-use Fx modules.
// For example, a logging package might export a simple option like this:
//
//	package logging
//
//	var Module = fx.Provide(func() *log.Logger {
//	  return log.New(os.Stdout, "", 0)
//	})
//
// A shared all-in-one microservice package could then use Options to bundle
// logging with similar metrics, tracing, and gRPC modules:
//
//	package server
//
//	var Module = fx.Options(
//	  logging.Module,
//	  metrics.Module,
//	  tracing.Module,
//	  grpc.Module,
//	)
//
// Since this all-in-one module has a minimal API surface, it's easy to add
// new functionality to it without breaking existing users. Individual
// applications can take advantage of all this functionality with only one
// line of code:
//
//	app := fx.New(server.Module)
//
// Use this pattern sparingly, since it limits the user's ability to customize
// their application.
func Options(opts ...Option) Option {
	return optionGroup(opts)
}

type optionGroup []Option

func (og optionGroup) apply(mod *module) {
	for _, opt := range og {
		opt.apply(mod)
	}
}

func (og optionGroup) String() string {
	items := make([]string, len(og))
	for i, opt := range og {
		items[i] = fmt.Sprint(opt)
	}
	return fmt.Sprintf("fx.Options(%s)", strings.Join(items, ", "))
}

// StartTimeout changes the application's start timeout.
func StartTimeout(v time.Duration) Option {
	return startTimeoutOption(v)
}

type startTimeoutOption time.Duration

func (t startTimeoutOption) apply(m *module) {
	if m.parent != nil {
		m.app.err = fmt.Errorf("fx.StartTimeout Option should be passed to top-level App, " +
			"not to fx.Module")
	} else {
		m.app.startTimeout = time.Duration(t)
	}
}

func (t startTimeoutOption) String() string {
	return fmt.Sprintf("fx.StartTimeout(%v)", time.Duration(t))
}

// StopTimeout changes the application's stop timeout.
func StopTimeout(v time.Duration) Option {
	return stopTimeoutOption(v)
}

type stopTimeoutOption time.Duration

func (t stopTimeoutOption) apply(m *module) {
	if m.parent != nil {
		m.app.err = fmt.Errorf("fx.StopTimeout Option should be passed to top-level App, " +
			"not to fx.Module")
	} else {
		m.app.stopTimeout = time.Duration(t)
	}
}

func (t stopTimeoutOption) String() string {
	return fmt.Sprintf("fx.StopTimeout(%v)", time.Duration(t))
}

// RecoverFromPanics causes panics that occur in functions given to [Provide],
// [Decorate], and [Invoke] to be recovered from.
// This error can be retrieved as any other error, by using (*App).Err().
func RecoverFromPanics() Option {
	return recoverFromPanicsOption{}
}

type recoverFromPanicsOption struct{}

func (o recoverFromPanicsOption) apply(m *module) {
	if m.parent != nil {
		m.app.err = fmt.Errorf("fx.RecoverFromPanics Option should be passed to top-level " +
			"App, not to fx.Module")
	} else {
		m.app.recoverFromPanics = true
	}
}

func (o recoverFromPanicsOption) String() string {
	return "fx.RecoverFromPanics()"
}

// WithLogger specifies how Fx should build an fxevent.Logger to log its events
// to. The argument must be a constructor with one of the following return
// types.
//
//	fxevent.Logger
//	(fxevent.Logger, error)
//
// For example,
//
//	WithLogger(func(logger *zap.Logger) fxevent.Logger {
//	  return &fxevent.ZapLogger{Logger: logger}
//	})
func WithLogger(constructor interface{}) Option {
	return withLoggerOption{
		constructor: constructor,
		Stack:       fxreflect.CallerStack(1, 0),
	}
}

type withLoggerOption struct {
	constructor interface{}
	Stack       fxreflect.Stack
}

func (l withLoggerOption) apply(m *module) {
	m.logConstructor = &provide{
		Target: l.constructor,
		Stack:  l.Stack,
	}
}

func (l withLoggerOption) String() string {
	return fmt.Sprintf("fx.WithLogger(%s)", fxreflect.FuncName(l.constructor))
}

// Printer is the interface required by Fx's logging backend. It's implemented
// by most loggers, including the one bundled with the standard library.
//
// Note, this will be deprecate with next release and you will need to implement
// fxevent.Logger interface instead.
type Printer interface {
	Printf(string, ...interface{})
}

// Logger redirects the application's log output to the provided printer.
// Deprecated: use WithLogger instead.
func Logger(p Printer) Option {
	return loggerOption{p}
}

type loggerOption struct{ p Printer }

func (l loggerOption) apply(m *module) {
	if m.parent != nil {
		m.app.err = fmt.Errorf("fx.Logger Option should be passed to top-level App, " +
			"not to fx.Module")
	} else {
		np := writerFromPrinter(l.p)
		m.log = fxlog.DefaultLogger(np) // assuming np is thread-safe.
	}
}

func (l loggerOption) String() string {
	return fmt.Sprintf("fx.Logger(%v)", l.p)
}

// NopLogger disables the application's log output. Note that this makes some
// failures difficult to debug, since no errors are printed to console.
var NopLogger = WithLogger(func() fxevent.Logger { return fxevent.NopLogger })

// An App is a modular application built around dependency injection. Most
// users will only need to use the New constructor and the all-in-one Run
// convenience method. In more unusual cases, users may need to use the Err,
// Start, Done, and Stop methods by hand instead of relying on Run.
//
// New creates and initializes an App. All applications begin with a
// constructor for the Lifecycle type already registered.
//
// In addition to that built-in functionality, users typically pass a handful
// of Provide options and one or more Invoke options. The Provide options
// teach the application how to instantiate a variety of types, and the Invoke
// options describe how to initialize the application.
//
// When created, the application immediately executes all the functions passed
// via Invoke options. To supply these functions with the parameters they
// need, the application looks for constructors that return the appropriate
// types; if constructors for any required types are missing or any
// invocations return an error, the application will fail to start (and Err
// will return a descriptive error message).
//
// Once all the invocations (and any required constructors) have been called,
// New returns and the application is ready to be started using Run or Start.
// On startup, it executes any OnStart hooks registered with its Lifecycle.
// OnStart hooks are executed one at a time, in order, and must all complete
// within a configurable deadline (by default, 15 seconds). For details on the
// order in which OnStart hooks are executed, see the documentation for the
// Start method.
//
// At this point, the application has successfully started up. If started via
// Run, it will continue operating until it receives a shutdown signal from
// Done (see the Done documentation for details); if started explicitly via
// Start, it will operate until the user calls Stop. On shutdown, OnStop hooks
// execute one at a time, in reverse order, and must all complete within a
// configurable deadline (again, 15 seconds by default).
type App struct {
	err       error
	clock     fxclock.Clock
	lifecycle *lifecycleWrapper

	container *dig.Container
	root      *module
	modules   []*module

	// Timeouts used
	startTimeout time.Duration
	stopTimeout  time.Duration
	// Decides how we react to errors when building the graph.
	errorHooks []ErrorHandler
	validate   bool
	// Whether to recover from panics in Dig container
	recoverFromPanics bool

	// Used to signal shutdowns.
	receivers signalReceivers

	osExit func(code int) // os.Exit override; used for testing only
}

// provide is a single constructor provided to Fx.
type provide struct {
	// Constructor provided to Fx. This may be an fx.Annotated.
	Target interface{}

	// Stack trace of where this provide was made.
	Stack fxreflect.Stack

	// IsSupply is true when the Target constructor was emitted by fx.Supply.
	IsSupply   bool
	SupplyType reflect.Type // set only if IsSupply

	// Set if the type should be provided at private scope.
	Private bool
}

// invoke is a single invocation request to Fx.
type invoke struct {
	// Function to invoke.
	Target interface{}

	// Stack trace of where this invoke was made.
	Stack fxreflect.Stack
}

// ErrorHandler handles Fx application startup errors.
type ErrorHandler interface {
	HandleError(error)
}

// ErrorHook registers error handlers that implement error handling functions.
// They are executed on invoke failures. Passing multiple ErrorHandlers appends
// the new handlers to the application's existing list.
func ErrorHook(funcs ...ErrorHandler) Option {
	return errorHookOption(funcs)
}

type errorHookOption []ErrorHandler

func (eho errorHookOption) apply(m *module) {
	m.app.errorHooks = append(m.app.errorHooks, eho...)
}

func (eho errorHookOption) String() string {
	items := make([]string, len(eho))
	for i, eh := range eho {
		items[i] = fmt.Sprint(eh)
	}
	return fmt.Sprintf("fx.ErrorHook(%v)", strings.Join(items, ", "))
}

type errorHandlerList []ErrorHandler

func (ehl errorHandlerList) HandleError(err error) {
	for _, eh := range ehl {
		eh.HandleError(err)
	}
}

// validate sets *App into validation mode without running invoked functions.
func validate(validate bool) Option {
	return &validateOption{
		validate: validate,
	}
}

type validateOption struct {
	validate bool
}

func (o validateOption) apply(m *module) {
	if m.parent != nil {
		m.app.err = fmt.Errorf("fx.validate Option should be passed to top-level App, " +
			"not to fx.Module")
	} else {
		m.app.validate = o.validate
	}
}

func (o validateOption) String() string {
	return fmt.Sprintf("fx.validate(%v)", o.validate)
}

// ValidateApp validates that supplied graph would run and is not missing any dependencies. This
// method does not invoke actual input functions.
func ValidateApp(opts ...Option) error {
	opts = append(opts, validate(true))
	app := New(opts...)

	return app.Err()
}

// New creates and initializes an App, immediately executing any functions
// registered via Invoke options. See the documentation of the App struct for
// details on the application's initialization, startup, and shutdown logic.
func New(opts ...Option) *App {
	logger := fxlog.DefaultLogger(os.Stderr)

	app := &App{
		clock:        fxclock.System,
		startTimeout: DefaultTimeout,
		stopTimeout:  DefaultTimeout,
		receivers:    newSignalReceivers(),
	}
	app.root = &module{
		app: app,
		// We start with a logger that writes to stderr. One of the
		// following three things can change this:
		//
		// - fx.Logger was provided to change the output stream
		// - fx.WithLogger was provided to change the logger
		//   implementation
		// - Both, fx.Logger and fx.WithLogger were provided
		//
		// The first two cases are straightforward: we use what the
		// user gave us. For the last case, however, we need to fall
		// back to what was provided to fx.Logger if fx.WithLogger
		// fails.
		log: logger,
	}
	app.modules = append(app.modules, app.root)

	for _, opt := range opts {
		opt.apply(app.root)
	}

	// There are a few levels of wrapping on the lifecycle here. To quickly
	// cover them:
	//
	// - lifecycleWrapper ensures that we don't unintentionally expose the
	//   Start and Stop methods of the internal lifecycle.Lifecycle type
	// - lifecycleWrapper also adapts the internal lifecycle.Hook type into
	//   the public fx.Hook type.
	// - appLogger ensures that the lifecycle always logs events to the
	//   "current" logger associated with the fx.App.
	app.lifecycle = &lifecycleWrapper{
		lifecycle.New(appLogger{app}, app.clock),
	}

	containerOptions := []dig.Option{
		dig.DeferAcyclicVerification(),
		dig.DryRun(app.validate),
	}

	if app.recoverFromPanics {
		containerOptions = append(containerOptions, dig.RecoverFromPanics())
	}

	app.container = dig.New(containerOptions...)

	for _, m := range app.modules {
		m.build(app, app.container)
	}

	for _, m := range app.modules {
		m.provideAll()
	}

	frames := fxreflect.CallerStack(0, 0) // include New in the stack for default Provides
	app.root.provide(provide{
		Target: func() Lifecycle { return app.lifecycle },
		Stack:  frames,
	})
	app.root.provide(provide{Target: app.shutdowner, Stack: frames})
	app.root.provide(provide{Target: app.dotGraph, Stack: frames})

	// Run decorators before executing any Invokes -- including the one
	// inside constructCustomLogger.
	app.err = multierr.Append(app.err, app.root.decorateAll())

	// If you are thinking about returning here after provides: do not (just yet)!
	// If a custom logger was being used, we're still buffering messages.
	// We'll want to flush them to the logger.

	// custom app logger will be initialized by the root module.
	for _, m := range app.modules {
		m.constructAllCustomLoggers()
	}

	// This error might have come from the provide loop above. We've
	// already flushed to the custom logger, so we can return.
	if app.err != nil {
		return app
	}

	if err := app.root.executeInvokes(); err != nil {
		app.err = err

		if dig.CanVisualizeError(err) {
			var b bytes.Buffer
			dig.Visualize(app.container, &b, dig.VisualizeError(err))
			err = errorWithGraph{
				graph: b.String(),
				err:   err,
			}
		}
		errorHandlerList(app.errorHooks).HandleError(err)
	}

	return app
}

func (app *App) log() fxevent.Logger {
	return app.root.log
}

// DotGraph contains a DOT language visualization of the dependency graph in
// an Fx application. It is provided in the container by default at
// initialization. On failure to build the dependency graph, it is attached
// to the error and if possible, colorized to highlight the root cause of the
// failure.
type DotGraph string

type errWithGraph interface {
	Graph() DotGraph
}

type errorWithGraph struct {
	graph string
	err   error
}

func (err errorWithGraph) Graph() DotGraph {
	return DotGraph(err.graph)
}

func (err errorWithGraph) Error() string {
	return err.err.Error()
}

// VisualizeError returns the visualization of the error if available.
func VisualizeError(err error) (string, error) {
	if e, ok := err.(errWithGraph); ok && e.Graph() != "" {
		return string(e.Graph()), nil
	}
	return "", errors.New("unable to visualize error")
}

// Exits the application with the given exit code.
func (app *App) exit(code int) {
	osExit := os.Exit
	if app.osExit != nil {
		osExit = app.osExit
	}
	osExit(code)
}

// Run starts the application, blocks on the signals channel, and then
// gracefully shuts the application down. It uses DefaultTimeout to set a
// deadline for application startup and shutdown, unless the user has
// configured different timeouts with the StartTimeout or StopTimeout options.
// It's designed to make typical applications simple to run.
//
// However, all of Run's functionality is implemented in terms of the exported
// Start, Done, and Stop methods. Applications with more specialized needs
// can use those methods directly instead of relying on Run.
func (app *App) Run() {
	// Historically, we do not os.Exit(0) even though most applications
	// cede control to Fx with they call app.Run. To avoid a breaking
	// change, never os.Exit for success.
	if code := app.run(app.Wait); code != 0 {
		app.exit(code)
	}
}

func (app *App) run(done func() <-chan ShutdownSignal) (exitCode int) {
	startCtx, cancel := app.clock.WithTimeout(context.Background(), app.StartTimeout())
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		return 1
	}

	sig := <-done()
	app.log().LogEvent(&fxevent.Stopping{Signal: sig.Signal})
	exitCode = sig.ExitCode

	stopCtx, cancel := app.clock.WithTimeout(context.Background(), app.StopTimeout())
	defer cancel()

	if err := app.Stop(stopCtx); err != nil {
		return 1
	}

	return exitCode
}

// Err returns any error encountered during New's initialization. See the
// documentation of the New method for details, but typical errors include
// missing constructors, circular dependencies, constructor errors, and
// invocation errors.
//
// Most users won't need to use this method, since both Run and Start
// short-circuit if initialization failed.
func (app *App) Err() error {
	return app.err
}

var (
	_onStartHook = "OnStart"
	_onStopHook  = "OnStop"
)

// Start kicks off all long-running goroutines, like network servers or
// message queue consumers. It does this by interacting with the application's
// Lifecycle.
//
// By taking a dependency on the Lifecycle type, some of the user-supplied
// functions called during initialization may have registered start and stop
// hooks. Because initialization calls constructors serially and in dependency
// order, hooks are naturally registered in serial and dependency order too.
//
// Start executes all OnStart hooks registered with the application's
// Lifecycle, one at a time and in order. This ensures that each constructor's
// start hooks aren't executed until all its dependencies' start hooks
// complete. If any of the start hooks return an error, Start short-circuits,
// calls Stop, and returns the inciting error.
//
// Note that Start short-circuits immediately if the New constructor
// encountered any errors in application initialization.
func (app *App) Start(ctx context.Context) (err error) {
	defer func() {
		app.log().LogEvent(&fxevent.Started{Err: err})
	}()

	if app.err != nil {
		// Some provides failed, short-circuit immediately.
		return app.err
	}

	return withTimeout(ctx, &withTimeoutParams{
		hook:      _onStartHook,
		callback:  app.start,
		lifecycle: app.lifecycle,
		log:       app.log(),
	})
}

// withRollback will execute an anonymous function with a given context.
// if the anon func returns an error, rollback methods will be called and related events emitted
func (app *App) withRollback(
	ctx context.Context,
	f func(context.Context) error,
) error {
	if err := f(ctx); err != nil {
		app.log().LogEvent(&fxevent.RollingBack{StartErr: err})

		stopErr := app.lifecycle.Stop(ctx)
		app.log().LogEvent(&fxevent.RolledBack{Err: stopErr})

		if stopErr != nil {
			return multierr.Append(err, stopErr)
		}

		return err
	}

	return nil
}

func (app *App) start(ctx context.Context) error {
	return app.withRollback(ctx, func(ctx context.Context) error {
		if err := app.lifecycle.Start(ctx); err != nil {
			return err
		}
		app.receivers.Start(ctx)
		return nil
	})
}

// Stop gracefully stops the application. It executes any registered OnStop
// hooks in reverse order, so that each constructor's stop hooks are called
// before its dependencies' stop hooks.
//
// If the application didn't start cleanly, only hooks whose OnStart phase was
// called are executed. However, all those hooks are executed, even if some
// fail.
func (app *App) Stop(ctx context.Context) (err error) {
	defer func() {
		app.log().LogEvent(&fxevent.Stopped{Err: err})
	}()

	cb := func(ctx context.Context) error {
		defer app.receivers.Stop(ctx)
		return app.lifecycle.Stop(ctx)
	}

	return withTimeout(ctx, &withTimeoutParams{
		hook:      _onStopHook,
		callback:  cb,
		lifecycle: app.lifecycle,
		log:       app.log(),
	})
}

// Done returns a channel of signals to block on after starting the
// application. Applications listen for the SIGINT and SIGTERM signals; during
// development, users can send the application SIGTERM by pressing Ctrl-C in
// the same terminal as the running process.
//
// Alternatively, a signal can be broadcast to all done channels manually by
// using the Shutdown functionality (see the Shutdowner documentation for details).
//
// Note: The channel Done returns will not receive a signal unless the application
// as been started via Start or Run.
func (app *App) Done() <-chan os.Signal {
	return app.receivers.Done()
}

// Wait returns a channel of [ShutdownSignal] to block on after starting the
// application and function, similar to [App.Done], but with a minor difference.
// Should an ExitCode be provided as a [ShutdownOption] to
// the Shutdowner Shutdown method, the exit code will be available as part
// of the ShutdownSignal struct.
//
// Should the app receive a SIGTERM or SIGINT, the given
// signal will be populated in the ShutdownSignal struct.
func (app *App) Wait() <-chan ShutdownSignal {
	return app.receivers.Wait()
}

// StartTimeout returns the configured startup timeout. Apps default to using
// DefaultTimeout, but users can configure this behavior using the
// StartTimeout option.
func (app *App) StartTimeout() time.Duration {
	return app.startTimeout
}

// StopTimeout returns the configured shutdown timeout. Apps default to using
// DefaultTimeout, but users can configure this behavior using the StopTimeout
// option.
func (app *App) StopTimeout() time.Duration {
	return app.stopTimeout
}

func (app *App) dotGraph() (DotGraph, error) {
	var b bytes.Buffer
	err := dig.Visualize(app.container, &b)
	return DotGraph(b.String()), err
}

type withTimeoutParams struct {
	log       fxevent.Logger
	hook      string
	callback  func(context.Context) error
	lifecycle *lifecycleWrapper
}

// errHookCallbackExited is returned when a hook callback does not finish executing
var errHookCallbackExited = errors.New("goroutine exited without returning")

func withTimeout(ctx context.Context, param *withTimeoutParams) error {
	c := make(chan error, 1)
	go func() {
		// If runtime.Goexit() is called from within the callback
		// then nothing is written to the chan.
		// However the defer will still be called, so we can write to the chan,
		// to avoid hanging until the timeout is reached.
		callbackExited := false
		defer func() {
			if !callbackExited {
				c <- errHookCallbackExited
			}
		}()

		c <- param.callback(ctx)
		callbackExited = true
	}()

	var err error

	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-c:
		// If the context finished at the same time as the callback
		// prefer the context error.
		// This eliminates non-determinism in select-case selection.
		if ctx.Err() != nil {
			err = ctx.Err()
		}
	}

	return err
}

// appLogger logs events to the given Fx app's "current" logger.
//
// Use this with lifecycle, for example, to ensure that events always go to the
// correct logger.
type appLogger struct{ app *App }

func (l appLogger) LogEvent(ev fxevent.Event) {
	l.app.log().LogEvent(ev)
}
```

## module.go

```go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"

	"go.uber.org/dig"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/internal/fxreflect"
	"go.uber.org/multierr"
)

// A container represents a set of constructors to provide
// dependencies, and a set of functions to invoke once all the
// dependencies have been initialized.
//
// This definition corresponds to the dig.Container and dig.Scope.
type container interface {
	Invoke(interface{}, ...dig.InvokeOption) error
	Provide(interface{}, ...dig.ProvideOption) error
	Decorate(interface{}, ...dig.DecorateOption) error
}

// Module is a named group of zero or more fx.Options.
// A Module creates a scope in which certain operations are taken
// place. For more information, see [Decorate], [Replace], or [Invoke].
func Module(name string, opts ...Option) Option {
	mo := moduleOption{
		name:    name,
		options: opts,
	}
	return mo
}

type moduleOption struct {
	name    string
	options []Option
}

func (o moduleOption) String() string {
	return fmt.Sprintf("fx.Module(%q, %v)", o.name, o.options)
}

func (o moduleOption) apply(mod *module) {
	// This get called on any submodules' that are declared
	// as part of another module.

	// 1. Create a new module with the parent being the specified
	// module.
	// 2. Apply child Options on the new module.
	// 3. Append it to the parent module.
	newModule := &module{
		name:   o.name,
		parent: mod,
		app:    mod.app,
	}
	for _, opt := range o.options {
		opt.apply(newModule)
	}
	mod.modules = append(mod.modules, newModule)
}

type module struct {
	parent         *module
	name           string
	scope          scope
	provides       []provide
	invokes        []invoke
	decorators     []decorator
	modules        []*module
	app            *App
	log            fxevent.Logger
	fallbackLogger fxevent.Logger
	logConstructor *provide
}

// scope is a private wrapper interface for dig.Container and dig.Scope.
// We can consider moving this into Fx using type constraints after Go 1.20
// is released and 1.17 is deprecated.
type scope interface {
	Decorate(f interface{}, opts ...dig.DecorateOption) error
	Invoke(f interface{}, opts ...dig.InvokeOption) error
	Provide(f interface{}, opts ...dig.ProvideOption) error
	Scope(name string, opts ...dig.ScopeOption) *dig.Scope
	String() string
}

// builds the Scopes using the App's Container. Note that this happens
// after applyModules' are called because the App's Container needs to
// be built for any Scopes to be initialized, and applys' should be called
// before the Container can get initialized.
func (m *module) build(app *App, root *dig.Container) {
	if m.parent == nil {
		m.scope = root
	} else {
		parentScope := m.parent.scope
		m.scope = parentScope.Scope(m.name)
		// use parent module's logger by default
		m.log = m.parent.log
	}

	if m.logConstructor != nil {
		// Since user supplied a custom logger, use a buffered logger
		// to hold all messages until user supplied logger is
		// instantiated. Then we flush those messages after fully
		// constructing the custom logger.
		m.fallbackLogger, m.log = m.log, new(logBuffer)
	}

	for _, mod := range m.modules {
		mod.build(app, root)
	}
}

func (m *module) provideAll() {
	for _, p := range m.provides {
		m.provide(p)
	}

	for _, m := range m.modules {
		m.provideAll()
	}
}

func (m *module) provide(p provide) {
	if m.app.err != nil {
		return
	}

	if p.IsSupply {
		m.supply(p)
		return
	}

	funcName := fxreflect.FuncName(p.Target)
	var info dig.ProvideInfo
	opts := []dig.ProvideOption{
		dig.FillProvideInfo(&info),
		dig.Export(!p.Private),
		dig.WithProviderCallback(func(ci dig.CallbackInfo) {
			m.log.LogEvent(&fxevent.Run{
				Name:       funcName,
				Kind:       "provide",
				ModuleName: m.name,
				Err:        ci.Error,
			})
		}),
	}

	if err := runProvide(m.scope, p, opts...); err != nil {
		m.app.err = err
	}
	outputNames := make([]string, len(info.Outputs))
	for i, o := range info.Outputs {
		outputNames[i] = o.String()
	}

	m.log.LogEvent(&fxevent.Provided{
		ConstructorName: funcName,
		StackTrace:      p.Stack.Strings(),
		ModuleName:      m.name,
		OutputTypeNames: outputNames,
		Err:             m.app.err,
		Private:         p.Private,
	})
}

func (m *module) supply(p provide) {
	typeName := p.SupplyType.String()
	opts := []dig.ProvideOption{
		dig.Export(!p.Private),
		dig.WithProviderCallback(func(ci dig.CallbackInfo) {
			m.log.LogEvent(&fxevent.Run{
				Name:       fmt.Sprintf("stub(%v)", typeName),
				Kind:       "supply",
				ModuleName: m.name,
			})
		}),
	}

	if err := runProvide(m.scope, p, opts...); err != nil {
		m.app.err = err
	}

	m.log.LogEvent(&fxevent.Supplied{
		TypeName:   typeName,
		StackTrace: p.Stack.Strings(),
		ModuleName: m.name,
		Err:        m.app.err,
	})
}

// Constructs custom loggers for all modules in the tree
func (m *module) constructAllCustomLoggers() {
	if m.logConstructor != nil {
		if buffer, ok := m.log.(*logBuffer); ok {
			// default to parent's logger if custom logger constructor fails
			if err := m.constructCustomLogger(buffer); err != nil {
				m.app.err = multierr.Append(m.app.err, err)
				m.log = m.fallbackLogger
				buffer.Connect(m.log)
			}
		}
		m.fallbackLogger = nil
	} else if m.parent != nil {
		m.log = m.parent.log
	}

	for _, mod := range m.modules {
		mod.constructAllCustomLoggers()
	}
}

// Mirroring the behavior of app.constructCustomLogger
func (m *module) constructCustomLogger(buffer *logBuffer) (err error) {
	p := m.logConstructor
	fname := fxreflect.FuncName(p.Target)
	defer func() {
		m.log.LogEvent(&fxevent.LoggerInitialized{
			Err:             err,
			ConstructorName: fname,
		})
	}()

	// TODO: Use dig.FillProvideInfo to inspect the provided constructor
	// and fail the application if its signature didn't match.
	if err := m.scope.Provide(p.Target); err != nil {
		return fmt.Errorf("fx.WithLogger(%v) from:\n%+v\nin Module: %q\nFailed: %w",
			fname, p.Stack, m.name, err)
	}

	return m.scope.Invoke(func(log fxevent.Logger) {
		m.log = log
		buffer.Connect(log)
	})
}

func (m *module) executeInvokes() error {
	for _, m := range m.modules {
		if err := m.executeInvokes(); err != nil {
			return err
		}
	}

	for _, invoke := range m.invokes {
		if err := m.executeInvoke(invoke); err != nil {
			return err
		}
	}

	return nil
}

func (m *module) executeInvoke(i invoke) (err error) {
	fnName := fxreflect.FuncName(i.Target)
	m.log.LogEvent(&fxevent.Invoking{
		FunctionName: fnName,
		ModuleName:   m.name,
	})
	err = runInvoke(m.scope, i)
	m.log.LogEvent(&fxevent.Invoked{
		FunctionName: fnName,
		ModuleName:   m.name,
		Err:          err,
		Trace:        fmt.Sprintf("%+v", i.Stack), // format stack trace as multi-line
	})
	return err
}

func (m *module) decorateAll() error {
	for _, d := range m.decorators {
		if err := m.decorate(d); err != nil {
			return err
		}
	}

	for _, m := range m.modules {
		if err := m.decorateAll(); err != nil {
			return err
		}
	}
	return nil
}

func (m *module) decorate(d decorator) (err error) {
	if d.IsReplace {
		return m.replace(d)
	}

	funcName := fxreflect.FuncName(d.Target)
	var info dig.DecorateInfo
	opts := []dig.DecorateOption{
		dig.FillDecorateInfo(&info),
		dig.WithDecoratorCallback(func(ci dig.CallbackInfo) {
			m.log.LogEvent(&fxevent.Run{
				Name:       funcName,
				Kind:       "decorate",
				ModuleName: m.name,
				Err:        ci.Error,
			})
		}),
	}

	err = runDecorator(m.scope, d, opts...)
	outputNames := make([]string, len(info.Outputs))
	for i, o := range info.Outputs {
		outputNames[i] = o.String()
	}

	m.log.LogEvent(&fxevent.Decorated{
		DecoratorName:   funcName,
		StackTrace:      d.Stack.Strings(),
		ModuleName:      m.name,
		OutputTypeNames: outputNames,
		Err:             err,
	})

	return err
}

func (m *module) replace(d decorator) error {
	typeName := d.ReplaceType.String()
	opts := []dig.DecorateOption{
		dig.WithDecoratorCallback(func(ci dig.CallbackInfo) {
			m.log.LogEvent(&fxevent.Run{
				Name:       fmt.Sprintf("stub(%v)", typeName),
				Kind:       "replace",
				ModuleName: m.name,
				Err:        ci.Error,
			})
		}),
	}

	err := runDecorator(m.scope, d, opts...)
	m.log.LogEvent(&fxevent.Replaced{
		ModuleName:      m.name,
		StackTrace:      d.Stack.Strings(),
		OutputTypeNames: []string{typeName},
		Err:             err,
	})
	return err
}
```

## lifecycle.go

```go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"context"

	"go.uber.org/fx/internal/lifecycle"
)

// A HookFunc is a function that can be used as a [Hook].
type HookFunc interface {
	~func() | ~func() error | ~func(context.Context) | ~func(context.Context) error
}

// Lifecycle allows constructors to register callbacks that are executed on
// application start and stop. See the documentation for App for details on Fx
// applications' initialization, startup, and shutdown logic.
type Lifecycle interface {
	Append(Hook)
}

// A Hook is a pair of start and stop callbacks, either of which can be nil.
// If a Hook's OnStart callback isn't executed (because a previous OnStart
// failure short-circuited application startup), its OnStop callback won't be
// executed.
type Hook struct {
	OnStart func(context.Context) error
	OnStop  func(context.Context) error

	onStartName string
	onStopName  string
}

// StartHook returns a new Hook with start as its [Hook.OnStart] function,
// wrapping its signature as needed. For example, given the following function:
//
//	func myhook() {
//	  fmt.Println("hook called")
//	}
//
// then calling:
//
//	lifecycle.Append(StartHook(myfunc))
//
// is functionally equivalent to calling:
//
//	lifecycle.Append(fx.Hook{
//	  OnStart: func(context.Context) error {
//	    myfunc()
//	    return nil
//	  },
//	})
//
// The same is true for all functions that satisfy the HookFunc constraint.
// Note that any context.Context parameter or error return will be propagated
// as expected. If propagation is not intended, users should instead provide a
// closure that discards the undesired value(s), or construct a Hook directly.
func StartHook[T HookFunc](start T) Hook {
	onstart, startname := lifecycle.Wrap(start)

	return Hook{
		OnStart:     onstart,
		onStartName: startname,
	}
}

// StopHook returns a new Hook with stop as its [Hook.OnStop] function,
// wrapping its signature as needed. For example, given the following function:
//
//	func myhook() {
//	  fmt.Println("hook called")
//	}
//
// then calling:
//
//	lifecycle.Append(StopHook(myfunc))
//
// is functionally equivalent to calling:
//
//	lifecycle.Append(fx.Hook{
//	  OnStop: func(context.Context) error {
//	    myfunc()
//	    return nil
//	  },
//	})
//
// The same is true for all functions that satisfy the HookFunc constraint.
// Note that any context.Context parameter or error return will be propagated
// as expected. If propagation is not intended, users should instead provide a
// closure that discards the undesired value(s), or construct a Hook directly.
func StopHook[T HookFunc](stop T) Hook {
	onstop, stopname := lifecycle.Wrap(stop)

	return Hook{
		OnStop:     onstop,
		onStopName: stopname,
	}
}

// StartStopHook returns a new Hook with start as its [Hook.OnStart] function
// and stop as its [Hook.OnStop] function, independently wrapping the signature
// of each as needed.
func StartStopHook[T1 HookFunc, T2 HookFunc](start T1, stop T2) Hook {
	var (
		onstart, startname = lifecycle.Wrap(start)
		onstop, stopname   = lifecycle.Wrap(stop)
	)

	return Hook{
		OnStart:     onstart,
		OnStop:      onstop,
		onStartName: startname,
		onStopName:  stopname,
	}
}

type lifecycleWrapper struct {
	*lifecycle.Lifecycle
}

func (l *lifecycleWrapper) Append(h Hook) {
	l.Lifecycle.Append(lifecycle.Hook{
		OnStart:     h.OnStart,
		OnStop:      h.OnStop,
		OnStartName: h.onStartName,
		OnStopName:  h.onStopName,
	})
}
```

## annotated.go

```go

// Copyright (c) 2020-2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/dig"
	"go.uber.org/fx/internal/fxreflect"
)

// Annotated annotates a constructor provided to Fx with additional options.
//
// For example,
//
//	func NewReadOnlyConnection(...) (*Connection, error)
//
//	fx.Provide(fx.Annotated{
//	  Name: "ro",
//	  Target: NewReadOnlyConnection,
//	})
//
// Is equivalent to,
//
//	type result struct {
//	  fx.Out
//
//	  Connection *Connection `name:"ro"`
//	}
//
//	fx.Provide(func(...) (result, error) {
//	  conn, err := NewReadOnlyConnection(...)
//	  return result{Connection: conn}, err
//	})
//
// Annotated cannot be used with constructors which produce fx.Out objects.
//
// When used with fx.Supply, the target is a value rather than a constructor function.
type Annotated struct {
	// If specified, this will be used as the name for all non-error values returned
	// by the constructor. For more information on named values, see the documentation
	// for the fx.Out type.
	//
	// A name option may not be provided if a group option is provided.
	Name string

	// If specified, this will be used as the group name for all non-error values returned
	// by the constructor. For more information on value groups, see the package documentation.
	//
	// A group option may not be provided if a name option is provided.
	//
	// Similar to group tags, the group name may be followed by a `,flatten`
	// option to indicate that each element in the slice returned by the
	// constructor should be injected into the value group individually.
	Group string

	// Target is the constructor or value being annotated with fx.Annotated.
	Target interface{}
}

func (a Annotated) String() string {
	var fields []string
	if len(a.Name) > 0 {
		fields = append(fields, fmt.Sprintf("Name: %q", a.Name))
	}
	if len(a.Group) > 0 {
		fields = append(fields, fmt.Sprintf("Group: %q", a.Group))
	}
	if a.Target != nil {
		fields = append(fields, fmt.Sprintf("Target: %v", fxreflect.FuncName(a.Target)))
	}
	return fmt.Sprintf("fx.Annotated{%v}", strings.Join(fields, ", "))
}

var (
	// field used for embedding fx.In type in generated struct.
	_inAnnotationField = reflect.StructField{
		Name:      "In",
		Type:      reflect.TypeOf(In{}),
		Anonymous: true,
	}
	// field used for embedding fx.Out type in generated struct.
	_outAnnotationField = reflect.StructField{
		Name:      "Out",
		Type:      reflect.TypeOf(Out{}),
		Anonymous: true,
	}
)

// Annotation can be passed to Annotate(f interface{}, anns ...Annotation)
// for annotating the parameter and result types of a function.
type Annotation interface {
	apply(*annotated) error
	build(*annotated) (interface{}, error)
}

var (
	_typeOfError reflect.Type = reflect.TypeOf((*error)(nil)).Elem()
	_nilError                 = reflect.Zero(_typeOfError)
)

// annotationError is a wrapper for an error that was encountered while
// applying annotation to a function. It contains the specific error
// that it encountered as well as the target interface that was attempted
// to be annotated.
type annotationError struct {
	target interface{}
	err    error
}

func (e *annotationError) Error() string {
	return e.err.Error()
}

// Unwrap the wrapped error.
func (e *annotationError) Unwrap() error {
	return e.err
}

type paramTagsAnnotation struct {
	tags []string
}

var _ Annotation = paramTagsAnnotation{}
var (
	errTagSyntaxSpace            = errors.New(`multiple tags are not separated by space`)
	errTagKeySyntax              = errors.New("tag key is invalid, Use group, name or optional as tag keys")
	errTagValueSyntaxQuote       = errors.New(`tag value should start with double quote. i.e. key:"value" `)
	errTagValueSyntaxEndingQuote = errors.New(`tag value should end in double quote. i.e. key:"value" `)
)

// Collections of key value pairs within a tag should be separated by a space.
// Eg: `group:"some" optional:"true"`.
func verifyTagsSpaceSeparated(tagIdx int, tag string) error {
	if tagIdx > 0 && tag != "" && tag[0] != ' ' {
		return errTagSyntaxSpace
	}
	return nil
}

// verify tag values are delimited with double quotes.
func verifyValueQuote(value string) (string, error) {
	// starting quote should be a double quote
	if value[0] != '"' {
		return "", errTagValueSyntaxQuote
	}
	// validate tag value is within quotes
	i := 1
	for i < len(value) && value[i] != '"' {
		if value[i] == '\\' {
			i++
		}
		i++
	}
	if i >= len(value) {
		return "", errTagValueSyntaxEndingQuote
	}
	return value[i+1:], nil

}

// Check whether the tag follows valid struct.
// format and returns an error if it's invalid. (i.e. not following
// tag:"value" space-separated list )
// Currently dig accepts only 'name', 'group', 'optional' as valid tag keys.
func verifyAnnotateTag(tag string) error {
	tagIdx := 0
	validKeys := map[string]struct{}{"group": {}, "optional": {}, "name": {}}
	for ; tag != ""; tagIdx++ {
		if err := verifyTagsSpaceSeparated(tagIdx, tag); err != nil {
			return err
		}
		i := 0
		if strings.TrimSpace(tag) == "" {
			return nil
		}
		// parsing the key i.e. till reaching colon :
		for i < len(tag) && tag[i] != ':' {
			i++
		}
		key := strings.TrimSpace(tag[:i])
		if _, ok := validKeys[key]; !ok {
			return errTagKeySyntax
		}
		value, err := verifyValueQuote(tag[i+1:])
		if err != nil {
			return err
		}
		tag = value
	}
	return nil

}

// Given func(T1, T2, T3, ..., TN), this generates a type roughly
// equivalent to,
//
//   struct {
//     fx.In
//
//     Field1 T1 `$tags[0]`
//     Field2 T2 `$tags[1]`
//     ...
//     FieldN TN `$tags[N-1]`
//   }
//
// If there has already been a ParamTag that was applied, this
// will return an error.
//
// If the tag is invalid and has mismatched quotation for example,
// (`tag_name:"tag_value') , this will return an error.

func (pt paramTagsAnnotation) apply(ann *annotated) error {
	if len(ann.ParamTags) > 0 {
		return errors.New("cannot apply more than one line of ParamTags")
	}
	for _, tag := range pt.tags {
		if err := verifyAnnotateTag(tag); err != nil {
			return err
		}
	}
	ann.ParamTags = pt.tags
	return nil
}

// build builds and returns a constructor after applying a ParamTags annotation
func (pt paramTagsAnnotation) build(ann *annotated) (interface{}, error) {
	paramTypes, remap := pt.parameters(ann)
	resultTypes, _ := ann.currentResultTypes()

	origFn := reflect.ValueOf(ann.Target)
	newFnType := reflect.FuncOf(paramTypes, resultTypes, false)
	newFn := reflect.MakeFunc(newFnType, func(args []reflect.Value) []reflect.Value {
		args = remap(args)
		return origFn.Call(args)
	})
	return newFn.Interface(), nil
}

// parameters returns the type for the parameters of the annotated function,
// and a function that maps the arguments of the annotated function
// back to the arguments of the target function.
func (pt paramTagsAnnotation) parameters(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
) {
	ft := reflect.TypeOf(ann.Target)
	types = make([]reflect.Type, ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		types[i] = ft.In(i)
	}

	// No parameter annotations. Return the original types
	// and an identity function.
	if len(pt.tags) == 0 {
		return types, func(args []reflect.Value) []reflect.Value {
			return args
		}
	}

	// Turn parameters into an fx.In struct.
	inFields := []reflect.StructField{_inAnnotationField}

	// there was a variadic argument, so it was pre-transformed
	if len(types) > 0 && isIn(types[0]) {
		paramType := types[0]

		for i := 1; i < paramType.NumField(); i++ {
			origField := paramType.Field(i)
			field := reflect.StructField{
				Name: origField.Name,
				Type: origField.Type,
				Tag:  origField.Tag,
			}
			if i-1 < len(pt.tags) {
				field.Tag = reflect.StructTag(pt.tags[i-1])
			}

			inFields = append(inFields, field)
		}

		types = []reflect.Type{reflect.StructOf(inFields)}
		return types, func(args []reflect.Value) []reflect.Value {
			param := args[0]
			args[0] = reflect.New(paramType).Elem()
			for i := 1; i < paramType.NumField(); i++ {
				args[0].Field(i).Set(param.Field(i))
			}
			return args
		}
	}

	for i, t := range types {
		field := reflect.StructField{
			Name: fmt.Sprintf("Field%d", i),
			Type: t,
		}
		if i < len(pt.tags) {
			field.Tag = reflect.StructTag(pt.tags[i])
		}

		inFields = append(inFields, field)
	}

	types = []reflect.Type{reflect.StructOf(inFields)}
	return types, func(args []reflect.Value) []reflect.Value {
		params := args[0]
		args = args[:0]
		for i := 0; i < ft.NumIn(); i++ {
			args = append(args, params.Field(i+1))
		}
		return args
	}
}

// ParamTags is an Annotation that annotates the parameter(s) of a function.
// When multiple tags are specified, each tag is mapped to the corresponding
// positional parameter.
//
// ParamTags cannot be used in a function that takes an fx.In struct as a
// parameter.
func ParamTags(tags ...string) Annotation {
	return paramTagsAnnotation{tags}
}

type resultTagsAnnotation struct {
	tags []string
}

var _ Annotation = resultTagsAnnotation{}

// Given func(T1, T2, T3, ..., TN), this generates a type roughly
// equivalent to,
//
//	struct {
//	  fx.Out
//
//	  Field1 T1 `$tags[0]`
//	  Field2 T2 `$tags[1]`
//	  ...
//	  FieldN TN `$tags[N-1]`
//	}
//
// If there has already been a ResultTag that was applied, this
// will return an error.
//
// If the tag is invalid and has mismatched quotation for example,
// (`tag_name:"tag_value') , this will return an error.
func (rt resultTagsAnnotation) apply(ann *annotated) error {
	if len(ann.ResultTags) > 0 {
		return errors.New("cannot apply more than one line of ResultTags")
	}
	for _, tag := range rt.tags {
		if err := verifyAnnotateTag(tag); err != nil {
			return err
		}
	}
	ann.ResultTags = rt.tags
	return nil
}

// build builds and returns a constructor after applying a ResultTags annotation
func (rt resultTagsAnnotation) build(ann *annotated) (interface{}, error) {
	paramTypes := ann.currentParamTypes()
	resultTypes, remapResults := rt.results(ann)
	origFn := reflect.ValueOf(ann.Target)
	newFnType := reflect.FuncOf(paramTypes, resultTypes, false)
	newFn := reflect.MakeFunc(newFnType, func(args []reflect.Value) []reflect.Value {
		results := origFn.Call(args)
		return remapResults(results)
	})
	return newFn.Interface(), nil
}

// results returns the types of the results of the annotated function,
// and a function that maps the results of the target function,
// into a result compatible with the annotated function.
func (rt resultTagsAnnotation) results(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
) {
	types, hasError := ann.currentResultTypes()

	if hasError {
		types = types[:len(types)-1]
	}

	// No result annotations. Return the original types
	// and an identity function.
	if len(rt.tags) == 0 {
		return types, func(results []reflect.Value) []reflect.Value {
			return results
		}
	}

	// if there's no Out struct among the return types, there was no As annotation applied
	// just replace original result types with an Out struct and apply tags
	var (
		newOut       outStructInfo
		existingOuts []reflect.Type
	)

	newOut.Fields = []reflect.StructField{_outAnnotationField}
	newOut.Offsets = []int{}

	for i, t := range types {
		if !isOut(t) {
			// this must be from the original function.
			// apply the tags
			field := reflect.StructField{
				Name: fmt.Sprintf("Field%d", i),
				Type: t,
			}
			if i < len(rt.tags) {
				field.Tag = reflect.StructTag(rt.tags[i])
			}
			newOut.Offsets = append(newOut.Offsets, len(newOut.Fields))
			newOut.Fields = append(newOut.Fields, field)
			continue
		}
		// this must be from an As annotation
		// apply the tags to the existing type
		taggedFields := make([]reflect.StructField, t.NumField())
		taggedFields[0] = _outAnnotationField
		for j, tag := range rt.tags {
			if j+1 < t.NumField() {
				field := t.Field(j + 1)
				taggedFields[j+1] = reflect.StructField{
					Name: field.Name,
					Type: field.Type,
					Tag:  reflect.StructTag(tag),
				}
			}
		}
		existingOuts = append(existingOuts, reflect.StructOf(taggedFields))
	}

	resType := reflect.StructOf(newOut.Fields)

	outTypes := []reflect.Type{resType}
	// append existing outs back to outTypes
	outTypes = append(outTypes, existingOuts...)
	if hasError {
		outTypes = append(outTypes, _typeOfError)
	}

	return outTypes, func(results []reflect.Value) []reflect.Value {
		var (
			outErr     error
			outResults []reflect.Value
		)
		outResults = append(outResults, reflect.New(resType).Elem())

		tIdx := 0
		for i, r := range results {
			if i == len(results)-1 && hasError {
				// If hasError and this is the last item,
				// we are guaranteed that this is an error
				// object.
				if err, _ := r.Interface().(error); err != nil {
					outErr = err
				}
				continue
			}
			if i < len(newOut.Offsets) {
				if fieldIdx := newOut.Offsets[i]; fieldIdx > 0 {
					// fieldIdx 0 is an invalid index
					// because it refers to uninitialized
					// outs and would point to fx.Out in the
					// struct definition. We need to check this
					// to prevent panic from setting fx.Out to
					// a value.
					outResults[0].Field(fieldIdx).Set(r)
				}
				continue
			}
			if isOut(r.Type()) {
				tIdx++
				if tIdx < len(outTypes) {
					newResult := reflect.New(outTypes[tIdx]).Elem()
					for j := 1; j < outTypes[tIdx].NumField(); j++ {
						newResult.Field(j).Set(r.Field(j))
					}
					outResults = append(outResults, newResult)
				}
			}
		}

		if hasError {
			if outErr != nil {
				outResults = append(outResults, reflect.ValueOf(outErr))
			} else {
				outResults = append(outResults, _nilError)
			}
		}

		return outResults
	}
}

// ResultTags is an Annotation that annotates the result(s) of a function.
// When multiple tags are specified, each tag is mapped to the corresponding
// positional result.
//
// ResultTags cannot be used on a function that returns an fx.Out struct.
func ResultTags(tags ...string) Annotation {
	return resultTagsAnnotation{tags}
}

type outStructInfo struct {
	Fields  []reflect.StructField // fields of the struct
	Offsets []int                 // Offsets[i] is the index of result i in Fields
}

type _lifecycleHookAnnotationType int

const (
	_unknownHookType _lifecycleHookAnnotationType = iota
	_onStartHookType
	_onStopHookType
)

type lifecycleHookAnnotation struct {
	Type   _lifecycleHookAnnotationType
	Target interface{}
}

var _ Annotation = (*lifecycleHookAnnotation)(nil)

func (la *lifecycleHookAnnotation) String() string {
	name := "UnknownHookAnnotation"
	switch la.Type {
	case _onStartHookType:
		name = _onStartHook
	case _onStopHookType:
		name = _onStopHook
	}
	return name
}

func (la *lifecycleHookAnnotation) apply(ann *annotated) error {
	if la.Target == nil {
		return fmt.Errorf(
			"cannot use nil function for %q hook annotation",
			la,
		)
	}

	for _, h := range ann.Hooks {
		if la.Type == h.Type {
			return fmt.Errorf(
				"cannot apply more than one %q hook annotation",
				la,
			)
		}
	}

	ft := reflect.TypeOf(la.Target)

	if ft.Kind() != reflect.Func {
		return fmt.Errorf(
			"must provide function for %q hook, got %v (%T)",
			la,
			la.Target,
			la.Target,
		)
	}

	if n := ft.NumOut(); n > 0 {
		if n > 1 || ft.Out(0) != _typeOfError {
			return fmt.Errorf(
				"optional hook return may only be an error, got %v (%T)",
				la.Target,
				la.Target,
			)
		}
	}

	if ft.IsVariadic() {
		return fmt.Errorf(
			"hooks must not accept variadic parameters, got %v (%T)",
			la.Target,
			la.Target,
		)
	}

	ann.Hooks = append(ann.Hooks, la)
	return nil
}

// build builds and returns a constructor after applying a lifecycle hook annotation.
func (la *lifecycleHookAnnotation) build(ann *annotated) (interface{}, error) {
	resultTypes, hasError := ann.currentResultTypes()
	if !hasError {
		resultTypes = append(resultTypes, _typeOfError)
	}

	hookInstaller, paramTypes, remapParams := la.buildHookInstaller(ann)

	origFn := reflect.ValueOf(ann.Target)
	newFnType := reflect.FuncOf(paramTypes, resultTypes, false)
	newFn := reflect.MakeFunc(newFnType, func(args []reflect.Value) []reflect.Value {
		// copy the original arguments before remapping the parameters
		// so that we can apply them to the hookInstaller.
		origArgs := make([]reflect.Value, len(args))
		copy(origArgs, args)
		args = remapParams(args)
		results := origFn.Call(args)
		if hasError {
			errVal := results[len(results)-1]
			results = results[:len(results)-1]
			if err, _ := errVal.Interface().(error); err != nil {
				// if constructor returned error, do not call hook installer
				return append(results, errVal)
			}
		}
		hookInstallerResults := hookInstaller.Call(append(results, origArgs...))
		results = append(results, hookInstallerResults[0])
		return results
	})
	return newFn.Interface(), nil
}

var (
	_typeOfLifecycle reflect.Type = reflect.TypeOf((*Lifecycle)(nil)).Elem()
	_typeOfContext   reflect.Type = reflect.TypeOf((*context.Context)(nil)).Elem()
)

// buildHookInstaller returns a function that appends a hook to Lifecycle when called,
// along with the new parameter types and a function that maps arguments to the annotated constructor
func (la *lifecycleHookAnnotation) buildHookInstaller(ann *annotated) (
	hookInstaller reflect.Value,
	paramTypes []reflect.Type,
	remapParams func([]reflect.Value) []reflect.Value, // function to remap parameters to function being annotated
) {
	paramTypes = ann.currentParamTypes()
	paramTypes, remapParams = injectLifecycle(paramTypes)

	resultTypes, hasError := ann.currentResultTypes()
	if hasError {
		resultTypes = resultTypes[:len(resultTypes)-1]
	}

	// look for the context.Context type from the original hook function
	// and then exclude it from the paramTypes of invokeFn because context.Context
	// will be injected by the lifecycle
	ctxPos := -1
	ctxStructPos := -1
	origHookFn := reflect.ValueOf(la.Target)
	origHookFnT := reflect.TypeOf(la.Target)
	invokeParamTypes := []reflect.Type{
		_typeOfLifecycle,
	}
	for i := 0; i < origHookFnT.NumIn(); i++ {
		t := origHookFnT.In(i)
		if t == _typeOfContext && ctxPos < 0 {
			ctxPos = i
			continue
		}
		if !isIn(t) {
			invokeParamTypes = append(invokeParamTypes, origHookFnT.In(i))
			continue
		}
		fields := []reflect.StructField{_inAnnotationField}
		for j := 1; j < t.NumField(); j++ {
			field := t.Field(j)
			if field.Type == _typeOfContext && ctxPos < 0 {
				ctxStructPos = i
				ctxPos = j
				continue
			}
			fields = append(fields, field)
		}
		invokeParamTypes = append(invokeParamTypes, reflect.StructOf(fields))

	}
	invokeFnT := reflect.FuncOf(invokeParamTypes, []reflect.Type{}, false)
	invokeFn := reflect.MakeFunc(invokeFnT, func(args []reflect.Value) (results []reflect.Value) {
		lc := args[0].Interface().(Lifecycle)
		args = args[1:]
		hookArgs := make([]reflect.Value, origHookFnT.NumIn())

		hookFn := func(ctx context.Context) (err error) {
			// If the hook function has multiple parameters, and the first
			// parameter is a context, inject the provided context.
			if ctxStructPos < 0 {
				offset := 0
				for i := 0; i < len(hookArgs); i++ {
					if i == ctxPos {
						hookArgs[i] = reflect.ValueOf(ctx)
						offset = 1
						continue
					}
					if i-offset >= 0 && i-offset < len(args) {
						hookArgs[i] = args[i-offset]
					}
				}
			} else {
				for i := 0; i < origHookFnT.NumIn(); i++ {
					if i != ctxStructPos {
						hookArgs[i] = args[i]
						continue
					}
					t := origHookFnT.In(i)
					v := reflect.New(t).Elem()
					for j := 1; j < t.NumField(); j++ {
						if j < ctxPos {
							v.Field(j).Set(args[i].Field(j))
						} else if j == ctxPos {
							v.Field(j).Set(reflect.ValueOf(ctx))
						} else {
							v.Field(j).Set(args[i].Field(j - 1))
						}
					}
					hookArgs[i] = v
				}
			}
			hookResults := origHookFn.Call(hookArgs)
			if len(hookResults) > 0 && hookResults[0].Type() == _typeOfError {
				err, _ = hookResults[0].Interface().(error)
			}
			return err
		}
		lc.Append(la.buildHook(hookFn))
		return results
	})

	installerType := reflect.FuncOf(append(resultTypes, paramTypes...), []reflect.Type{_typeOfError}, false)
	hookInstaller = reflect.MakeFunc(installerType, func(args []reflect.Value) (results []reflect.Value) {
		// build a private scope for hook function
		var scope *dig.Scope
		switch la.Type {
		case _onStartHookType:
			scope = ann.container.Scope("onStartHookScope")
		case _onStopHookType:
			scope = ann.container.Scope("onStopHookScope")
		}

		// provide the private scope with the current dependencies and results of the annotated function
		results = []reflect.Value{_nilError}
		ctor := makeHookScopeCtor(paramTypes, resultTypes, args)
		if err := scope.Provide(ctor); err != nil {
			results[0] = reflect.ValueOf(fmt.Errorf("error providing possible parameters for hook installer: %w", err))
			return results
		}

		// invoking invokeFn appends the hook function to lifecycle
		if err := scope.Invoke(invokeFn.Interface()); err != nil {
			results[0] = reflect.ValueOf(fmt.Errorf("error invoking hook installer: %w", err))
			return results
		}
		return results
	})
	return hookInstaller, paramTypes, remapParams
}

var (
	_nameTag  = "name"
	_groupTag = "group"
)

// makeHookScopeCtor makes a constructor that provides all possible parameters
// that the lifecycle hook being appended can depend on. It also deduplicates
// duplicate param and result types, which is possible when using fx.Decorate,
// and uses values from results for providing the deduplicated types.
func makeHookScopeCtor(paramTypes []reflect.Type, resultTypes []reflect.Type, args []reflect.Value) interface{} {
	type key struct {
		t     reflect.Type
		name  string
		group string
	}
	seen := map[key]struct{}{}
	outTypes := make([]reflect.Type, len(resultTypes))
	for i, t := range resultTypes {
		outTypes[i] = t
		if isOut(t) {
			for j := 1; j < t.NumField(); j++ {
				field := t.Field(j)
				seen[key{
					t:     field.Type,
					name:  field.Tag.Get(_nameTag),
					group: field.Tag.Get(_groupTag),
				}] = struct{}{}
			}
			continue
		}
		seen[key{t: t}] = struct{}{}
	}

	fields := []reflect.StructField{_outAnnotationField}

	skippedParams := make([][]int, len(paramTypes))

	for i, t := range paramTypes {
		skippedParams[i] = []int{}
		if isIn(t) {
			for j := 1; j < t.NumField(); j++ {
				origField := t.Field(j)
				k := key{
					t:     origField.Type,
					name:  origField.Tag.Get(_nameTag),
					group: origField.Tag.Get(_groupTag),
				}

				if _, ok := seen[k]; ok {
					skippedParams[i] = append(skippedParams[i], j)
					continue
				}

				field := reflect.StructField{
					Name: fmt.Sprintf("Field%d", j-1),
					Type: origField.Type,
					Tag:  origField.Tag,
				}
				fields = append(fields, field)
			}
			continue
		}
		k := key{t: t}

		if _, ok := seen[k]; ok {
			skippedParams[i] = append(skippedParams[i], i)
			continue
		}
		field := reflect.StructField{
			Name: fmt.Sprintf("Field%d", i),
			Type: t,
		}
		fields = append(fields, field)
	}

	outTypes = append(outTypes, reflect.StructOf(fields))
	ctorType := reflect.FuncOf([]reflect.Type{}, outTypes, false)
	ctor := reflect.MakeFunc(ctorType, func(_ []reflect.Value) []reflect.Value {
		nOut := len(outTypes)
		results := make([]reflect.Value, nOut)
		for i := 0; i < nOut-1; i++ {
			results[i] = args[i]
		}

		v := reflect.New(outTypes[nOut-1]).Elem()
		fieldIdx := 1
		for i := nOut - 1; i < len(args); i++ {
			paramIdx := i - (nOut - 1)
			if isIn(paramTypes[paramIdx]) {
				skippedIdx := 0
				for j := 1; j < paramTypes[paramIdx].NumField(); j++ {
					if len(skippedParams[paramIdx]) > 0 && skippedParams[paramIdx][skippedIdx] == j {
						// skip
						skippedIdx++
						continue
					}
					v.Field(fieldIdx).Set(args[i].Field(j))
					fieldIdx++
				}
			} else {
				if len(skippedParams[paramIdx]) > 0 && skippedParams[paramIdx][0] == paramIdx {
					continue
				}
				v.Field(fieldIdx).Set(args[i])
				fieldIdx++
			}
		}
		results[nOut-1] = v

		return results
	})
	return ctor.Interface()
}

func injectLifecycle(paramTypes []reflect.Type) ([]reflect.Type, func([]reflect.Value) []reflect.Value) {
	// since lifecycle already exists in param types, no need to inject again
	if lifecycleExists(paramTypes) {
		return paramTypes, func(args []reflect.Value) []reflect.Value {
			return args
		}
	}
	// If params are tagged or there's an untagged variadic argument,
	// add a Lifecycle field to the param struct
	if len(paramTypes) > 0 && isIn(paramTypes[0]) {
		taggedParam := paramTypes[0]
		fields := []reflect.StructField{
			taggedParam.Field(0),
			{
				Name: "Lifecycle",
				Type: _typeOfLifecycle,
			},
		}
		for i := 1; i < taggedParam.NumField(); i++ {
			fields = append(fields, taggedParam.Field(i))
		}
		newParamType := reflect.StructOf(fields)
		return []reflect.Type{newParamType}, func(args []reflect.Value) []reflect.Value {
			param := args[0]
			args[0] = reflect.New(taggedParam).Elem()
			for i := 1; i < taggedParam.NumField(); i++ {
				args[0].Field(i).Set(param.Field(i + 1))
			}
			return args
		}
	}

	return append([]reflect.Type{_typeOfLifecycle}, paramTypes...), func(args []reflect.Value) []reflect.Value {
		return args[1:]
	}
}

func lifecycleExists(paramTypes []reflect.Type) bool {
	for _, t := range paramTypes {
		if t == _typeOfLifecycle {
			return true
		}
		if isIn(t) {
			for i := 1; i < t.NumField(); i++ {
				if t.Field(i).Type == _typeOfLifecycle {
					return true
				}
			}
		}
	}
	return false
}

func (la *lifecycleHookAnnotation) buildHook(fn func(context.Context) error) (hook Hook) {
	switch la.Type {
	case _onStartHookType:
		hook.OnStart = fn
	case _onStopHookType:
		hook.OnStop = fn
	}
	return hook
}

// OnStart is an Annotation that appends an OnStart Hook to the application
// Lifecycle when that function is called. This provides a way to create
// Lifecycle OnStart (see Lifecycle type documentation) hooks without building a
// function that takes a dependency on the Lifecycle type.
//
//	fx.Provide(
//		fx.Annotate(
//			NewServer,
//			fx.OnStart(func(ctx context.Context, server Server) error {
//				return server.Listen(ctx)
//			}),
//		)
//	)
//
// Which is functionally the same as:
//
//	 fx.Provide(
//	   func(lifecycle fx.Lifecycle, p Params) Server {
//	     server := NewServer(p)
//	     lifecycle.Append(fx.Hook{
//		      OnStart: func(ctx context.Context) error {
//			    return server.Listen(ctx)
//		      },
//	     })
//		 return server
//	   }
//	 )
//
// It is also possible to use OnStart annotation with other parameter and result
// annotations, provided that the parameter of the function passed to OnStart
// matches annotated parameters and results.
//
// For example, the following is possible:
//
//	fx.Provide(
//		fx.Annotate(
//			func (a A) B {...},
//			fx.ParamTags(`name:"A"`),
//			fx.ResultTags(`name:"B"`),
//			fx.OnStart(func (p OnStartParams) {...}),
//		),
//	)
//
// As long as OnStartParams looks like the following and has no other dependencies
// besides Context or Lifecycle:
//
//	type OnStartParams struct {
//		fx.In
//		FieldA A `name:"A"`
//		FieldB B `name:"B"`
//	}
//
// Only one OnStart annotation may be applied to a given function at a time,
// however functions may be annotated with other types of lifecycle Hooks, such
// as OnStop. The hook function passed into OnStart cannot take any arguments
// outside of the annotated constructor's existing dependencies or results, except
// a context.Context.
func OnStart(onStart interface{}) Annotation {
	return &lifecycleHookAnnotation{
		Type:   _onStartHookType,
		Target: onStart,
	}
}

// OnStop is an Annotation that appends an OnStop Hook to the application
// Lifecycle when that function is called. This provides a way to create
// Lifecycle OnStop (see Lifecycle type documentation) hooks without building a
// function that takes a dependency on the Lifecycle type.
//
//	fx.Provide(
//		fx.Annotate(
//			NewServer,
//			fx.OnStop(func(ctx context.Context, server Server) error {
//				return server.Shutdown(ctx)
//			}),
//		)
//	)
//
// Which is functionally the same as:
//
//	 fx.Provide(
//	   func(lifecycle fx.Lifecycle, p Params) Server {
//	     server := NewServer(p)
//	     lifecycle.Append(fx.Hook{
//		      OnStop: func(ctx context.Context) error {
//			    return server.Shutdown(ctx)
//		      },
//	     })
//		 return server
//	   }
//	 )
//
// It is also possible to use OnStop annotation with other parameter and result
// annotations, provided that the parameter of the function passed to OnStop
// matches annotated parameters and results.
//
// For example, the following is possible:
//
//	fx.Provide(
//		fx.Annotate(
//			func (a A) B {...},
//			fx.ParamTags(`name:"A"`),
//			fx.ResultTags(`name:"B"`),
//			fx.OnStop(func (p OnStopParams) {...}),
//		),
//	)
//
// As long as OnStopParams looks like the following and has no other dependencies
// besides Context or Lifecycle:
//
//	type OnStopParams struct {
//		fx.In
//		FieldA A `name:"A"`
//		FieldB B `name:"B"`
//	}
//
// Only one OnStop annotation may be applied to a given function at a time,
// however functions may be annotated with other types of lifecycle Hooks, such
// as OnStart. The hook function passed into OnStop cannot take any arguments
// outside of the annotated constructor's existing dependencies or results, except
// a context.Context.
func OnStop(onStop interface{}) Annotation {
	return &lifecycleHookAnnotation{
		Type:   _onStopHookType,
		Target: onStop,
	}
}

type asAnnotation struct {
	targets []interface{}
	types   []reflect.Type
}

func isOut(t reflect.Type) bool {
	return (t.Kind() == reflect.Struct &&
		dig.IsOut(reflect.New(t).Elem().Interface()))
}

func isIn(t reflect.Type) bool {
	return (t.Kind() == reflect.Struct &&
		dig.IsIn(reflect.New(t).Elem().Interface()))
}

var _ Annotation = (*asAnnotation)(nil)

// As is an Annotation that annotates the result of a function (i.e. a
// constructor) to be provided as another interface.
//
// For example, the following code specifies that the return type of
// bytes.NewBuffer (bytes.Buffer) should be provided as io.Writer type:
//
//	fx.Provide(
//	  fx.Annotate(bytes.NewBuffer(...), fx.As(new(io.Writer)))
//	)
//
// In other words, the code above is equivalent to:
//
//	fx.Provide(func() io.Writer {
//	  return bytes.NewBuffer()
//	  // provides io.Writer instead of *bytes.Buffer
//	})
//
// Note that the bytes.Buffer type is provided as an io.Writer type, so this
// constructor does NOT provide both bytes.Buffer and io.Writer type; it just
// provides io.Writer type.
//
// When multiple values are returned by the annotated function, each type
// gets mapped to corresponding positional result of the annotated function.
//
// For example,
//
//	func a() (bytes.Buffer, bytes.Buffer) {
//	  ...
//	}
//	fx.Provide(
//	  fx.Annotate(a, fx.As(new(io.Writer), new(io.Reader)))
//	)
//
// Is equivalent to,
//
//	fx.Provide(func() (io.Writer, io.Reader) {
//	  w, r := a()
//	  return w, r
//	}
//
// As annotation cannot be used in a function that returns an [Out] struct as a return type.
func As(interfaces ...interface{}) Annotation {
	return &asAnnotation{targets: interfaces}
}

func (at *asAnnotation) apply(ann *annotated) error {
	at.types = make([]reflect.Type, len(at.targets))
	for i, typ := range at.targets {
		t := reflect.TypeOf(typ)
		if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Interface {
			return fmt.Errorf("fx.As: argument must be a pointer to an interface: got %v", t)
		}
		t = t.Elem()
		at.types[i] = t
	}

	ann.As = append(ann.As, at.types)
	return nil
}

// build implements Annotation
func (at *asAnnotation) build(ann *annotated) (interface{}, error) {
	paramTypes := ann.currentParamTypes()

	resultTypes, remapResults, err := at.results(ann)
	if err != nil {
		return nil, err
	}

	origFn := reflect.ValueOf(ann.Target)
	newFnType := reflect.FuncOf(paramTypes, resultTypes, false)
	newFn := reflect.MakeFunc(newFnType, func(args []reflect.Value) []reflect.Value {
		results := origFn.Call(args)
		return remapResults(results)
	})
	return newFn.Interface(), nil
}

func (at *asAnnotation) results(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
	err error,
) {
	types, hasError := ann.currentResultTypes()
	fields := []reflect.StructField{_outAnnotationField}
	if hasError {
		types = types[:len(types)-1]
	}
	resultFields, getResult := extractResultFields(types)

	for i, f := range resultFields {
		t := f.Type
		field := reflect.StructField{
			Name: fmt.Sprintf("Field%d", i),
			Type: t,
			Tag:  f.Tag,
		}
		if i < len(at.types) {
			if !t.Implements(at.types[i]) {
				return nil, nil, fmt.Errorf("invalid fx.As: %v does not implement %v", t, at.types[i])
			}
			field.Type = at.types[i]
		}
		fields = append(fields, field)
	}
	resType := reflect.StructOf(fields)

	var outTypes []reflect.Type
	outTypes = append(types, resType)
	if hasError {
		outTypes = append(outTypes, _typeOfError)
	}

	return outTypes, func(results []reflect.Value) []reflect.Value {
		var (
			outErr     error
			outResults []reflect.Value
		)

		for i, r := range results {
			if i == len(results)-1 && hasError {
				// If hasError and this is the last item,
				// we are guaranteed that this is an error
				// object.
				if err, _ := r.Interface().(error); err != nil {
					outErr = err
				}
				continue
			}
			outResults = append(outResults, r)
		}

		newOutResult := reflect.New(resType).Elem()
		for i := 1; i < resType.NumField(); i++ {
			newOutResult.Field(i).Set(getResult(i, results))
		}
		outResults = append(outResults, newOutResult)

		if hasError {
			if outErr != nil {
				outResults = append(outResults, reflect.ValueOf(outErr))
			} else {
				outResults = append(outResults, _nilError)
			}
		}

		return outResults
	}, nil
}

func extractResultFields(types []reflect.Type) ([]reflect.StructField, func(int, []reflect.Value) reflect.Value) {
	var resultFields []reflect.StructField
	if len(types) > 0 && isOut(types[0]) {
		for i := 1; i < types[0].NumField(); i++ {
			resultFields = append(resultFields, types[0].Field(i))
		}
		return resultFields, func(idx int, results []reflect.Value) reflect.Value {
			return results[0].Field(idx)
		}
	}
	for i, t := range types {
		if isOut(t) {
			continue
		}
		field := reflect.StructField{
			Name: fmt.Sprintf("Field%d", i),
			Type: t,
		}
		resultFields = append(resultFields, field)
	}
	return resultFields, func(idx int, results []reflect.Value) reflect.Value {
		return results[idx-1]
	}
}

type fromAnnotation struct {
	targets []interface{}
	types   []reflect.Type
}

var _ Annotation = (*fromAnnotation)(nil)

// From is an [Annotation] that annotates the parameter(s) for a function (i.e. a
// constructor) to be accepted from other provided types. It is analogous to the
// [As] for parameter types to the constructor.
//
// For example,
//
//	type Runner interface { Run() }
//	func NewFooRunner() *FooRunner // implements Runner
//	func NewRunnerWrap(r Runner) *RunnerWrap
//
//	fx.Provide(
//	  fx.Annotate(
//	    NewRunnerWrap,
//	    fx.From(new(*FooRunner)),
//	  ),
//	)
//
// Is equivalent to,
//
//	fx.Provide(func(r *FooRunner) *RunnerWrap {
//	  // need *FooRunner instead of Runner
//	  return NewRunnerWrap(r)
//	})
//
// When the annotated function takes in multiple parameters, each type gets
// mapped to corresponding positional parameter of the annotated function
//
// For example,
//
//	func NewBarRunner() *BarRunner // implements Runner
//	func NewRunnerWraps(r1 Runner, r2 Runner) *RunnerWraps
//
//	fx.Provide(
//	  fx.Annotate(
//	    NewRunnerWraps,
//	    fx.From(new(*FooRunner), new(*BarRunner)),
//	  ),
//	)
//
// Is equivalent to,
//
//	fx.Provide(func(r1 *FooRunner, r2 *BarRunner) *RunnerWraps {
//	  return NewRunnerWraps(r1, r2)
//	})
//
// From annotation cannot be used in a function that takes an [In] struct as a
// parameter.
func From(interfaces ...interface{}) Annotation {
	return &fromAnnotation{targets: interfaces}
}

func (fr *fromAnnotation) apply(ann *annotated) error {
	if len(ann.From) > 0 {
		return errors.New("cannot apply more than one line of From")
	}
	ft := reflect.TypeOf(ann.Target)
	fr.types = make([]reflect.Type, len(fr.targets))
	for i, typ := range fr.targets {
		if ft.IsVariadic() && i == ft.NumIn()-1 {
			return errors.New("fx.From: cannot annotate a variadic argument")
		}
		t := reflect.TypeOf(typ)
		if t == nil || t.Kind() != reflect.Ptr {
			return fmt.Errorf("fx.From: argument must be a pointer to a type that implements some interface: got %v", t)
		}
		fr.types[i] = t.Elem()
	}
	ann.From = fr.types
	return nil
}

// build builds and returns a constructor after applying a From annotation
func (fr *fromAnnotation) build(ann *annotated) (interface{}, error) {
	paramTypes, remap, err := fr.parameters(ann)
	if err != nil {
		return nil, err
	}
	resultTypes, _ := ann.currentResultTypes()

	origFn := reflect.ValueOf(ann.Target)
	newFnType := reflect.FuncOf(paramTypes, resultTypes, false)
	newFn := reflect.MakeFunc(newFnType, func(args []reflect.Value) []reflect.Value {
		args = remap(args)
		return origFn.Call(args)
	})
	return newFn.Interface(), nil
}

// parameters returns the type for the parameters of the annotated function,
// and a function that maps the arguments of the annotated function
// back to the arguments of the target function.
func (fr *fromAnnotation) parameters(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
	err error,
) {
	ft := reflect.TypeOf(ann.Target)
	types = make([]reflect.Type, ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		types[i] = ft.In(i)
	}

	// No parameter annotations. Return the original types
	// and an identity function.
	if len(fr.targets) == 0 {
		return types, func(args []reflect.Value) []reflect.Value {
			return args
		}, nil
	}

	// Turn parameters into an fx.In struct.
	inFields := []reflect.StructField{_inAnnotationField}

	// The following situations may occur:
	// 1. there was a variadic argument, so it was pre-transformed.
	// 2. another parameter annotation was transformed (ex: ParamTags).
	// so need to visit fields of the fx.In struct.
	if len(types) > 0 && isIn(types[0]) {
		paramType := types[0]

		for i := 1; i < paramType.NumField(); i++ {
			origField := paramType.Field(i)
			field := reflect.StructField{
				Name: origField.Name,
				Type: origField.Type,
				Tag:  origField.Tag,
			}
			if i-1 < len(fr.types) {
				t := fr.types[i-1]
				if !t.Implements(field.Type) {
					return nil, nil, fmt.Errorf("invalid fx.From: %v does not implement %v", t, field.Type)
				}
				field.Type = t
			}

			inFields = append(inFields, field)
		}

		types = []reflect.Type{reflect.StructOf(inFields)}
		return types, func(args []reflect.Value) []reflect.Value {
			param := args[0]
			args[0] = reflect.New(paramType).Elem()
			for i := 1; i < paramType.NumField(); i++ {
				args[0].Field(i).Set(param.Field(i))
			}
			return args
		}, nil
	}

	for i, t := range types {
		field := reflect.StructField{
			Name: fmt.Sprintf("Field%d", i),
			Type: t,
		}
		if i < len(fr.types) {
			t := fr.types[i]
			if !t.Implements(field.Type) {
				return nil, nil, fmt.Errorf("invalid fx.From: %v does not implement %v", t, field.Type)
			}
			field.Type = t
		}

		inFields = append(inFields, field)
	}

	types = []reflect.Type{reflect.StructOf(inFields)}
	return types, func(args []reflect.Value) []reflect.Value {
		params := args[0]
		args = args[:0]
		for i := 0; i < ft.NumIn(); i++ {
			args = append(args, params.Field(i+1))
		}
		return args
	}, nil
}

type annotated struct {
	Target      interface{}
	Annotations []Annotation
	ParamTags   []string
	ResultTags  []string
	As          [][]reflect.Type
	From        []reflect.Type
	FuncPtr     uintptr
	Hooks       []*lifecycleHookAnnotation
	// container is used to build private scopes for lifecycle hook functions
	// added via fx.OnStart and fx.OnStop annotations.
	container *dig.Container
}

func (ann annotated) String() string {
	var sb strings.Builder
	sb.WriteString("fx.Annotate(")
	sb.WriteString(fxreflect.FuncName(ann.Target))
	if tags := ann.ParamTags; len(tags) > 0 {
		fmt.Fprintf(&sb, ", fx.ParamTags(%q)", tags)
	}
	if tags := ann.ResultTags; len(tags) > 0 {
		fmt.Fprintf(&sb, ", fx.ResultTags(%q)", tags)
	}
	if as := ann.As; len(as) > 0 {
		fmt.Fprintf(&sb, ", fx.As(%v)", as)
	}
	if from := ann.From; len(from) > 0 {
		fmt.Fprintf(&sb, ", fx.From(%v)", from)
	}
	return sb.String()
}

// Build builds and returns a constructor based on fx.In/fx.Out params and
// results wrapping the original constructor passed to fx.Annotate.
func (ann *annotated) Build() (interface{}, error) {
	ann.container = dig.New()
	ft := reflect.TypeOf(ann.Target)
	if ft.Kind() != reflect.Func {
		return nil, fmt.Errorf("must provide constructor function, got %v (%T)", ann.Target, ann.Target)
	}

	if err := ann.typeCheckOrigFn(); err != nil {
		return nil, fmt.Errorf("invalid annotation function %T: %w", ann.Target, err)
	}

	ann.applyOptionalTag()

	var (
		err        error
		lcHookAnns []*lifecycleHookAnnotation
	)
	for _, annotation := range ann.Annotations {
		if lcHookAnn, ok := annotation.(*lifecycleHookAnnotation); ok {
			lcHookAnns = append(lcHookAnns, lcHookAnn)
			continue
		}
		if ann.Target, err = annotation.build(ann); err != nil {
			return nil, err
		}
	}

	// need to call cleanUpAsResults before applying lifecycle annotations
	// to exclude the original results from the hook's scope if any
	// fx.As annotations were applied
	ann.cleanUpAsResults()

	for _, la := range lcHookAnns {
		if ann.Target, err = la.build(ann); err != nil {
			return nil, err
		}
	}
	return ann.Target, nil
}

// applyOptionalTag checks if function being annotated is variadic
// and applies optional tag to the variadic argument before
// applying any other annotations
func (ann *annotated) applyOptionalTag() {
	ft := reflect.TypeOf(ann.Target)
	if !ft.IsVariadic() {
		return
	}

	resultTypes, _ := ann.currentResultTypes()

	fields := []reflect.StructField{_inAnnotationField}
	for i := 0; i < ft.NumIn(); i++ {
		field := reflect.StructField{
			Name: fmt.Sprintf("Field%d", i),
			Type: ft.In(i),
		}
		if i == ft.NumIn()-1 {
			// Mark a variadic argument optional by default
			// so that just wrapping a function in fx.Annotate does not
			// suddenly introduce a required []arg dependency.
			field.Tag = reflect.StructTag(`optional:"true"`)
		}
		fields = append(fields, field)
	}
	paramType := reflect.StructOf(fields)
	origFn := reflect.ValueOf(ann.Target)
	newFnType := reflect.FuncOf([]reflect.Type{paramType}, resultTypes, false)
	newFn := reflect.MakeFunc(newFnType, func(args []reflect.Value) []reflect.Value {
		params := args[0]
		args = args[:0]
		for i := 0; i < ft.NumIn(); i++ {
			args = append(args, params.Field(i+1))
		}
		return origFn.CallSlice(args)
	})
	ann.Target = newFn.Interface()
}

// cleanUpAsResults does a check to see if an As annotation was applied.
// If there was any fx.As annotation applied, cleanUpAsResults wraps the
// function one more time to remove the results from the original function.
func (ann *annotated) cleanUpAsResults() {
	// clean up orig function results if there were any As annotations
	if len(ann.As) < 1 {
		return
	}
	paramTypes := ann.currentParamTypes()
	resultTypes, hasError := ann.currentResultTypes()
	numRes := len(ann.As)
	if hasError {
		numRes++
	}
	newResultTypes := resultTypes[len(resultTypes)-numRes:]
	origFn := reflect.ValueOf(ann.Target)
	newFnType := reflect.FuncOf(paramTypes, newResultTypes, false)
	newFn := reflect.MakeFunc(newFnType, func(args []reflect.Value) (results []reflect.Value) {
		results = origFn.Call(args)
		results = results[len(results)-numRes:]
		return
	})
	ann.Target = newFn.Interface()
}

// checks and returns a non-nil error if the target function:
// - returns an fx.Out struct as a result and has either a ResultTags or an As annotation
// - takes in an fx.In struct as a parameter and has either a ParamTags or a From annotation
// - has an error result not as the last result.
func (ann *annotated) typeCheckOrigFn() error {
	ft := reflect.TypeOf(ann.Target)
	numOut := ft.NumOut()
	for i := 0; i < numOut; i++ {
		ot := ft.Out(i)
		if ot == _typeOfError && i != numOut-1 {
			return fmt.Errorf(
				"only the last result can be an error: "+
					"%v (%v) returns error as result %d",
				fxreflect.FuncName(ann.Target), ft, i)
		}
		if ot.Kind() != reflect.Struct {
			continue
		}
		if !dig.IsOut(reflect.New(ft.Out(i)).Elem().Interface()) {
			continue
		}
		if len(ann.ResultTags) > 0 || len(ann.As) > 0 {
			return errors.New("fx.Out structs cannot be annotated with fx.ResultTags or fx.As")
		}
	}
	for i := 0; i < ft.NumIn(); i++ {
		it := ft.In(i)
		if it.Kind() != reflect.Struct {
			continue
		}
		if !dig.IsIn(reflect.New(ft.In(i)).Elem().Interface()) {
			continue
		}
		if len(ann.ParamTags) > 0 || len(ann.From) > 0 {
			return errors.New("fx.In structs cannot be annotated with fx.ParamTags or fx.From")
		}
	}
	return nil
}

func (ann *annotated) currentResultTypes() (resultTypes []reflect.Type, hasError bool) {
	ft := reflect.TypeOf(ann.Target)
	numOut := ft.NumOut()
	resultTypes = make([]reflect.Type, numOut)

	for i := 0; i < numOut; i++ {
		resultTypes[i] = ft.Out(i)
		if resultTypes[i] == _typeOfError && i == numOut-1 {
			hasError = true
		}
	}
	return resultTypes, hasError
}

func (ann *annotated) currentParamTypes() []reflect.Type {
	ft := reflect.TypeOf(ann.Target)
	paramTypes := make([]reflect.Type, ft.NumIn())

	for i := 0; i < ft.NumIn(); i++ {
		paramTypes[i] = ft.In(i)
	}
	return paramTypes
}

// Annotate lets you annotate a function's parameters and returns
// without you having to declare separate struct definitions for them.
//
// For example,
//
//	func NewGateway(ro, rw *db.Conn) *Gateway { ... }
//	fx.Provide(
//	  fx.Annotate(
//	    NewGateway,
//	    fx.ParamTags(`name:"ro" optional:"true"`, `name:"rw"`),
//	    fx.ResultTags(`name:"foo"`),
//	  ),
//	)
//
// Is equivalent to,
//
//	type params struct {
//	  fx.In
//
//	  RO *db.Conn `name:"ro" optional:"true"`
//	  RW *db.Conn `name:"rw"`
//	}
//
//	type result struct {
//	  fx.Out
//
//	  GW *Gateway `name:"foo"`
//	}
//
//	fx.Provide(func(p params) result {
//	   return result{GW: NewGateway(p.RO, p.RW)}
//	})
//
// Using the same annotation multiple times is invalid.
// For example, the following will fail with an error:
//
//	fx.Provide(
//	  fx.Annotate(
//	    NewGateWay,
//	    fx.ParamTags(`name:"ro" optional:"true"`),
//	    fx.ParamTags(`name:"rw"), // ERROR: ParamTags was already used above
//	    fx.ResultTags(`name:"foo"`)
//	  )
//	)
//
// is considered an invalid usage and will not apply any of the
// Annotations to NewGateway.
//
// If more tags are given than the number of parameters/results, only
// the ones up to the number of parameters/results will be applied.
//
// # Variadic functions
//
// If the provided function is variadic, Annotate treats its parameter as a
// slice. For example,
//
//	fx.Annotate(func(w io.Writer, rs ...io.Reader) {
//	  // ...
//	}, ...)
//
// Is equivalent to,
//
//	fx.Annotate(func(w io.Writer, rs []io.Reader) {
//	  // ...
//	}, ...)
//
// You can use variadic parameters with Fx's value groups.
// For example,
//
//	fx.Annotate(func(mux *http.ServeMux, handlers ...http.Handler) {
//	  // ...
//	}, fx.ParamTags(``, `group:"server"`))
//
// If we provide the above to the application,
// any constructor in the Fx application can inject its HTTP handlers
// by using fx.Annotate, fx.Annotated, or fx.Out.
//
//	fx.Annotate(
//	  func(..) http.Handler { ... },
//	  fx.ResultTags(`group:"server"`),
//	)
//
//	fx.Annotated{
//	  Target: func(..) http.Handler { ... },
//	  Group:  "server",
//	}
func Annotate(t interface{}, anns ...Annotation) interface{} {
	result := annotated{Target: t}
	for _, ann := range anns {
		if err := ann.apply(&result); err != nil {
			return annotationError{
				target: t,
				err:    err,
			}
		}
	}
	result.Annotations = anns
	return result
}
```

## decorate.go

```go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/dig"
	"go.uber.org/fx/internal/fxreflect"
)

// Decorate specifies one or more decorator functions to an Fx application.
//
// # Decorator functions
//
// Decorator functions let users augment objects in the graph.
// They can take in zero or more dependencies that must be provided to the
// application with fx.Provide, and produce one or more values that can be used
// by other fx.Provide and fx.Invoke calls.
//
//	fx.Decorate(func(log *zap.Logger) *zap.Logger {
//	  return log.Named("myapp")
//	})
//	fx.Invoke(func(log *zap.Logger) {
//	  log.Info("hello")
//	  // Output:
//	  // {"level": "info","logger":"myapp","msg":"hello"}
//	})
//
// The following decorator accepts multiple dependencies from the graph,
// augments and returns one of them.
//
//	fx.Decorate(func(log *zap.Logger, cfg *Config) *zap.Logger {
//	  return log.Named(cfg.Name)
//	})
//
// Similar to fx.Provide, functions passed to fx.Decorate may optionally return
// an error as their last result.
// If a decorator returns a non-nil error, it will halt application startup.
//
//	fx.Decorate(func(conn *sql.DB, cfg *Config) (*sql.DB, error) {
//	  if err := conn.Ping(); err != nil {
//	    return sql.Open("driver-name", cfg.FallbackDB)
//	  }
//	  return conn, nil
//	})
//
// Decorators support both, fx.In and fx.Out structs, similar to fx.Provide and
// fx.Invoke.
//
//	type Params struct {
//	  fx.In
//
//	  Client usersvc.Client `name:"readOnly"`
//	}
//
//	type Result struct {
//	  fx.Out
//
//	  Client usersvc.Client `name:"readOnly"`
//	}
//
//	fx.Decorate(func(p Params) Result {
//	  ...
//	})
//
// Decorators can be annotated with the fx.Annotate function, but not with the
// fx.Annotated type. Refer to documentation on fx.Annotate() to learn how to
// use it for annotating functions.
//
//	fx.Decorate(
//	  fx.Annotate(
//	    func(client usersvc.Client) usersvc.Client {
//	      // ...
//	    },
//	    fx.ParamTags(`name:"readOnly"`),
//	    fx.ResultTags(`name:"readOnly"`),
//	  ),
//	)
//
// Decorators support augmenting, filtering, or replacing value groups.
// To decorate a value group, expect the entire value group slice and produce
// the new slice.
//
//	type HandlerParam struct {
//	  fx.In
//
//	  Log      *zap.Logger
//	  Handlers []Handler `group:"server"
//	}
//
//	type HandlerResult struct {
//	  fx.Out
//
//	  Handlers []Handler `group:"server"
//	}
//
//	fx.Decorate(func(p HandlerParam) HandlerResult {
//	  var r HandlerResult
//	  for _, handler := range p.Handlers {
//	    r.Handlers = append(r.Handlers, wrapWithLogger(p.Log, handler))
//	  }
//	  return r
//	}),
//
// # Decorator scope
//
// Modifications made to the Fx graph with fx.Decorate are scoped to the
// deepest fx.Module inside which the decorator was specified.
//
//	fx.Module("mymodule",
//	  fx.Decorate(func(log *zap.Logger) *zap.Logger {
//	    return log.Named("myapp")
//	  }),
//	  fx.Invoke(func(log *zap.Logger) {
//	    log.Info("decorated logger")
//	    // Output:
//	    // {"level": "info","logger":"myapp","msg":"decorated logger"}
//	  }),
//	),
//	fx.Invoke(func(log *zap.Logger) {
//	  log.Info("plain logger")
//	  // Output:
//	  // {"level": "info","msg":"plain logger"}
//	}),
//
// Decorations specified in the top-level fx.New call apply across the
// application and chain with module-specific decorators.
//
//	fx.New(
//	  // ...
//	  fx.Decorate(func(log *zap.Logger) *zap.Logger {
//	    return log.With(zap.Field("service", "myservice"))
//	  }),
//	  // ...
//	  fx.Invoke(func(log *zap.Logger) {
//	    log.Info("outer decorator")
//	    // Output:
//	    // {"level": "info","service":"myservice","msg":"outer decorator"}
//	  }),
//	  // ...
//	  fx.Module("mymodule",
//	    fx.Decorate(func(log *zap.Logger) *zap.Logger {
//	      return log.Named("myapp")
//	    }),
//	    fx.Invoke(func(log *zap.Logger) {
//	      log.Info("inner decorator")
//	      // Output:
//	      // {"level": "info","logger":"myapp","service":"myservice","msg":"inner decorator"}
//	    }),
//	  ),
//	)
func Decorate(decorators ...interface{}) Option {
	return decorateOption{
		Targets: decorators,
		Stack:   fxreflect.CallerStack(1, 0),
	}
}

type decorateOption struct {
	Targets []interface{}
	Stack   fxreflect.Stack
}

func (o decorateOption) apply(mod *module) {
	for _, target := range o.Targets {
		mod.decorators = append(mod.decorators, decorator{
			Target: target,
			Stack:  o.Stack,
		})
	}
}

func (o decorateOption) String() string {
	items := make([]string, len(o.Targets))
	for i, f := range o.Targets {
		items[i] = fxreflect.FuncName(f)
	}
	return fmt.Sprintf("fx.Decorate(%s)", strings.Join(items, ", "))
}

// decorator is a single decorator used in Fx.
type decorator struct {
	// Decorator provided to Fx.
	Target interface{}

	// Stack trace of where this provide was made.
	Stack fxreflect.Stack

	// Whether this decorator was specified via fx.Replace
	IsReplace   bool
	ReplaceType reflect.Type // set only if IsReplace
}

func runDecorator(c container, d decorator, opts ...dig.DecorateOption) (err error) {
	decorator := d.Target
	defer func() {
		if err != nil {
			err = fmt.Errorf("fx.Decorate(%v) from:\n%+vFailed: %v", decorator, d.Stack, err)
		}
	}()

	switch decorator := decorator.(type) {
	case annotated:
		if dcor, derr := decorator.Build(); derr == nil {
			err = c.Decorate(dcor, opts...)
		}
	default:
		err = c.Decorate(decorator, opts...)
	}
	return
}
```

## extract.go

```go

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"
	"reflect"
	"unicode"
	"unicode/utf8"
)

var _typeOfIn = reflect.TypeOf(In{})

// Extract fills the given struct with values from the dependency injection
// container on application initialization. The target MUST be a pointer to a
// struct. Only exported fields will be filled.
//
// Deprecated: Use Populate instead.
func Extract(target interface{}) Option {
	v := reflect.ValueOf(target)

	if t := v.Type(); t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return Error(fmt.Errorf("Extract expected a pointer to a struct, got a %v", t))
	}

	v = v.Elem()
	t := v.Type()

	// We generate a function which accepts a single fx.In struct as an
	// argument. This struct contains all exported fields of the target
	// struct.

	// Fields of the generated fx.In struct.
	fields := make([]reflect.StructField, 0, t.NumField()+1)

	// Anonymous dig.In field.
	fields = append(fields, reflect.StructField{
		Name:      _typeOfIn.Name(),
		Anonymous: true,
		Type:      _typeOfIn,
	})

	// List of values in the target struct aligned with the fields of the
	// generated struct.
	//
	// So for example, if the target is,
	//
	// 	var target struct {
	// 		Foo io.Reader
	// 		bar []byte
	// 		Baz io.Writer
	// 	}
	//
	// The generated struct has the shape,
	//
	// 	struct {
	// 		fx.In
	//
	// 		F0 io.Reader
	// 		F2 io.Writer
	// 	}
	//
	// And `targets` is,
	//
	// 	[
	// 		target.Field(0),  // Foo io.Reader
	// 		target.Field(2),  // Baz io.Writer
	// 	]
	//
	// As we iterate through the fields of the generated struct, we can copy
	// the value into the corresponding value in the targets list.
	targets := make([]reflect.Value, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		// Skip unexported fields.
		if f.Anonymous {
			// If embedded, StructField.PkgPath is not a reliable indicator of
			// whether the field is exported. See
			// https://github.com/golang/go/issues/21122

			t := f.Type
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}

			if !isExported(t.Name()) {
				continue
			}
		} else if f.PkgPath != "" {
			continue
		}

		// We don't copy over names or embedded semantics.
		fields = append(fields, reflect.StructField{
			Name: fmt.Sprintf("F%d", i),
			Type: f.Type,
			Tag:  f.Tag,
		})
		targets = append(targets, v.Field(i))
	}

	// Equivalent to,
	//
	// 	func(r struct {
	// 		fx.In
	//
	// 		F1 Foo
	// 		F2 Bar
	// 	}) {
	// 		target.Foo = r.F1
	// 		target.Bar = r.F2
	// 	}

	fn := reflect.MakeFunc(
		reflect.FuncOf(
			[]reflect.Type{reflect.StructOf(fields)},
			nil,   /* results */
			false, /* variadic */
		),
		func(args []reflect.Value) []reflect.Value {
			result := args[0]
			for i := 1; i < result.NumField(); i++ {
				targets[i-1].Set(result.Field(i))
			}
			return nil
		},
	)

	return Invoke(fn.Interface())
}

// isExported reports whether the identifier is exported.
func isExported(id string) bool {
	r, _ := utf8.DecodeRuneInString(id)
	return unicode.IsUpper(r)
}
```

## inout.go

```go
 
// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import "go.uber.org/dig"

// In can be embedded in a constructor's parameter struct to take advantage of
// advanced dependency injection features.
//
// Modules should take a single parameter struct that embeds an In in order to
// provide a forward-compatible API: since adding fields to a struct is
// backward-compatible, modules can then add optional dependencies in minor
// releases.
//
// # Parameter Structs
//
// Fx constructors declare their dependencies as function parameters. This can
// quickly become unreadable if the constructor has a lot of dependencies.
//
//	func NewHandler(users *UserGateway, comments *CommentGateway, posts *PostGateway, votes *VoteGateway, authz *AuthZGateway) *Handler {
//	  // ...
//	}
//
// To improve the readability of constructors like this, create a struct that
// lists all the dependencies as fields and change the function to accept that
// struct instead. The new struct is called a parameter struct.
//
// Fx has first class support for parameter structs: any struct embedding
// fx.In gets treated as a parameter struct, so the individual fields in the
// struct are supplied via dependency injection. Using a parameter struct, we
// can make the constructor above much more readable:
//
//	type HandlerParams struct {
//	  fx.In
//
//	  Users    *UserGateway
//	  Comments *CommentGateway
//	  Posts    *PostGateway
//	  Votes    *VoteGateway
//	  AuthZ    *AuthZGateway
//	}
//
//	func NewHandler(p HandlerParams) *Handler {
//	  // ...
//	}
//
// Though it's rarely a good idea, constructors can receive any combination of
// parameter structs and parameters.
//
//	func NewHandler(p HandlerParams, l *log.Logger) *Handler {
//	  // ...
//	}
//
// # Optional Dependencies
//
// Constructors often have optional dependencies on some types: if those types are
// missing, they can operate in a degraded state. Fx supports optional
// dependencies via the `optional:"true"` tag to fields on parameter structs.
//
//	type UserGatewayParams struct {
//	  fx.In
//
//	  Conn  *sql.DB
//	  Cache *redis.Client `optional:"true"`
//	}
//
// If an optional field isn't available in the container, the constructor
// receives the field's zero value.
//
//	func NewUserGateway(p UserGatewayParams, log *log.Logger) (*UserGateway, error) {
//	  if p.Cache == nil {
//	    log.Print("Caching disabled")
//	  }
//	  // ...
//	}
//
// Constructors that declare optional dependencies MUST gracefully handle
// situations in which those dependencies are absent.
//
// The optional tag also allows adding new dependencies without breaking
// existing consumers of the constructor.
//
// # Named Values
//
// Some use cases require the application container to hold multiple values of
// the same type. For details on producing named values, see the documentation
// for the Out type.
//
// Fx allows functions to consume named values via the `name:".."` tag on
// parameter structs. Note that both the name AND type of the fields on the
// parameter struct must match the corresponding result struct.
//
//	type GatewayParams struct {
//	  fx.In
//
//	  WriteToConn  *sql.DB `name:"rw"`
//	  ReadFromConn *sql.DB `name:"ro"`
//	}
//
// The name tag may be combined with the optional tag to declare the
// dependency optional.
//
//	type GatewayParams struct {
//	  fx.In
//
//	  WriteToConn  *sql.DB `name:"rw"`
//	  ReadFromConn *sql.DB `name:"ro" optional:"true"`
//	}
//
//	func NewCommentGateway(p GatewayParams, log *log.Logger) (*CommentGateway, error) {
//	  if p.ReadFromConn == nil {
//	    log.Print("Warning: Using RW connection for reads")
//	    p.ReadFromConn = p.WriteToConn
//	  }
//	  // ...
//	}
//
// # Value Groups
//
// To make it easier to produce and consume many values of the same type, Fx
// supports named, unordered collections called value groups. For details on
// producing value groups, see the documentation for the Out type.
//
// Functions can depend on a value group by requesting a slice tagged with
// `group:".."`. This will execute all constructors that provide a value to
// that group in an unspecified order, then collect all the results into a
// single slice. Keep in mind that this makes the types of the parameter and
// result struct fields different: if a group of constructors each returns
// type T, parameter structs consuming the group must use a field of type []T.
//
//	type ServerParams struct {
//	  fx.In
//
//	  Handlers []Handler `group:"server"`
//	}
//
//	func NewServer(p ServerParams) *Server {
//	  server := newServer()
//	  for _, h := range p.Handlers {
//	    server.Register(h)
//	  }
//	  return server
//	}
//
// Note that values in a value group are unordered. Fx makes no guarantees
// about the order in which these values will be produced.
//
// # Soft Value Groups
//
// A soft value group can be thought of as a best-attempt at populating the
// group with values from constructors that have already run. In other words,
// if a constructor's output type is only consumed by a soft value group,
// it will not be run.
//
// Note that Fx does not guarantee precise execution order of constructors
// or invokers, which means that the change in code that affects execution
// ordering of other constructors or functions will affect the values
// populated in this group.
//
// To declare a soft relationship between a group and its constructors, use
// the `soft` option on the group tag (`group:"[groupname],soft"`).
// This option is only valid for input parameters.
//
//	type Params struct {
//	  fx.In
//
//	  Handlers []Handler `group:"server,soft"`
//	  Logger   *zap.Logger
//	}
//
//	NewHandlerAndLogger := func() (Handler, *zap.Logger) { ... }
//	NewHandler := func() Handler { ... }
//	Foo := func(Params) { ... }
//
//	app := fx.New(
//	  fx.Provide(fx.Annotate(NewHandlerAndLogger, fx.ResultTags(`group:"server"`))),
//	  fx.Provide(fx.Annotate(NewHandler, fx.ResultTags(`group::"server"`))),
//	  fx.Invoke(Foo),
//	)
//
// The only constructor called is `NewHandlerAndLogger`, because this also provides
// `*zap.Logger` needed in the `Params` struct received by `Foo`. The Handlers
// group will be populated with a single Handler returned by `NewHandlerAndLogger`.
//
// In the next example, the slice `s` isn't populated as the provider would be
// called only because `strings` soft group value is its only consumer.
//
//	 app := fx.New(
//	   fx.Provide(
//	     fx.Annotate(
//	       func() (string, int) { return "hello", 42 },
//	       fx.ResultTags(`group:"strings"`),
//	     ),
//	   ),
//	   fx.Invoke(
//	     fx.Annotate(func(s []string) {
//	       // s will be an empty slice
//	     }, fx.ParamTags(`group:"strings,soft"`)),
//	   ),
//	)
//
// In the next example, the slice `s` will be populated because there is a
// consumer for the same type which is not a `soft` dependency.
//
//	 app := fx.New(
//	   fx.Provide(
//	     fx.Annotate(
//	       func() string { "hello" },
//	       fx.ResultTags(`group:"strings"`),
//	     ),
//	   ),
//	   fx.Invoke(
//	     fx.Annotate(func(b []string) {
//	       // b is []string{"hello"}
//	     }, fx.ParamTags(`group:"strings"`)),
//	   ),
//	   fx.Invoke(
//	     fx.Annotate(func(s []string) {
//	       // s is []string{"hello"}
//	     }, fx.ParamTags(`group:"strings,soft"`)),
//	   ),
//	)
//
// # Unexported fields
//
// By default, a type that embeds fx.In may not have any unexported fields. The
// following will return an error if used with Fx.
//
//	type Params struct {
//	  fx.In
//
//	  Logger *zap.Logger
//	  mu     sync.Mutex
//	}
//
// If you have need of unexported fields on such a type, you may opt-into
// ignoring unexported fields by adding the ignore-unexported struct tag to the
// fx.In. For example,
//
//	type Params struct {
//	  fx.In `ignore-unexported:"true"`
//
//	  Logger *zap.Logger
//	  mu     sync.Mutex
//	}
type In = dig.In

// Out is the inverse of In: it can be embedded in result structs to take
// advantage of advanced features.
//
// Modules should return a single result struct that embeds an Out in order to
// provide a forward-compatible API: since adding fields to a struct is
// backward-compatible, minor releases can provide additional types.
//
// # Result Structs
//
// Result structs are the inverse of parameter structs (discussed in the In
// documentation). These structs represent multiple outputs from a
// single function as fields. Fx treats all structs embedding fx.Out as result
// structs, so other constructors can rely on the result struct's fields
// directly.
//
// Without result structs, we sometimes have function definitions like this:
//
//	func SetupGateways(conn *sql.DB) (*UserGateway, *CommentGateway, *PostGateway, error) {
//	  // ...
//	}
//
// With result structs, we can make this both more readable and easier to
// modify in the future:
//
//	type Gateways struct {
//	  fx.Out
//
//	  Users    *UserGateway
//	  Comments *CommentGateway
//	  Posts    *PostGateway
//	}
//
//	func SetupGateways(conn *sql.DB) (Gateways, error) {
//	  // ...
//	}
//
// # Named Values
//
// Some use cases require the application container to hold multiple values of
// the same type. For details on consuming named values, see the documentation
// for the In type.
//
// A constructor that produces a result struct can tag any field with
// `name:".."` to have the corresponding value added to the graph under the
// specified name. An application may contain at most one unnamed value of a
// given type, but may contain any number of named values of the same type.
//
//	type ConnectionResult struct {
//	  fx.Out
//
//	  ReadWrite *sql.DB `name:"rw"`
//	  ReadOnly  *sql.DB `name:"ro"`
//	}
//
//	func ConnectToDatabase(...) (ConnectionResult, error) {
//	  // ...
//	  return ConnectionResult{ReadWrite: rw, ReadOnly:  ro}, nil
//	}
//
// # Value Groups
//
// To make it easier to produce and consume many values of the same type, Fx
// supports named, unordered collections called value groups. For details on
// consuming value groups, see the documentation for the In type.
//
// Constructors can send values into value groups by returning a result struct
// tagged with `group:".."`.
//
//	type HandlerResult struct {
//	  fx.Out
//
//	  Handler Handler `group:"server"`
//	}
//
//	func NewHelloHandler() HandlerResult {
//	  // ...
//	}
//
//	func NewEchoHandler() HandlerResult {
//	  // ...
//	}
//
// Any number of constructors may provide values to this named collection, but
// the ordering of the final collection is unspecified. Keep in mind that
// value groups require parameter and result structs to use fields with
// different types: if a group of constructors each returns type T, parameter
// structs consuming the group must use a field of type []T.
//
// To provide multiple values for a group from a result struct, produce a
// slice and use the `,flatten` option on the group tag. This indicates that
// each element in the slice should be injected into the group individually.
//
//	type IntResult struct {
//	  fx.Out
//
//	  Handler []int `group:"server"`         // Consume as [][]int
//	  Handler []int `group:"server,flatten"` // Consume as []int
//	}
type Out = dig.Out
```

## invoke.go

```go

// Copyright (c) 2019-2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"
	"strings"

	"go.uber.org/fx/internal/fxreflect"
)

// Invoke registers functions that are executed eagerly on application start.
// Arguments for these invocations are built using the constructors registered
// by Provide. Passing multiple Invoke options appends the new invocations to
// the application's existing list.
//
// Unlike constructors, invocations are always executed, and they're always
// run in order. Invocations may have any number of returned values.
// If the final returned object is an error, it indicates whether the operation
// was successful.
// All other returned values are discarded.
//
// Invokes registered in [Module]s are run before the ones registered at the
// scope of the parent. Invokes within the same Module is run in the order
// they were provided. For example,
//
//	fx.New(
//		fx.Invoke(func3),
//		fx.Module("someModule",
//			fx.Invoke(func1),
//			fx.Invoke(func2),
//		),
//		fx.Invoke(func4),
//	)
//
// invokes func1, func2, func3, func4 in that order.
//
// Typically, invoked functions take a handful of high-level objects (whose
// constructors depend on lower-level objects) and introduce them to each
// other. This kick-starts the application by forcing it to instantiate a
// variety of types.
//
// To see an invocation in use, read through the package-level example. For
// advanced features, including optional parameters and named instances, see
// the documentation of the In and Out types.
func Invoke(funcs ...interface{}) Option {
	return invokeOption{
		Targets: funcs,
		Stack:   fxreflect.CallerStack(1, 0),
	}
}

type invokeOption struct {
	Targets []interface{}
	Stack   fxreflect.Stack
}

func (o invokeOption) apply(mod *module) {
	for _, target := range o.Targets {
		mod.invokes = append(mod.invokes, invoke{
			Target: target,
			Stack:  o.Stack,
		})
	}
}

func (o invokeOption) String() string {
	items := make([]string, len(o.Targets))
	for i, f := range o.Targets {
		items[i] = fxreflect.FuncName(f)
	}
	return fmt.Sprintf("fx.Invoke(%s)", strings.Join(items, ", "))
}
func runInvoke(c container, i invoke) error {
	fn := i.Target
	switch fn := fn.(type) {
	case Option:
		return fmt.Errorf("fx.Option should be passed to fx.New directly, "+
			"not to fx.Invoke: fx.Invoke received %v from:\n%+v",
			fn, i.Stack)

	case annotated:
		af, err := fn.Build()
		if err != nil {
			return err
		}

		return c.Invoke(af)
	default:
		return c.Invoke(fn)
	}
}
```

## log.go

```go

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"go.uber.org/fx/fxevent"
)

// logBuffer will buffer all messages until a logger has been
// initialized.
type logBuffer struct {
	events []fxevent.Event
	logger fxevent.Logger
}

// LogEvent buffers or logs an event.
func (l *logBuffer) LogEvent(event fxevent.Event) {
	if l.logger == nil {
		l.events = append(l.events, event)
	} else {
		l.logger.LogEvent(event)
	}
}

// Connect flushes out all buffered events to a logger and resets them.
func (l *logBuffer) Connect(logger fxevent.Logger) {
	l.logger = logger
	for _, e := range l.events {
		logger.LogEvent(e)
	}
	l.events = nil
}
```

## populate.go

```go

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"
	"reflect"
)

// Populate sets targets with values from the dependency injection container
// during application initialization. All targets must be pointers to the
// values that must be populated. Pointers to structs that embed In are
// supported, which can be used to populate multiple values in a struct.
//
// Annotating each pointer with ParamTags is also supported as a shorthand
// to passing a pointer to a struct that embeds In with field tags. For example:
//
//	 var a A
//	 var b B
//	 fx.Populate(
//		fx.Annotate(
//				&a,
//				fx.ParamTags(`name:"A"`)
//	 	),
//		fx.Annotate(
//				&b,
//				fx.ParamTags(`name:"B"`)
//	 	)
//	 )
//
// Code above is equivalent to the following:
//
//	type Target struct {
//		fx.In
//
//		a A `name:"A"`
//		b B `name:"B"`
//	}
//	var target Target
//	...
//	fx.Populate(&target)
//
// This is most helpful in unit tests: it lets tests leverage Fx's automatic
// constructor wiring to build a few structs, but then extract those structs
// for further testing.
func Populate(targets ...interface{}) Option {
	// Validate all targets are non-nil pointers.
	fields := make([]reflect.StructField, len(targets)+1)
	fields[0] = reflect.StructField{
		Name:      "In",
		Type:      reflect.TypeOf(In{}),
		Anonymous: true,
	}
	for i, t := range targets {
		if t == nil {
			return Error(fmt.Errorf("failed to Populate: target %v is nil", i+1))
		}
		var (
			rt  reflect.Type
			tag reflect.StructTag
		)
		switch t := t.(type) {
		case annotated:
			rt = reflect.TypeOf(t.Target)
			tag = reflect.StructTag(t.ParamTags[0])
			targets[i] = t.Target
		default:
			rt = reflect.TypeOf(t)
		}
		if rt.Kind() != reflect.Ptr {
			return Error(fmt.Errorf("failed to Populate: target %v is not a pointer type, got %T", i+1, t))
		}
		fields[i+1] = reflect.StructField{
			Name: fmt.Sprintf("Field%d", i),
			Type: rt.Elem(),
			Tag:  tag,
		}
	}

	// Build a function that looks like:
	//
	// func(t1 T1, t2 T2, ...) {
	//   *targets[0] = t1
	//   *targets[1] = t2
	//   [...]
	// }
	//
	fnType := reflect.FuncOf([]reflect.Type{reflect.StructOf(fields)}, nil, false /* variadic */)
	fn := reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		arg := args[0]
		for i, target := range targets {
			reflect.ValueOf(target).Elem().Set(arg.Field(i + 1))
		}
		return nil
	})
	return Invoke(fn.Interface())
}
```

## provide.go

```go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/dig"
	"go.uber.org/fx/internal/fxreflect"
)

// Provide registers any number of constructor functions, teaching the
// application how to instantiate various types. The supplied constructor
// function(s) may depend on other types available in the application, must
// return one or more objects, and may return an error. For example:
//
//	// Constructs type *C, depends on *A and *B.
//	func(*A, *B) *C
//
//	// Constructs type *C, depends on *A and *B, and indicates failure by
//	// returning an error.
//	func(*A, *B) (*C, error)
//
//	// Constructs types *B and *C, depends on *A, and can fail.
//	func(*A) (*B, *C, error)
//
// The order in which constructors are provided doesn't matter, and passing
// multiple Provide options appends to the application's collection of
// constructors. Constructors are called only if one or more of their returned
// types are needed, and their results are cached for reuse (so instances of a
// type are effectively singletons within an application). Taken together,
// these properties make it perfectly reasonable to Provide a large number of
// constructors even if only a fraction of them are used.
//
// See the documentation of the In and Out types for advanced features,
// including optional parameters and named instances.
//
// See the documentation for [Private] for restricting access to constructors.
//
// Constructor functions should perform as little external interaction as
// possible, and should avoid spawning goroutines. Things like server listen
// loops, background timer loops, and background processing goroutines should
// instead be managed using Lifecycle callbacks.
func Provide(constructors ...interface{}) Option {
	return provideOption{
		Targets: constructors,
		Stack:   fxreflect.CallerStack(1, 0),
	}
}

type provideOption struct {
	Targets []interface{}
	Stack   fxreflect.Stack
}

func (o provideOption) apply(mod *module) {
	var private bool

	targets := make([]interface{}, 0, len(o.Targets))
	for _, target := range o.Targets {
		if _, ok := target.(privateOption); ok {
			private = true
			continue
		}
		targets = append(targets, target)
	}

	for _, target := range targets {
		mod.provides = append(mod.provides, provide{
			Target:  target,
			Stack:   o.Stack,
			Private: private,
		})
	}
}

type privateOption struct{}

// Private is an option that can be passed as an argument to [Provide] to
// restrict access to the constructors being provided. Specifically,
// corresponding constructors can only be used within the current module
// or modules the current module contains. Other modules that contain this
// module won't be able to use the constructor.
//
// For example, the following would fail because the app doesn't have access
// to the inner module's constructor.
//
//	fx.New(
//		fx.Module("SubModule", fx.Provide(func() int { return 0 }, fx.Private)),
//		fx.Invoke(func(a int) {}),
//	)
var Private = privateOption{}

func (o provideOption) String() string {
	items := make([]string, len(o.Targets))
	for i, c := range o.Targets {
		items[i] = fxreflect.FuncName(c)
	}
	return fmt.Sprintf("fx.Provide(%s)", strings.Join(items, ", "))
}

func runProvide(c container, p provide, opts ...dig.ProvideOption) error {
	constructor := p.Target
	if _, ok := constructor.(Option); ok {
		return fmt.Errorf("fx.Option should be passed to fx.New directly, "+
			"not to fx.Provide: fx.Provide received %v from:\n%+v",
			constructor, p.Stack)
	}

	switch constructor := constructor.(type) {
	case annotationError:
		// fx.Annotate failed. Turn it into an Fx error.
		return fmt.Errorf(
			"encountered error while applying annotation using fx.Annotate to %s: %w",
			fxreflect.FuncName(constructor.target), constructor.err)

	case annotated:
		ctor, err := constructor.Build()
		if err != nil {
			return fmt.Errorf("fx.Provide(%v) from:\n%+vFailed: %w", constructor, p.Stack, err)
		}

		opts = append(opts, dig.LocationForPC(constructor.FuncPtr))
		if err := c.Provide(ctor, opts...); err != nil {
			return fmt.Errorf("fx.Provide(%v) from:\n%+vFailed: %w", constructor, p.Stack, err)
		}

	case Annotated:
		ann := constructor
		switch {
		case len(ann.Group) > 0 && len(ann.Name) > 0:
			return fmt.Errorf(
				"fx.Annotated may specify only one of Name or Group: received %v from:\n%+v",
				ann, p.Stack)
		case len(ann.Name) > 0:
			opts = append(opts, dig.Name(ann.Name))
		case len(ann.Group) > 0:
			opts = append(opts, dig.Group(ann.Group))
		}

		if err := c.Provide(ann.Target, opts...); err != nil {
			return fmt.Errorf("fx.Provide(%v) from:\n%+vFailed: %w", ann, p.Stack, err)
		}

	default:
		if reflect.TypeOf(constructor).Kind() == reflect.Func {
			ft := reflect.ValueOf(constructor).Type()

			for i := 0; i < ft.NumOut(); i++ {
				t := ft.Out(i)

				if t == reflect.TypeOf(Annotated{}) {
					return fmt.Errorf(
						"fx.Annotated should be passed to fx.Provide directly, "+
							"it should not be returned by the constructor: "+
							"fx.Provide received %v from:\n%+v",
						fxreflect.FuncName(constructor), p.Stack)
				}
			}
		}

		if err := c.Provide(constructor, opts...); err != nil {
			return fmt.Errorf("fx.Provide(%v) from:\n%+vFailed: %w", fxreflect.FuncName(constructor), p.Stack, err)
		}
	}
	return nil
}
```

## replace.go

```go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/fx/internal/fxreflect"
)

// Replace provides instantiated values for graph modification as if
// they had been provided using a decorator with fx.Decorate.
// The most specific type of each value (as determined by reflection) is used.
//
// Refer to the documentation on fx.Decorate to see how graph modifications
// work with fx.Module.
//
// This serves a purpose similar to what fx.Supply does for fx.Provide.
//
// For example, given,
//
//	var log *zap.Logger = ...
//
// The following two forms are equivalent.
//
//	fx.Replace(log)
//
//	fx.Decorate(
//		func() *zap.Logger {
//			return log
//		},
//	)
//
// Replace panics if a value (or annotation target) is an untyped nil or an error.
//
// # Replace Caveats
//
// As mentioned above, Replace uses the most specific type of the provided
// value. For interface values, this refers to the type of the implementation,
// not the interface. So if you try to replace an io.Writer, fx.Replace will
// use the type of the implementation.
//
//	var stderr io.Writer = os.Stderr
//	fx.Replace(stderr)
//
// Is equivalent to,
//
//	fx.Decorate(func() *os.File { return os.Stderr })
//
// This is typically NOT what you intended. To replace the io.Writer in the
// container with the value above, we need to use the fx.Annotate function with
// the fx.As annotation.
//
//	fx.Replace(
//		fx.Annotate(os.Stderr, fx.As(new(io.Writer)))
//	)
func Replace(values ...interface{}) Option {
	decorators := make([]interface{}, len(values)) // one function per value
	types := make([]reflect.Type, len(values))
	for i, value := range values {
		switch value := value.(type) {
		case annotated:
			var typ reflect.Type
			value.Target, typ = newReplaceDecorator(value.Target)
			decorators[i] = value
			types[i] = typ
		default:
			decorators[i], types[i] = newReplaceDecorator(value)
		}
	}

	return replaceOption{
		Targets: decorators,
		Types:   types,
		Stack:   fxreflect.CallerStack(1, 0),
	}
}

type replaceOption struct {
	Targets []interface{}
	Types   []reflect.Type // type of value produced by constructor[i]
	Stack   fxreflect.Stack
}

func (o replaceOption) apply(m *module) {
	for i, target := range o.Targets {
		m.decorators = append(m.decorators, decorator{
			Target:      target,
			Stack:       o.Stack,
			IsReplace:   true,
			ReplaceType: o.Types[i],
		})
	}
}

func (o replaceOption) String() string {
	items := make([]string, 0, len(o.Targets))
	for _, typ := range o.Types {
		items = append(items, typ.String())
	}
	return fmt.Sprintf("fx.Replace(%s)", strings.Join(items, ", "))
}

// Returns a function that takes no parameters, and returns the given value.
func newReplaceDecorator(value interface{}) (interface{}, reflect.Type) {
	switch value.(type) {
	case nil:
		panic("untyped nil passed to fx.Replace")
	case error:
		panic("error value passed to fx.Replace")
	}

	typ := reflect.TypeOf(value)
	returnTypes := []reflect.Type{typ}
	returnValues := []reflect.Value{reflect.ValueOf(value)}

	ft := reflect.FuncOf([]reflect.Type{}, returnTypes, false)
	fv := reflect.MakeFunc(ft, func([]reflect.Value) []reflect.Value {
		return returnValues
	})

	return fv.Interface(), typ
}
```

## shutdown.go

```go

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"time"
)

// Shutdowner provides a method that can manually trigger the shutdown of the
// application by sending a signal to all open Done channels. Shutdowner works
// on applications using Run as well as Start, Done, and Stop. The Shutdowner is
// provided to all Fx applications.
type Shutdowner interface {
	Shutdown(...ShutdownOption) error
}

// ShutdownOption provides a way to configure properties of the shutdown
// process. Currently, no options have been implemented.
type ShutdownOption interface {
	apply(*shutdowner)
}

type exitCodeOption int

func (code exitCodeOption) apply(s *shutdowner) {
	s.exitCode = int(code)
}

var _ ShutdownOption = exitCodeOption(0)

// ExitCode is a [ShutdownOption] that may be passed to the Shutdown method of the
// [Shutdowner] interface.
// The given integer exit code will be broadcasted to any receiver waiting
// on a [ShutdownSignal] from the [Wait] method.
func ExitCode(code int) ShutdownOption {
	return exitCodeOption(code)
}

type shutdownTimeoutOption time.Duration

func (shutdownTimeoutOption) apply(*shutdowner) {}

var _ ShutdownOption = shutdownTimeoutOption(0)

// ShutdownTimeout is a [ShutdownOption] that allows users to specify a timeout
// for a given call to Shutdown method of the [Shutdowner] interface. As the
// Shutdown method will block while waiting for a signal receiver relay
// goroutine to stop.
//
// Deprecated: This option has no effect. Shutdown is not a blocking operation.
func ShutdownTimeout(timeout time.Duration) ShutdownOption {
	return shutdownTimeoutOption(timeout)
}

type shutdowner struct {
	app      *App
	exitCode int
}

// Shutdown broadcasts a signal to all of the application's Done channels
// and begins the Stop process. Applications can be shut down only after they
// have finished starting up.
func (s *shutdowner) Shutdown(opts ...ShutdownOption) error {
	for _, opt := range opts {
		opt.apply(s)
	}

	return s.app.receivers.Broadcast(ShutdownSignal{
		Signal:   _sigTERM,
		ExitCode: s.exitCode,
	})
}

func (app *App) shutdowner() Shutdowner {
	return &shutdowner{app: app}
}
```

## signal.go

```go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPSignalE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// ShutdownSignal represents a signal to be written to Wait or Done.
// Should a user call the Shutdown method via the Shutdowner interface with
// a provided ExitCode, that exit code will be populated in the ExitCode field.
//
// Should the application receive an operating system signal,
// the Signal field will be populated with the received os.Signal.
type ShutdownSignal struct {
	Signal   os.Signal
	ExitCode int
}

// String will render a ShutdownSignal type as a string suitable for printing.
func (sig ShutdownSignal) String() string {
	return fmt.Sprintf("%v", sig.Signal)
}

func newSignalReceivers() signalReceivers {
	return signalReceivers{
		notify:  signal.Notify,
		signals: make(chan os.Signal, 1),
	}
}

type signalReceivers struct {
	// this mutex protects writes and reads of this struct to prevent
	// race conditions in a parallel execution pattern
	m sync.Mutex

	// our os.Signal channel we relay from
	signals chan os.Signal
	// when written to, will instruct the signal relayer to shutdown
	shutdown chan struct{}
	// is written to when signal relay has finished shutting down
	finished chan struct{}

	// this stub allows us to unit test signal relay functionality
	notify func(c chan<- os.Signal, sig ...os.Signal)

	// last will contain a pointer to the last ShutdownSignal received, or
	// nil if none, if a new channel is created by Wait or Done, this last
	// signal will be immediately written to, this allows Wait or Done state
	// to be read after application stop
	last *ShutdownSignal

	// contains channels created by Done
	done []chan os.Signal

	// contains channels created by Wait
	wait []chan ShutdownSignal
}

func (recv *signalReceivers) relayer(ctx context.Context) {
	defer func() {
		recv.finished <- struct{}{}
	}()

	select {
	case <-recv.shutdown:
		return
	case signal := <-recv.signals:
		recv.Broadcast(ShutdownSignal{
			Signal: signal,
		})
	}
}

// running returns true if the the signal relay go-routine is running.
// this method must be invoked under locked mutex to avoid race condition.
func (recv *signalReceivers) running() bool {
	return recv.shutdown != nil && recv.finished != nil
}

func (recv *signalReceivers) Start(ctx context.Context) {
	recv.m.Lock()
	defer recv.m.Unlock()

	// if the receiver has already been started; don't start it again
	if recv.running() {
		return
	}

	recv.finished = make(chan struct{}, 1)
	recv.shutdown = make(chan struct{}, 1)
	recv.notify(recv.signals, os.Interrupt, _sigINT, _sigTERM)
	go recv.relayer(ctx)
}

func (recv *signalReceivers) Stop(ctx context.Context) error {
	recv.m.Lock()
	defer recv.m.Unlock()

	// if the relayer is not running; return nil error
	if !recv.running() {
		return nil
	}

	recv.shutdown <- struct{}{}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-recv.finished:
		close(recv.shutdown)
		close(recv.finished)
		recv.shutdown = nil
		recv.finished = nil
		recv.last = nil
		return nil
	}
}

func (recv *signalReceivers) Done() <-chan os.Signal {
	recv.m.Lock()
	defer recv.m.Unlock()

	ch := make(chan os.Signal, 1)

	// If we had received a signal prior to the call of done, send it's
	// os.Signal to the new channel.
	// However we still want to have the operating system notify signals to this
	// channel should the application receive another.
	if recv.last != nil {
		ch <- recv.last.Signal
	}

	recv.done = append(recv.done, ch)
	return ch
}

func (recv *signalReceivers) Wait() <-chan ShutdownSignal {
	recv.m.Lock()
	defer recv.m.Unlock()

	ch := make(chan ShutdownSignal, 1)

	if recv.last != nil {
		ch <- *recv.last
	}

	recv.wait = append(recv.wait, ch)
	return ch
}

func (recv *signalReceivers) Broadcast(signal ShutdownSignal) error {
	recv.m.Lock()
	defer recv.m.Unlock()

	recv.last = &signal

	channels, unsent := recv.broadcast(
		signal,
		recv.broadcastDone,
		recv.broadcastWait,
	)

	if unsent != 0 {
		return &unsentSignalError{
			Signal: signal,
			Total:  channels,
			Unsent: unsent,
		}
	}

	return nil
}

func (recv *signalReceivers) broadcast(
	signal ShutdownSignal,
	anchors ...func(ShutdownSignal) (int, int),
) (int, int) {
	var channels, unsent int

	for _, anchor := range anchors {
		c, u := anchor(signal)
		channels += c
		unsent += u
	}

	return channels, unsent
}

func (recv *signalReceivers) broadcastDone(signal ShutdownSignal) (int, int) {
	var unsent int

	for _, reader := range recv.done {
		select {
		case reader <- signal.Signal:
		default:
			unsent++
		}
	}

	return len(recv.done), unsent
}

func (recv *signalReceivers) broadcastWait(signal ShutdownSignal) (int, int) {
	var unsent int

	for _, reader := range recv.wait {
		select {
		case reader <- signal:
		default:
			unsent++
		}
	}

	return len(recv.wait), unsent
}

type unsentSignalError struct {
	Signal ShutdownSignal
	Unsent int
	Total  int
}

func (err *unsentSignalError) Error() string {
	return fmt.Sprintf(
		"send %v signal: %v/%v channels are blocked",
		err.Signal,
		err.Unsent,
		err.Total,
	)
}
```

## supply.go

```go

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import (
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/fx/internal/fxreflect"
)

// Supply provides instantiated values for dependency injection as if
// they had been provided using a constructor that simply returns them.
// The most specific type of each value (as determined by reflection) is used.
//
// This serves a purpose similar to what fx.Replace does for fx.Decorate.
//
// For example, given:
//
//	type (
//		TypeA struct{}
//		TypeB struct{}
//		TypeC struct{}
//	)
//
//	var a, b, c = &TypeA{}, TypeB{}, &TypeC{}
//
// The following two forms are equivalent:
//
//	fx.Supply(a, b, fx.Annotated{Target: c})
//
//	fx.Provide(
//		func() *TypeA { return a },
//		func() TypeB { return b },
//		fx.Annotated{Target: func() *TypeC { return c }},
//	)
//
// Supply panics if a value (or annotation target) is an untyped nil or an error.
//
// # Supply Caveats
//
// As mentioned above, Supply uses the most specific type of the provided
// value. For interface values, this refers to the type of the implementation,
// not the interface. So if you supply an http.Handler, fx.Supply will use the
// type of the implementation.
//
//	var handler http.Handler = http.HandlerFunc(f)
//	fx.Supply(handler)
//
// Is equivalent to,
//
//	fx.Provide(func() http.HandlerFunc { return f })
//
// This is typically NOT what you intended. To supply the handler above as an
// http.Handler, we need to use the fx.Annotate function with the fx.As
// annotation.
//
//	fx.Supply(
//		fx.Annotate(handler, fx.As(new(http.Handler))),
//	)
func Supply(values ...interface{}) Option {
	constructors := make([]interface{}, len(values)) // one function per value
	types := make([]reflect.Type, len(values))
	for i, value := range values {
		switch value := value.(type) {
		case annotated:
			var typ reflect.Type
			value.Target, typ = newSupplyConstructor(value.Target)
			constructors[i] = value
			types[i] = typ
		case Annotated:
			var typ reflect.Type
			value.Target, typ = newSupplyConstructor(value.Target)
			constructors[i] = value
			types[i] = typ
		default:
			constructors[i], types[i] = newSupplyConstructor(value)
		}
	}

	return supplyOption{
		Targets: constructors,
		Types:   types,
		Stack:   fxreflect.CallerStack(1, 0),
	}
}

type supplyOption struct {
	Targets []interface{}
	Types   []reflect.Type // type of value produced by constructor[i]
	Stack   fxreflect.Stack
}

func (o supplyOption) apply(m *module) {
	for i, target := range o.Targets {
		m.provides = append(m.provides, provide{
			Target:     target,
			Stack:      o.Stack,
			IsSupply:   true,
			SupplyType: o.Types[i],
		})
	}
}

func (o supplyOption) String() string {
	items := make([]string, 0, len(o.Targets))
	for _, typ := range o.Types {
		items = append(items, typ.String())
	}
	return fmt.Sprintf("fx.Supply(%s)", strings.Join(items, ", "))
}

// Returns a function that takes no parameters, and returns the given value.
func newSupplyConstructor(value interface{}) (interface{}, reflect.Type) {
	switch value.(type) {
	case nil:
		panic("untyped nil passed to fx.Supply")
	case error:
		panic("error value passed to fx.Supply")
	}

	typ := reflect.TypeOf(value)
	returnTypes := []reflect.Type{typ}
	returnValues := []reflect.Value{reflect.ValueOf(value)}

	ft := reflect.FuncOf([]reflect.Type{}, returnTypes, false)
	fv := reflect.MakeFunc(ft, func([]reflect.Value) []reflect.Value {
		return returnValues
	})

	return fv.Interface(), typ
}
```

## fxevent

### console.go

```go

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxevent

import (
	"fmt"
	"io"
	"strings"
)

// ConsoleLogger is an Fx event logger that attempts to write human-readable
// messages to the console.
//
// Use this during development.
type ConsoleLogger struct {
	W io.Writer
}

var _ Logger = (*ConsoleLogger)(nil)

func (l *ConsoleLogger) logf(msg string, args ...interface{}) {
	fmt.Fprintf(l.W, "[Fx] "+msg+"\n", args...)
}

// LogEvent logs the given event to the provided Zap logger.
func (l *ConsoleLogger) LogEvent(event Event) {
	switch e := event.(type) {
	case *OnStartExecuting:
		l.logf("HOOK OnStart\t\t%s executing (caller: %s)", e.FunctionName, e.CallerName)
	case *OnStartExecuted:
		if e.Err != nil {
			l.logf("HOOK OnStart\t\t%s called by %s failed in %s: %+v", e.FunctionName, e.CallerName, e.Runtime, e.Err)
		} else {
			l.logf("HOOK OnStart\t\t%s called by %s ran successfully in %s", e.FunctionName, e.CallerName, e.Runtime)
		}
	case *OnStopExecuting:
		l.logf("HOOK OnStop\t\t%s executing (caller: %s)", e.FunctionName, e.CallerName)
	case *OnStopExecuted:
		if e.Err != nil {
			l.logf("HOOK OnStop\t\t%s called by %s failed in %s: %+v", e.FunctionName, e.CallerName, e.Runtime, e.Err)
		} else {
			l.logf("HOOK OnStop\t\t%s called by %s ran successfully in %s", e.FunctionName, e.CallerName, e.Runtime)
		}
	case *Supplied:
		if e.Err != nil {
			l.logf("ERROR\tFailed to supply %v: %+v", e.TypeName, e.Err)
		} else if e.ModuleName != "" {
			l.logf("SUPPLY\t%v from module %q", e.TypeName, e.ModuleName)
		} else {
			l.logf("SUPPLY\t%v", e.TypeName)
		}
	case *Provided:
		var privateStr string
		if e.Private {
			privateStr = " (PRIVATE)"
		}
		for _, rtype := range e.OutputTypeNames {
			if e.ModuleName != "" {
				l.logf("PROVIDE%v\t%v <= %v from module %q", privateStr, rtype, e.ConstructorName, e.ModuleName)
			} else {
				l.logf("PROVIDE%v\t%v <= %v", privateStr, rtype, e.ConstructorName)
			}
		}
		if e.Err != nil {
			l.logf("Error after options were applied: %+v", e.Err)
		}
	case *Replaced:
		for _, rtype := range e.OutputTypeNames {
			if e.ModuleName != "" {
				l.logf("REPLACE\t%v from module %q", rtype, e.ModuleName)
			} else {
				l.logf("REPLACE\t%v", rtype)
			}
		}
		if e.Err != nil {
			l.logf("ERROR\tFailed to replace: %+v", e.Err)
		}
	case *Decorated:
		for _, rtype := range e.OutputTypeNames {
			if e.ModuleName != "" {
				l.logf("DECORATE\t%v <= %v from module %q", rtype, e.DecoratorName, e.ModuleName)
			} else {
				l.logf("DECORATE\t%v <= %v", rtype, e.DecoratorName)
			}
		}
		if e.Err != nil {
			l.logf("Error after options were applied: %+v", e.Err)
		}
	case *Run:
		var moduleStr string
		if e.ModuleName != "" {
			moduleStr = fmt.Sprintf(" from module %q", e.ModuleName)
		}
		l.logf("RUN\t%v: %v%v", e.Kind, e.Name, moduleStr)
		if e.Err != nil {
			l.logf("Error returned: %+v", e.Err)
		}

	case *Invoking:
		if e.ModuleName != "" {
			l.logf("INVOKE\t\t%s from module %q", e.FunctionName, e.ModuleName)
		} else {
			l.logf("INVOKE\t\t%s", e.FunctionName)
		}
	case *Invoked:
		if e.Err != nil {
			l.logf("ERROR\t\tfx.Invoke(%v) called from:\n%+vFailed: %+v", e.FunctionName, e.Trace, e.Err)
		}
	case *Stopping:
		l.logf("%v", strings.ToUpper(e.Signal.String()))
	case *Stopped:
		if e.Err != nil {
			l.logf("ERROR\t\tFailed to stop cleanly: %+v", e.Err)
		}
	case *RollingBack:
		l.logf("ERROR\t\tStart failed, rolling back: %+v", e.StartErr)
	case *RolledBack:
		if e.Err != nil {
			l.logf("ERROR\t\tCouldn't roll back cleanly: %+v", e.Err)
		}
	case *Started:
		if e.Err != nil {
			l.logf("ERROR\t\tFailed to start: %+v", e.Err)
		} else {
			l.logf("RUNNING")
		}
	case *LoggerInitialized:
		if e.Err != nil {
			l.logf("ERROR\t\tFailed to initialize custom logger: %+v", e.Err)
		} else {
			l.logf("LOGGER\tInitialized custom logger from %v", e.ConstructorName)
		}
	}
}
```

### event.go

```go

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxevent

import (
	"os"
	"time"
)

// Event defines an event emitted by fx.
type Event interface {
	event() // Only fxlog can implement this interface.
}

// Passing events by type to make Event hashable in the future.
func (*OnStartExecuting) event()  {}
func (*OnStartExecuted) event()   {}
func (*OnStopExecuting) event()   {}
func (*OnStopExecuted) event()    {}
func (*Supplied) event()          {}
func (*Provided) event()          {}
func (*Replaced) event()          {}
func (*Decorated) event()         {}
func (*Run) event()               {}
func (*Invoking) event()          {}
func (*Invoked) event()           {}
func (*Stopping) event()          {}
func (*Stopped) event()           {}
func (*RollingBack) event()       {}
func (*RolledBack) event()        {}
func (*Started) event()           {}
func (*LoggerInitialized) event() {}

// OnStartExecuting is emitted before an OnStart hook is executed.
type OnStartExecuting struct {
	// FunctionName is the name of the function that will be executed.
	FunctionName string

	// CallerName is the name of the function that scheduled the hook for
	// execution.
	CallerName string
}

// OnStartExecuted is emitted after an OnStart hook has been executed.
type OnStartExecuted struct {
	// FunctionName is the name of the function that was executed.
	FunctionName string

	// CallerName is the name of the function that scheduled the hook for
	// execution.
	CallerName string

	// Method specifies the kind of the hook. This is one of "OnStart" and
	// "OnStop".
	Method string

	// Runtime specifies how long it took to run this hook.
	Runtime time.Duration

	// Err is non-nil if the hook failed to execute.
	Err error
}

// OnStopExecuting is emitted before an OnStop hook is executed.
type OnStopExecuting struct {
	// FunctionName is the name of the function that will be executed.
	FunctionName string

	// CallerName is the name of the function that scheduled the hook for
	// execution.
	CallerName string
}

// OnStopExecuted is emitted after an OnStop hook has been executed.
type OnStopExecuted struct {
	// FunctionName is the name of the function that was executed.
	FunctionName string

	// CallerName is the name of the function that scheduled the hook for
	// execution.
	CallerName string

	// Runtime specifies how long it took to run this hook.
	Runtime time.Duration

	// Err is non-nil if the hook failed to execute.
	Err error
}

// Supplied is emitted after a value is added with fx.Supply.
type Supplied struct {
	// TypeName is the name of the type of value that was added.
	TypeName string

	// StackTrace is the stack trace of the call to Supply.
	StackTrace []string

	// ModuleName is the name of the module in which the value was added to.
	ModuleName string

	// Err is non-nil if we failed to supply the value.
	Err error
}

// Provided is emitted when a constructor is provided to Fx.
type Provided struct {
	// ConstructorName is the name of the constructor that was provided to
	// Fx.
	ConstructorName string

	// StackTrace is the stack trace of where the constructor was provided to Fx.
	StackTrace []string

	// OutputTypeNames is a list of names of types that are produced by
	// this constructor.
	OutputTypeNames []string

	// ModuleName is the name of the module in which the constructor was
	// provided to.
	ModuleName string

	// Err is non-nil if we failed to provide this constructor.
	Err error

	// Private denotes whether the provided constructor is a [Private] constructor.
	Private bool
}

// Replaced is emitted when a value replaces a type in Fx.
type Replaced struct {
	// OutputTypeNames is a list of names of types that were replaced.
	OutputTypeNames []string

	// StackTrace is the stack trace of the call to Replace.
	StackTrace []string

	// ModuleName is the name of the module in which the value was added to.
	ModuleName string

	// Err is non-nil if we failed to supply the value.
	Err error
}

// Decorated is emitted when a decorator is executed in Fx.
type Decorated struct {
	// DecoratorName is the name of the decorator function that was
	// provided to Fx.
	DecoratorName string

	// StackTrace is the stack trace of where the decorator was given to Fx.
	StackTrace []string

	// ModuleName is the name of the module in which the value was added to.
	ModuleName string

	// OutputTypeNames is a list of names of types that are decorated by
	// this decorator.
	OutputTypeNames []string

	// Err is non-nil if we failed to run this decorator.
	Err error
}

// Run is emitted after a constructor, decorator, or supply/replace stub is run by Fx.
type Run struct {
	// Name is the name of the function that was run.
	Name string

	// Kind indicates which Fx option was used to pass along the function.
	// It is either "provide", "decorate", "supply", or "replace".
	Kind string

	// ModuleName is the name of the module in which the function belongs.
	ModuleName string

	// Err is non-nil if the function returned an error.
	// If fx.RecoverFromPanics is used, this will include panics.
	Err error
}

// Invoking is emitted before we invoke a function specified with fx.Invoke.
type Invoking struct {
	// FunctionName is the name of the function that will be invoked.
	FunctionName string

	// ModuleName is the name of the module in which the value was added to.
	ModuleName string
}

// Invoked is emitted after we invoke a function specified with fx.Invoke,
// whether it succeeded or failed.
type Invoked struct {
	// Functionname is the name of the function that was invoked.
	FunctionName string

	// ModuleName is the name of the module in which the value was added to.
	ModuleName string

	// Err is non-nil if the function failed to execute.
	Err error

	// Trace records information about where the fx.Invoke call was made.
	// Note that this is NOT a stack trace of the error itself.
	Trace string
}

// Started is emitted when an application is started successfully and/or it
// errored.
type Started struct {
	// Err is non-nil if the application failed to start successfully.
	Err error
}

// Stopping is emitted when the application receives a signal to shut down
// after starting. This may happen with fx.Shutdowner or by sending a signal to
// the application on the command line.
type Stopping struct {
	// Signal is the signal that caused this shutdown.
	Signal os.Signal
}

// Stopped is emitted when the application has finished shutting down, whether
// successfully or not.
type Stopped struct {
	// Err is non-nil if errors were encountered during shutdown.
	Err error
}

// RollingBack is emitted when the application failed to start up due to an
// error, and is being rolled back.
type RollingBack struct {
	// StartErr is the error that caused this rollback.
	StartErr error
}

// RolledBack is emitted after a service has been rolled back, whether it
// succeeded or not.
type RolledBack struct {
	// Err is non-nil if the rollback failed.
	Err error
}

// LoggerInitialized is emitted when a logger supplied with fx.WithLogger is
// instantiated, or if it fails to instantiate.
type LoggerInitialized struct {
	// ConstructorName is the name of the constructor that builds this
	// logger.
	ConstructorName string

	// Err is non-nil if the logger failed to build.
	Err error
}
```

### logger.go

```go

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxevent

// Logger defines interface used for logging.
type Logger interface {
	// LogEvent is called when a logging event is emitted.
	LogEvent(Event)
}

// NopLogger is an Fx event logger that ignores all messages.
var NopLogger = nopLogger{}

type nopLogger struct{}

var _ Logger = nopLogger{}

func (nopLogger) LogEvent(Event) {}

func (nopLogger) String() string { return "NopLogger" }
```

### zap.go

```go


// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxevent

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is an Fx event logger that logs events to Zap.
type ZapLogger struct {
	Logger *zap.Logger

	logLevel   zapcore.Level // default: zapcore.InfoLevel
	errorLevel *zapcore.Level
}

var _ Logger = (*ZapLogger)(nil)

// UseErrorLevel sets the level of error logs emitted by Fx to level.
func (l *ZapLogger) UseErrorLevel(level zapcore.Level) {
	l.errorLevel = &level
}

// UseLogLevel sets the level of non-error logs emitted by Fx to level.
func (l *ZapLogger) UseLogLevel(level zapcore.Level) {
	l.logLevel = level
}

func (l *ZapLogger) logEvent(msg string, fields ...zap.Field) {
	l.Logger.Log(l.logLevel, msg, fields...)
}

func (l *ZapLogger) logError(msg string, fields ...zap.Field) {
	lvl := zapcore.ErrorLevel
	if l.errorLevel != nil {
		lvl = *l.errorLevel
	}
	l.Logger.Log(lvl, msg, fields...)
}

// LogEvent logs the given event to the provided Zap logger.
func (l *ZapLogger) LogEvent(event Event) {
	switch e := event.(type) {
	case *OnStartExecuting:
		l.logEvent("OnStart hook executing",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *OnStartExecuted:
		if e.Err != nil {
			l.logError("OnStart hook failed",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.logEvent("OnStart hook executed",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *OnStopExecuting:
		l.logEvent("OnStop hook executing",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *OnStopExecuted:
		if e.Err != nil {
			l.logError("OnStop hook failed",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.logEvent("OnStop hook executed",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *Supplied:
		if e.Err != nil {
			l.logError("error encountered while applying options",
				zap.String("type", e.TypeName),
				zap.Strings("stacktrace", e.StackTrace),
				moduleField(e.ModuleName),
				zap.Error(e.Err))
		} else {
			l.logEvent("supplied",
				zap.String("type", e.TypeName),
				zap.Strings("stacktrace", e.StackTrace),
				moduleField(e.ModuleName),
			)
		}
	case *Provided:
		for _, rtype := range e.OutputTypeNames {
			l.logEvent("provided",
				zap.String("constructor", e.ConstructorName),
				zap.Strings("stacktrace", e.StackTrace),
				moduleField(e.ModuleName),
				zap.String("type", rtype),
				maybeBool("private", e.Private),
			)
		}
		if e.Err != nil {
			l.logError("error encountered while applying options",
				moduleField(e.ModuleName),
				zap.Strings("stacktrace", e.StackTrace),
				zap.Error(e.Err))
		}
	case *Replaced:
		for _, rtype := range e.OutputTypeNames {
			l.logEvent("replaced",
				zap.Strings("stacktrace", e.StackTrace),
				moduleField(e.ModuleName),
				zap.String("type", rtype),
			)
		}
		if e.Err != nil {
			l.logError("error encountered while replacing",
				zap.Strings("stacktrace", e.StackTrace),
				moduleField(e.ModuleName),
				zap.Error(e.Err))
		}
	case *Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.logEvent("decorated",
				zap.String("decorator", e.DecoratorName),
				zap.Strings("stacktrace", e.StackTrace),
				moduleField(e.ModuleName),
				zap.String("type", rtype),
			)
		}
		if e.Err != nil {
			l.logError("error encountered while applying options",
				zap.Strings("stacktrace", e.StackTrace),
				moduleField(e.ModuleName),
				zap.Error(e.Err))
		}
	case *Run:
		if e.Err != nil {
			l.logError("error returned",
				zap.String("name", e.Name),
				zap.String("kind", e.Kind),
				moduleField(e.ModuleName),
				zap.Error(e.Err),
			)
		} else {
			l.logEvent("run",
				zap.String("name", e.Name),
				zap.String("kind", e.Kind),
				moduleField(e.ModuleName),
			)
		}
	case *Invoking:
		// Do not log stack as it will make logs hard to read.
		l.logEvent("invoking",
			zap.String("function", e.FunctionName),
			moduleField(e.ModuleName),
		)
	case *Invoked:
		if e.Err != nil {
			l.logError("invoke failed",
				zap.Error(e.Err),
				zap.String("stack", e.Trace),
				zap.String("function", e.FunctionName),
				moduleField(e.ModuleName),
			)
		}
	case *Stopping:
		l.logEvent("received signal",
			zap.String("signal", strings.ToUpper(e.Signal.String())))
	case *Stopped:
		if e.Err != nil {
			l.logError("stop failed", zap.Error(e.Err))
		}
	case *RollingBack:
		l.logError("start failed, rolling back", zap.Error(e.StartErr))
	case *RolledBack:
		if e.Err != nil {
			l.logError("rollback failed", zap.Error(e.Err))
		}
	case *Started:
		if e.Err != nil {
			l.logError("start failed", zap.Error(e.Err))
		} else {
			l.logEvent("started")
		}
	case *LoggerInitialized:
		if e.Err != nil {
			l.logError("custom logger initialization failed", zap.Error(e.Err))
		} else {
			l.logEvent("initialized custom fxevent.Logger", zap.String("function", e.ConstructorName))
		}
	}
}

func moduleField(name string) zap.Field {
	if len(name) == 0 {
		return zap.Skip()
	}
	return zap.String("module", name)
}

func maybeBool(name string, b bool) zap.Field {
	if b {
		return zap.Bool(name, true)
	}
	return zap.Skip()
}
```

## internal

### fxclock

#### clock.go

```go


// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxclock

import (
	"context"
	"time"
)

// Clock defines how Fx accesses time.
// The interface is pretty minimal but it matches github.com/benbjohnson/clock.
// We intentionally don't use that interface directly;
// this keeps it a test dependency for us.
type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
	Sleep(time.Duration)
	WithTimeout(context.Context, time.Duration) (context.Context, context.CancelFunc)
}

// System is the default implementation of Clock based on real time.
var System Clock = systemClock{}

type systemClock struct{}

func (systemClock) Now() time.Time {
	return time.Now()
}

func (systemClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (systemClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (systemClock) WithTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, d)
}
```

### fxlog

#### default.go

```go

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxlog

import (
	"io"

	"go.uber.org/fx/fxevent"
)

// DefaultLogger constructs a Logger out of io.Writer.
func DefaultLogger(w io.Writer) fxevent.Logger {
	return &fxevent.ConsoleLogger{W: w}
}
```

#### spy.go

```go

// Copyright (c) 2020-2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxlog

import (
	"reflect"
	"sync"

	"go.uber.org/fx/fxevent"
)

// Events is a list of events captured by fxlog.Spy.
type Events []fxevent.Event

// Len returns the number of events in this list.
func (es Events) Len() int { return len(es) }

// SelectByTypeName returns a new list with only events matching the specified
// type.
func (es Events) SelectByTypeName(name string) Events {
	var out Events
	for _, e := range es {
		if reflect.TypeOf(e).Elem().Name() == name {
			out = append(out, e)
		}
	}
	return out
}

// Spy is an Fx event logger that captures emitted events and/or logged
// statements. It may be used in tests of Fx logs.
type Spy struct {
	mu     sync.RWMutex
	events Events
}

var _ fxevent.Logger = &Spy{}

// LogEvent appends an Event.
func (s *Spy) LogEvent(event fxevent.Event) {
	s.mu.Lock()
	s.events = append(s.events, event)
	s.mu.Unlock()
}

// Events returns all captured events.
func (s *Spy) Events() Events {
	s.mu.RLock()
	defer s.mu.RUnlock()

	events := make(Events, len(s.events))
	copy(events, s.events)
	return events
}

// EventTypes returns all captured event types.
func (s *Spy) EventTypes() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	types := make([]string, len(s.events))
	for i, e := range s.events {
		types[i] = reflect.TypeOf(e).Elem().Name()
	}
	return types
}

// Reset clears all messages and events from the Spy.
func (s *Spy) Reset() {
	s.mu.Lock()
	s.events = s.events[:0]
	s.mu.Unlock()
}
```

### fxreflect

#### fxreflect.go

```go

// Copyright (c) 2019-2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxreflect

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

// Match from beginning of the line until the first `vendor/` (non-greedy)
var vendorRe = regexp.MustCompile("^.*?/vendor/")

// sanitize makes the function name suitable for logging display. It removes
// url-encoded elements from the `dot.git` package names and shortens the
// vendored paths.
func sanitize(function string) string {
	// Use the stdlib to un-escape any package import paths which can happen
	// in the case of the "dot-git" postfix. Seems like a bug in stdlib =/
	if unescaped, err := url.QueryUnescape(function); err == nil {
		function = unescaped
	}

	// strip everything prior to the vendor
	return vendorRe.ReplaceAllString(function, "vendor/")
}

// Caller returns the formatted calling func name
func Caller() string {
	return CallerStack(1, 0).CallerName()
}

// FuncName returns a funcs formatted name
func FuncName(fn interface{}) string {
	fnV := reflect.ValueOf(fn)
	if fnV.Kind() != reflect.Func {
		return fmt.Sprint(fn)
	}

	function := runtime.FuncForPC(fnV.Pointer()).Name()
	return fmt.Sprintf("%s()", sanitize(function))
}

// Ascend the call stack until we leave the Fx production code. This allows us
// to avoid hard-coding a frame skip, which makes this code work well even
// when it's wrapped.
func shouldIgnoreFrame(f Frame) bool {
	// Treat test files as leafs.
	if strings.Contains(f.File, "_test.go") {
		return false
	}

	// The unique, fully-qualified name for all functions begins with
	// "{{importPath}}.". We'll ignore Fx and its subpackages.
	s := strings.TrimPrefix(f.Function, "go.uber.org/fx")
	if len(s) > 0 && s[0] == '.' || s[0] == '/' {
		// We want to match,
		//   go.uber.org/fx.Foo
		//   go.uber.org/fx/something.Foo
		// But not, go.uber.org/fxfoo
		return true
	}

	return false
}
```

#### stack.go

```go

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fxreflect

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

// Frame holds information about a single frame in the call stack.
type Frame struct {
	// Unique, package path-qualified name for the function of this call
	// frame.
	Function string

	// File and line number of our location in the frame.
	//
	// Note that the line number does not refer to where the function was
	// defined but where in the function the next call was made.
	File string
	Line int
}

func (f Frame) String() string {
	// This takes the following forms.
	//  (path/to/file.go)
	//  (path/to/file.go:42)
	//  path/to/package.MyFunction
	//  path/to/package.MyFunction (path/to/file.go)
	//  path/to/package.MyFunction (path/to/file.go:42)

	var sb strings.Builder
	sb.WriteString(f.Function)
	if len(f.File) > 0 {
		if sb.Len() > 0 {
			sb.WriteRune(' ')
		}
		fmt.Fprintf(&sb, "(%v", f.File)
		if f.Line > 0 {
			fmt.Fprintf(&sb, ":%d", f.Line)
		}
		sb.WriteRune(')')
	}

	if sb.Len() == 0 {
		return "unknown"
	}

	return sb.String()
}

const _defaultCallersDepth = 8

// Stack is a stack of call frames.
//
// Formatted with %v, the output is in a single-line, in the form,
//
//	foo/bar.Baz() (path/to/foo.go:42); bar/baz.Qux() (bar/baz/qux.go:12); ...
//
// Formatted with %+v, the output is in the form,
//
//	foo/bar.Baz()
//		path/to/foo.go:42
//	bar/baz.Qux()
//		bar/baz/qux.go:12
type Stack []Frame

// String returns a single-line, semi-colon representation of a Stack.
// For a list of strings where each represents one frame, use Strings.
// For a cleaner multi-line representation, use %+v.
func (fs Stack) String() string {
	return strings.Join(fs.Strings(), "; ")
}

// Strings returns a list of strings, each representing a frame in the stack.
// Each line will be in the form,
//
//	foo/bar.Baz() (path/to/foo.go:42)
func (fs Stack) Strings() []string {
	items := make([]string, len(fs))
	for i, f := range fs {
		items[i] = f.String()
	}
	return items
}

// Format implements fmt.Formatter to handle "%+v".
func (fs Stack) Format(w fmt.State, c rune) {
	if !w.Flag('+') {
		// Without %+v, fall back to String().
		io.WriteString(w, fs.String())
		return
	}

	for _, f := range fs {
		fmt.Fprintln(w, f.Function)
		fmt.Fprintf(w, "\t%v:%v\n", f.File, f.Line)
	}
}

// CallerName returns the name of the first caller in this stack that isn't
// owned by the Fx library.
func (fs Stack) CallerName() string {
	for _, f := range fs {
		if shouldIgnoreFrame(f) {
			continue
		}
		return f.Function
	}
	return "n/a"
}

// CallerStack returns the call stack for the calling function, up to depth frames
// deep, skipping the provided number of frames, not including Callers itself.
//
// If zero, depth defaults to 8.
func CallerStack(skip, depth int) Stack {
	if depth <= 0 {
		depth = _defaultCallersDepth
	}

	pcs := make([]uintptr, depth)

	// +2 to skip this frame and runtime.Callers.
	n := runtime.Callers(skip+2, pcs)
	pcs = pcs[:n] // truncate to number of frames actually read

	result := make([]Frame, 0, n)
	frames := runtime.CallersFrames(pcs)
	for f, more := frames.Next(); more; f, more = frames.Next() {
		result = append(result, Frame{
			Function: sanitize(f.Function),
			File:     f.File,
			Line:     f.Line,
		})
	}
	return result
}
```

### lifecycle

#### lifecycle.go

```go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
	"time"

	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/internal/fxclock"
	"go.uber.org/fx/internal/fxreflect"
	"go.uber.org/multierr"
)

// Reflection types for each of the supported hook function signatures. These
// are used in cases in which the Callable constraint matches a user-defined
// function type that cannot be converted to an underlying function type with
// a conventional conversion or type switch.
var (
	_reflFunc             = reflect.TypeOf(Func(nil))
	_reflErrorFunc        = reflect.TypeOf(ErrorFunc(nil))
	_reflContextFunc      = reflect.TypeOf(ContextFunc(nil))
	_reflContextErrorFunc = reflect.TypeOf(ContextErrorFunc(nil))
)

// Discrete function signatures that are allowed as part of a [Callable].
type (
	// A Func can be converted to a ContextErrorFunc.
	Func = func()
	// An ErrorFunc can be converted to a ContextErrorFunc.
	ErrorFunc = func() error
	// A ContextFunc can be converted to a ContextErrorFunc.
	ContextFunc = func(context.Context)
	// A ContextErrorFunc is used as a [Hook.OnStart] or [Hook.OnStop]
	// function.
	ContextErrorFunc = func(context.Context) error
)

// A Callable is a constraint that matches functions that are, or can be
// converted to, functions suitable for a Hook.
//
// Callable must be identical to [fx.HookFunc].
type Callable interface {
	~Func | ~ErrorFunc | ~ContextFunc | ~ContextErrorFunc
}

// Wrap wraps x into a ContextErrorFunc suitable for a Hook.
func Wrap[T Callable](x T) (ContextErrorFunc, string) {
	if x == nil {
		return nil, ""
	}

	switch fn := any(x).(type) {
	case Func:
		return func(context.Context) error {
			fn()
			return nil
		}, fxreflect.FuncName(x)
	case ErrorFunc:
		return func(context.Context) error {
			return fn()
		}, fxreflect.FuncName(x)
	case ContextFunc:
		return func(ctx context.Context) error {
			fn(ctx)
			return nil
		}, fxreflect.FuncName(x)
	case ContextErrorFunc:
		return fn, fxreflect.FuncName(x)
	}

	// Since (1) we're already using reflect in Fx, (2) we're not particularly
	// concerned with performance, and (3) unsafe would require discrete build
	// targets for appengine (etc), just use reflect to convert user-defined
	// function types to their underlying function types and then call Wrap
	// again with the converted value.
	reflVal := reflect.ValueOf(x)
	switch {
	case reflVal.CanConvert(_reflFunc):
		return Wrap(reflVal.Convert(_reflFunc).Interface().(Func))
	case reflVal.CanConvert(_reflErrorFunc):
		return Wrap(reflVal.Convert(_reflErrorFunc).Interface().(ErrorFunc))
	case reflVal.CanConvert(_reflContextFunc):
		return Wrap(reflVal.Convert(_reflContextFunc).Interface().(ContextFunc))
	default:
		// Is already convertible to ContextErrorFunc.
		return Wrap(reflVal.Convert(_reflContextErrorFunc).Interface().(ContextErrorFunc))
	}
}

// A Hook is a pair of start and stop callbacks, either of which can be nil,
// plus a string identifying the supplier of the hook.
type Hook struct {
	OnStart     func(context.Context) error
	OnStop      func(context.Context) error
	OnStartName string
	OnStopName  string

	callerFrame fxreflect.Frame
}

type appState int

const (
	stopped appState = iota
	starting
	incompleteStart
	started
	stopping
)

func (as appState) String() string {
	switch as {
	case stopped:
		return "stopped"
	case starting:
		return "starting"
	case incompleteStart:
		return "incompleteStart"
	case started:
		return "started"
	case stopping:
		return "stopping"
	default:
		return "invalidState"
	}
}

// Lifecycle coordinates application lifecycle hooks.
type Lifecycle struct {
	clock        fxclock.Clock
	logger       fxevent.Logger
	state        appState
	hooks        []Hook
	numStarted   int
	startRecords HookRecords
	stopRecords  HookRecords
	runningHook  Hook
	mu           sync.Mutex
}

// New constructs a new Lifecycle.
func New(logger fxevent.Logger, clock fxclock.Clock) *Lifecycle {
	return &Lifecycle{logger: logger, clock: clock}
}

// Append adds a Hook to the lifecycle.
func (l *Lifecycle) Append(hook Hook) {
	// Save the caller's stack frame to report file/line number.
	if f := fxreflect.CallerStack(2, 0); len(f) > 0 {
		hook.callerFrame = f[0]
	}
	l.hooks = append(l.hooks, hook)
}

// Start runs all OnStart hooks, returning immediately if it encounters an
// error.
func (l *Lifecycle) Start(ctx context.Context) error {
	if ctx == nil {
		return errors.New("called OnStart with nil context")
	}

	l.mu.Lock()
	if l.state != stopped {
		defer l.mu.Unlock()
		return fmt.Errorf("attempted to start lifecycle when in state: %v", l.state)
	}
	l.numStarted = 0
	l.state = starting

	l.startRecords = make(HookRecords, 0, len(l.hooks))
	l.mu.Unlock()

	var returnState appState = incompleteStart
	defer func() {
		l.mu.Lock()
		l.state = returnState
		l.mu.Unlock()
	}()

	for _, hook := range l.hooks {
		// if ctx has cancelled, bail out of the loop.
		if err := ctx.Err(); err != nil {
			return err
		}

		if hook.OnStart != nil {
			l.mu.Lock()
			l.runningHook = hook
			l.mu.Unlock()

			runtime, err := l.runStartHook(ctx, hook)
			if err != nil {
				return err
			}

			l.mu.Lock()
			l.startRecords = append(l.startRecords, HookRecord{
				CallerFrame: hook.callerFrame,
				Func:        hook.OnStart,
				Runtime:     runtime,
			})
			l.mu.Unlock()
		}
		l.numStarted++
	}

	returnState = started
	return nil
}

func (l *Lifecycle) runStartHook(ctx context.Context, hook Hook) (runtime time.Duration, err error) {
	funcName := hook.OnStartName
	if len(funcName) == 0 {
		funcName = fxreflect.FuncName(hook.OnStart)
	}

	l.logger.LogEvent(&fxevent.OnStartExecuting{
		CallerName:   hook.callerFrame.Function,
		FunctionName: funcName,
	})
	defer func() {
		l.logger.LogEvent(&fxevent.OnStartExecuted{
			CallerName:   hook.callerFrame.Function,
			FunctionName: funcName,
			Runtime:      runtime,
			Err:          err,
		})
	}()

	begin := l.clock.Now()
	err = hook.OnStart(ctx)
	return l.clock.Since(begin), err
}

// Stop runs any OnStop hooks whose OnStart counterpart succeeded. OnStop
// hooks run in reverse order.
func (l *Lifecycle) Stop(ctx context.Context) error {
	if ctx == nil {
		return errors.New("called OnStop with nil context")
	}

	l.mu.Lock()
	if l.state != started && l.state != incompleteStart && l.state != starting {
		defer l.mu.Unlock()
		return nil
	}
	l.state = stopping
	l.mu.Unlock()

	defer func() {
		l.mu.Lock()
		l.state = stopped
		l.mu.Unlock()
	}()

	l.mu.Lock()
	l.stopRecords = make(HookRecords, 0, l.numStarted)
	// Take a snapshot of hook state to avoid races.
	allHooks := l.hooks[:]
	numStarted := l.numStarted
	l.mu.Unlock()

	// Run backward from last successful OnStart.
	var errs []error
	for ; numStarted > 0; numStarted-- {
		if err := ctx.Err(); err != nil {
			return err
		}
		hook := allHooks[numStarted-1]
		if hook.OnStop == nil {
			continue
		}

		l.mu.Lock()
		l.runningHook = hook
		l.mu.Unlock()

		runtime, err := l.runStopHook(ctx, hook)
		if err != nil {
			// For best-effort cleanup, keep going after errors.
			errs = append(errs, err)
		}

		l.mu.Lock()
		l.stopRecords = append(l.stopRecords, HookRecord{
			CallerFrame: hook.callerFrame,
			Func:        hook.OnStop,
			Runtime:     runtime,
		})
		l.mu.Unlock()
	}

	return multierr.Combine(errs...)
}

func (l *Lifecycle) runStopHook(ctx context.Context, hook Hook) (runtime time.Duration, err error) {
	funcName := hook.OnStopName
	if len(funcName) == 0 {
		funcName = fxreflect.FuncName(hook.OnStop)
	}

	l.logger.LogEvent(&fxevent.OnStopExecuting{
		CallerName:   hook.callerFrame.Function,
		FunctionName: funcName,
	})
	defer func() {
		l.logger.LogEvent(&fxevent.OnStopExecuted{
			CallerName:   hook.callerFrame.Function,
			FunctionName: funcName,
			Runtime:      runtime,
			Err:          err,
		})
	}()

	begin := l.clock.Now()
	err = hook.OnStop(ctx)
	return l.clock.Since(begin), err
}

// RunningHookCaller returns the name of the hook that was running when a Start/Stop
// hook timed out.
func (l *Lifecycle) RunningHookCaller() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.runningHook.callerFrame.Function
}

// HookRecord keeps track of each Hook's execution time, the caller that appended the Hook, and function that ran as the Hook.
type HookRecord struct {
	CallerFrame fxreflect.Frame             // stack frame of the caller
	Func        func(context.Context) error // function that ran as sanitized name
	Runtime     time.Duration               // how long the hook ran
}

// HookRecords is a Stringer wrapper of HookRecord slice.
type HookRecords []HookRecord

func (rs HookRecords) Len() int {
	return len(rs)
}

func (rs HookRecords) Less(i, j int) bool {
	// Sort by runtime, greater ones at top.
	return rs[i].Runtime > rs[j].Runtime
}

func (rs HookRecords) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// Used for logging startup errors.
func (rs HookRecords) String() string {
	var b strings.Builder
	for _, r := range rs {
		fmt.Fprintf(&b, "%s took %v from %s",
			fxreflect.FuncName(r.Func), r.Runtime, r.CallerFrame)
	}
	return b.String()
}

// Format implements fmt.Formatter to handle "%+v".
func (rs HookRecords) Format(w fmt.State, c rune) {
	if !w.Flag('+') {
		// Without %+v, fall back to String().
		io.WriteString(w, rs.String())
		return
	}

	for _, r := range rs {
		fmt.Fprintf(w, "\n%s took %v from:\n\t%+v",
			fxreflect.FuncName(r.Func),
			r.Runtime,
			r.CallerFrame)
	}
	fmt.Fprintf(w, "\n")
}
```
{% endraw %}
