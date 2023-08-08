---
layout: post
title: "Microservices with FX, KrakenD and Consul"
mermaid: true
categories: [ "tech", "golang" ]
---
I want to simple tools to build reliable and effective microservices. At the current time, I choose the three main tools:
1. Microservice Framework: FX framework created by Uber. I value its design philosophy.
2. Configuration and services center: consul created by Hashicorp.
3. Gateway: KrakenD. I value its flex configuration and api aggregation.

## architecture

<div class="mermaid">
flowchart LR
    subgraph arch["microservices with fx, krakend and consul"]
    krakend --http--> S1 & S2 & dots[...] & Sn -. trace .-> zipkin
    krakend -. trace .-> zipkin
    subgraph gateway["api gateway"]
    krakend
    style krakend fill:#ceb,color:#a2c
    style gateway color:#e2f
    end
    subgraph center["configuration and service center"]
    krakend <--http--> consul
    style consul fill:#f8f,color:#d2d
    style center color:#f12
    end
    subgraph microservices["microservices built with fx"]
    consul <--http--> S1 & S2 & dots[...] & Sn
    style S1 fill:#aef,color:#e2e
    style S2 fill:#aef,color:#e2e
    style Sn fill:#aef,color:#e2e
    style dots fill:none,stroke:none
    style microservices color:#e12
    end
    subgraph tracer
    zipkin
    style zipkin fill:#fce,color:#b2b
    style tracer color:#e2d
    end
    style arch fill:#fff,color:#f84
    end
    linkStyle 0,1,2,3 stroke:green,color:red;
    linkStyle 4,5,6,7,8 stroke:red,color:green;
    linkStyle 9,10,11,12,13 stroke:purple,color:blue;
</div>
<script>
    const config = {
      startOnLoad: true,
      // Available curve styles include basis, bumpX, bumpY, cardinal, catmullRom, linear, monotoneX, monotoneY, natural, step, stepAfter, and stepBefore
      flowchart: { useMaxWidth: true, htmlLabels: true, curve: 'bumpY' },
      securityLevel: 'loose',
    };
    mermaid.initialize(config);
</script>

## Krakend

Krakend is built on lura, which is an open framework to assemble ultra performance API Gateways with middlewares. So firstly let's study lura.

### lura

#### config

config/config.go

```go
// ServiceConfig defines the lura service
type ServiceConfig struct {
        // name of the service
        Name string `mapstructure:"name"`
        // set of endpoint definitions
        Endpoints []*EndpointConfig `mapstructure:"endpoints"`
        // set of async agent definitions
        AsyncAgents []*AsyncAgent `mapstructure:"async_agent"`
        // defafult timeout
        Timeout time.Duration `mapstructure:"timeout"`
        // default TTL for GET
        CacheTTL time.Duration `mapstructure:"cache_ttl"`
        // default set of hosts
        Host []string `mapstructure:"host"`
        // port to bind the lura service
        Port int `mapstructure:"port"`
        // version code of the configuration
        Version int `mapstructure:"version"`
        // OutputEncoding defines the default encoding strategy 
        // to use for the endpoint responses
        OutputEncoding string `mapstructure:"output_encoding"`
        // Extra configuration for customized behaviour
        ExtraConfig ExtraConfig `mapstructure:"extra_config"`

        // ReadTimeout is the maximum duration for reading the entire
        // request, including the body.
        //
        // Because ReadTimeout does not let Handlers make per-request
        // decisions on each request body's acceptable deadline or
        // upload rate, most users will prefer to use
        // ReadHeaderTimeout. It is valid to use them both.
        ReadTimeout time.Duration `mapstructure:"read_timeout"`
        // WriteTimeout is the maximum duration before timing out
        // writes of the response. It is reset whenever a new
        // request's header is read. Like ReadTimeout, it does not
        // let Handlers make decisions on a per-request basis.
        WriteTimeout time.Duration `mapstructure:"write_timeout"`
        // IdleTimeout is the maximum amount of time to wait for the
        // next request when keep-alives are enabled. If IdleTimeout
        // is zero, the value of ReadTimeout is used. If both are
        // zero, ReadHeaderTimeout is used.
        IdleTimeout time.Duration `mapstructure:"idle_timeout"`
        // ReadHeaderTimeout is the amount of time allowed to read
        // request headers. The connection's read deadline is reset
        // after reading the headers and the Handler can decide what
        // is considered too slow for the body.
        ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`

        // DisableKeepAlives, if true, prevents re-use of TCP connections
        // between different HTTP requests.
        DisableKeepAlives bool `mapstructure:"disable_keep_alives"`
        // DisableCompression, if true, prevents the Transport from
        // requesting compression with an "Accept-Encoding: gzip"
        // request header when the Request contains no existing
        // Accept-Encoding value. If the Transport requests gzip on
        // its own and gets a gzipped response, it's transparently
        // decoded in the Response.Body. However, if the user
        // explicitly requested gzip it is not automatically
        // uncompressed.
        DisableCompression bool `mapstructure:"disable_compression"`
        // MaxIdleConns controls the maximum number of idle (keep-alive)
        // connections across all hosts. Zero means no limit.
        MaxIdleConns int `mapstructure:"max_idle_connections"`
        // MaxIdleConnsPerHost, if non-zero, controls the maximum idle
        // (keep-alive) connections to keep per-host. If zero,
        // DefaultMaxIdleConnsPerHost is used.
        MaxIdleConnsPerHost int `mapstructure:"max_idle_connections_per_host"`
        // IdleConnTimeout is the maximum amount of time an idle
        // (keep-alive) connection will remain idle before closing
