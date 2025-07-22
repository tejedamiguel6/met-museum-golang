package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"www.example.com/met/models"
)

func GetCollectionObjects() (models.CollectionObjects, error) {
	req, err := http.NewRequest("GET",
		"https://collectionapi.metmuseum.org/public/collection/v1/objects", nil)
	if err != nil {
		fmt.Println("error making request")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error making request")
	}
	defer res.Body.Close()

	var collectionObjectsResponse models.CollectionObjects

	if err := json.NewDecoder(res.Body).Decode(&collectionObjectsResponse); err != nil {
		fmt.Println("new error in decoder", err)
	}

	ids := collectionObjectsResponse.ObjectIDs
	if len(ids) > 5 {
		collectionObjectsResponse.ObjectIDs = ids[:5]
	}
	fmt.Println("collectionObjectsResponse:", collectionObjectsResponse)
	return collectionObjectsResponse, nil

}

func GetCollectionObjectItem(id int) (*models.CollectionObjectItem, error) {
	url := fmt.Sprintf("https://collectionapi.metmuseum.org/public/collection/v1/objects/%d", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error at GETCOLLECTION OBJECT:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error at GETCOLLECTION OBJECT:", err)
	}
	defer res.Body.Close()

	var collectionObjectResponse models.CollectionObjectItem

	if err := json.NewDecoder(res.Body).Decode(&collectionObjectResponse); err != nil {
		fmt.Println("new error in decoder", err)
	}

	fmt.Println("id:", id)
	fmt.Println("RESULTS-->:", &collectionObjectResponse)

	return &collectionObjectResponse, nil

}
