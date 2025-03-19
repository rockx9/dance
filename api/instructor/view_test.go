package instructor

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestDeleteInstructorHandle(t *testing.T) {
	mockResponse := `{"code":0,"data":null,"msg":"ok"}`
	r := SetUpRouter()
	r.DELETE("/:id", DeleteInstructorHandle)
	req, _ := http.NewRequest("DELETE", "/11", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
