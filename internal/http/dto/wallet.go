package dto

type CreateWalletRequest struct {
	TraderID string `json:"traderId"`
}

type CreateWalletResponse struct {
	Address string `json:"address"`
}