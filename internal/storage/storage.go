package storage

import "errors"

var (
	ErrTransactionNotFound = errors.New("payment transaction not found")
	ErrTransactionExists   = errors.New("payment transaction already exists")
)
