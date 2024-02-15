package models

import (
	"fmt"
	"log"

	"github.com/UniversityRadioYork/myradio-go"
)

func GetAllSeasonsForTerm(session *myradio.Session) ([]*myradio.Season, error) {
	// Your implementation here
}

func main() {
	// Replace "your-api-key" with your actual MyRadio API key
	session, err := myradio.NewSession("uctarCLCbU7Br9ZGMXBy2uZssxf2K648")
	if err != nil {
		log.Fatal("Failed to create MyRadio session:", err)
	}

	// Test the GetAllSeasonsForTerm function
	seasons, err := GetAllSeasonsForTerm(session)
	if err != nil {
		log.Fatal("Failed to get all seasons for term:", err)
	}

	// Print the fetched seasons
	for _, season := range seasons {
		fmt.Println(*season)
	}
}
