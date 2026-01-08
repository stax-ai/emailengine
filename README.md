# EmailEngine Go SDK

Typed Go client and ergonomic SDK for EmailEngine with Zap logging, OpenTelemetry tracing, retries, and pagination helpers.

## Install

```bash
go get github.com/stax-ai/emailengine
```

## Quickstart

```go
package main

import (
  "context"
  "os"
  "github.com/stax-ai/emailengine"
)

func main() {
  cli, _ := emailengine.New(
    emailengine.WithToken(os.Getenv("EMAILENGINE_TOKEN")),
  )
  // List accounts
  _, _ = cli.ListAccounts(context.Background(), 0, 20)
}
```

## Features

- Zap logging and OpenTelemetry tracing middleware
- Automatic retries for 429/5xx and transient network errors
- Context deadline -> X-EE-Timeout header mapping
- Simple helpers: `ListAccounts`, `SendMessage`, and full raw client access via `Raw()`
- Generic paginator helper

## Examples

See `examples/basic`.

## Regenerate API (from swagger.json)

```bash
scripts/generate.sh
```


