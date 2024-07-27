package responses

type MapResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONWebResponse(code int, status string, msg string, data interface{}) MapResponse {
	return MapResponse{
		Code:    code,
		Status:  status,
		Message: msg,
		Data:    data,
	}
}

type PaginatedResponse struct {
    Code    int         `json:"code"`
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Meta    interface{} `json:"meta,omitempty"`
}

// Helper function to create a paginated response
func PaginatedJSONResponse(code int, status string, msg string, data interface{}, meta interface{}) PaginatedResponse {
    return PaginatedResponse{
        Code:    code,
        Status:  status,
        Message: msg,
        Data:    data,
        Meta:    meta,
    }
}