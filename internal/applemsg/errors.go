package applemsg

import "errors"

var (
	ErrDatabaseUnavailable = errors.New(
		"messages database unavailable",
	)

	ErrPermissionDenied = errors.New(
		"messages database permission denied",
	)
)
