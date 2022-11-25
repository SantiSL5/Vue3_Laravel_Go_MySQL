#!/bin/bash
chown -R go:go /go
go mod tidy | goreload main.go
exec "$@";