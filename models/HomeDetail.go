package models

import (
	"errors"
	"fmt"
	"reflect"
)

type HomeDetail struct {
	LotArea      int    `json:"lotArea"`
	BldgType     string `json:"bldgType"`
	HouseStyle   string `json:"houseStyle"`
	OverallCond  int    `json:"overallCond"`
	YearBuilt    int    `json:"yearBuilt"`
	RoofStyle    string `json:"roofStyle"`
	ExterCond    string `json:"exterCond"`
	Foundation   string `json:"foundation"`
	BsmtCond     string `json:"bsmtCond"`
	Heating      string `json:"heating"`
	HeatingQC    string `json:"heatingQC"`
	CentralAir   string `json:"centralAir"`
	Electrical   string `json:"electrical"`
	FullBath     int    `json:"fullBath"`
	HalfBath     int    `json:"halfBath"`
	BedroomAbvGr int    `json:"bedroomAbvGr"`
	KitchenAbvGr int    `json:"kitchenAbvGr"`
	KitchenQual  string `json:"kitchenQual"`
	TotRmsAbvGrd int    `json:"totRmsAbvGrd"`
	Fireplaces   int    `json:"fireplaces"`
	FireplaceQu  string `json:"fireplaceQu"`
	GarageType   string `json:"garageType"`
	GarageFinish string `json:"garageFinish"`
	GarageCars   int    `json:"garageCars"`
	GarageCond   string `json:"garageCond"`
	PoolArea     int    `json:"poolArea"`
	PoolQC       string `json:"poolQC"`
	Fence        string `json:"fence"`
	MoSold       int    `json:"moSold"`
	YrSold       int    `json:"yrSold"`
	SalePrice    string `json:"salePrice"`
}

type HomeDetails []HomeDetail

// SetField : set fields of struct
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

// FillStruct : fills the stuct with values
func (s *HomeDetail) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
