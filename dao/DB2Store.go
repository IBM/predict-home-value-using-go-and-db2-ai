package dao

import (
	"database/sql"
	"db2wml/models"
	"fmt"
	"log"
)

// Store interface to run CRUD
// on the database
type Store interface {
	GetData(sqlStr string, getResults bool) (map[string]interface{}, error)
	GetPredictionResult(sqlStr string, id int32) (map[string]interface{}, error)
	InsertData(sqlStr string, maxID int32, homeDetail *models.HomeDetail) error
	DropTable(sqlStr string) error
}

type Db2Store struct {
	Db *sql.DB
}

// GetData : Runs SQL select query and returns a map
func (store *Db2Store) GetData(sqlStr string, getResults bool) (map[string]interface{}, error) {
	stmt, err := store.Db.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Error while running ", sqlStr, err)
		return nil, err
	}
	// if no result expected just return nil
	if !getResults {
		defer rows.Close()
		return nil, nil
	}

	m, err := getMapFromRows(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return m, nil
}

// GetPredictionResult : Runs SQL select query with argumetns and returns a map
func (store *Db2Store) GetPredictionResult(sqlStr string, id int32) (map[string]interface{}, error) {
	stmt, err := store.Db.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal("Error while running ", sqlStr, err)
		return nil, err
	}

	m, err := getMapFromRows(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return m, nil
}

// getMapFromRows : get Map from Rows
func getMapFromRows(rows *sql.Rows) (map[string]interface{}, error) {
	cols, _ := rows.Columns()
	m := make(map[string]interface{})
	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
	}
	return m, nil
}

// InsertData : Runs insert statments to store to be  preicted data
func (store *Db2Store) InsertData(sqlStr string, maxID int32, homeDetail *models.HomeDetail) error {
	stmt, err := store.Db.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Executin insert statement")
	res, err := stmt.Exec(
		maxID,
		homeDetail.LotArea,
		homeDetail.BldgType,
		homeDetail.HouseStyle,
		homeDetail.OverallCond,
		homeDetail.YearBuilt,
		homeDetail.RoofStyle,
		homeDetail.ExterCond,
		homeDetail.Foundation,
		homeDetail.BsmtCond,
		homeDetail.Heating,
		homeDetail.HeatingQC,
		homeDetail.CentralAir,
		homeDetail.Electrical,
		homeDetail.FullBath,
		homeDetail.HalfBath,
		homeDetail.BedroomAbvGr,
		homeDetail.KitchenAbvGr,
		homeDetail.KitchenQual,
		homeDetail.TotRmsAbvGrd,
		homeDetail.Fireplaces,
		homeDetail.FireplaceQu,
		homeDetail.GarageType,
		homeDetail.GarageFinish,
		homeDetail.GarageCars,
		homeDetail.GarageCond,
		homeDetail.PoolArea,
		homeDetail.PoolQC,
		homeDetail.Fence,
		homeDetail.MoSold,
		homeDetail.YrSold)

	if err != nil {
		return err
	}
	log.Print("Rows Affected")
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("affected = %d\n", rowCnt)
	return nil
}

// DropTable : drops table
func (store *Db2Store) DropTable(sqlStr string) error {
	_, err := store.Db.Exec(sqlStr)
	if err != nil {
		return err
	}
	fmt.Println("TABLE DROPPED")
	return nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

// InitStore We will need to call the InitStore method to initialize the store. This will
// typically be done at the beginning of our application (in this case, when the server starts up)
// This can also be used to set up the store as a mock, which we will be observing
// later on
func InitStore(s Store) {
	store = s
}