// itself.
        // Zero means no limit.
        IdleConnTimeout time.Duration `mapstructure:"idle_connection_timeout"`
        // ResponseHeaderTimeout, if non-zero, specifies the amount of
        // time to wait for a server's response headers after fully
        // writing the request (including its body, if any). This
        // time does not include the time to read the response body.
        ResponseHeaderTimeout time.Duration `mapstructure:"response_header_timeout"`
        // ExpectContinueTimeout, if non-zero, specifies the amount of
        // time to wait for a server's first response headers after fully
        // writing the request headers if the request has an
        // "Expect: 100-continue" header. Zero means no timeout and
        // causes the body to be sent immediately, without
        // waiting for the server to approve.
        // This time does not include the time to send the request header.
        ExpectContinueTimeout time.Duration `mapstructure:"expect_continue_timeout"`
        // DialerTimeout is the maximum amount of time a dial will wait for
        // a connect to complete. If Deadline is also set, it may fail
        // earlier.
        //
        // The default is no timeout.
//
        // When using TCP and dialing a host name with multiple IP
        // addresses, the timeout may be divided between them.
        //
        // With or without a timeout, the operating system may impose
        // its own earlier timeout. For instance, TCP timeouts are
        // often around 3 minutes.
        DialerTimeout time.Duration `mapstructure:"dialer_timeout"`
        // DialerFallbackDelay specifies the length of time to wait before
        // spawning a fallback connection, when DualStack is enabled.
        // If zero, a default delay of 300ms is used.
        DialerFallbackDelay time.Duration `mapstructure:"dialer_fallback_delay"`
        // DialerKeepAlive specifies the keep-alive period for an active
        // network connection.
        // If zero, keep-alives are not enabled. Network protocols
        // that do not support keep-alives ignore this field.
        DialerKeepAlive time.Duration `mapstructure:"dialer_keep_alive"`

        // DisableStrictREST flags if the REST enforcement is disabled
        DisableStrictREST bool `mapstructure:"disable_rest"`

        // Plugin defines the configuration for the plugin loader
        Plugin *Plugin `mapstructure:"plugin"`

        // TLS defines the configuration params for enabling TLS (HTTPS & HTTP/2) at
        // the router layer
        TLS *TLS `mapstructure:"tls"`

        // run lura in debug mode
        Debug     bool `mapstructure:"debug_endpoint"`
        Echo      bool `mapstructure:"echo_endpoint"`
        uriParser URIParser

        // SequentialStart flags if the agents should be started sequentially
        // before starting the router
        SequentialStart bool `mapstructure:"sequential_start"`

        // AllowInsecureConnections sets the http client tls configuration to allow
        // insecure connections to the backends for development
        //  (enables InsecureSkipVerify)
        AllowInsecureConnections bool `mapstructure:"allow_insecure_connections"`

        // ClientTLS is used to configure the http default transport
        // with TLS parameters
        ClientTLS *ClientTLS `mapstructure:"client_tls"`
}
```

```go
// AsyncAgent defines the configuration of a single subscriber/consumer
// to be initialized and maintained by the lura service
type AsyncAgent struct {
        Name       string     `mapstructure:"name"`
        Connection Connection `mapstructure:"connection"`
        Consumer   Consumer   `mapstructure:"consumer"`
        // the encoding format
        Encoding string `mapstructure:"encoding"`
        // set of definitions of the backends to be linked to this endpoint
        Backend []*Backend `mapstructure:"backend"`

        // Endpoint Extra configuration for customized behaviour
        ExtraConfig ExtraConfig `mapstructure:"extra_config"`
}

type Consumer struct {
        // timeout of the pipe defined by this subscriber
        Timeout time.Duration `mapstructure:"timeout"`
        Workers int           `mapstructure:"workers"`
        Topic   string        `mapstructure:"topic"`
        MaxRate float64       `mapstructure:"max_rate"`
}

