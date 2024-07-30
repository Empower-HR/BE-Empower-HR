package handler

type AnnRequest struct {
	CompanyID uint `json:"company_id" form:"company_id"`
	Title  string `json:"title" form:"title"`
	Description  string `json:"description" form:"description"`
	Attachment   string `json:"attachment" form:"attachment"`
}
