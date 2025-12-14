package utils

type ApiResponse struct {
	Data          interface{} `json:"data"`
	Error         interface{} `json:"error,omitempty"`
	StatusMessage string      `json:"status_message"`
	StatusCode    int64       `json:"status_code"`
}

func NewApiResponse(data interface{}, statusMessage string, statusCode int64) *ApiResponse {
	return &ApiResponse{
		Data:          data,
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
	}
}

func NewApiResponseNoData(statusMessage string, statusCode int64) *ApiResponse {
	return &ApiResponse{
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
		Data:          make([]interface{}, 0),
	}
}

func NewApiResponseWithError(error interface{}, statusMessage string, statusCode int64) *ApiResponse {
	return &ApiResponse{
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
		Data:          make([]interface{}, 0),
		Error:         error,
	}
}
