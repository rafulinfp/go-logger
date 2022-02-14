package main

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/rafulinfp/gologger/dto"
	"github.com/rafulinfp/gologger/models"
)

var secret = "mysecret"

var dao = dto.LogDAO{}

func init() {
	dao.Connect()
}

func main() {
	route := gin.Default()

	// configure routes
	v1 := route.Group("/api/v1/log")
	{
		v1.GET("/", getLogs)
		v1.GET("/:id", getLog)
		v1.POST("/", addLog)
	}
	route.Run()
}

func getLog(c *gin.Context) {
	// prepare dummy data
	var data = models.LogEntry{ID: "1", Type: "info", Message: "all good", Timestamp: time.Now()}

	// populate response.
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "data": data})
	return
}

// getLogs ... get all logs
func getLogs(c *gin.Context) {

	// return unauthorized if no secret.
	if c.GetHeader("secret") != secret {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
		return
	}

	// prepare dummy data
	slice, _ := dao.FindAll()

	// populate response.
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": slice})
	return
}

// addLog ... save log entry into mongo db.
func addLog(c *gin.Context) {

	// return unauthorized if no secret.
	if c.GetHeader("secret") != secret {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
		return
	}

	var entry models.LogEntry
	// bind body to 'entry'
	if c.ShouldBind(&entry) == nil {
		// add the current datetime
		entry.Timestamp = time.Now()
		entry.ID = bson.NewObjectId()
		if err := dao.Insert(entry); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Failed to log entry", "resourceId": entry.ID})
			return
		}

		// prepare response
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Log entry saved", "resourceId": entry.ID})
	} else {
		c.JSON(http.StatusBadRequest, entry)
	}
	return
}
