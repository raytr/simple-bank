package response

type BaseResponse struct {
	Message string `json:"message"`
}

var SuccessResponse = BaseResponse{
	Message: "success",
}
