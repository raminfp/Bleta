package renders

import (
	"html/template"
	"log"
	"net/http"
	"github.com/raminfp/Bleta/middleware/security"
	"github.com/raminfp/Bleta/settings"
)

type Context struct {
	Temppath string
	Data interface{}
}

func (objtext Context) Render(w http.ResponseWriter) error {

	SecureHeaders(w)
	var indexTemplate = template.Must(template.ParseFiles(objtext.Temppath))
	if err := indexTemplate.Execute(w, objtext.Data); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SecureHeaders(w http.ResponseWriter) bool {

	if settings.StatusSecureHeader == true {
		header := security.SecureConfig{
			XSSProtection:      "1; mode=block",
			ContentTypeNosniff: "nosniff",
			XFrameOptions:      "SAMEORIGIN",
		}
		w.Header().Set("x-frame-options",header.XFrameOptions)
		w.Header().Set("x-content-type-options",header.ContentTypeNosniff)
		w.Header().Set("x-xss-protection",header.XSSProtection)
		return true
	}
	return false
}
