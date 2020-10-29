module github.com/lgadban/gloo-access-logger

go 1.14

require (
	github.com/envoyproxy/go-control-plane v0.9.7
	github.com/golang/protobuf v1.4.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/solo-io/go-utils v0.16.5
	go.opencensus.io v0.21.0
	go.uber.org/zap v1.9.1
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	google.golang.org/grpc v1.27.0
)

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
