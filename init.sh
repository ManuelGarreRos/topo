#!/bin/bash

if [ "$TOPO_GO_GET" = true ] && [ "$TOPO_GO_SERVER" != true ]
then
	cd /src && tail -f /dev/null
fi

if [ "$TOPO_GO_VENDOR" = true ]
then
	echo "STARTING VENDORING"
	cd /src && go mod vendor
fi

if [ "$TOPO_GO_SERVER" = true ] && [ "$TOPO_GO_DEBUG" = true ]
then
	echo "STARTING SEVER WITH DEBUG"
	cd /src && air && dlv --listen=:40000 --headless=true --check-go-version=false --api-version=2 --accept-multiclient exec ./tmp/main
	# cd /src && go build -gcflags="all=-N -l" -o /dst/server -mod vendor ./cmd && cd /dst && dlv --listen=:40000 --headless=true --check-go-version=false --api-version=2 --accept-multiclient exec ./server
fi

if [ "$TOPO_GO_SERVER" = true ] && [ "$TOPO_GO_DEBUG" != true ]
then
	echo "STARTING SEVER"
	cd /src && air
	# cd /src && go build -o /dst/server -mod vendor ./cmd && cd /dst && ./server
fi