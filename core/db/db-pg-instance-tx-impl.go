/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package db

import (
	"database/sql"
	"log"
	"math/rand"
	"time"
)

func (slf *pgInstanceTx) Commit() (err error) {
	if slf == nil {
		return nil
	}

	if slf.isCommit || slf.isRollback || slf.tx == nil {
		return nil
	}

	err = slf.tx.Commit()
	slf.isCommit = err == nil
	return
}

func (slf *pgInstanceTx) Rollback() {
	const (
		min    int64 = 100
		max    int64 = 300
		maxTry       = 3
	)

	var err error

	if slf == nil {
		return
	}

	if slf.isCommit || slf.isRollback || slf.tx == nil {
		if slf.errCommit != nil && slf.isCommit {
			log.Printf("db.tx.rollback: isCommit = %v\n", slf.isCommit)
		}

		if slf.isRollback {
			log.Printf("db.tx.rollback: isRollback = %v\n", slf.isRollback)
		}

		return
	}

	iteration := -1
	for {
		iteration++
		err = slf.tx.Rollback()

		if err == nil {
			break
		}

		if err == sql.ErrTxDone {
			err = nil
			break
		}

		if iteration >= maxTry {
			break
		}

		time.Sleep(time.Microsecond * time.Duration(rand.Int63n(max-min)+min))
	}

	slf.isRollback = err == nil
}
