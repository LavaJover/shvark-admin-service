package usecase

import (
	"github.com/LavaJover/shvark-admin-service/internal/domain"
	"github.com/LavaJover/shvark-admin-service/internal/grpcclients"
	"github.com/LavaJover/shvark-admin-service/internal/httpclients"
	ssopb "github.com/LavaJover/shvark-sso-service/proto/gen"
)

type TraderUsecase struct {
	ssoClient 		*grpcclients.SSOClient
	walletClient 	*httpclients.WalletHTTPClient
}

func NewTraderUsecase(ssoClient *grpcclients.SSOClient, walletClient *httpclients.WalletHTTPClient) *TraderUsecase {
	return &TraderUsecase{ssoClient: ssoClient, walletClient: walletClient}
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