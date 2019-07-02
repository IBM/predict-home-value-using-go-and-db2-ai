package main

import (
	"database/sql"
	"db2wml/dao"
	"db2wml/routers"

	_ "github.com/ibmdb/go_ibm_db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/opentracing/opentracing-go"
	// "github.com/opentracing/opentracing-go/ext"
	// "github.com/uber/jaeger-client-go"
	// jaegerprom "github.com/uber/jaeger-lib/metrics/prometheus"

	"os"

	log "github.com/sirupsen/logrus"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	//Connectionto the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	var host = os.Getenv("DB2_HOST")
	var db2port = os.Getenv("DB2_PORT")
	var database = os.Getenv("DB2_DBNAME")
	var username = os.Getenv("DB2_USER")
	var password = os.Getenv("DB2_PASS")

	con := "HOSTNAME=" + host + ";PORT=" + db2port + ";DATABASE=" + database + ";UID=" + username + ";PWD=" + password
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Info("Successfully connected to the database")
	// Initialize the DB2Store
	dao.InitStore(&dao.Db2Store{Db: db})

	router := gin.Default()
	router.RedirectTrailingSlash = false

	//CORS
	router.Use(CORSMiddleware())

	router.GET("/", routers.Index)
	router.NoRoute(routers.NotFoundError)
	router.GET("/500", routers.InternalServerError)
	router.GET("/health", routers.HealthGET)
	router.POST("/predict", routers.PredictPOST)

	log.Info("Starting DB2WML on port " + port())

	router.Run(port())
}
