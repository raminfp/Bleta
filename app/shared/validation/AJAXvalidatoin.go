package validation

import (
	"net/http"
)

// check request if request is AJAX return true else return false
func Is_Ajax(req *http.Request) bool {
	if req.Header.Get("X-Requested-With") == "XMLHttpRequest"{
		return true
	}
	return false
}
