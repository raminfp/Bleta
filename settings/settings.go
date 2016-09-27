package settings

import (
	"flag"
	"github.com/gorilla/sessions"
	"github.com/raminfp/Bleta/middleware/session"
)

// Define Static Path on your projects

const staticpathfile string =  "D:/Github Project/Go Api/packages/src/github.com/raminfp/go_rest_api/static"

// define path project static files (image , javascript , css , media)
func StaticFilesDirs() string  {
	var staticpathdir string
	flag.StringVar(&staticpathdir, "dir", staticpathfile , "")
	flag.Parse()
	return staticpathdir
}

// Session Configuration

func ConfigSessions() {

	option := sessions.Options{Path: "/" , HttpOnly:true}
	ses := session.Session{Options:option,Name:"session_id",SecretKey:"123456"}
	session.Configure(ses)
}


// click jacking , ContentTypeNosniff protection
var StatusSecureHeader bool = true