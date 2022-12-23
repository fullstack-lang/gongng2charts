// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Vertice__dummysDeclaration__ models.Vertice
var __Vertice_time__dummyDeclaration time.Duration

// An VerticeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVertice updateVertice deleteVertice
type VerticeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// VerticeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVertice updateVertice
type VerticeInput struct {
	// The Vertice to submit or modify
	// in: body
	Vertice *orm.VerticeAPI
}

// GetVertices
//
// swagger:route GET /vertices vertices getVertices
//
// # Get all vertices
//
// Responses:
// default: genericError
//
//	200: verticeDBResponse
func GetVertices(c *gin.Context) {
	db := orm.BackRepo.BackRepoVertice.GetDB()

	// source slice
	var verticeDBs []orm.VerticeDB
	query := db.Find(&verticeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	verticeAPIs := make([]orm.VerticeAPI, 0)

	// for each vertice, update fields from the database nullable fields
	for idx := range verticeDBs {
		verticeDB := &verticeDBs[idx]
		_ = verticeDB
		var verticeAPI orm.VerticeAPI

		// insertion point for updating fields
		verticeAPI.ID = verticeDB.ID
		verticeDB.CopyBasicFieldsToVertice(&verticeAPI.Vertice)
		verticeAPI.VerticePointersEnconding = verticeDB.VerticePointersEnconding
		verticeAPIs = append(verticeAPIs, verticeAPI)
	}

	c.JSON(http.StatusOK, verticeAPIs)
}

// PostVertice
//
// swagger:route POST /vertices vertices postVertice
//
// Creates a vertice
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostVertice(c *gin.Context) {
	db := orm.BackRepo.BackRepoVertice.GetDB()

	// Validate input
	var input orm.VerticeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create vertice
	verticeDB := orm.VerticeDB{}
	verticeDB.VerticePointersEnconding = input.VerticePointersEnconding
	verticeDB.CopyBasicFieldsFromVertice(&input.Vertice)

	query := db.Create(&verticeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoVertice.CheckoutPhaseOneInstance(&verticeDB)
	vertice := (*orm.BackRepo.BackRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID]

	if vertice != nil {
		models.AfterCreateFromFront(&models.Stage, vertice)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, verticeDB)
}

// GetVertice
//
// swagger:route GET /vertices/{ID} vertices getVertice
//
// Gets the details for a vertice.
//
// Responses:
// default: genericError
//
//	200: verticeDBResponse
func GetVertice(c *gin.Context) {
	db := orm.BackRepo.BackRepoVertice.GetDB()

	// Get verticeDB in DB
	var verticeDB orm.VerticeDB
	if err := db.First(&verticeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var verticeAPI orm.VerticeAPI
	verticeAPI.ID = verticeDB.ID
	verticeAPI.VerticePointersEnconding = verticeDB.VerticePointersEnconding
	verticeDB.CopyBasicFieldsToVertice(&verticeAPI.Vertice)

	c.JSON(http.StatusOK, verticeAPI)
}

// UpdateVertice
//
// swagger:route PATCH /vertices/{ID} vertices updateVertice
//
// # Update a vertice
//
// Responses:
// default: genericError
//
//	200: verticeDBResponse
func UpdateVertice(c *gin.Context) {
	db := orm.BackRepo.BackRepoVertice.GetDB()

	// Get model if exist
	var verticeDB orm.VerticeDB

	// fetch the vertice
	query := db.First(&verticeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.VerticeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	verticeDB.CopyBasicFieldsFromVertice(&input.Vertice)
	verticeDB.VerticePointersEnconding = input.VerticePointersEnconding

	query = db.Model(&verticeDB).Updates(verticeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	verticeNew := new(models.Vertice)
	verticeDB.CopyBasicFieldsToVertice(verticeNew)

	// get stage instance from DB instance, and call callback function
	verticeOld := (*orm.BackRepo.BackRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID]
	if verticeOld != nil {
		models.AfterUpdateFromFront(&models.Stage, verticeOld, verticeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the verticeDB
	c.JSON(http.StatusOK, verticeDB)
}

// DeleteVertice
//
// swagger:route DELETE /vertices/{ID} vertices deleteVertice
//
// # Delete a vertice
//
// default: genericError
//
//	200: verticeDBResponse
func DeleteVertice(c *gin.Context) {
	db := orm.BackRepo.BackRepoVertice.GetDB()

	// Get model if exist
	var verticeDB orm.VerticeDB
	if err := db.First(&verticeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&verticeDB)

	// get an instance (not staged) from DB instance, and call callback function
	verticeDeleted := new(models.Vertice)
	verticeDB.CopyBasicFieldsToVertice(verticeDeleted)

	// get stage instance from DB instance, and call callback function
	verticeStaged := (*orm.BackRepo.BackRepoVertice.Map_VerticeDBID_VerticePtr)[verticeDB.ID]
	if verticeStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, verticeStaged, verticeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
