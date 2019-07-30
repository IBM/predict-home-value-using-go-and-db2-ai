package routers

import (
	"db2wml/models"
	"db2wml/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PredictPOST API to get prediction result
func PredictPOST(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var t *models.HomeDetail
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res2B, _ := json.Marshal(t)
	fmt.Println(string(res2B))

	prediciton, error := service.GetPredictions(t)
	if error != nil {
		fmt.Println(fmt.Errorf("Error: %v", error))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, prediciton)

}
