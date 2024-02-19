package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		http.Post("https://api.vercel.com/v1/integrations/deploy/prj_2bDFRwtPnefmEsrGNZmTC4yZniaz/FB0sf3qqqI", "application/json", nil)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
