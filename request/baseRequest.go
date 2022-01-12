package request

// SizeInput 偏移量
type SizeInput struct {
	Size int `json:"size" validate:"required,min=1" form:"size,default=15" comment:"条目"`
}

// PageInput 页码
type PageInput struct {
	Page int `json:"page" validate:"required,min=1" form:"page,default=1" comment:"页码"`
}