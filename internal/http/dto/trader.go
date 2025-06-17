package dto

type RegisterTraderRequest struct {
	Login 	 string `json:"login"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterTraderResponse struct {}

type Trader struct {
	ID 			string 	`json:"id"`
	Username 	string 	`json:"username"`
	Login 		string 	`json:"login"`
	Password	string 	`json:"password"`
}

type GetTradersResponse struct {
	TotalPages 	int64 		`json:"total_pages"`
	Traders 	[]*Trader 	`json:"traders"`
}