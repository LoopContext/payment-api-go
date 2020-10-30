package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Apply method
func (s WalletSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("wallets"), sorts, joins)
}

// ApplyWithAlias method
func (s WalletSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UserID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("userId"), Direction: s.UserID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UserIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("userId") + ")", Direction: s.UserIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UserIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("userId") + ")", Direction: s.UserIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Balance != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("balance"), Direction: s.Balance.String()}
		*sorts = append(*sorts, sort)
	}

	if s.BalanceMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("balance") + ")", Direction: s.BalanceMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.BalanceMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("balance") + ")", Direction: s.BalanceMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.BalanceAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("balance") + ")", Direction: s.BalanceAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletTypeID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("walletTypeId"), Direction: s.WalletTypeID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.WalletTypeIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("walletTypeId") + ")", Direction: s.WalletTypeIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletTypeIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("walletTypeId") + ")", Direction: s.WalletTypeIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletType != nil {
		_alias := alias + "_walletType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallet_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletTypeId"))
		err := s.WalletType.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Accounts != nil {
		_alias := alias + "_accounts"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("accounts"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("walletId")+" = "+dialect.Quote(alias)+".id")
		err := s.Accounts.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Payments != nil {
		_alias := alias + "_payments"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("walletId")+" = "+dialect.Quote(alias)+".id")
		err := s.Payments.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s WalletTypeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("wallet_types"), sorts, joins)
}

// ApplyWithAlias method
func (s WalletTypeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("walletId"), Direction: s.WalletID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.WalletIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("walletId") + ")", Direction: s.WalletIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("walletId") + ")", Direction: s.WalletIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Wallet != nil {
		_alias := alias + "_wallet"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallets"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletId"))
		err := s.Wallet.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s AccountProviderTypeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("account_provider_types"), sorts, joins)
}

// ApplyWithAlias method
func (s AccountProviderTypeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("accountProviderId"), Direction: s.AccountProviderID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("accountProviderId") + ")", Direction: s.AccountProviderIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("accountProviderId") + ")", Direction: s.AccountProviderIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProvider != nil {
		_alias := alias + "_accountProvider"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("account_providers"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountProviderId"))
		err := s.AccountProvider.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s AccountProviderSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("account_providers"), sorts, joins)
}

// ApplyWithAlias method
func (s AccountProviderSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Address != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("address"), Direction: s.Address.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AddressMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("address") + ")", Direction: s.AddressMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AddressMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("address") + ")", Direction: s.AddressMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Phone != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("phone"), Direction: s.Phone.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PhoneMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("phone") + ")", Direction: s.PhoneMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PhoneMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("phone") + ")", Direction: s.PhoneMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderTypeID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("accountProviderTypeId"), Direction: s.AccountProviderTypeID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderTypeIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("accountProviderTypeId") + ")", Direction: s.AccountProviderTypeIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderTypeIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("accountProviderTypeId") + ")", Direction: s.AccountProviderTypeIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Accounts != nil {
		_alias := alias + "_accounts"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("accounts"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("accountProviderId")+" = "+dialect.Quote(alias)+".id")
		err := s.Accounts.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.AccountProviderType != nil {
		_alias := alias + "_accountProviderType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("account_provider_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountProviderTypeId"))
		err := s.AccountProviderType.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s AccountSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("accounts"), sorts, joins)
}

// ApplyWithAlias method
func (s AccountSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountNumber != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("accountNumber"), Direction: s.AccountNumber.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AccountNumberMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("accountNumber") + ")", Direction: s.AccountNumberMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountNumberMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("accountNumber") + ")", Direction: s.AccountNumberMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Balance != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("balance"), Direction: s.Balance.String()}
		*sorts = append(*sorts, sort)
	}

	if s.BalanceMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("balance") + ")", Direction: s.BalanceMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.BalanceMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("balance") + ")", Direction: s.BalanceMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.BalanceAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("balance") + ")", Direction: s.BalanceAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("accountProviderId"), Direction: s.AccountProviderID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("accountProviderId") + ")", Direction: s.AccountProviderIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProviderIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("accountProviderId") + ")", Direction: s.AccountProviderIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("walletId"), Direction: s.WalletID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.WalletIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("walletId") + ")", Direction: s.WalletIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("walletId") + ")", Direction: s.WalletIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountProvider != nil {
		_alias := alias + "_accountProvider"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("account_providers"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountProviderId"))
		err := s.AccountProvider.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Wallet != nil {
		_alias := alias + "_wallet"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallets"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletId"))
		err := s.Wallet.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Payments != nil {
		_alias := alias + "_payments"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("accountId")+" = "+dialect.Quote(alias)+".id")
		err := s.Payments.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s PaymentChannelSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("payment_channels"), sorts, joins)
}

// ApplyWithAlias method
func (s PaymentChannelSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentId"), Direction: s.PaymentID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentId") + ")", Direction: s.PaymentIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentId") + ")", Direction: s.PaymentIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Payment != nil {
		_alias := alias + "_payment"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentId"))
		err := s.Payment.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s PaymentTypeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("payment_types"), sorts, joins)
}

// ApplyWithAlias method
func (s PaymentTypeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Description != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("description"), Direction: s.Description.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DescriptionMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("description") + ")", Direction: s.DescriptionMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentId"), Direction: s.PaymentID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentId") + ")", Direction: s.PaymentIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentId") + ")", Direction: s.PaymentIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Payment != nil {
		_alias := alias + "_payment"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentId"))
		err := s.Payment.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s PaymentSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("payments"), sorts, joins)
}

// ApplyWithAlias method
func (s PaymentSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentRef != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentRef"), Direction: s.PaymentRef.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentRefMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentRef") + ")", Direction: s.PaymentRefMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentRefMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentRef") + ")", Direction: s.PaymentRefMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Amount != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("amount"), Direction: s.Amount.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AmountMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AmountMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AmountAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("amount") + ")", Direction: s.AmountAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Concept != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("concept"), Direction: s.Concept.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ConceptMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("concept") + ")", Direction: s.ConceptMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ConceptMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("concept") + ")", Direction: s.ConceptMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("walletId"), Direction: s.WalletID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.WalletIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("walletId") + ")", Direction: s.WalletIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.WalletIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("walletId") + ")", Direction: s.WalletIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("accountId"), Direction: s.AccountID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AccountIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("accountId") + ")", Direction: s.AccountIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AccountIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("accountId") + ")", Direction: s.AccountIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentChannelID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentChannelId"), Direction: s.PaymentChannelID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentChannelIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentChannelId") + ")", Direction: s.PaymentChannelIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentChannelIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentChannelId") + ")", Direction: s.PaymentChannelIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentTypeID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("paymentTypeId"), Direction: s.PaymentTypeID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentTypeIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("paymentTypeId") + ")", Direction: s.PaymentTypeIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PaymentTypeIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("paymentTypeId") + ")", Direction: s.PaymentTypeIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Wallet != nil {
		_alias := alias + "_wallet"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallets"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletId"))
		err := s.Wallet.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Account != nil {
		_alias := alias + "_account"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("accounts"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountId"))
		err := s.Account.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.PaymentChannel != nil {
		_alias := alias + "_paymentChannel"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payment_channels"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentChannelId"))
		err := s.PaymentChannel.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.PaymentType != nil {
		_alias := alias + "_paymentType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payment_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentTypeId"))
		err := s.PaymentType.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
