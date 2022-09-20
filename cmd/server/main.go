package main

import (
	"database/sql"

	"github.com/iamstivgo/integrador-back-go.git/cmd/server/handler"
	"github.com/iamstivgo/integrador-back-go.git/internal/dentist"
	"github.com/iamstivgo/integrador-back-go.git/internal/patient"
	"github.com/iamstivgo/integrador-back-go.git/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/integrador")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storageDentist := store.NewDentistSqlStore(db)
	repoDentist := dentist.NewRepository(storageDentist)
	serviceDentist := dentist.NewService(repoDentist)
	handlerDentist := handler.dentistHandler.NewDentistHandler(serviceDentist)

	storagePatient := store.NewPatientSqlStore(db)
	repoPatient := patient.NewRepository(storagePatient)
	servicePatient := patient.NewService(repoPatient)
	handlerPatient := handler.patientHandler.NewPatientHandler(servicePatient)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
	dentists := router.Group("/dentists")
	{
		dentists.GET("/:id", handlerDentist.GetByID())
		dentists.POST("/", handlerDentist.Post())
		dentists.PUT("/:id", handlerDentist.Put())
		dentists.DELETE("/:id", handlerDentist.Delete())
	}

	patients := router.Group("/patients")
	{
		patients.GET("/:id", handlerPatient.GetByID())
		patients.POST("/", handlerPatient.Post())
		patients.PUT("/:id", handlerPatient.Put())
		patients.DELETE("/:id", handlerPatient.Delete())
	}

	router.Run(":8080")
}