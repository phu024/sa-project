package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phu024/sa-64/controller"
	"github.com/phu024/sa-64/entity"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	//Gender Routes
	r.GET("/genders", controller.ListGenders)
	r.GET("/gender/:id", controller.GetGender)
	r.POST("/genders", controller.CreateGender)
	r.PATCH("/genders", controller.UpdateGender)
	r.DELETE("/genders/:id", controller.DeleteGender)

	// allergy Routes
	r.GET("/allergys", controller.ListAllergys)
	r.GET("/allergy/:id", controller.GetAllergy)
	r.POST("/allergys", controller.CreateAllergy)
	r.PATCH("/allergys", controller.UpdateAllergy)
	r.DELETE("/allergys/:id", controller.DeleteAllergy)

	// underlying_disease Routes
	r.GET("/underlying_diseases", controller.ListUnderlying_diseases)
	r.GET("/underlying_disease/:id", controller.GetUnderlying_disease)
	r.POST("/underlying_diseases", controller.CreateUnderlying_disease)
	r.PATCH("/underlying_diseases", controller.UpdateUnderlying_disease)
	r.DELETE("/underlying_diseases/:id", controller.DeleteUnderlying_disease)

	// recorder Routes
	r.GET("/recorders", controller.ListRecorder)
	r.GET("/recorder/:id", controller.GetRecorder)
	r.POST("/recorders", controller.CreateRecorder)
	r.PATCH("/recorders", controller.UpdateRecorder)
	r.DELETE("/recorders/:id", controller.DeleteRecorder)

	// patient Routes
	r.GET("/patients", controller.ListPatients)
	r.GET("/patient/:id", controller.GetPatient)
	r.POST("/patients", controller.CreatePatient)
	r.PATCH("/patients", controller.UpdatePatient)
	r.DELETE("/patients/:id", controller.DeletePatient)

	r.Run()
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
