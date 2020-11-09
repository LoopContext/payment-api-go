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
	if !gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "wallets")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		if filter != nil {
			filter.UserID = &jwtClaims.Subject
		} else {
			filter = &gen.WalletFilterType{
				UserID: &jwtClaims.Subject,
			}
		}
	}
	return r.GeneratedQueryResolver.Wallets(ctx, offset, limit, q, sort, filter)
}

// CreateWallet method
func (r *MutationResolver) CreateWallet(ctx context.Context, input map[string]interface{}) (item *gen.Wallet, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateWallet(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "wallets")
}

// ReadWallet method
func (r *QueryResolver) Wallet(ctx context.Context, id *string, q *string, filter *gen.WalletFilterType) (*gen.Wallet, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "wallets")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		qw, err := r.GeneratedResolver.Handlers.QueryWallet(ctx, r.GeneratedResolver,
			gen.QueryWalletHandlerOptions{
				Filter: &gen.WalletFilterType{
					UserID: &jwtClaims.Subject,
					ID:     id,
				},
			})
		if err != nil || qw == nil {
			return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "wallets")
		}
	}
	return r.GeneratedQueryResolver.Wallet(ctx, id, q, filter)
}

// UpdateWallet method
func (r *MutationResolver) UpdateWallet(ctx context.Context, id string, input map[string]interface{}) (item *gen.Wallet, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateWallet(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "wallets")
}

// DeleteWallet method
func (r *MutationResolver) DeleteWallet(ctx context.Context, id string) (item *gen.Wallet, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteWallet(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallets")
}

// DeleteAllWallets method
func (r *MutationResolver) DeleteAllWallets(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallets", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllWallets(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallets")
}

// CreateWalletType method
func (r *MutationResolver) CreateWalletType(ctx context.Context, input map[string]interface{}) (item *gen.WalletType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateWalletType(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "wallet_types")
}

// ReadWalletType method
func (r *QueryResolver) WalletType(ctx context.Context, id *string, q *string, filter *gen.WalletTypeFilterType) (*gen.WalletType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.WalletType(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "wallet_types")
}

// UpdateWalletType method
func (r *MutationResolver) UpdateWalletType(ctx context.Context, id string, input map[string]interface{}) (item *gen.WalletType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateWalletType(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "wallet_types")
}

// DeleteWalletType method
func (r *MutationResolver) DeleteWalletType(ctx context.Context, id string) (item *gen.WalletType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteWalletType(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallet_types")
}

// DeleteAllWalletTypes method
func (r *MutationResolver) DeleteAllWalletTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "wallet_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllWalletTypes(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "wallet_types")
}

// CreateAccountProviderType method
func (r *MutationResolver) CreateAccountProviderType(ctx context.Context, input map[string]interface{}) (item *gen.AccountProviderType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateAccountProviderType(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "account_provider_types")
}

// ReadAccountProviderType method
func (r *QueryResolver) AccountProviderType(ctx context.Context, id *string, q *string, filter *gen.AccountProviderTypeFilterType) (*gen.AccountProviderType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.AccountProviderType(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "account_provider_types")
}

// UpdateAccountProviderType method
func (r *MutationResolver) UpdateAccountProviderType(ctx context.Context, id string, input map[string]interface{}) (item *gen.AccountProviderType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateAccountProviderType(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "account_provider_types")
}

// DeleteAccountProviderType method
func (r *MutationResolver) DeleteAccountProviderType(ctx context.Context, id string) (item *gen.AccountProviderType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAccountProviderType(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_provider_types")
}

// DeleteAllAccountProviderTypes method
func (r *MutationResolver) DeleteAllAccountProviderTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_provider_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllAccountProviderTypes(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_provider_types")
}

// CreateAccountProvider method
func (r *MutationResolver) CreateAccountProvider(ctx context.Context, input map[string]interface{}) (item *gen.AccountProvider, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateAccountProvider(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "account_providers")
}

// ReadAccountProvider method
func (r *QueryResolver) AccountProvider(ctx context.Context, id *string, q *string, filter *gen.AccountProviderFilterType) (*gen.AccountProvider, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.AccountProvider(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "account_providers")
}

// UpdateAccountProvider method
func (r *MutationResolver) UpdateAccountProvider(ctx context.Context, id string, input map[string]interface{}) (item *gen.AccountProvider, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateAccountProvider(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "account_providers")
}

// DeleteAccountProvider method
func (r *MutationResolver) DeleteAccountProvider(ctx context.Context, id string) (item *gen.AccountProvider, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAccountProvider(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_providers")
}

// DeleteAllAccountProviders method
func (r *MutationResolver) DeleteAllAccountProviders(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "account_providers", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllAccountProviders(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "account_providers")
}

// Accounts method
func (r *QueryResolver) Accounts(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.AccountSortType, filter *gen.AccountFilterType) (*gen.AccountResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "accounts")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		if filter != nil && filter.Wallet != nil {
			filter.Wallet.UserID = &jwtClaims.Subject
		} else {
			filter = &gen.AccountFilterType{
				Wallet: &gen.WalletFilterType{
					UserID: &jwtClaims.Subject,
				},
			}
		}
	}
	return r.GeneratedQueryResolver.Accounts(ctx, offset, limit, q, sort, filter)
}

// CreateAccount method
func (r *MutationResolver) CreateAccount(ctx context.Context, input map[string]interface{}) (item *gen.Account, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreateAccount(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "accounts")
}

