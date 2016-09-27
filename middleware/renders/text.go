package renders

import (
	"net/http"
	"fmt"
)
type String struct  {
	Message string
}

func (str String) HttpResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w,str.Message)
}