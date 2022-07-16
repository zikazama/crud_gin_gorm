package main

import (
	"crud_gin_gorm/config"
	routes_student "crud_gin_gorm/routes/student"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()

	router.GET("/", index)

	// route user
	router.POST("/student", routes_student.StoreStudent)
	router.GET("/student", routes_student.ReadStudent)
	router.GET("/student/:id", routes_student.ReadStudent)
	router.PUT("/student/:id", routes_student.UpdateStudent)
	router.DELETE("/student/:id", routes_student.DeleteStudent)

	router.Run()
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Halo ini API gin with gorm by Fauzi",
	})
}
