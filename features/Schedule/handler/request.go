package handler

type ScheduleRequest struct {
	Company       uint   `json:"company_id"`
	Name          string `json:"name"`
	EffectiveDate string `json:"effective_date"`
	ScheduleIn    string `json:"schedule_in"`
	ScheduleOut   string `json:"schedule_out"`
	BreakStart    string `json:"break_start"`
	BreakEnd      string `json:"break_end"`
	Days          string `json:"repeat_until"`
	Description   string `json:"description"`
}
