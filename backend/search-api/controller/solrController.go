package controller

import (
	"search-api/services"
)

func AddFromId(id string) error {

	err := services.SolrService.AddFromId(id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromId(id string) error {
	err := services.SolrService.DeleteFromId(id)
	if err != nil {
		return err
	}
	return nil
}

//654513064fd3abb838f8c48d