type Connection struct {
        MaxRetries      int           `mapstructure:"max_retries"`
        BackoffStrategy string        `mapstructure:"backoff_strategy"`
        HealthInterval  time.Duration `mapstructure:"health_interval"`
}
```

```go
// EndpointConfig defines the configuration of a single endpoint to be exposed
// by the lura service
type EndpointConfig struct {
        // url pattern to be registered and exposed to the world
        Endpoint string `mapstructure:"endpoint"`
        // HTTP method of the endpoint (GET, POST, PUT, etc)
        Method string `mapstructure:"method"`
        // set of definitions of the backends to be linked to this endpoint
        Backend []*Backend `mapstructure:"backend"`
        // number of concurrent calls this endpoint must send to the backends
        ConcurrentCalls int `mapstructure:"concurrent_calls"`
        // timeout of this endpoint
        Timeout time.Duration `mapstructure:"timeout"`
        // duration of the cache header
        CacheTTL time.Duration `mapstructure:"cache_ttl"`
        // list of query string params to be extracted from the URI
        QueryString []string `mapstructure:"input_query_strings"`
        // Endpoint Extra configuration for customized behaviour
        ExtraConfig ExtraConfig `mapstructure:"extra_config"`
        // HeadersToPass defines the list of headers to pass to the backends
        HeadersToPass []string `mapstructure:"input_headers"`
        // OutputEncoding defines the encoding strategy to use
        // for the endpoint responses
        OutputEncoding string `mapstructure:"output_encoding"`
}
```

```go
// Backend defines how lura should connect to the backend service
// (the API resource to consume)
// and how it should process the received response
type Backend struct {
        // Group defines the name of the property the response
        // should be moved to. If empty, the response is not changed
        Group string `mapstructure:"group"`
        // Method defines the HTTP method of the request to send to the backend
        Method string `mapstructure:"method"`
        // Host is a set of hosts of the API
        Host []string `mapstructure:"host"`
        // HostSanitizationDisabled can be set to false 
        // if the hostname should be sanitized
        HostSanitizationDisabled bool `mapstructure:"disable_host_sanitize"`
        // URLPattern is the URL pattern to use to locate 
        // the resource to be consumed
        URLPattern string `mapstructure:"url_pattern"`
        // AllowList is a set of response fields to allow. 
        // If empty, the filter id not used
        AllowList []string `mapstructure:"allow"`
        // DenyList is a set of response fields to remove. 
        // If empty, the filter id not used
AllowList []string `mapstructure:"allow"`
        // DenyList is a set of response fields to remove. 
        // If empty, the filter id not used
        DenyList []string `mapstructure:"deny"`
        // map of response fields to be renamed and their new names
        Mapping map[string]string `mapstructure:"mapping"`
        // the encoding format
        Encoding string `mapstructure:"encoding"`
        // the response to process is a collection
        IsCollection bool `mapstructure:"is_collection"`
        // name of the field to extract to the root. 
        // If empty, the formater will do nothing
        Target string `mapstructure:"target"`
        // name of the service discovery driver to use
        SD string `mapstructure:"sd"`
        // scheme to use for servers fetched from
        SDScheme string `mapstructure:"sd_scheme"`

        // list of keys to be replaced in the URLPattern
        URLKeys []string
        // number of concurrent calls this endpoint must send to the API
        URLKeys []string
        // number of concurrent calls this endpoint must send to the API
        ConcurrentCalls int
        // timeout of this backend
        Timeout time.Duration
        // decoder to use in order to parse the received response from the API
        Decoder encoding.Decoder `json:"-"`
        // Backend Extra configuration for customized behaviours
        ExtraConfig ExtraConfig `mapstructure:"extra_config"`
        // HeadersToPass defines the list of headers to pass to this backend
        HeadersToPass []string `mapstructure:"input_headers"`
}
```

```go
// Plugin contains the config required by the plugin module
type Plugin struct {
        Folder  string `mapstructure:"folder"`
        Pattern string `mapstructure:"pattern"`
}

// TLS defines the configuration params for enabling TLS 
// (HTTPS & HTTP/2) at the router layer
type TLS struct {
        IsDisabled               bool     `mapstructure:"disabled"`
        PublicKey                string   `mapstructure:"public_key"`
        PrivateKey               string   `mapstructure:"private_key"`
        CaCerts                  []string `mapstructure:"ca_certs"`
        MinVersion               string   `mapstructure:"min_version"`
        MaxVersion               string   `mapstructure:"max_version"`
        CurvePreferences         []uint16 `mapstructure:"curve_preferences"`
    PreferServerCipherSuites bool     `mapstructure:"prefer_server_cipher_suites"`
        EnableMTLS               bool     `mapstructure:"enable_mtls"`
        DisableSystemCaPool      bool     `mapstructure:"disable_system_ca_pool"`
}

// ClientTLS defines the configuration params for an HTTP Client
type ClientTLS struct {
    AllowInsecureConnections bool     `mapstructure:"allow_insecure_connections"`
        CaCerts                  []string `mapstructure:"ca_certs"`
        DisableSystemCaPool      bool     `mapstructure:"disable_system_ca_pool"`
        MinVersion               string   `mapstructure:"min_version"`
        MaxVersion               string   `mapstructure:"max_version"`
        CurvePreferences         []uint16 `mapstructure:"curve_preferences"`
        CipherSuites             []uint16 `mapstructure:"cipher_suites"`
}
```
