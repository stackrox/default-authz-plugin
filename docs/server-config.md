# Server Configuration

This document describes how the server for the Default Authorization Plugin can
be configured. Here, "server" refers to the functionality that is not related to
evaluating access rules (for that, see [this document](writing-gval-rules.md) instead).

The server configuration controls:
- which interface to bind to
- which port to listen on
- whether to use TLS (and if yes, which certificates to use)
- how authentication is performed

The server configuration is stored in a JSON file. This file contains the server configuration
as a JSON object. The following example gives an overview of all the possible configuration
options:
```
{
  // The address to bind to; can be used to listen only on specific interfaces.
  "bindAddress": "0.0.0.0",
  // The port to listen on.
  "port": 8080,
  // The TLS configuration. If this is not specified, "disableTls" must be set
  // to true.
  "tls": {
    // File storing the PEM-encoded TLS certificate private key.
    "keyFile": "/etc/certs/server-key.pem",
    // File storing the PEM-encoded TLS certificate.
    "certFile": "/etc/certs/server-cert.pem"
  },
  // Whether to disable TLS and use plain HTTP only. For secuity reasons,
  // this must be explicitly specified.
  "disableTLS": false,
  // Authentication configuration.
  "auth": {
    // Exactly one of the following values must be set (in this example we set
    // multiple for demonstration purposes).
    
    // Whether to allow anonymous access.
    "allowAnonymous": false,
    // PEM-encoded CA certificate for client certificate authentication.
    "clientCACertFile": "/etc/certs/client-ca.pem",
    // .htpasswd file for HTTP basic authentication. Only bcrypt may be used/
    "htpasswdFile": "/etc/authz-plugin/.htpasswd"
  }
}
```

Further examples can be found in the `examples/config` subdirectory of the source root.

## Specifying a Config File

The server configuration file is passed to the Default Authorization Plugin binary via
the `-server-config <file.json>` flag.
