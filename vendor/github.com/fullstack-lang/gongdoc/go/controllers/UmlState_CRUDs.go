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
var __UmlState__dummysDeclaration__ models.UmlState
var __UmlState_time__dummyDeclaration time.Duration

// An UmlStateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUmlState updateUmlState deleteUmlState
type UmlStateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UmlStateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUmlState updateUmlState
type UmlStateInput struct {
	// The UmlState to submit or modify
	// in: body
	UmlState *orm.UmlStateAPI
}

// GetUmlStates
//
// swagger:route GET /umlstates umlstates getUmlStates
//
// Get all umlstates
//
// Responses:
//    default: genericError
//        200: umlstateDBsResponse
func GetUmlStates(c *gin.Context) {
	db := orm.BackRepo.BackRepoUmlState.GetDB()

	// source slice
	var umlstateDBs []orm.UmlStateDB
	query := db.Find(&umlstateDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	umlstateAPIs := make([]orm.UmlStateAPI, 0)

	// for each umlstate, update fields from the database nullable fields
	for idx := range umlstateDBs {
		umlstateDB := &umlstateDBs[idx]
		_ = umlstateDB
		var umlstateAPI orm.UmlStateAPI

		// insertion point for updating fields
		umlstateAPI.ID = umlstateDB.ID
		umlstateDB.CopyBasicFieldsToUmlState(&umlstateAPI.UmlState)
		umlstateAPI.UmlStatePointersEnconding = umlstateDB.UmlStatePointersEnconding
		umlstateAPIs = append(umlstateAPIs, umlstateAPI)
	}

	c.JSON(http.StatusOK, umlstateAPIs)
}

// PostUmlState
//
// swagger:route POST /umlstates umlstates postUmlState
//
// Creates a umlstate
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: umlstateDBResponse
func PostUmlState(c *gin.Context) {
	db := orm.BackRepo.BackRepoUmlState.GetDB()

	// Validate input
	var input orm.UmlStateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create umlstate
	umlstateDB := orm.UmlStateDB{}
	umlstateDB.UmlStatePointersEnconding = input.UmlStatePointersEnconding
	umlstateDB.CopyBasicFieldsFromUmlState(&input.UmlState)

	query := db.Create(&umlstateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, umlstateDB)
}

// GetUmlState
//
// swagger:route GET /umlstates/{ID} umlstates getUmlState
//
// Gets the details for a umlstate.
//
// Responses:
//    default: genericError
//        200: umlstateDBResponse
func GetUmlState(c *gin.Context) {
	db := orm.BackRepo.BackRepoUmlState.GetDB()

	// Get umlstateDB in DB
	var umlstateDB orm.UmlStateDB
	if err := db.First(&umlstateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var umlstateAPI orm.UmlStateAPI
	umlstateAPI.ID = umlstateDB.ID
	umlstateAPI.UmlStatePointersEnconding = umlstateDB.UmlStatePointersEnconding
	umlstateDB.CopyBasicFieldsToUmlState(&umlstateAPI.UmlState)

	c.JSON(http.StatusOK, umlstateAPI)
}

// UpdateUmlState
//
// swagger:route PATCH /umlstates/{ID} umlstates updateUmlState
//
// Update a umlstate
//
// Responses:
//    default: genericError
//        200: umlstateDBResponse
func UpdateUmlState(c *gin.Context) {
	db := orm.BackRepo.BackRepoUmlState.GetDB()

	// Get model if exist
	var umlstateDB orm.UmlStateDB

	// fetch the umlstate
	query := db.First(&umlstateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.UmlStateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	umlstateDB.CopyBasicFieldsFromUmlState(&input.UmlState)
	umlstateDB.UmlStatePointersEnconding = input.UmlStatePointersEnconding

	query = db.Model(&umlstateDB).Updates(umlstateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the umlstateDB
	c.JSON(http.StatusOK, umlstateDB)
}

// DeleteUmlState
//
// swagger:route DELETE /umlstates/{ID} umlstates deleteUmlState
//
// Delete a umlstate
//
// Responses:
//    default: genericError
func DeleteUmlState(c *gin.Context) {
	db := orm.BackRepo.BackRepoUmlState.GetDB()

	// Get model if exist
	var umlstateDB orm.UmlStateDB
	if err := db.First(&umlstateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&umlstateDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
