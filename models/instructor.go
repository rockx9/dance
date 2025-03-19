package models

type Instructor struct {
	ID           int     `json:"id" form:"id" binding:"required"`
	Name         string  `json:"name" form:"name" binding:"required"`
	Bio          *string `json:"bio" form:"bio"`
	Specialty    *string `json:"specialty" form:"specialty"`
	Availability *bool   `json:"availability" form:"availability"`
}
