package renders

import (
	"net/http"
	"encoding/json"
)

type Json struct {
	Message string
}

func (str Json) JsonResponse(w http.ResponseWriter) error{
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonBytes, err := json.MarshalIndent(str.Message, "", "")
	if err != nil {
		return err
	}

	w.Write(jsonBytes)
	return nil
}


