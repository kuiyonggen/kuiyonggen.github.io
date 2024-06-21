package main

import (
    "fxmoduledemo/bundlefx" 
    "fxmoduledemo/httphandler" 
    "go.uber.org/fx"
)

func main() { 
    fx.New(
        bundlefx.Module,
        fx.Invoke(httphandler.New),
    ).Run()
}
