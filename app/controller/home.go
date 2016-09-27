package controller


import (
	"net/http"
	//"fmt"
	//"log"
	//"github.com/raminfp/go_rest_api/middleware/renders"
	//"encoding/json"
	"github.com/raminfp/Bleta/middleware/renders"
	"github.com/raminfp/Bleta/middleware/csrf"
	"github.com/raminfp/Bleta/middleware/session"

)


type Index struct {
	Title string
	Body  string
	CSRF_Token string
}

// Action Index
func IndexController(w http.ResponseWriter , r *http.Request){



	samin_token, err := csrf.GenerateRandomString(46)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
	}
	// Get session
	sess := session.NewSession(r)
	sess.Values["samin_token"] = "test"

	if r.Method == "GET"{

		data := &Index{
			Title: "Hello",
			Body:  "Are you ok?",
			CSRF_Token: samin_token,

		}
		sess.Save(r,w)
		html := renders.Context{Temppath:"templates/index.tmpl",Data:data}
		html.Render(w)
	}

	// Define path for Redirect
	//path := renders.RedirectTo{Request:r,StatusCode:302,URLString:"/ramin"}
	//path.RenderRedirect(w)

	// Define render Http Message
	//text := renders.String{Message:"Hello"}
	//text.HttpResponse(w)

	// Define JSON render
	//json := renders.Json{Message:"ramin"}
	//json.JsonResponse(w)
}
