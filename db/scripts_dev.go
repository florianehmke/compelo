// +build dev

package db

import (
	"net/http"
)

var Scripts = http.Dir("db/scripts")
