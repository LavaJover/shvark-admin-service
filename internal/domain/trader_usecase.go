package domain

type TraderUsecase interface {
	RegisterNewTrader(trader *Trader) error
	GetTraders(page, limit int64) ([]*Trader, int64, error)
}

type Trader struct {
	ID 			string
	Username 	string
	Login 		string
	Password 	string
}