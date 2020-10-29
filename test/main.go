package main

import (
	"context"
	"log"

	v2 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
	alsv2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	log.Println("connected...")
	alsClient := alsv2.NewAccessLogServiceClient(conn)
	stream, err := alsClient.StreamAccessLogs(context.Background())
	if err != nil {
		log.Fatalf("err creating stream: %v", err)
	}

	accessLog := alsv2.StreamAccessLogsMessage{
		LogEntries: &alsv2.StreamAccessLogsMessage_HttpLogs{
			HttpLogs: &alsv2.StreamAccessLogsMessage_HTTPAccessLogEntries{
				LogEntry: []*v2.HTTPAccessLogEntry{
					&v2.HTTPAccessLogEntry{
						CommonProperties: &v2.AccessLogCommon{
							UpstreamCluster: "foo",
						},
						Request: &v2.HTTPRequestProperties{
							RequestHeaders: map[string]string{
								"x-user-id": "foo",
							},
						},
					},
				},
			},
		},
	}

	err = stream.Send(&accessLog)
	if err != nil {
		log.Fatalf("err sending accessLog: %v", err)
	}
	log.Println("message sent")
}
