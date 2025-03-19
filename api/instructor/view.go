package instructor

import "github.com/gin-gonic/gin"

// @Summary	List instructors
// @Description	List all instructors
// @Accept	json
// @Produce	json
// @Tags	instructors
// @Param 	Authorization	header	string	true	"JWT token"
// @Success 200 {object} []models.Instructor "successful"
// @Router /instructors [GET]
func GetInstructorListHandle(ctx *gin.Context) {
	NewInstructorViewset(ctx).ListInstances()
}

// @Summary	Get one instructor
// @Description	Get instructor by ID
// @Accept	json
// @Produce	json
// @Tags	instructors
// @Param 	Authorization	header	string	true	"JWT token"
// @Param   id		path    int     true        "instructor ID"
// @Success 200 {object} models.Instructor "successful"
// @Router /instructors/{id} [GET]
func GetInstructorHandle(ctx *gin.Context) {
	NewInstructorViewset(ctx).GetInstance()
}

// @Summary	Update instructor
// @Description	Update instructor
// @Accept	json
// @Produce	json
// @Tags	instructors
// @Param 	Authorization	header	string	true	"JWT token"
// @Param   id		path    int     true        "instructor ID"
// @Param   form	body	types.UpdateInstructorRequest	true	"form""
// @Success 200 {object} models.Instructor "successful"
// @Router /instructors/{id} [PUT]
func UpdateInstructorHandle(ctx *gin.Context) {
	NewInstructorViewset(ctx).UpdateInstance()
}

// @Summary	Create instructor
// @Description	Create a instructor
// @Accept	json
// @Produce	json
// @Tags	instructors
// @Param 	Authorization	header	string	true	"JWT token"
// @Param   form	body	types.CreateInstructorRequest	true	"form""
// @Success 200 {object} models.Instructor "successful"
// @Router /instructors [POST]
func CreateInstructorHandle(ctx *gin.Context) {
	NewInstructorViewset(ctx).CreateInstance()
}

// @Summary	Delete instructor
// @Description	Delete a instructor
// @Accept	json
// @Produce	json
// @Tags	instructors
// @Param 	Authorization	header	string	true	"JWT token"
// @Param   id		path    int     true        "instructor ID"
// @Success 200 {object} nil "successful"
// @Router /instructors/{id} [DELETE]
func DeleteInstructorHandle(ctx *gin.Context) {
	NewInstructorViewset(ctx).DeleteInstance()
}
