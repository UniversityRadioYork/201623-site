package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/UniversityRadioYork/2016-site/models"
	"github.com/UniversityRadioYork/2016-site/structs"
	"github.com/UniversityRadioYork/myradio-go"
)

type ShowGridController struct {
	session *myradio.Session
	config  *structs.Config
}

func NewShowGridController(session *myradio.Session, config *structs.Config) *ShowGridController {
	return &ShowGridController{
		session: session,
		config:  config,
	}
}

func (c *ShowGridController) Get(w http.ResponseWriter, r *http.Request) {
	// Fetch the seasons
	seasons, err := models.GetAllSeasonsForTerm(c.session)
	if err != nil {
		log.Println("Error fetching seasons:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Prepare the data for the template
	data := struct {
		PageTitle string
		Seasons   []*myradio.Season
	}{
		PageTitle: "Shows Grid",
		Seasons:   seasons,
	}

	// Parse and render the template
	tmpl, err := template.ParseFiles("views/shows_grid.tmpl")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "shows_grid", data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
