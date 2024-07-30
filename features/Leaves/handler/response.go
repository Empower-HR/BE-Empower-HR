package handler

type LeaveHistory struct {
	LeaveID      uint   `json:"leave_id,omitempty"`
	EmployeeID   uint   `json:"personal_id"`
	PersonalName string `json:"name"`
	JobPosition  string `json:"job_position"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Reason       string `json:"reason"`
	Status       string `json:"status"`
	TotalLeave   int    `json:"total_leave"`
}
