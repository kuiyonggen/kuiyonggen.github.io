# Get started with Fx

This introduces you to the basics of Fx.
In this tutorial you will:

- [start an empty application](#create-a-minimal-application)
- [add an HTTP server to it](#add-an-http-server)
- [register a handler with the server](#register-a-handler)
- [add logging to your application](#add-a-logger)
- [refactor to loosen coupling to your handler](#decouple-registration)
- [add another handler to the server](#register-another-handler)
- [generalize your implementation](#register-many-handlers)

First, get set up for the rest of the tutorial.

1. Start a new empty project.

   ```bash
   mkdir fxdemo
   cd fxdemo
   go mod init example.com/fxdemo
   ```

2. Install the latest version of Fx.

   ```bash
   go get go.uber.org/fx@latest
   ```

Now begin by [creating a minimal application](#create-a-minimal-application).
