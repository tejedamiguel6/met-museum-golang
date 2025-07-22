package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"www.example.com/met/models"
)

func CallMetAPI() ([]models.Department, error) {
	req, err := http.NewRequest("GET",
		"https://collectionapi.metmuseum.org/public/collection/v1/departments", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error calling api ", err)
	}

	defer res.Body.Close()
	// read the response body
	fmt.Println("API call completed")
	var departmentsResponse models.DepartmentsResponse

	if err := json.NewDecoder(res.Body).Decode(&departmentsResponse); err != nil {
		fmt.Println("new error in decoder", err)
	}
	// Just return the first department for now (for testing)
	fmt.Println("here are the results", departmentsResponse.Departments)
	return departmentsResponse.Departments, nil

}
