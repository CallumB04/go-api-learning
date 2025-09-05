package data

import (
	"encoding/json"
	"os"

	"github.com/callumb04/go-api-learning/internal/models"
)

// JSON File Paths
const usersJSONPath string = "data/users.json"

// Read users.json file and return a slice of User objects.
func LoadUsers() ([]models.User, error) {
	// Read users from file
	data, err := os.ReadFile(usersJSONPath)

	// Return nil if error loading users.
	if err != nil {
		return nil, err
	}

	// Create empty slice for users.
	var users []models.User

	// Parse raw JSON data into users slice.
	// Return nil if fails.
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}
