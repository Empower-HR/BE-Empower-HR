package handler

type AnnRequest struct {
	CompanyID uint `json:"company_id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
	Attachment   string `json:"atchment"`
}
