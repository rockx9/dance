package instructor

import (
	"dance/models"
	"dance/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestDeleteInstructorHandle(t *testing.T) {
	var resp types.BaseResponse
	r := SetUpRouter()
	r.DELETE("/instructors/:id", DeleteInstructorHandle)
	req, _ := http.NewRequest("DELETE", "/instructors/11", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if err := types.BuildResponse(w.Body, &resp); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	if err := resp.Verify(); err != nil {
		assert.Fail(t, err.Error())
	}
}

type InstructorListResponse struct {
	types.BaseResponse
	Data []interface{} `json:"data"`
}

func TestGetInstructorListHandle(t *testing.T) {
	var resp InstructorListResponse
	r := SetUpRouter()
	r.GET("/instructors", GetInstructorListHandle)
	req, _ := http.NewRequest("GET", "/instructors", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if err := types.BuildResponse(w.Body, &resp); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	if err := resp.Verify(); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	if len(resp.Data) == 0 {
		assert.Fail(t, "No record in Instructor list.")
	}

	//types.PrintJson(resp.Data)
}

type GetInstructorResponse struct {
	types.BaseResponse
	Data models.Instructor `json:"data"`
}

func TestGetInstructorHandle(t *testing.T) {
	var resp GetInstructorResponse
	r := SetUpRouter()
	r.GET("/instructors/:id", GetInstructorHandle)
	req, _ := http.NewRequest("GET", "/instructors/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if err := types.BuildResponse(w.Body, &resp); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	if err := resp.Verify(); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	if resp.Data.ID != 1 {
		assert.Fail(t, "Fail to get Instructor")
	}

}

func TestUpdateInstructorHandle(t *testing.T) {
	var resp types.BaseResponse
	form := types.UpdateInstructorRequest{
		Name: types.GetStringPtr("test123"),
	}
	br, _ := types.StructToReader(form)
	r := SetUpRouter()
	r.PUT("/instructors/:id", UpdateInstructorHandle)

	req, _ := http.NewRequest("PUT", "/instructors/1", br)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if err := types.BuildResponse(w.Body, &resp); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	if err := resp.Verify(); err != nil {
		assert.Fail(t, err.Error())
	}
}

func TestCreateInstructorHandle(t *testing.T) {
	var resp types.BaseResponse
	form := types.CreateInstructorRequest{
		models.Instructor{
			ID:   4,
			Name: "test123",
		},
	}
	br, _ := types.StructToReader(form)
	r := SetUpRouter()
	r.POST("/instructors", CreateInstructorHandle)

	req, _ := http.NewRequest("POST", "/instructors", br)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if err := types.BuildResponse(w.Body, &resp); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	if err := resp.Verify(); err != nil {
		assert.Fail(t, err.Error())
	}
}
