// +build dev

package db

import (
	"net/http"
)

var Scripts = http.Dir("internal/db/scripts")
