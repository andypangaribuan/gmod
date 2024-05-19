/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package net

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/mol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func (*stuNet) IsPortUsed(port int, host ...string) bool {
	var (
		targetHost = "127.0.0.1"
		timeout    = time.Second * 3
	)

	if len(host) > 0 && host[0] != "" {
		targetHost = host[0]
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(targetHost, strconv.Itoa(port)), timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}

func (*stuNet) GrpcConnection(address string, opt ...mol.NetOpt) (grpc.ClientConnInterface, error) {
	var (
		cred    = insecure.NewCredentials()
		dialOpt = make([]grpc.DialOption, 0)
	)

	ls := strings.Split(address, ":")
	if len(ls) != 2 {
		return nil, errors.New("invalid grpc address")
	}

	port, err := strconv.Atoi(ls[1])
	if err != nil {
		return nil, errors.New("port must be number")
	}

	if port != 443 && port < 1000 {
		return nil, errors.New("port cannot less than 1000")
	}

	if port == 443 {
		cred = credentials.NewTLS(&tls.Config{})
	}

	if len(opt) > 0 {
		o := opt[0]

		if o.UsingClientLB != nil {
			resolver.SetDefaultScheme("dns")
			address = fmt.Sprintf("dns:///%v", address)
			dialOpt = append(dialOpt, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
		}

		if o.TlsConfig != nil {
			cred = credentials.NewTLS(o.TlsConfig)
		}

		if cred == nil && o.TlsCertFile != nil && o.TlsServerNameOverride != nil {
			c, err := credentials.NewClientTLSFromFile(*o.TlsCertFile, *o.TlsServerNameOverride)
			if err != nil {
				return nil, err
			}

			cred = c
		}
	}

	dialOpt = append(dialOpt, grpc.WithTransportCredentials(cred))

	conn, err := grpc.Dial(address, dialOpt...)
	if err != nil {
		return nil, err
	}

	//make sure we connected to the grpc service
	c, err := net.DialTimeout("tcp", address, time.Second*5)
	if err != nil {
		return nil, err
	}
	c.Close()

	return conn, nil
}
