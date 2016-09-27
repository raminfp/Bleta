package renders

import (
	"net/http"
	"fmt"
)

type RedirectTo struct {
	StatusCode     int
	Request  *http.Request
	URLString string
}

type errors struct {
	InvalidStatusCode int
}


func (r RedirectTo) RenderRedirect(w http.ResponseWriter) error  {

	if (r.StatusCode < 300 || r.StatusCode > 308) && r.StatusCode != 201 {
		panic(fmt.Sprintf("Error! Invalid StatusCode %d", r.StatusCode))
	}
	http.Redirect(w, r.Request, r.URLString, r.StatusCode)
	return nil
}

func (e *errors) Error() int {
	return e.InvalidStatusCode
}
