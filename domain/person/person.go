package person

import (
	"encoding/json"
	"fmt"
	"io"
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
			err = createEmptyFile(dbFilePath)
			if err != nil {
				return Service{}, err
			}
			return Service{
				dbFilePath: dbFilePath,
				people: domain.People{},
			}, nil
		}
	}

	jsonFile, err := os.Open(dbFilePath)
	if err != nil {
		return Service{}, fmt.Errorf("Error trying to open file that contains all people: %s", err.Error())
	}

	jsonFileContentByte, err := io.ReadAll(jsonFile)
	if err != nil {
		return Service{}, fmt.Errorf("Error trying to read file: %s", err.Error())
	}

	var allPeople domain.People
	json.Unmarshal(jsonFileContentByte, &allPeople)

	return Service{
		dbFilePath: dbFilePath,
		people: allPeople,
	}, nil
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

func (s Service) Create(person domain.Person) error {

	return nil
}

func (s Service) exists(person domain.Person) bool {
	for index, currentPerson := range
	return false
}