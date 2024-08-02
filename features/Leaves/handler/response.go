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
}

type LeaveHistoryResponse struct {
	Code           int            `json:"code"`
	Status         string         `json:"status"`
	Message        string         `json:"message"`
	Data           []LeaveHistory `json:"data"`
	TotalEmployees int            `json:"total_employe"`
	TotalLeaves    int            `json:"total_leaves"`
}

type LeaveHistoryEmployeeResponse struct {
	Code     int            `json:"code"`
	Status   string         `json:"status"`
	Message  string         `json:"message"`
	Names    string         `json:"names"`
	Data     []LeaveHistory `json:"data"`
	UsedCuti int            `json:"used"`
}
