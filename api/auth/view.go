package auth

import "github.com/gin-gonic/gin"

// @Summary	login
// @Description	login
// @Accept	json
// @Produce	json
// @Tags	auth
// @Param   form	body	types.LoginRequest	true	"form""
// @Success 200 {object} models.Instructor "successful"
// @Router /auth/login [POST]
func LoginHandle(ctx *gin.Context) {
	NewAuthViewset(ctx).Login()
}
