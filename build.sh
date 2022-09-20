#!/bin/sh

npm --prefix=./ui ci && npm --prefix=./ui run build
goreleaser build --snapshot --rm-dist