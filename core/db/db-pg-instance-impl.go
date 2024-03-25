/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 * All Rights Reserved.
 */

package db

import (
	"errors"

	"github.com/andypangaribuan/gmod/gm"
)

func (slf *pgInstance) crw() (*srConnection, error) {
	connWriteLocking.Lock()
	defer connWriteLocking.Unlock()

	if slf.rw.sx == nil {
		sx, err := createConnection(slf.rw)
		if err != nil {
			return nil, err
		}

		slf.rw.sx = sx
	}

	return slf.rw, nil
}

func (slf *pgInstance) cro() (*srConnection, error) {
	connReadLocking.Lock()
	defer connReadLocking.Unlock()

	if slf.ro == nil {
		return nil, errors.New("database configuration doesn't have read only connection")
	}

	if slf.ro.sx == nil {
		sx, err := createConnection(slf.ro)
		if err != nil {
			return nil, err
		}

		slf.ro.sx = sx
	}

	return slf.ro, nil
}

func (slf *pgInstance) Ping() (string, error) {
	conn, err := slf.crw()
	if err != nil {
		return conn.conf.Host, err
	}

	return conn.conf.Host, conn.sx.Ping()
}

func (slf *pgInstance) PingRead() (string, error) {
	conn, err := slf.cro()
	if err != nil {
		return conn.conf.Host, err
	}

	return conn.conf.Host, conn.sx.Ping()
}

func (slf *pgInstance) Execute(query string, args ...interface{}) (string, error) {
	conn, err := slf.crw()
	if err != nil {
		return conn.conf.Host, err
	}

	startTime := gm.Util.Timenow()
	_, err = execute(conn, nil, query, args...)
	printSql(conn, startTime, query, args...)

	return conn.conf.Host, err
}

func (slf *pgInstance) ExecuteRID(query string, args ...interface{}) (*int64, string, error) {
	conn, err := slf.crw()
	if err != nil {
		return nil, conn.conf.Host, err
	}

	startTime := gm.Util.Timenow()
	id, _, err := executeRID(conn, nil, query, args...)
	printSql(conn, startTime, query, args...)

	return id, conn.conf.Host, err
}
