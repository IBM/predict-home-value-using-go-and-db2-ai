package dao

import (
	"db2wml/constants"
	"db2wml/models"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// DropPredictionOutputTable : Drop prediction output table
func DropPredictionOutputTable() error {
	m, tableExistError := store.GetData(constants.TABLE_IF_EXISTS, true)
	if tableExistError != nil {
		fmt.Println(fmt.Errorf("Error: %v", tableExistError))
		return tableExistError
	}

	if m == nil || len(m) == 0 {
		log.Info("Table doesn't exist. No need to drop.")
		return nil
	}

	error := store.DropTable(constants.DropOutputTable)
	if error != nil {
		fmt.Println(fmt.Errorf("Error: %v", error))
		return error
	}
	return nil
}

// SaveHomeDetail : Save home detail for prediction
func SaveHomeDetail(homeDetail *models.HomeDetail) (int32, error) {
	id, error := GetMaxID()
	if error != nil {
		return 0, error
	}
	log.Info("Max ID: ", id)
	err := store.InsertData(constants.INSERT_SQL, id, homeDetail)
	if err != nil {
		log.Error("Error while saving data: ", err)
		return 0, err
	}
	return id, nil
}

// PredictHomeSalePrice :
func PredictHomeSalePrice() error {
	_, error := store.GetData(constants.PREDICT_LR_SQL, false)
	if error != nil {
		return error
	}
	return nil
}

// GetMaxID : get max id from the table.
func GetMaxID() (int32, error) {
	m, error := store.GetData(constants.MAXID, true)
	if error != nil {
		log.Error("Error while getting max id: ", error)
		return 0, error
	}
	newID := m["MAXID"].(int32)
	return newID, nil
}

// GetPredicitonResult : returns prediciton result
func GetPredicitonResult(id int32) (map[string]interface{}, error) {
	m, error := store.GetPredictionResult(constants.PREDICTION_RESULT, id)
	if error != nil {
		return nil, error
	}
	return m, nil
}
