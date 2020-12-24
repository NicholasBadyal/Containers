package main

import (
	"github.com/NicholasBadyal/Containers/rediboard/v1/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	ListenAddr = "0.0.0.0:8080"
	RedisAddr  = "redis-server:6379"
	//RedisAddr  = "localhost:6379"
)

func initRouter(database *db.Database) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home_page", gin.H{})
	})

	r.POST("/points", func(c *gin.Context) {
		_, err := c.MultipartForm()
		if err != nil {
			log.Println("failed to get form: ", err)
		}

		username := c.PostForm("username")
		points, err := strconv.Atoi(c.PostForm("points"))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "internal_error", gin.H{"Error": "failed to parse score"})
			return
		}

		user := &db.User{
			Username: username,
			Points:   points,
		}
		err = database.SaveUser(user)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "internal_error", gin.H{"Error": "failed to save user"})
			return
		}

		c.HTML(http.StatusOK, "home_page", user)
	})

	r.GET("/points/:username", func(c *gin.Context) {
		username := c.Param("username")
		user, err := database.GetUser(username)
		if err != nil {
			if err.Error() == db.ErrNil.Error() {
				c.HTML(http.StatusNotFound, "internal_error", gin.H{"Error": "No record found for " + username})
				return
			}
			c.HTML(http.StatusInternalServerError, "internal_error", gin.H{"Error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "home_page", user)
	})

	r.GET("/leaderboard", func(c *gin.Context) {
		leaderboard, err := database.GetLeaderboard()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "internal_error", gin.H{"Error": err})
			return
		}
		c.HTML(http.StatusOK, "leaderboard", leaderboard.Users)
	})

	r.GET("/crash", func(c *gin.Context) {
		os.Exit(0)
	})

	return r
}

func main() {
	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err.Error())
	}

	router := initRouter(database)
	if err = router.Run(ListenAddr); err != nil {
		log.Fatalf("failed to start server: %v", err.Error())
	}
	log.Println("rediboard server start")
}
