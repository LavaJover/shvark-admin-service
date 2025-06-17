package usecase

import (
	"github.com/LavaJover/shvark-admin-service/internal/domain"
	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	"github.com/LavaJover/shvark-admin-service/internal/httpclients"
	ssopb "github.com/LavaJover/shvark-sso-service/proto/gen"
	userpb "github.com/LavaJover/shvark-user-service/proto/gen"
)

type TraderUsecase struct {
	ssoClient 		*grpcclients.SSOClient
	walletClient 	*httpclients.WalletHTTPClient
	usersClient  	*grpcclients.UserClient
}

func NewTraderUsecase(ssoClient *grpcclients.SSOClient, walletClient *httpclients.WalletHTTPClient, userClient *grpcclients.UserClient) *TraderUsecase {
	return &TraderUsecase{ssoClient: ssoClient, walletClient: walletClient, usersClient: userClient}
}

func (traderUc *TraderUsecase) RegisterNewTrader(trader *domain.Trader) error {
	// Register trader in SSO-service by credentials
	ssoResponse, err := traderUc.ssoClient.Register(&ssopb.RegisterRequest{
		Login: trader.Login,
		Username: trader.Login,
		Password: trader.Password,
	})
	if err != nil {
		return err
	}
	trader.ID = ssoResponse.UserId

	// Create trader crypto-wallet
	_, err = traderUc.walletClient.CreateTraderWallet(&httpclients.CreateWalletRequest{
		TraderID: trader.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (traderUc *TraderUsecase) GetTraders(page, limit int64) ([]*domain.Trader, int64, error) {
	userResponse, err := traderUc.usersClient.GetUsers(&userpb.GetUsersRequest{
		Page: page,
		Limit: limit,
	})
	if err != nil {
		return nil, 0, err
	}

	var traders []*domain.Trader
	for _, userRecord := range userResponse.Users {
		traders = append(traders, &domain.Trader{
			ID: userRecord.UserId,
			Username: userRecord.Username,
			Login: userRecord.Username,
			Password: userRecord.Password,
		})
	}
	totalPages := userResponse.TotalPages

	return traders, int64(totalPages), nil
}