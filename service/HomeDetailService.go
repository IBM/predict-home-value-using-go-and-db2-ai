package service

import (
	"db2wml/dao"
	"db2wml/models"
	"log"
)

// GetPredictions : returns prediciton result
func GetPredictions(homeDetail *models.HomeDetail) (*models.HomeDetail, error) {
	//drop the out table.
	err := dao.DropPredictionOutputTable()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Finished dropping result table.")
	id, saveErr := dao.SaveHomeDetail(homeDetail)
	if saveErr != nil {
		log.Fatal(saveErr)
		return nil, saveErr
	}
	log.Println("Finished saving home data with id: ", id)
	predictErr := dao.PredictHomeSalePrice()
	if predictErr != nil {
		log.Fatal(predictErr)
		return nil, predictErr
	}
	log.Println("Finished predicting the results.")
	result, resErr := dao.GetPredicitonResult(id)
	if resErr != nil {
		log.Fatal(resErr)
		return nil, resErr
	}
	log.Println("Finished getting prediction results")

	//create Homedetail struct from map.

	homeDetailStruct := &models.HomeDetail{}
	strErr := homeDetailStruct.FillStruct(result)
	if strErr != nil {
		log.Fatal(strErr)
		return nil, strErr
	}

	return homeDetailStruct, nil
}
