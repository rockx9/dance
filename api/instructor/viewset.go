package instructor

import (
	"dance/api/middleware"
	"dance/service"
	"dance/types"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type InstructorApiHandle interface {
	ListInstances()
	GetInstance()
	CreateInstance()
	UpdateInstance()
	DeleteInstance()
}

type InstructorViewset struct {
	middleware.Context
	s service.IInstructor
}

func NewInstructorViewset(c *gin.Context) InstructorApiHandle {
	v := &InstructorViewset{}
	v.SetGinContext(c)
	v.s = service.NewInstructor()
	return v
}

func (v InstructorViewset) ListInstances() {
	data, err := v.s.ListAllInstances()
	if err != nil {
		v.ResponseError(err)
		return
	}
	v.ResponseOK(data)
}

func (v InstructorViewset) getPathIdParameter() (int, error) {
	id := v.Param("id")
	instanceId, err := strconv.Atoi(id)
	if err != nil {
		return instanceId, errors.New("Instructor ID Error")
	}
	return instanceId, nil
}

func (v InstructorViewset) GetInstance() {
	id, err := v.getPathIdParameter()
	if err != nil {
		v.ResponseError(err)
		return
	}
	obj, err := v.s.GetInstance(id)
	if err != nil {
		v.ResponseError(err)
		return
	}
	v.ResponseOK(obj)
}

func (v InstructorViewset) CreateInstance() {
	var form types.CreateInstructorRequest
	if err := v.Bind(&form); err != nil {
		v.ResponseError(err)
		return
	}
	if err := v.s.CreateInstance(form.Instructor); err != nil {
		v.ResponseError(err)
		return
	}
	v.ResponseOK(nil)
}

func (v InstructorViewset) UpdateInstance() {
	var (
		id   int
		form types.UpdateInstructorRequest
		err  error
	)
	if id, err = v.getPathIdParameter(); err != nil {
		v.ResponseError(err)
		return
	}

	if err = v.Bind(&form); err != nil {
		v.ResponseError(err)
		return
	}

	if err = v.s.UpdateInstance(id, form); err != nil {
		v.ResponseError(err)
		return
	}
	v.ResponseOK(nil)
}

func (v InstructorViewset) DeleteInstance() {
	id, err := v.getPathIdParameter()
	if err != nil {
		v.ResponseError(err)
		return
	}
	if err = v.s.DeleteInstance(id); err != nil {
		v.ResponseError(err)
		return
	}
	v.ResponseOK(nil)
}
