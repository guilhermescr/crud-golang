package person

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/guilhermescr/crud-golang/domain"
)

type Service struct {
	dbFilePath string
	people     domain.People
}

func NewService(dbFilePath string) (Service, error) {
	// check if file exists
	 _, err := os.Stat(dbFilePath)
	 if err != nil {
		if os.IsNotExist(err) {
			// create an empty file
		}
	 }

	// if it doesn't, create empty file
	// if it does, read the file and update the variable people from Service with people from the file
}

func createEmptyFile(dbFilePath string) error {
	var people domain.People = domain.People{
		People: []domain.Person{},
	}

	peopleJSON, err := json.Marshal(people)
	if err != nil {
		return fmt.Errorf("Error trying to encode people as JSON: %s", err.Error())
	}

	err = os.WriteFile(dbFilePath, peopleJSON, 0755)
	if err != nil {
		return fmt.Errorf("Error trying to write file. Error: %s", err.Error())
	}

	return nil
}