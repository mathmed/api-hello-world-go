#!/bin/bash
cd /usr/local/go/src/api-hello-world-go && /usr/local/go/bin/go mod tidy
exec /usr/local/go/bin/go run /usr/local/go/src/api-hello-world-go/main.go
