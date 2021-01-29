#!/bin/bash
cd /app && /usr/local/go/bin/go get github.com/joho/godotenv && /usr/local/go/bin/go get -u github.com/gorilla/mux
exec /usr/local/go/bin/go run /app/main.go