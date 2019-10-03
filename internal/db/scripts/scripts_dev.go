// +build dev

package scripts

import (
	"net/http"
)

var Scripts = http.Dir("internal/db/scripts")