// ReadAccount method
func (r *QueryResolver) Account(ctx context.Context, id *string, q *string, filter *gen.AccountFilterType) (*gen.Account, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "accounts")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		qa, err := r.GeneratedResolver.Handlers.QueryAccount(ctx, r.GeneratedResolver,
			gen.QueryAccountHandlerOptions{
				Filter: &gen.AccountFilterType{
					Wallet: &gen.WalletFilterType{
						UserID: &jwtClaims.Subject,
					},
					ID: id,
				},
			})
		if err != nil || qa == nil {
			return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "accounts")
		}
	}
	return r.GeneratedQueryResolver.Account(ctx, id, q, filter)
}

// UpdateAccount method
func (r *MutationResolver) UpdateAccount(ctx context.Context, id string, input map[string]interface{}) (item *gen.Account, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdateAccount(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "accounts")
}

// DeleteAccount method
func (r *MutationResolver) DeleteAccount(ctx context.Context, id string) (item *gen.Account, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAccount(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "accounts")
}

// DeleteAllAccounts method
func (r *MutationResolver) DeleteAllAccounts(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "accounts", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllAccounts(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "accounts")
}

// CreatePaymentChannel method
func (r *MutationResolver) CreatePaymentChannel(ctx context.Context, input map[string]interface{}) (item *gen.PaymentChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreatePaymentChannel(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_channels")
}

// ReadPaymentChannel method
func (r *QueryResolver) PaymentChannel(ctx context.Context, id *string, q *string, filter *gen.PaymentChannelFilterType) (*gen.PaymentChannel, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.PaymentChannel(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_channels")
}

// UpdatePaymentChannel method
func (r *MutationResolver) UpdatePaymentChannel(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdatePaymentChannel(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_channels")
}

// DeletePaymentChannel method
func (r *MutationResolver) DeletePaymentChannel(ctx context.Context, id string) (item *gen.PaymentChannel, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeletePaymentChannel(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_channels")
}

// DeleteAllPaymentChannels method
func (r *MutationResolver) DeleteAllPaymentChannels(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_channels", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllPaymentChannels(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_channels")
}

// CreatePaymentType method
func (r *MutationResolver) CreatePaymentType(ctx context.Context, input map[string]interface{}) (item *gen.PaymentType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreatePaymentType(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payment_types")
}

// ReadPaymentType method
func (r *QueryResolver) PaymentType(ctx context.Context, id *string, q *string, filter *gen.PaymentTypeFilterType) (*gen.PaymentType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstRead[:1]) {
		return r.GeneratedQueryResolver.PaymentType(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payment_types")
}

// UpdatePaymentType method
func (r *MutationResolver) UpdatePaymentType(ctx context.Context, id string, input map[string]interface{}) (item *gen.PaymentType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdatePaymentType(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payment_types")
}

// DeletePaymentType method
func (r *MutationResolver) DeletePaymentType(ctx context.Context, id string) (item *gen.PaymentType, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeletePaymentType(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_types")
}

// DeleteAllPaymentTypes method
func (r *MutationResolver) DeleteAllPaymentTypes(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payment_types", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllPaymentTypes(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payment_types")
}

// Payments method
func (r *QueryResolver) Payments(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PaymentSortType, filter *gen.PaymentFilterType) (*gen.PaymentResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstList[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "payments")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		if filter != nil && filter.Wallet != nil {
			filter.Wallet.UserID = &jwtClaims.Subject
		} else {
			filter = &gen.PaymentFilterType{
				Wallet: &gen.WalletFilterType{
					UserID: &jwtClaims.Subject,
				},
			}
		}
	}
	return r.GeneratedQueryResolver.Payments(ctx, offset, limit, q, sort, filter)
}

// CreatePayment method
func (r *MutationResolver) CreatePayment(ctx context.Context, input map[string]interface{}) (item *gen.Payment, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstCreate[:1]) {
		return r.GeneratedMutationResolver.CreatePayment(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "payments")
}

// ReadPayment method
func (r *QueryResolver) Payment(ctx context.Context, id *string, q *string, filter *gen.PaymentFilterType) (*gen.Payment, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstRead[:1]) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "payments")
	}
	if !gen.HasRole(jwtClaims, "admin") {
		qp, err := r.GeneratedResolver.Handlers.QueryPayment(ctx, r.GeneratedResolver,
			gen.QueryPaymentHandlerOptions{
				Filter: &gen.PaymentFilterType{
					Wallet: &gen.WalletFilterType{
						UserID: &jwtClaims.Subject,
					},
					ID: id,
				},
			})
		if err != nil || qp == nil {
			return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payments")
		}
	}
	return r.GeneratedQueryResolver.Payment(ctx, id, q, filter)
}

// UpdatePayment method
func (r *MutationResolver) UpdatePayment(ctx context.Context, id string, input map[string]interface{}) (item *gen.Payment, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstUpdate[:1]) {
		return r.GeneratedMutationResolver.UpdatePayment(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "payments")
}

// DeletePayment method
func (r *MutationResolver) DeletePayment(ctx context.Context, id string) (item *gen.Payment, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeletePayment(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payments")
}

// DeleteAllPayments method
func (r *MutationResolver) DeleteAllPayments(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "payments", gen.JWTPermissionConstDelete[:1]) {
		return r.GeneratedMutationResolver.DeleteAllPayments(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "payments")
}
