package main

import (
	"os"

	router "github.com/Modifa/KO_TICKET_LOGGER/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	gin.DisableConsoleColor()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	// r.Use(CORSMiddleware())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// OPTIONS method for ReactJS
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "x-access-token", "content-type", "Content-Length", "Authorization", "Cache-Control"}
	config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Init(r)

	return r
}

func setupConfigs() {
	//NGrok For Testing Purposes
	os.Setenv("CURRENTDOMAIN", "https://3ef7-160-19-36-38.ngrok.io")

	//Reddis Details
	os.Setenv("REDISSERVER_HOST", "redis-19714.c124.us-central1-1.gce.cloud.redislabs.com")
	os.Setenv("REDISSERVER_PORT", "19714")
	os.Setenv("REDISSERVER_PASSWORD", "ULXGpAVRYk1G9tBxi9D4jkksGQLA7A9Q")

	//PostgresSQL Connection String
	os.Setenv("PostgresConnectionString","postgres://mxwkzkok:NSletRKB0o_eG6UBO-RZBMpyyOLiL5Xl@berry.db.elephantsql.com/mxwkzkok")
	// Server Port 	
	os.Setenv("WEBSERVER_PORT", "8080")

}

func main() {
	//Uncommented When Not Debugging
	// gin.SetMode(gin.ReleaseMode)
	// export GIN_MODE=release

	// gocron.Start()
	// s := gocron.NewScheduler()
	// gocron.Every(2).Seconds().Do(c.CheckNewUser)
	//  <-s.Start()

	r := setupRouter()

	setupConfigs()

	r.Run(":" + os.Getenv("WEBSERVER_PORT"))
}
