package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/webdevfuel/third-party-integration-saas-api/db"
	"github.com/webdevfuel/third-party-integration-saas-api/integration"
)

func main() {
	conn, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	r.Get("/tags/{app}", func(w http.ResponseWriter, r *http.Request) {
		app := chi.URLParam(r, "app")
		if app == "activecampaign" {
			activeCampaignIntegration := integration.NewActiveCampaignIntegration()
			tags, err := integration.GetIntegrationTags(activeCampaignIntegration)
			if err != nil {
				log.Println(err)
			}
			j, _ := json.Marshal(tags)
			w.Write(j)
			return
		}
		if app == "convertkit" {
			convertKitIntegration := integration.NewConvertKitIntegration()
			tags, err := integration.GetIntegrationTags(convertKitIntegration)
			if err != nil {
				log.Println(err)
			}
			j, _ := json.Marshal(tags)
			w.Write(j)
			return
		}
	})
	http.ListenAndServe("localhost:8080", r)
}
