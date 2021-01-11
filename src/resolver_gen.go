package src

import (
	"context"

	"github.com/loopcontext/payment-api-go/gen"
)

// NewResolver ...
func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{gen.NewGeneratedResolver(handlers, db, ec)}
}

// Resolver ...
type Resolver struct {
	*gen.GeneratedResolver
}

// MutationResolver ...
type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

// BeginTransaction ...
func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

// QueryResolver ...
type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

// Mutation ...
func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{
		GeneratedMutationResolver: &gen.GeneratedMutationResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// Query ...
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{
		GeneratedQueryResolver: &gen.GeneratedQueryResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// WalletResultTypeResolver struct
type WalletResultTypeResolver struct {
	*gen.GeneratedWalletResultTypeResolver
}

// WalletResultType ...
func (r *Resolver) WalletResultType() gen.WalletResultTypeResolver {
	return &WalletResultTypeResolver{
		GeneratedWalletResultTypeResolver: &gen.GeneratedWalletResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// WalletResolver struct
type WalletResolver struct {
	*gen.GeneratedWalletResolver
}

// Wallet ...
func (r *Resolver) Wallet() gen.WalletResolver {
	return &WalletResolver{
		GeneratedWalletResolver: &gen.GeneratedWalletResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// WalletTypeResultTypeResolver struct
type WalletTypeResultTypeResolver struct {
	*gen.GeneratedWalletTypeResultTypeResolver
}

// WalletTypeResultType ...
func (r *Resolver) WalletTypeResultType() gen.WalletTypeResultTypeResolver {
	return &WalletTypeResultTypeResolver{
		GeneratedWalletTypeResultTypeResolver: &gen.GeneratedWalletTypeResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// WalletTypeResolver struct
type WalletTypeResolver struct {
	*gen.GeneratedWalletTypeResolver
}

// WalletType ...
func (r *Resolver) WalletType() gen.WalletTypeResolver {
	return &WalletTypeResolver{
		GeneratedWalletTypeResolver: &gen.GeneratedWalletTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// AccountProviderTypeResultTypeResolver struct
type AccountProviderTypeResultTypeResolver struct {
	*gen.GeneratedAccountProviderTypeResultTypeResolver
}

// AccountProviderTypeResultType ...
func (r *Resolver) AccountProviderTypeResultType() gen.AccountProviderTypeResultTypeResolver {
	return &AccountProviderTypeResultTypeResolver{
		GeneratedAccountProviderTypeResultTypeResolver: &gen.GeneratedAccountProviderTypeResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// AccountProviderTypeResolver struct
type AccountProviderTypeResolver struct {
	*gen.GeneratedAccountProviderTypeResolver
}

// AccountProviderType ...
func (r *Resolver) AccountProviderType() gen.AccountProviderTypeResolver {
	return &AccountProviderTypeResolver{
		GeneratedAccountProviderTypeResolver: &gen.GeneratedAccountProviderTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// AccountProviderResultTypeResolver struct
type AccountProviderResultTypeResolver struct {
	*gen.GeneratedAccountProviderResultTypeResolver
}

// AccountProviderResultType ...
func (r *Resolver) AccountProviderResultType() gen.AccountProviderResultTypeResolver {
	return &AccountProviderResultTypeResolver{
		GeneratedAccountProviderResultTypeResolver: &gen.GeneratedAccountProviderResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// AccountProviderResolver struct
type AccountProviderResolver struct {
	*gen.GeneratedAccountProviderResolver
}

// AccountProvider ...
func (r *Resolver) AccountProvider() gen.AccountProviderResolver {
	return &AccountProviderResolver{
		GeneratedAccountProviderResolver: &gen.GeneratedAccountProviderResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// AccountResultTypeResolver struct
type AccountResultTypeResolver struct {
	*gen.GeneratedAccountResultTypeResolver
}

// AccountResultType ...
func (r *Resolver) AccountResultType() gen.AccountResultTypeResolver {
	return &AccountResultTypeResolver{
		GeneratedAccountResultTypeResolver: &gen.GeneratedAccountResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// AccountResolver struct
type AccountResolver struct {
	*gen.GeneratedAccountResolver
}

// Account ...
func (r *Resolver) Account() gen.AccountResolver {
	return &AccountResolver{
		GeneratedAccountResolver: &gen.GeneratedAccountResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PaymentChannelResultTypeResolver struct
type PaymentChannelResultTypeResolver struct {
	*gen.GeneratedPaymentChannelResultTypeResolver
}

// PaymentChannelResultType ...
func (r *Resolver) PaymentChannelResultType() gen.PaymentChannelResultTypeResolver {
	return &PaymentChannelResultTypeResolver{
		GeneratedPaymentChannelResultTypeResolver: &gen.GeneratedPaymentChannelResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PaymentChannelResolver struct
type PaymentChannelResolver struct {
	*gen.GeneratedPaymentChannelResolver
}

// PaymentChannel ...
func (r *Resolver) PaymentChannel() gen.PaymentChannelResolver {
	return &PaymentChannelResolver{
		GeneratedPaymentChannelResolver: &gen.GeneratedPaymentChannelResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PaymentTypeResultTypeResolver struct
type PaymentTypeResultTypeResolver struct {
	*gen.GeneratedPaymentTypeResultTypeResolver
}

// PaymentTypeResultType ...
func (r *Resolver) PaymentTypeResultType() gen.PaymentTypeResultTypeResolver {
	return &PaymentTypeResultTypeResolver{
		GeneratedPaymentTypeResultTypeResolver: &gen.GeneratedPaymentTypeResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PaymentTypeResolver struct
type PaymentTypeResolver struct {
	*gen.GeneratedPaymentTypeResolver
}

// PaymentType ...
func (r *Resolver) PaymentType() gen.PaymentTypeResolver {
	return &PaymentTypeResolver{
		GeneratedPaymentTypeResolver: &gen.GeneratedPaymentTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PaymentResultTypeResolver struct
type PaymentResultTypeResolver struct {
	*gen.GeneratedPaymentResultTypeResolver
}

// PaymentResultType ...
func (r *Resolver) PaymentResultType() gen.PaymentResultTypeResolver {
	return &PaymentResultTypeResolver{
		GeneratedPaymentResultTypeResolver: &gen.GeneratedPaymentResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PaymentResolver struct
type PaymentResolver struct {
	*gen.GeneratedPaymentResolver
}

// Payment ...
func (r *Resolver) Payment() gen.PaymentResolver {
	return &PaymentResolver{
		GeneratedPaymentResolver: &gen.GeneratedPaymentResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}
