package types

type PaginatorRequest struct {
	Page     int64 `json:"page" form:"page" default:"1"`
	PageSize int64 `json:"pageSize" form:"pageSize" default:"10"`
}

func GetStringPtr(s string) *string {
	return &s
}

func GetStringVal(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
