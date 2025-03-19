package types

import "dance/models"

type CreateInstructorRequest struct {
	models.Instructor
}

type UpdateInstructorRequest struct {
	Name         *string `json:"name" form:"name"`
	Bio          *string `json:"bio" form:"bio"`
	Specialty    *string `json:"specialty" form:"specialty"`
	Availability *bool   `json:"availability" form:"availability"`
}
