package handler

type AttRequest struct {
	Clock_in string `json:"clock_in"`
	Date  string `json:"date"`
	Long  string `json:"long"`
	Lat   string `json:"lat"`
	Notes string `json:"notes"`
}
type AttUpdateReq struct {
	Clock_out string `json:"clock_out"`
	Status    string    `json:"status"`
	Long  string `json:"long"`
	Lat   string `json:"lat"`
	Notes string `json:"notes"`
}