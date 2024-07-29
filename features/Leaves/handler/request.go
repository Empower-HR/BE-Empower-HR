package handler

type LeaveRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Reason    string `json:"reason"`
}
