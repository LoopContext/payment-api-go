package src

import (
	"context"
	"fmt"

	"github.com/loopcontext/payment-api-go/gen"
)

const (
	jwtTokenPermissionErrMsg = "You don't have permission to %s the entity %s"
)

// Wallets method
func (r *QueryResolver) Wallets(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.WalletSortType, filter *gen.WalletFilterType) (*gen.WalletResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "wallets")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Wallets to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Wallets(ctx, offset, limit, q, sort, filter)
}

// CreateWallet method
func (r *MutationResolver) CreateWallet(ctx context.Context, input map[string]interface{}) (item *gen.Wallet, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "wallets")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Wallets to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateWallet(ctx, input)
}

// ReadWallet method
func (r *QueryResolver) Wallet(ctx context.Context, id *string, q *string, filter *gen.WalletFilterType) (*gen.Wallet, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "wallets")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Wallets to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Wallet(ctx, id, q, filter)
}

// UpdateWallet method
func (r *MutationResolver) UpdateWallet(ctx context.Context, id string, input map[string]interface{}) (item *gen.Wallet, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "wallets")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Wallets to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateWallet(ctx, id, input)
}

// DeleteWallet method
func (r *MutationResolver) DeleteWallet(ctx context.Context, id string) (item *gen.Wallet, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallets")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Wallets to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteWallet(ctx, id)
}

// DeleteAllWallets method
func (r *MutationResolver) DeleteAllWallets(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallets")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Wallets to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllWallets(ctx)
}

// WalletTypes method
func (r *QueryResolver) WalletTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.WalletTypeSortType, filter *gen.WalletTypeFilterType) (*gen.WalletTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "wallet_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope WalletTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.WalletTypes(ctx, offset, limit, q, sort, filter)
}

// CreateWalletType method
func (r *MutationResolver) CreateWalletType(ctx context.Context, input map[string]interface{}) (item *gen.WalletType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "wallet_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope WalletTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateWalletType(ctx, input)
}

// ReadWalletType method
func (r *QueryResolver) WalletType(ctx context.Context, id *string, q *string, filter *gen.WalletTypeFilterType) (*gen.WalletType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "wallet_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope WalletTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.WalletType(ctx, id, q, filter)
}

// UpdateWalletType method
func (r *MutationResolver) UpdateWalletType(ctx context.Context, id string, input map[string]interface{}) (item *gen.WalletType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "wallet_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope WalletTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateWalletType(ctx, id, input)
}

// DeleteWalletType method
func (r *MutationResolver) DeleteWalletType(ctx context.Context, id string) (item *gen.WalletType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallet_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope WalletTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteWalletType(ctx, id)
}

// DeleteAllWalletTypes method
func (r *MutationResolver) DeleteAllWalletTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallet_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope WalletTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllWalletTypes(ctx)
}

// AccountProviderTypes method
func (r *QueryResolver) AccountProviderTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.AccountProviderTypeSortType, filter *gen.AccountProviderTypeFilterType) (*gen.AccountProviderTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "account_provider_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviderTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.AccountProviderTypes(ctx, offset, limit, q, sort, filter)
}

// CreateAccountProviderType method
func (r *MutationResolver) CreateAccountProviderType(ctx context.Context, input map[string]interface{}) (item *gen.AccountProviderType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "account_provider_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviderTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateAccountProviderType(ctx, input)
}

// ReadAccountProviderType method
func (r *QueryResolver) AccountProviderType(ctx context.Context, id *string, q *string, filter *gen.AccountProviderTypeFilterType) (*gen.AccountProviderType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "account_provider_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviderTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.AccountProviderType(ctx, id, q, filter)
}

// UpdateAccountProviderType method
func (r *MutationResolver) UpdateAccountProviderType(ctx context.Context, id string, input map[string]interface{}) (item *gen.AccountProviderType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "account_provider_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviderTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateAccountProviderType(ctx, id, input)
}

// DeleteAccountProviderType method
func (r *MutationResolver) DeleteAccountProviderType(ctx context.Context, id string) (item *gen.AccountProviderType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_provider_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviderTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAccountProviderType(ctx, id)
}

// DeleteAllAccountProviderTypes method
func (r *MutationResolver) DeleteAllAccountProviderTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_provider_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviderTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllAccountProviderTypes(ctx)
}

// AccountProviders method
func (r *QueryResolver) AccountProviders(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.AccountProviderSortType, filter *gen.AccountProviderFilterType) (*gen.AccountProviderResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "account_providers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviders to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.AccountProviders(ctx, offset, limit, q, sort, filter)
}

// CreateAccountProvider method
func (r *MutationResolver) CreateAccountProvider(ctx context.Context, input map[string]interface{}) (item *gen.AccountProvider, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "account_providers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviders to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateAccountProvider(ctx, input)
}

// ReadAccountProvider method
func (r *QueryResolver) AccountProvider(ctx context.Context, id *string, q *string, filter *gen.AccountProviderFilterType) (*gen.AccountProvider, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "account_providers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviders to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.AccountProvider(ctx, id, q, filter)
}

// UpdateAccountProvider method
func (r *MutationResolver) UpdateAccountProvider(ctx context.Context, id string, input map[string]interface{}) (item *gen.AccountProvider, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "account_providers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviders to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateAccountProvider(ctx, id, input)
}

