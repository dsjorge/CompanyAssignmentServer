package main

import (
	"log"
	"net/http"

	"company/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/company", controllers.GetAll)
	router.GET("/api/v1/resume", controllers.GetAllResume)
	router.GET("/api/v1/company/:id", controllers.GetById)
	router.GET("/api/v1/available/:id", controllers.GetAvailableOwners)
	router.POST("/api/v1/company", controllers.AddCompany)
	router.PUT("/api/v1/company", controllers.UpdateCompany)
	log.Fatal(http.ListenAndServe(":8080", router))
}
