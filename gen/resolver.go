package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/loopcontext/go-graphql-orm/events"
)

// ResolutionHandlers struct
type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateWallet     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Wallet, err error)
	UpdateWallet     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Wallet, err error)
	DeleteWallet     func(ctx context.Context, r *GeneratedResolver, id string) (item *Wallet, err error)
	DeleteAllWallets func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryWallet      func(ctx context.Context, r *GeneratedResolver, opts QueryWalletHandlerOptions) (*Wallet, error)
	QueryWallets     func(ctx context.Context, r *GeneratedResolver, opts QueryWalletsHandlerOptions) (*WalletResultType, error)

	WalletWalletType func(ctx context.Context, r *GeneratedResolver, obj *Wallet) (res *WalletType, err error)

	WalletAccounts func(ctx context.Context, r *GeneratedResolver, obj *Wallet) (res []*Account, err error)

	WalletPayments func(ctx context.Context, r *GeneratedResolver, obj *Wallet) (res []*Payment, err error)

	CreateWalletType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *WalletType, err error)
	UpdateWalletType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *WalletType, err error)
	DeleteWalletType     func(ctx context.Context, r *GeneratedResolver, id string) (item *WalletType, err error)
	DeleteAllWalletTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryWalletType      func(ctx context.Context, r *GeneratedResolver, opts QueryWalletTypeHandlerOptions) (*WalletType, error)
	QueryWalletTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryWalletTypesHandlerOptions) (*WalletTypeResultType, error)

	WalletTypeWallet func(ctx context.Context, r *GeneratedResolver, obj *WalletType) (res *Wallet, err error)

	CreateAccountProviderType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *AccountProviderType, err error)
	UpdateAccountProviderType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *AccountProviderType, err error)
	DeleteAccountProviderType     func(ctx context.Context, r *GeneratedResolver, id string) (item *AccountProviderType, err error)
	DeleteAllAccountProviderTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryAccountProviderType      func(ctx context.Context, r *GeneratedResolver, opts QueryAccountProviderTypeHandlerOptions) (*AccountProviderType, error)
	QueryAccountProviderTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryAccountProviderTypesHandlerOptions) (*AccountProviderTypeResultType, error)

	AccountProviderTypeAccountProvider func(ctx context.Context, r *GeneratedResolver, obj *AccountProviderType) (res *AccountProvider, err error)

	CreateAccountProvider     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *AccountProvider, err error)
	UpdateAccountProvider     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *AccountProvider, err error)
	DeleteAccountProvider     func(ctx context.Context, r *GeneratedResolver, id string) (item *AccountProvider, err error)
	DeleteAllAccountProviders func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryAccountProvider      func(ctx context.Context, r *GeneratedResolver, opts QueryAccountProviderHandlerOptions) (*AccountProvider, error)
	QueryAccountProviders     func(ctx context.Context, r *GeneratedResolver, opts QueryAccountProvidersHandlerOptions) (*AccountProviderResultType, error)

	AccountProviderAccounts func(ctx context.Context, r *GeneratedResolver, obj *AccountProvider) (res []*Account, err error)

	AccountProviderAccountProviderType func(ctx context.Context, r *GeneratedResolver, obj *AccountProvider) (res *AccountProviderType, err error)

	CreateAccount     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Account, err error)
	UpdateAccount     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Account, err error)
	DeleteAccount     func(ctx context.Context, r *GeneratedResolver, id string) (item *Account, err error)
	DeleteAllAccounts func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryAccount      func(ctx context.Context, r *GeneratedResolver, opts QueryAccountHandlerOptions) (*Account, error)
	QueryAccounts     func(ctx context.Context, r *GeneratedResolver, opts QueryAccountsHandlerOptions) (*AccountResultType, error)

	AccountAccountProvider func(ctx context.Context, r *GeneratedResolver, obj *Account) (res *AccountProvider, err error)

	AccountWallet func(ctx context.Context, r *GeneratedResolver, obj *Account) (res *Wallet, err error)

	AccountPayments func(ctx context.Context, r *GeneratedResolver, obj *Account) (res []*Payment, err error)

	CreatePaymentChannel     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentChannel, err error)
	UpdatePaymentChannel     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentChannel, err error)
	DeletePaymentChannel     func(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentChannel, err error)
	DeleteAllPaymentChannels func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPaymentChannel      func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentChannelHandlerOptions) (*PaymentChannel, error)
	QueryPaymentChannels     func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentChannelsHandlerOptions) (*PaymentChannelResultType, error)

	PaymentChannelPayment func(ctx context.Context, r *GeneratedResolver, obj *PaymentChannel) (res *Payment, err error)

	CreatePaymentType     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentType, err error)
	UpdatePaymentType     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentType, err error)
	DeletePaymentType     func(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentType, err error)
	DeleteAllPaymentTypes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPaymentType      func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentTypeHandlerOptions) (*PaymentType, error)
	QueryPaymentTypes     func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentTypesHandlerOptions) (*PaymentTypeResultType, error)

	PaymentTypePayment func(ctx context.Context, r *GeneratedResolver, obj *PaymentType) (res *Payment, err error)

	CreatePayment     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Payment, err error)
	UpdatePayment     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Payment, err error)
	DeletePayment     func(ctx context.Context, r *GeneratedResolver, id string) (item *Payment, err error)
	DeleteAllPayments func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPayment      func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentHandlerOptions) (*Payment, error)
	QueryPayments     func(ctx context.Context, r *GeneratedResolver, opts QueryPaymentsHandlerOptions) (*PaymentResultType, error)

	PaymentWallet func(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *Wallet, err error)

	PaymentAccount func(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *Account, err error)

	PaymentPaymentChannel func(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *PaymentChannel, err error)

	PaymentPaymentType func(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *PaymentType, err error)
}

