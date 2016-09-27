package main

import (
	"net/http"
	"fmt"
	"github.com/raminfp/Bleta/route"
	"log"
	"time"
	"github.com/raminfp/Bleta/settings"
)

func main() {
	settings.ConfigSessions()
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP ")
	log.Fatal(http.ListenAndServe(":8000",route.Url()))
}
