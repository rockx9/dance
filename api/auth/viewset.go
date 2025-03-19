package auth

import (
	"dance/api/middleware"
	"dance/service"
	"dance/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthView interface {
	Login()
}

type AuthViewset struct {
	middleware.Context
	s service.IAuth
}

func NewAuthViewset(c *gin.Context) IAuthView {
	v := &AuthViewset{}
	v.SetGinContext(c)
	v.s = service.NewAuth()
	return v
}

func (v AuthViewset) Login() {
	var form types.LoginRequest
	if err := v.Bind(&form); err != nil {
		v.ResponseError(err)
		return
	}
	token, err := v.s.Verify(form.Username, form.Password)
	if err != nil {
		v.ResponseError(err)
		return
	}
	http.SetCookie(v.GC.Writer, &http.Cookie{
		Name:  "token",
		Value: token,
	})
	v.ResponseOK(token)
}
