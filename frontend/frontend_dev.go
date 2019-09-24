// +build dev

package frontend

import (
	"net/http"
)

var Frontend = http.Dir("frontend/compelo/dist")
