package auth

import (
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

type LoginResponse struct {
	types.BaseResponse
	Data interface{} `json:"data"`
}

func TestLoginHandle(t *testing.T) {
	var resp LoginResponse
	form := types.LoginRequest{
		Username: "admin",
		Password: "abcd.1234",
	}
	br, _ := types.StructToReader(form)
	r := SetUpRouter()
	r.POST("/login", LoginHandle)

	req, _ := http.NewRequest("POST", "/login", br)
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
	//types.PrintJson(resp)
}
