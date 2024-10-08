/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func (slf *stuFuseGRouter) AutoRecover(autoRecover bool) {
	slf.withAutoRecover = autoRecover
}

func (slf *stuFuseGRouter) Server() *grpc.Server {
	return slf.fnGetServer()
}

// reference
// - https://github.com/grpc/grpc-go/blob/master/examples/features/health/server/main.go
// - https://github.com/hashicorp/go-plugin/blob/main/grpc_server.go
func (slf *stuFuseGRouter) RunHealthCheck() {
	// server := health.NewServer()
	// grpc_health_v1.RegisterHealthServer(slf.Server(), server)

	// go func() {
	// 	var (
	// 		duration = time.Second
	// 		next     = grpc_health_v1.HealthCheckResponse_SERVING
	// 	)

	// 	for {
	// 		server.SetServingStatus("", next)
	// 		if next == grpc_health_v1.HealthCheckResponse_SERVING {
	// 			next = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	// 		} else {
	// 			next = grpc_health_v1.HealthCheckResponse_SERVING
	// 		}

	// 		time.Sleep(duration)
	// 	}
	// }()

	server := health.NewServer()
	server.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(slf.Server(), server)
}
