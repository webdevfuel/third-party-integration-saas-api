package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.Get("/integrations/{id}/tags", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		app, err := integration.GetIntegrationApp(id, conn)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if app == "activecampaign" {
			activeCampaignIntegration, _ := integration.NewActiveCampaignIntegration(id, conn)
			tags, err := integration.GetIntegrationTags(activeCampaignIntegration)
			if err != nil {
				log.Println(err)
			}
			j, _ := json.Marshal(tags)
			w.Write(j)
			return
		}
		if app == "convertkit" {
			convertKitIntegration, _ := integration.NewConvertKitIntegration(id, conn)
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