// DeleteAccountProvider method
func (r *MutationResolver) DeleteAccountProvider(ctx context.Context, id string) (item *gen.AccountProvider, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_providers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviders to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAccountProvider(ctx, id)
}

// DeleteAllAccountProviders method
func (r *MutationResolver) DeleteAllAccountProviders(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_providers")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope AccountProviders to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllAccountProviders(ctx)
}

// Accounts method
func (r *QueryResolver) Accounts(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.AccountSortType, filter *gen.AccountFilterType) (*gen.AccountResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "accounts")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Accounts to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Accounts(ctx, offset, limit, q, sort, filter)
}

// CreateAccount method
func (r *MutationResolver) CreateAccount(ctx context.Context, input map[string]interface{}) (item *gen.Account, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "accounts")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Accounts to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateAccount(ctx, input)
}

// ReadAccount method
func (r *QueryResolver) Account(ctx context.Context, id *string, q *string, filter *gen.AccountFilterType) (*gen.Account, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "accounts")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Accounts to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Account(ctx, id, q, filter)
}

// UpdateAccount method
func (r *MutationResolver) UpdateAccount(ctx context.Context, id string, input map[string]interface{}) (item *gen.Account, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "accounts")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Accounts to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateAccount(ctx, id, input)
}

// DeleteAccount method
func (r *MutationResolver) DeleteAccount(ctx context.Context, id string) (item *gen.Account, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "accounts")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Accounts to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAccount(ctx, id)
}

// DeleteAllAccounts method
func (r *MutationResolver) DeleteAllAccounts(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "accounts")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Accounts to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllAccounts(ctx)
}

// PaymentChannels method
func (r *QueryResolver) PaymentChannels(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentChannelSortType, filter *gen.PaymentChannelFilterType) (*gen.PaymentChannelResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payment_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentChannels(ctx, offset, limit, q, sort, filter)
}

// CreatePaymentChannel method
func (r *MutationResolver) CreatePaymentChannel(ctx context.Context, input map[string]interface{}) (item *gen.PaymentChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePaymentChannel(ctx, input)
}

// ReadPaymentChannel method
func (r *QueryResolver) PaymentChannel(ctx context.Context, id *string, q *string, filter *gen.PaymentChannelFilterType) (*gen.PaymentChannel, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentChannel(ctx, id, q, filter)
}

// UpdatePaymentChannel method
func (r *MutationResolver) UpdatePaymentChannel(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePaymentChannel(ctx, id, input)
}

// DeletePaymentChannel method
func (r *MutationResolver) DeletePaymentChannel(ctx context.Context, id string) (item *gen.PaymentChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePaymentChannel(ctx, id)
}

// DeleteAllPaymentChannels method
func (r *MutationResolver) DeleteAllPaymentChannels(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_channels")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentChannels to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPaymentChannels(ctx)
}

// PaymentTypes method
func (r *QueryResolver) PaymentTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentTypeSortType, filter *gen.PaymentTypeFilterType) (*gen.PaymentTypeResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payment_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentTypes(ctx, offset, limit, q, sort, filter)
}

// CreatePaymentType method
func (r *MutationResolver) CreatePaymentType(ctx context.Context, input map[string]interface{}) (item *gen.PaymentType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePaymentType(ctx, input)
}

// ReadPaymentType method
func (r *QueryResolver) PaymentType(ctx context.Context, id *string, q *string, filter *gen.PaymentTypeFilterType) (*gen.PaymentType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PaymentType(ctx, id, q, filter)
}

// UpdatePaymentType method
func (r *MutationResolver) UpdatePaymentType(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePaymentType(ctx, id, input)
}

// DeletePaymentType method
func (r *MutationResolver) DeletePaymentType(ctx context.Context, id string) (item *gen.PaymentType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePaymentType(ctx, id)
}

// DeleteAllPaymentTypes method
func (r *MutationResolver) DeleteAllPaymentTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_types")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PaymentTypes to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPaymentTypes(ctx)
}

// Payments method
func (r *QueryResolver) Payments(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentSortType, filter *gen.PaymentFilterType) (*gen.PaymentResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payments")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Payments to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Payments(ctx, offset, limit, q, sort, filter)
}

// CreatePayment method
func (r *MutationResolver) CreatePayment(ctx context.Context, input map[string]interface{}) (item *gen.Payment, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstCreate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payments")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Payments to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePayment(ctx, input)
}

// ReadPayment method
func (r *QueryResolver) Payment(ctx context.Context, id *string, q *string, filter *gen.PaymentFilterType) (*gen.Payment, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payments")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Payments to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Payment(ctx, id, q, filter)
}

// UpdatePayment method
func (r *MutationResolver) UpdatePayment(ctx context.Context, id string, input map[string]interface{}) (item *gen.Payment, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstUpdate[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payments")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Payments to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePayment(ctx, id, input)
}

// DeletePayment method
func (r *MutationResolver) DeletePayment(ctx context.Context, id string) (item *gen.Payment, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") && !gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstDelete[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payments")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Payments to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePayment(ctx, id)
}

// DeleteAllPayments method
func (r *MutationResolver) DeleteAllPayments(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstDelete[:1]) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payments")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Payments to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPayments(ctx)
}