// DefaultResolutionHandlers ...
func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateWallet:     CreateWalletHandler,
		UpdateWallet:     UpdateWalletHandler,
		DeleteWallet:     DeleteWalletHandler,
		DeleteAllWallets: DeleteAllWalletsHandler,
		QueryWallet:      QueryWalletHandler,
		QueryWallets:     QueryWalletsHandler,

		WalletWalletType: WalletWalletTypeHandler,

		WalletAccounts: WalletAccountsHandler,

		WalletPayments: WalletPaymentsHandler,

		CreateWalletType:     CreateWalletTypeHandler,
		UpdateWalletType:     UpdateWalletTypeHandler,
		DeleteWalletType:     DeleteWalletTypeHandler,
		DeleteAllWalletTypes: DeleteAllWalletTypesHandler,
		QueryWalletType:      QueryWalletTypeHandler,
		QueryWalletTypes:     QueryWalletTypesHandler,

		WalletTypeWallet: WalletTypeWalletHandler,

		CreateAccountProviderType:     CreateAccountProviderTypeHandler,
		UpdateAccountProviderType:     UpdateAccountProviderTypeHandler,
		DeleteAccountProviderType:     DeleteAccountProviderTypeHandler,
		DeleteAllAccountProviderTypes: DeleteAllAccountProviderTypesHandler,
		QueryAccountProviderType:      QueryAccountProviderTypeHandler,
		QueryAccountProviderTypes:     QueryAccountProviderTypesHandler,

		AccountProviderTypeAccountProvider: AccountProviderTypeAccountProviderHandler,

		CreateAccountProvider:     CreateAccountProviderHandler,
		UpdateAccountProvider:     UpdateAccountProviderHandler,
		DeleteAccountProvider:     DeleteAccountProviderHandler,
		DeleteAllAccountProviders: DeleteAllAccountProvidersHandler,
		QueryAccountProvider:      QueryAccountProviderHandler,
		QueryAccountProviders:     QueryAccountProvidersHandler,

		AccountProviderAccounts: AccountProviderAccountsHandler,

		AccountProviderAccountProviderType: AccountProviderAccountProviderTypeHandler,

		CreateAccount:     CreateAccountHandler,
		UpdateAccount:     UpdateAccountHandler,
		DeleteAccount:     DeleteAccountHandler,
		DeleteAllAccounts: DeleteAllAccountsHandler,
		QueryAccount:      QueryAccountHandler,
		QueryAccounts:     QueryAccountsHandler,

		AccountAccountProvider: AccountAccountProviderHandler,

		AccountWallet: AccountWalletHandler,

		AccountPayments: AccountPaymentsHandler,

		CreatePaymentChannel:     CreatePaymentChannelHandler,
		UpdatePaymentChannel:     UpdatePaymentChannelHandler,
		DeletePaymentChannel:     DeletePaymentChannelHandler,
		DeleteAllPaymentChannels: DeleteAllPaymentChannelsHandler,
		QueryPaymentChannel:      QueryPaymentChannelHandler,
		QueryPaymentChannels:     QueryPaymentChannelsHandler,

		PaymentChannelPayment: PaymentChannelPaymentHandler,

		CreatePaymentType:     CreatePaymentTypeHandler,
		UpdatePaymentType:     UpdatePaymentTypeHandler,
		DeletePaymentType:     DeletePaymentTypeHandler,
		DeleteAllPaymentTypes: DeleteAllPaymentTypesHandler,
		QueryPaymentType:      QueryPaymentTypeHandler,
		QueryPaymentTypes:     QueryPaymentTypesHandler,

		PaymentTypePayment: PaymentTypePaymentHandler,

		CreatePayment:     CreatePaymentHandler,
		UpdatePayment:     UpdatePaymentHandler,
		DeletePayment:     DeletePaymentHandler,
		DeleteAllPayments: DeleteAllPaymentsHandler,
		QueryPayment:      QueryPaymentHandler,
		QueryPayments:     QueryPaymentsHandler,

		PaymentWallet: PaymentWalletHandler,

		PaymentAccount: PaymentAccountHandler,

		PaymentPaymentChannel: PaymentPaymentChannelHandler,

		PaymentPaymentType: PaymentPaymentTypeHandler,
	}
	return handlers
}

// GeneratedResolver struct
type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
