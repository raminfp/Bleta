package origin

import (
	"net/http"
)

type StatusOrigin struct {
	Active bool
}

func (OriginStatus StatusOrigin) OriginDomain(rw http.ResponseWriter, req *http.Request) bool {
	if OriginStatus.Active == true {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	return false
}
