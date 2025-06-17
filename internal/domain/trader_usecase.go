package domain

type TraderUsecase interface {
	RegisterNewTrader(trader *Trader) error
}

type Trader struct {
	ID 			string
	Username 	string
	Login 		string
	Password 	string
}