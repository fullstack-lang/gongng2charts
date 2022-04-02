package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongng2charts/go/orm"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongng2charts/go")
	{ // insertion point for registrations
		v1.GET("/v1/chartconfigurations", GetChartConfigurations)
		v1.GET("/v1/chartconfigurations/:id", GetChartConfiguration)
		v1.POST("/v1/chartconfigurations", PostChartConfiguration)
		v1.PATCH("/v1/chartconfigurations/:id", UpdateChartConfiguration)
		v1.PUT("/v1/chartconfigurations/:id", UpdateChartConfiguration)
		v1.DELETE("/v1/chartconfigurations/:id", DeleteChartConfiguration)

		v1.GET("/v1/datapoints", GetDataPoints)
		v1.GET("/v1/datapoints/:id", GetDataPoint)
		v1.POST("/v1/datapoints", PostDataPoint)
		v1.PATCH("/v1/datapoints/:id", UpdateDataPoint)
		v1.PUT("/v1/datapoints/:id", UpdateDataPoint)
		v1.DELETE("/v1/datapoints/:id", DeleteDataPoint)

		v1.GET("/v1/datasets", GetDatasets)
		v1.GET("/v1/datasets/:id", GetDataset)
		v1.POST("/v1/datasets", PostDataset)
		v1.PATCH("/v1/datasets/:id", UpdateDataset)
		v1.PUT("/v1/datasets/:id", UpdateDataset)
		v1.DELETE("/v1/datasets/:id", DeleteDataset)

		v1.GET("/v1/labels", GetLabels)
		v1.GET("/v1/labels/:id", GetLabel)
		v1.POST("/v1/labels", PostLabel)
		v1.PATCH("/v1/labels/:id", UpdateLabel)
		v1.PUT("/v1/labels/:id", UpdateLabel)
		v1.DELETE("/v1/labels/:id", DeleteLabel)

		v1.GET("/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
