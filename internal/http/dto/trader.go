package dto

type RegisterTraderRequest struct {
	Login 	 string `json:"login"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterTraderResponse struct {
	
}