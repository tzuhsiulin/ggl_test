package dto

type CommonErrorResponse struct {
	Status  string `json:"status"`
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}
