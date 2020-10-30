package gen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/v2/ast"
)

// GeneratedQueryResolver struct
type GeneratedQueryResolver struct{ *GeneratedResolver }

// QueryWalletHandlerOptions struct
type QueryWalletHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *WalletFilterType
}

// Wallet ...
func (r *GeneratedQueryResolver) Wallet(ctx context.Context, id *string, q *string, filter *WalletFilterType) (*Wallet, error) {
	opts := QueryWalletHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryWallet(ctx, r.GeneratedResolver, opts)
}

// QueryWalletHandler handler
func QueryWalletHandler(ctx context.Context, r *GeneratedResolver, opts QueryWalletHandlerOptions) (*Wallet, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := WalletQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &WalletResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("wallets")+".id = ?", *opts.ID)
	}

	var items []*Wallet
	giOpts := GetItemsOptions{
		Alias:      TableName("wallets"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryWalletsHandlerOptions struct
type QueryWalletsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*WalletSortType
	Filter *WalletFilterType
}

// Wallets ...
func (r *GeneratedQueryResolver) Wallets(ctx context.Context, offset *int, limit *int, q *string, sort []*WalletSortType, filter *WalletFilterType) (*WalletResultType, error) {
	opts := QueryWalletsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryWallets(ctx, r.GeneratedResolver, opts)
}

// QueryWalletsHandler handler
func QueryWalletsHandler(ctx context.Context, r *GeneratedResolver, opts QueryWalletsHandlerOptions) (*WalletResultType, error) {
	query := WalletQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &WalletResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedWalletResultTypeResolver struct
type GeneratedWalletResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedWalletResultTypeResolver) Items(ctx context.Context, obj *WalletResultType) (items []*Wallet, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("wallets"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Wallet{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedWalletResultTypeResolver) Count(ctx context.Context, obj *WalletResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("wallets"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Wallet{})
}

// GeneratedWalletResolver struct
type GeneratedWalletResolver struct{ *GeneratedResolver }

// WalletType ...
func (r *GeneratedWalletResolver) WalletType(ctx context.Context, obj *Wallet) (res *WalletType, err error) {
	return r.Handlers.WalletWalletType(ctx, r.GeneratedResolver, obj)
}

// WalletWalletTypeHandler handler
func WalletWalletTypeHandler(ctx context.Context, r *GeneratedResolver, obj *Wallet) (res *WalletType, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.WalletTypeID != nil {
		item, _err := loaders["WalletType"].Load(ctx, dataloader.StringKey(*obj.WalletTypeID))()
		res, _ = item.(*WalletType)

		if res == nil {
			_err = fmt.Errorf("WalletType with id '%s' not found", *obj.WalletTypeID)
		}
		err = _err
	}

	return
}

// Accounts ...
func (r *GeneratedWalletResolver) Accounts(ctx context.Context, obj *Wallet) (res []*Account, err error) {
	return r.Handlers.WalletAccounts(ctx, r.GeneratedResolver, obj)
}

// WalletAccountsHandler handler
func WalletAccountsHandler(ctx context.Context, r *GeneratedResolver, obj *Wallet) (res []*Account, err error) {

	items := []*Account{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Accounts").Error
	res = items

	return
}

// AccountsIds ...
func (r *GeneratedWalletResolver) AccountsIds(ctx context.Context, obj *Wallet) (ids []string, err error) {
	ids = []string{}

	items := []*Account{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("accounts")+".id").Related(&items, "Accounts").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// AccountsConnection method
func (r *GeneratedWalletResolver) AccountsConnection(ctx context.Context, obj *Wallet, offset *int, limit *int, q *string, sort []*AccountSortType, filter *AccountFilterType) (res *AccountResultType, err error) {
	f := &AccountFilterType{
		Wallet: &WalletFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &AccountFilterType{
			And: []*AccountFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryAccountsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryAccounts(ctx, r.GeneratedResolver, opts)
}

// Payments ...
func (r *GeneratedWalletResolver) Payments(ctx context.Context, obj *Wallet) (res []*Payment, err error) {
	return r.Handlers.WalletPayments(ctx, r.GeneratedResolver, obj)
}

// WalletPaymentsHandler handler
func WalletPaymentsHandler(ctx context.Context, r *GeneratedResolver, obj *Wallet) (res []*Payment, err error) {

	items := []*Payment{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Payments").Error
	res = items

	return
}

// PaymentsIds ...
func (r *GeneratedWalletResolver) PaymentsIds(ctx context.Context, obj *Wallet) (ids []string, err error) {
	ids = []string{}

	items := []*Payment{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("payments")+".id").Related(&items, "Payments").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// PaymentsConnection method
func (r *GeneratedWalletResolver) PaymentsConnection(ctx context.Context, obj *Wallet, offset *int, limit *int, q *string, sort []*PaymentSortType, filter *PaymentFilterType) (res *PaymentResultType, err error) {
	f := &PaymentFilterType{
		Wallet: &WalletFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &PaymentFilterType{
			And: []*PaymentFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryPaymentsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPayments(ctx, r.GeneratedResolver, opts)
}

// QueryWalletTypeHandlerOptions struct
type QueryWalletTypeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *WalletTypeFilterType
}

// WalletType ...
func (r *GeneratedQueryResolver) WalletType(ctx context.Context, id *string, q *string, filter *WalletTypeFilterType) (*WalletType, error) {
	opts := QueryWalletTypeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryWalletType(ctx, r.GeneratedResolver, opts)
}

// QueryWalletTypeHandler handler
func QueryWalletTypeHandler(ctx context.Context, r *GeneratedResolver, opts QueryWalletTypeHandlerOptions) (*WalletType, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := WalletTypeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &WalletTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("wallet_types")+".id = ?", *opts.ID)
	}

	var items []*WalletType
	giOpts := GetItemsOptions{
		Alias:      TableName("wallet_types"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryWalletTypesHandlerOptions struct
type QueryWalletTypesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*WalletTypeSortType
	Filter *WalletTypeFilterType
}

// WalletTypes ...
func (r *GeneratedQueryResolver) WalletTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*WalletTypeSortType, filter *WalletTypeFilterType) (*WalletTypeResultType, error) {
	opts := QueryWalletTypesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryWalletTypes(ctx, r.GeneratedResolver, opts)
}

// QueryWalletTypesHandler handler
func QueryWalletTypesHandler(ctx context.Context, r *GeneratedResolver, opts QueryWalletTypesHandlerOptions) (*WalletTypeResultType, error) {
	query := WalletTypeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &WalletTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedWalletTypeResultTypeResolver struct
type GeneratedWalletTypeResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedWalletTypeResultTypeResolver) Items(ctx context.Context, obj *WalletTypeResultType) (items []*WalletType, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("wallet_types"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*WalletType{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedWalletTypeResultTypeResolver) Count(ctx context.Context, obj *WalletTypeResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("wallet_types"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &WalletType{})
}

// GeneratedWalletTypeResolver struct
type GeneratedWalletTypeResolver struct{ *GeneratedResolver }

// Wallet ...
func (r *GeneratedWalletTypeResolver) Wallet(ctx context.Context, obj *WalletType) (res *Wallet, err error) {
	return r.Handlers.WalletTypeWallet(ctx, r.GeneratedResolver, obj)
}

// WalletTypeWalletHandler handler
func WalletTypeWalletHandler(ctx context.Context, r *GeneratedResolver, obj *WalletType) (res *Wallet, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.WalletID != nil {
		item, _err := loaders["Wallet"].Load(ctx, dataloader.StringKey(*obj.WalletID))()
		res, _ = item.(*Wallet)

		err = _err
	}

	return
}

// QueryAccountProviderTypeHandlerOptions struct
type QueryAccountProviderTypeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *AccountProviderTypeFilterType
}

// AccountProviderType ...
func (r *GeneratedQueryResolver) AccountProviderType(ctx context.Context, id *string, q *string, filter *AccountProviderTypeFilterType) (*AccountProviderType, error) {
	opts := QueryAccountProviderTypeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryAccountProviderType(ctx, r.GeneratedResolver, opts)
}

// QueryAccountProviderTypeHandler handler
func QueryAccountProviderTypeHandler(ctx context.Context, r *GeneratedResolver, opts QueryAccountProviderTypeHandlerOptions) (*AccountProviderType, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := AccountProviderTypeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &AccountProviderTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("account_provider_types")+".id = ?", *opts.ID)
	}

	var items []*AccountProviderType
	giOpts := GetItemsOptions{
		Alias:      TableName("account_provider_types"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryAccountProviderTypesHandlerOptions struct
type QueryAccountProviderTypesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*AccountProviderTypeSortType
	Filter *AccountProviderTypeFilterType
}

// AccountProviderTypes ...
func (r *GeneratedQueryResolver) AccountProviderTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*AccountProviderTypeSortType, filter *AccountProviderTypeFilterType) (*AccountProviderTypeResultType, error) {
	opts := QueryAccountProviderTypesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryAccountProviderTypes(ctx, r.GeneratedResolver, opts)
}

// QueryAccountProviderTypesHandler handler
func QueryAccountProviderTypesHandler(ctx context.Context, r *GeneratedResolver, opts QueryAccountProviderTypesHandlerOptions) (*AccountProviderTypeResultType, error) {
	query := AccountProviderTypeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &AccountProviderTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedAccountProviderTypeResultTypeResolver struct
type GeneratedAccountProviderTypeResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedAccountProviderTypeResultTypeResolver) Items(ctx context.Context, obj *AccountProviderTypeResultType) (items []*AccountProviderType, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("account_provider_types"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*AccountProviderType{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedAccountProviderTypeResultTypeResolver) Count(ctx context.Context, obj *AccountProviderTypeResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("account_provider_types"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &AccountProviderType{})
}

// GeneratedAccountProviderTypeResolver struct
type GeneratedAccountProviderTypeResolver struct{ *GeneratedResolver }

// AccountProvider ...
func (r *GeneratedAccountProviderTypeResolver) AccountProvider(ctx context.Context, obj *AccountProviderType) (res *AccountProvider, err error) {
	return r.Handlers.AccountProviderTypeAccountProvider(ctx, r.GeneratedResolver, obj)
}

// AccountProviderTypeAccountProviderHandler handler
func AccountProviderTypeAccountProviderHandler(ctx context.Context, r *GeneratedResolver, obj *AccountProviderType) (res *AccountProvider, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AccountProviderID != nil {
		item, _err := loaders["AccountProvider"].Load(ctx, dataloader.StringKey(*obj.AccountProviderID))()
		res, _ = item.(*AccountProvider)

		err = _err
	}

	return
}

// QueryAccountProviderHandlerOptions struct
type QueryAccountProviderHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *AccountProviderFilterType
}

// AccountProvider ...
func (r *GeneratedQueryResolver) AccountProvider(ctx context.Context, id *string, q *string, filter *AccountProviderFilterType) (*AccountProvider, error) {
	opts := QueryAccountProviderHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryAccountProvider(ctx, r.GeneratedResolver, opts)
}

// QueryAccountProviderHandler handler
func QueryAccountProviderHandler(ctx context.Context, r *GeneratedResolver, opts QueryAccountProviderHandlerOptions) (*AccountProvider, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := AccountProviderQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &AccountProviderResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("account_providers")+".id = ?", *opts.ID)
	}

	var items []*AccountProvider
	giOpts := GetItemsOptions{
		Alias:      TableName("account_providers"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryAccountProvidersHandlerOptions struct
type QueryAccountProvidersHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*AccountProviderSortType
	Filter *AccountProviderFilterType
}

// AccountProviders ...
func (r *GeneratedQueryResolver) AccountProviders(ctx context.Context, offset *int, limit *int, q *string, sort []*AccountProviderSortType, filter *AccountProviderFilterType) (*AccountProviderResultType, error) {
	opts := QueryAccountProvidersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryAccountProviders(ctx, r.GeneratedResolver, opts)
}

// QueryAccountProvidersHandler handler
func QueryAccountProvidersHandler(ctx context.Context, r *GeneratedResolver, opts QueryAccountProvidersHandlerOptions) (*AccountProviderResultType, error) {
	query := AccountProviderQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &AccountProviderResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedAccountProviderResultTypeResolver struct
type GeneratedAccountProviderResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedAccountProviderResultTypeResolver) Items(ctx context.Context, obj *AccountProviderResultType) (items []*AccountProvider, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("account_providers"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*AccountProvider{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedAccountProviderResultTypeResolver) Count(ctx context.Context, obj *AccountProviderResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("account_providers"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &AccountProvider{})
}

// GeneratedAccountProviderResolver struct
type GeneratedAccountProviderResolver struct{ *GeneratedResolver }

// Accounts ...
func (r *GeneratedAccountProviderResolver) Accounts(ctx context.Context, obj *AccountProvider) (res []*Account, err error) {
	return r.Handlers.AccountProviderAccounts(ctx, r.GeneratedResolver, obj)
}

// AccountProviderAccountsHandler handler
func AccountProviderAccountsHandler(ctx context.Context, r *GeneratedResolver, obj *AccountProvider) (res []*Account, err error) {

	items := []*Account{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Accounts").Error
	res = items

	return
}

// AccountsIds ...
func (r *GeneratedAccountProviderResolver) AccountsIds(ctx context.Context, obj *AccountProvider) (ids []string, err error) {
	ids = []string{}

	items := []*Account{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("accounts")+".id").Related(&items, "Accounts").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// AccountsConnection method
func (r *GeneratedAccountProviderResolver) AccountsConnection(ctx context.Context, obj *AccountProvider, offset *int, limit *int, q *string, sort []*AccountSortType, filter *AccountFilterType) (res *AccountResultType, err error) {
	f := &AccountFilterType{
		AccountProvider: &AccountProviderFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &AccountFilterType{
			And: []*AccountFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryAccountsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryAccounts(ctx, r.GeneratedResolver, opts)
}

// AccountProviderType ...
func (r *GeneratedAccountProviderResolver) AccountProviderType(ctx context.Context, obj *AccountProvider) (res *AccountProviderType, err error) {
	return r.Handlers.AccountProviderAccountProviderType(ctx, r.GeneratedResolver, obj)
}

// AccountProviderAccountProviderTypeHandler handler
func AccountProviderAccountProviderTypeHandler(ctx context.Context, r *GeneratedResolver, obj *AccountProvider) (res *AccountProviderType, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AccountProviderTypeID != nil {
		item, _err := loaders["AccountProviderType"].Load(ctx, dataloader.StringKey(*obj.AccountProviderTypeID))()
		res, _ = item.(*AccountProviderType)

		if res == nil {
			_err = fmt.Errorf("AccountProviderType with id '%s' not found", *obj.AccountProviderTypeID)
		}
		err = _err
	}

	return
}

// QueryAccountHandlerOptions struct
type QueryAccountHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *AccountFilterType
}

// Account ...
func (r *GeneratedQueryResolver) Account(ctx context.Context, id *string, q *string, filter *AccountFilterType) (*Account, error) {
	opts := QueryAccountHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryAccount(ctx, r.GeneratedResolver, opts)
}

// QueryAccountHandler handler
func QueryAccountHandler(ctx context.Context, r *GeneratedResolver, opts QueryAccountHandlerOptions) (*Account, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := AccountQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &AccountResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("accounts")+".id = ?", *opts.ID)
	}

	var items []*Account
	giOpts := GetItemsOptions{
		Alias:      TableName("accounts"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryAccountsHandlerOptions struct
type QueryAccountsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*AccountSortType
	Filter *AccountFilterType
}

// Accounts ...
func (r *GeneratedQueryResolver) Accounts(ctx context.Context, offset *int, limit *int, q *string, sort []*AccountSortType, filter *AccountFilterType) (*AccountResultType, error) {
	opts := QueryAccountsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryAccounts(ctx, r.GeneratedResolver, opts)
}

// QueryAccountsHandler handler
func QueryAccountsHandler(ctx context.Context, r *GeneratedResolver, opts QueryAccountsHandlerOptions) (*AccountResultType, error) {
	query := AccountQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &AccountResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedAccountResultTypeResolver struct
type GeneratedAccountResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedAccountResultTypeResolver) Items(ctx context.Context, obj *AccountResultType) (items []*Account, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("accounts"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Account{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedAccountResultTypeResolver) Count(ctx context.Context, obj *AccountResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("accounts"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Account{})
}

// GeneratedAccountResolver struct
type GeneratedAccountResolver struct{ *GeneratedResolver }

// AccountProvider ...
func (r *GeneratedAccountResolver) AccountProvider(ctx context.Context, obj *Account) (res *AccountProvider, err error) {
	return r.Handlers.AccountAccountProvider(ctx, r.GeneratedResolver, obj)
}

// AccountAccountProviderHandler handler
func AccountAccountProviderHandler(ctx context.Context, r *GeneratedResolver, obj *Account) (res *AccountProvider, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AccountProviderID != nil {
		item, _err := loaders["AccountProvider"].Load(ctx, dataloader.StringKey(*obj.AccountProviderID))()
		res, _ = item.(*AccountProvider)

		if res == nil {
			_err = fmt.Errorf("AccountProvider with id '%s' not found", *obj.AccountProviderID)
		}
		err = _err
	}

	return
}

// Wallet ...
func (r *GeneratedAccountResolver) Wallet(ctx context.Context, obj *Account) (res *Wallet, err error) {
	return r.Handlers.AccountWallet(ctx, r.GeneratedResolver, obj)
}

// AccountWalletHandler handler
func AccountWalletHandler(ctx context.Context, r *GeneratedResolver, obj *Account) (res *Wallet, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.WalletID != nil {
		item, _err := loaders["Wallet"].Load(ctx, dataloader.StringKey(*obj.WalletID))()
		res, _ = item.(*Wallet)

		if res == nil {
			_err = fmt.Errorf("Wallet with id '%s' not found", *obj.WalletID)
		}
		err = _err
	}

	return
}

// Payments ...
func (r *GeneratedAccountResolver) Payments(ctx context.Context, obj *Account) (res []*Payment, err error) {
	return r.Handlers.AccountPayments(ctx, r.GeneratedResolver, obj)
}

// AccountPaymentsHandler handler
func AccountPaymentsHandler(ctx context.Context, r *GeneratedResolver, obj *Account) (res []*Payment, err error) {

	items := []*Payment{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Payments").Error
	res = items

	return
}

// PaymentsIds ...
func (r *GeneratedAccountResolver) PaymentsIds(ctx context.Context, obj *Account) (ids []string, err error) {
	ids = []string{}

	items := []*Payment{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("payments")+".id").Related(&items, "Payments").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// PaymentsConnection method
func (r *GeneratedAccountResolver) PaymentsConnection(ctx context.Context, obj *Account, offset *int, limit *int, q *string, sort []*PaymentSortType, filter *PaymentFilterType) (res *PaymentResultType, err error) {
	f := &PaymentFilterType{
		Account: &AccountFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &PaymentFilterType{
			And: []*PaymentFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryPaymentsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPayments(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentChannelHandlerOptions struct
type QueryPaymentChannelHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PaymentChannelFilterType
}

// PaymentChannel ...
func (r *GeneratedQueryResolver) PaymentChannel(ctx context.Context, id *string, q *string, filter *PaymentChannelFilterType) (*PaymentChannel, error) {
	opts := QueryPaymentChannelHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentChannel(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentChannelHandler handler
func QueryPaymentChannelHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentChannelHandlerOptions) (*PaymentChannel, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PaymentChannelQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PaymentChannelResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("payment_channels")+".id = ?", *opts.ID)
	}

	var items []*PaymentChannel
	giOpts := GetItemsOptions{
		Alias:      TableName("payment_channels"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryPaymentChannelsHandlerOptions struct
type QueryPaymentChannelsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PaymentChannelSortType
	Filter *PaymentChannelFilterType
}

// PaymentChannels ...
func (r *GeneratedQueryResolver) PaymentChannels(ctx context.Context, offset *int, limit *int, q *string, sort []*PaymentChannelSortType, filter *PaymentChannelFilterType) (*PaymentChannelResultType, error) {
	opts := QueryPaymentChannelsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentChannels(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentChannelsHandler handler
func QueryPaymentChannelsHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentChannelsHandlerOptions) (*PaymentChannelResultType, error) {
	query := PaymentChannelQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &PaymentChannelResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedPaymentChannelResultTypeResolver struct
type GeneratedPaymentChannelResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPaymentChannelResultTypeResolver) Items(ctx context.Context, obj *PaymentChannelResultType) (items []*PaymentChannel, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("payment_channels"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*PaymentChannel{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedPaymentChannelResultTypeResolver) Count(ctx context.Context, obj *PaymentChannelResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("payment_channels"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &PaymentChannel{})
}

// GeneratedPaymentChannelResolver struct
type GeneratedPaymentChannelResolver struct{ *GeneratedResolver }

// Payment ...
func (r *GeneratedPaymentChannelResolver) Payment(ctx context.Context, obj *PaymentChannel) (res *Payment, err error) {
	return r.Handlers.PaymentChannelPayment(ctx, r.GeneratedResolver, obj)
}

// PaymentChannelPaymentHandler handler
func PaymentChannelPaymentHandler(ctx context.Context, r *GeneratedResolver, obj *PaymentChannel) (res *Payment, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PaymentID != nil {
		item, _err := loaders["Payment"].Load(ctx, dataloader.StringKey(*obj.PaymentID))()
		res, _ = item.(*Payment)

		err = _err
	}

	return
}

// QueryPaymentTypeHandlerOptions struct
type QueryPaymentTypeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PaymentTypeFilterType
}

// PaymentType ...
func (r *GeneratedQueryResolver) PaymentType(ctx context.Context, id *string, q *string, filter *PaymentTypeFilterType) (*PaymentType, error) {
	opts := QueryPaymentTypeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentType(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentTypeHandler handler
func QueryPaymentTypeHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentTypeHandlerOptions) (*PaymentType, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PaymentTypeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PaymentTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("payment_types")+".id = ?", *opts.ID)
	}

	var items []*PaymentType
	giOpts := GetItemsOptions{
		Alias:      TableName("payment_types"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryPaymentTypesHandlerOptions struct
type QueryPaymentTypesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PaymentTypeSortType
	Filter *PaymentTypeFilterType
}

// PaymentTypes ...
func (r *GeneratedQueryResolver) PaymentTypes(ctx context.Context, offset *int, limit *int, q *string, sort []*PaymentTypeSortType, filter *PaymentTypeFilterType) (*PaymentTypeResultType, error) {
	opts := QueryPaymentTypesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPaymentTypes(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentTypesHandler handler
func QueryPaymentTypesHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentTypesHandlerOptions) (*PaymentTypeResultType, error) {
	query := PaymentTypeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &PaymentTypeResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedPaymentTypeResultTypeResolver struct
type GeneratedPaymentTypeResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPaymentTypeResultTypeResolver) Items(ctx context.Context, obj *PaymentTypeResultType) (items []*PaymentType, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("payment_types"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*PaymentType{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedPaymentTypeResultTypeResolver) Count(ctx context.Context, obj *PaymentTypeResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("payment_types"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &PaymentType{})
}

// GeneratedPaymentTypeResolver struct
type GeneratedPaymentTypeResolver struct{ *GeneratedResolver }

// Payment ...
func (r *GeneratedPaymentTypeResolver) Payment(ctx context.Context, obj *PaymentType) (res *Payment, err error) {
	return r.Handlers.PaymentTypePayment(ctx, r.GeneratedResolver, obj)
}

// PaymentTypePaymentHandler handler
func PaymentTypePaymentHandler(ctx context.Context, r *GeneratedResolver, obj *PaymentType) (res *Payment, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PaymentID != nil {
		item, _err := loaders["Payment"].Load(ctx, dataloader.StringKey(*obj.PaymentID))()
		res, _ = item.(*Payment)

		err = _err
	}

	return
}

// QueryPaymentHandlerOptions struct
type QueryPaymentHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PaymentFilterType
}

// Payment ...
func (r *GeneratedQueryResolver) Payment(ctx context.Context, id *string, q *string, filter *PaymentFilterType) (*Payment, error) {
	opts := QueryPaymentHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPayment(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentHandler handler
func QueryPaymentHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentHandlerOptions) (*Payment, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PaymentQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PaymentResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("payments")+".id = ?", *opts.ID)
	}

	var items []*Payment
	giOpts := GetItemsOptions{
		Alias:      TableName("payments"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryPaymentsHandlerOptions struct
type QueryPaymentsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PaymentSortType
	Filter *PaymentFilterType
}

// Payments ...
func (r *GeneratedQueryResolver) Payments(ctx context.Context, offset *int, limit *int, q *string, sort []*PaymentSortType, filter *PaymentFilterType) (*PaymentResultType, error) {
	opts := QueryPaymentsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPayments(ctx, r.GeneratedResolver, opts)
}

// QueryPaymentsHandler handler
func QueryPaymentsHandler(ctx context.Context, r *GeneratedResolver, opts QueryPaymentsHandlerOptions) (*PaymentResultType, error) {
	query := PaymentQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &PaymentResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedPaymentResultTypeResolver struct
type GeneratedPaymentResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPaymentResultTypeResolver) Items(ctx context.Context, obj *PaymentResultType) (items []*Payment, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("payments"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Payment{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedPaymentResultTypeResolver) Count(ctx context.Context, obj *PaymentResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("payments"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Payment{})
}

// GeneratedPaymentResolver struct
type GeneratedPaymentResolver struct{ *GeneratedResolver }

// Wallet ...
func (r *GeneratedPaymentResolver) Wallet(ctx context.Context, obj *Payment) (res *Wallet, err error) {
	return r.Handlers.PaymentWallet(ctx, r.GeneratedResolver, obj)
}

// PaymentWalletHandler handler
func PaymentWalletHandler(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *Wallet, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.WalletID != nil {
		item, _err := loaders["Wallet"].Load(ctx, dataloader.StringKey(*obj.WalletID))()
		res, _ = item.(*Wallet)

		err = _err
	}

	return
}

// Account ...
func (r *GeneratedPaymentResolver) Account(ctx context.Context, obj *Payment) (res *Account, err error) {
	return r.Handlers.PaymentAccount(ctx, r.GeneratedResolver, obj)
}

// PaymentAccountHandler handler
func PaymentAccountHandler(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *Account, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AccountID != nil {
		item, _err := loaders["Account"].Load(ctx, dataloader.StringKey(*obj.AccountID))()
		res, _ = item.(*Account)

		if res == nil {
			_err = fmt.Errorf("Account with id '%s' not found", *obj.AccountID)
		}
		err = _err
	}

	return
}

// PaymentChannel ...
func (r *GeneratedPaymentResolver) PaymentChannel(ctx context.Context, obj *Payment) (res *PaymentChannel, err error) {
	return r.Handlers.PaymentPaymentChannel(ctx, r.GeneratedResolver, obj)
}

// PaymentPaymentChannelHandler handler
func PaymentPaymentChannelHandler(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *PaymentChannel, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PaymentChannelID != nil {
		item, _err := loaders["PaymentChannel"].Load(ctx, dataloader.StringKey(*obj.PaymentChannelID))()
		res, _ = item.(*PaymentChannel)

		if res == nil {
			_err = fmt.Errorf("PaymentChannel with id '%s' not found", *obj.PaymentChannelID)
		}
		err = _err
	}

	return
}

// PaymentType ...
func (r *GeneratedPaymentResolver) PaymentType(ctx context.Context, obj *Payment) (res *PaymentType, err error) {
	return r.Handlers.PaymentPaymentType(ctx, r.GeneratedResolver, obj)
}

// PaymentPaymentTypeHandler handler
func PaymentPaymentTypeHandler(ctx context.Context, r *GeneratedResolver, obj *Payment) (res *PaymentType, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.PaymentTypeID != nil {
		item, _err := loaders["PaymentType"].Load(ctx, dataloader.StringKey(*obj.PaymentTypeID))()
		res, _ = item.(*PaymentType)

		if res == nil {
			_err = fmt.Errorf("PaymentType with id '%s' not found", *obj.PaymentTypeID)
		}
		err = _err
	}

	return
}
