package midware

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type RegisterRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type FeedbackRequest struct{
	Context string `json:"context"`
}