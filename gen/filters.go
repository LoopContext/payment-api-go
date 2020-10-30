package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// IsEmpty ...
func (f *WalletFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *WalletFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("wallets"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *WalletFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.WalletType != nil {
		_alias := alias + "_walletType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallet_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletTypeId"))
		err := f.WalletType.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Accounts != nil {
		_alias := alias + "_accounts"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("accounts"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("walletId")+" = "+dialect.Quote(alias)+".id")
		err := f.Accounts.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Payments != nil {
		_alias := alias + "_payments"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("walletId")+" = "+dialect.Quote(alias)+".id")
		err := f.Payments.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *WalletFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.UserID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" = ?")
		values = append(values, f.UserID)
	}

	if f.UserIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" != ?")
		values = append(values, f.UserIDNe)
	}

	if f.UserIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" > ?")
		values = append(values, f.UserIDGt)
	}

	if f.UserIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" < ?")
		values = append(values, f.UserIDLt)
	}

	if f.UserIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" >= ?")
		values = append(values, f.UserIDGte)
	}

	if f.UserIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" <= ?")
		values = append(values, f.UserIDLte)
	}

	if f.UserIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IN (?)")
		values = append(values, f.UserIDIn)
	}

	if f.UserIDNull != nil {
		if *f.UserIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("userId")+" IS NOT NULL")
		}
	}

	if f.Balance != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" = ?")
		values = append(values, f.Balance)
	}

	if f.BalanceNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" != ?")
		values = append(values, f.BalanceNe)
	}

	if f.BalanceGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" > ?")
		values = append(values, f.BalanceGt)
	}

	if f.BalanceLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" < ?")
		values = append(values, f.BalanceLt)
	}

	if f.BalanceGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" >= ?")
		values = append(values, f.BalanceGte)
	}

	if f.BalanceLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" <= ?")
		values = append(values, f.BalanceLte)
	}

	if f.BalanceIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" IN (?)")
		values = append(values, f.BalanceIn)
	}

	if f.BalanceNull != nil {
		if *f.BalanceNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" IS NOT NULL")
		}
	}

	if f.WalletTypeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" = ?")
		values = append(values, f.WalletTypeID)
	}

	if f.WalletTypeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" != ?")
		values = append(values, f.WalletTypeIDNe)
	}

	if f.WalletTypeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" > ?")
		values = append(values, f.WalletTypeIDGt)
	}

	if f.WalletTypeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" < ?")
		values = append(values, f.WalletTypeIDLt)
	}

	if f.WalletTypeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" >= ?")
		values = append(values, f.WalletTypeIDGte)
	}

	if f.WalletTypeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" <= ?")
		values = append(values, f.WalletTypeIDLte)
	}

	if f.WalletTypeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" IN (?)")
		values = append(values, f.WalletTypeIDIn)
	}

	if f.WalletTypeIDNull != nil {
		if *f.WalletTypeIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletTypeId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *WalletFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.UserIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") = ?")
		values = append(values, f.UserIDMin)
	}

	if f.UserIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") = ?")
		values = append(values, f.UserIDMax)
	}

	if f.UserIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") != ?")
		values = append(values, f.UserIDMinNe)
	}

	if f.UserIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") != ?")
		values = append(values, f.UserIDMaxNe)
	}

	if f.UserIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") > ?")
		values = append(values, f.UserIDMinGt)
	}

	if f.UserIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") > ?")
		values = append(values, f.UserIDMaxGt)
	}

	if f.UserIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") < ?")
		values = append(values, f.UserIDMinLt)
	}

	if f.UserIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") < ?")
		values = append(values, f.UserIDMaxLt)
	}

	if f.UserIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") >= ?")
		values = append(values, f.UserIDMinGte)
	}

	if f.UserIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") >= ?")
		values = append(values, f.UserIDMaxGte)
	}

	if f.UserIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") <= ?")
		values = append(values, f.UserIDMinLte)
	}

	if f.UserIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") <= ?")
		values = append(values, f.UserIDMaxLte)
	}

	if f.UserIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("userId")+") IN (?)")
		values = append(values, f.UserIDMinIn)
	}

	if f.UserIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("userId")+") IN (?)")
		values = append(values, f.UserIDMaxIn)
	}

	if f.BalanceMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") = ?")
		values = append(values, f.BalanceMin)
	}

	if f.BalanceMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") = ?")
		values = append(values, f.BalanceMax)
	}

	if f.BalanceAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") = ?")
		values = append(values, f.BalanceAvg)
	}

	if f.BalanceMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") != ?")
		values = append(values, f.BalanceMinNe)
	}

	if f.BalanceMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") != ?")
		values = append(values, f.BalanceMaxNe)
	}

	if f.BalanceAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") != ?")
		values = append(values, f.BalanceAvgNe)
	}

	if f.BalanceMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") > ?")
		values = append(values, f.BalanceMinGt)
	}

	if f.BalanceMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") > ?")
		values = append(values, f.BalanceMaxGt)
	}

	if f.BalanceAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") > ?")
		values = append(values, f.BalanceAvgGt)
	}

	if f.BalanceMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") < ?")
		values = append(values, f.BalanceMinLt)
	}

	if f.BalanceMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") < ?")
		values = append(values, f.BalanceMaxLt)
	}

	if f.BalanceAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") < ?")
		values = append(values, f.BalanceAvgLt)
	}

	if f.BalanceMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") >= ?")
		values = append(values, f.BalanceMinGte)
	}

	if f.BalanceMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") >= ?")
		values = append(values, f.BalanceMaxGte)
	}

	if f.BalanceAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") >= ?")
		values = append(values, f.BalanceAvgGte)
	}

	if f.BalanceMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") <= ?")
		values = append(values, f.BalanceMinLte)
	}

	if f.BalanceMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") <= ?")
		values = append(values, f.BalanceMaxLte)
	}

	if f.BalanceAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") <= ?")
		values = append(values, f.BalanceAvgLte)
	}

	if f.BalanceMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") IN (?)")
		values = append(values, f.BalanceMinIn)
	}

	if f.BalanceMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") IN (?)")
		values = append(values, f.BalanceMaxIn)
	}

	if f.BalanceAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") IN (?)")
		values = append(values, f.BalanceAvgIn)
	}

	if f.WalletTypeIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletTypeId")+") = ?")
		values = append(values, f.WalletTypeIDMin)
	}

	if f.WalletTypeIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletTypeId")+") = ?")
		values = append(values, f.WalletTypeIDMax)
	}

	if f.WalletTypeIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletTypeId")+") != ?")
		values = append(values, f.WalletTypeIDMinNe)
	}

	if f.WalletTypeIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletTypeId")+") != ?")
		values = append(values, f.WalletTypeIDMaxNe)
	}

	if f.WalletTypeIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletTypeId")+") > ?")
		values = append(values, f.WalletTypeIDMinGt)
	}

	if f.WalletTypeIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletTypeId")+") > ?")
		values = append(values, f.WalletTypeIDMaxGt)
	}

	if f.WalletTypeIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletTypeId")+") < ?")
		values = append(values, f.WalletTypeIDMinLt)
	}

	if f.WalletTypeIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletTypeId")+") < ?")
		values = append(values, f.WalletTypeIDMaxLt)
	}

	if f.WalletTypeIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletTypeId")+") >= ?")
		values = append(values, f.WalletTypeIDMinGte)
	}

	if f.WalletTypeIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletTypeId")+") >= ?")
		values = append(values, f.WalletTypeIDMaxGte)
	}

	if f.WalletTypeIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletTypeId")+") <= ?")
		values = append(values, f.WalletTypeIDMinLte)
	}

	if f.WalletTypeIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletTypeId")+") <= ?")
		values = append(values, f.WalletTypeIDMaxLte)
	}

	if f.WalletTypeIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletTypeId")+") IN (?)")
		values = append(values, f.WalletTypeIDMinIn)
	}

	if f.WalletTypeIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletTypeId")+") IN (?)")
		values = append(values, f.WalletTypeIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *WalletFilterType) AndWith(f2 ...*WalletFilterType) *WalletFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &WalletFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *WalletFilterType) OrWith(f2 ...*WalletFilterType) *WalletFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &WalletFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *WalletTypeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *WalletTypeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("wallet_types"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *WalletTypeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Wallet != nil {
		_alias := alias + "_wallet"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallets"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletId"))
		err := f.Wallet.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *WalletTypeFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.WalletID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" = ?")
		values = append(values, f.WalletID)
	}

	if f.WalletIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" != ?")
		values = append(values, f.WalletIDNe)
	}

	if f.WalletIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" > ?")
		values = append(values, f.WalletIDGt)
	}

	if f.WalletIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" < ?")
		values = append(values, f.WalletIDLt)
	}

	if f.WalletIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" >= ?")
		values = append(values, f.WalletIDGte)
	}

	if f.WalletIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" <= ?")
		values = append(values, f.WalletIDLte)
	}

	if f.WalletIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IN (?)")
		values = append(values, f.WalletIDIn)
	}

	if f.WalletIDNull != nil {
		if *f.WalletIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *WalletTypeFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.WalletIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") = ?")
		values = append(values, f.WalletIDMin)
	}

	if f.WalletIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") = ?")
		values = append(values, f.WalletIDMax)
	}

	if f.WalletIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") != ?")
		values = append(values, f.WalletIDMinNe)
	}

	if f.WalletIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") != ?")
		values = append(values, f.WalletIDMaxNe)
	}

	if f.WalletIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") > ?")
		values = append(values, f.WalletIDMinGt)
	}

	if f.WalletIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") > ?")
		values = append(values, f.WalletIDMaxGt)
	}

	if f.WalletIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") < ?")
		values = append(values, f.WalletIDMinLt)
	}

	if f.WalletIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") < ?")
		values = append(values, f.WalletIDMaxLt)
	}

	if f.WalletIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") >= ?")
		values = append(values, f.WalletIDMinGte)
	}

	if f.WalletIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") >= ?")
		values = append(values, f.WalletIDMaxGte)
	}

	if f.WalletIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") <= ?")
		values = append(values, f.WalletIDMinLte)
	}

	if f.WalletIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") <= ?")
		values = append(values, f.WalletIDMaxLte)
	}

	if f.WalletIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") IN (?)")
		values = append(values, f.WalletIDMinIn)
	}

	if f.WalletIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") IN (?)")
		values = append(values, f.WalletIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *WalletTypeFilterType) AndWith(f2 ...*WalletTypeFilterType) *WalletTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &WalletTypeFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *WalletTypeFilterType) OrWith(f2 ...*WalletTypeFilterType) *WalletTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &WalletTypeFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *AccountProviderTypeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *AccountProviderTypeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("account_provider_types"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *AccountProviderTypeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.AccountProvider != nil {
		_alias := alias + "_accountProvider"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("account_providers"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountProviderId"))
		err := f.AccountProvider.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *AccountProviderTypeFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.AccountProviderID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" = ?")
		values = append(values, f.AccountProviderID)
	}

	if f.AccountProviderIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" != ?")
		values = append(values, f.AccountProviderIDNe)
	}

	if f.AccountProviderIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" > ?")
		values = append(values, f.AccountProviderIDGt)
	}

	if f.AccountProviderIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" < ?")
		values = append(values, f.AccountProviderIDLt)
	}

	if f.AccountProviderIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" >= ?")
		values = append(values, f.AccountProviderIDGte)
	}

	if f.AccountProviderIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" <= ?")
		values = append(values, f.AccountProviderIDLte)
	}

	if f.AccountProviderIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" IN (?)")
		values = append(values, f.AccountProviderIDIn)
	}

	if f.AccountProviderIDNull != nil {
		if *f.AccountProviderIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *AccountProviderTypeFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.AccountProviderIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") = ?")
		values = append(values, f.AccountProviderIDMin)
	}

	if f.AccountProviderIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") = ?")
		values = append(values, f.AccountProviderIDMax)
	}

	if f.AccountProviderIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") != ?")
		values = append(values, f.AccountProviderIDMinNe)
	}

	if f.AccountProviderIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") != ?")
		values = append(values, f.AccountProviderIDMaxNe)
	}

	if f.AccountProviderIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") > ?")
		values = append(values, f.AccountProviderIDMinGt)
	}

	if f.AccountProviderIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") > ?")
		values = append(values, f.AccountProviderIDMaxGt)
	}

	if f.AccountProviderIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") < ?")
		values = append(values, f.AccountProviderIDMinLt)
	}

	if f.AccountProviderIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") < ?")
		values = append(values, f.AccountProviderIDMaxLt)
	}

	if f.AccountProviderIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") >= ?")
		values = append(values, f.AccountProviderIDMinGte)
	}

	if f.AccountProviderIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") >= ?")
		values = append(values, f.AccountProviderIDMaxGte)
	}

	if f.AccountProviderIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") <= ?")
		values = append(values, f.AccountProviderIDMinLte)
	}

	if f.AccountProviderIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") <= ?")
		values = append(values, f.AccountProviderIDMaxLte)
	}

	if f.AccountProviderIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") IN (?)")
		values = append(values, f.AccountProviderIDMinIn)
	}

	if f.AccountProviderIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") IN (?)")
		values = append(values, f.AccountProviderIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *AccountProviderTypeFilterType) AndWith(f2 ...*AccountProviderTypeFilterType) *AccountProviderTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AccountProviderTypeFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *AccountProviderTypeFilterType) OrWith(f2 ...*AccountProviderTypeFilterType) *AccountProviderTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AccountProviderTypeFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *AccountProviderFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *AccountProviderFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("account_providers"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *AccountProviderFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Accounts != nil {
		_alias := alias + "_accounts"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("accounts"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("accountProviderId")+" = "+dialect.Quote(alias)+".id")
		err := f.Accounts.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.AccountProviderType != nil {
		_alias := alias + "_accountProviderType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("account_provider_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountProviderTypeId"))
		err := f.AccountProviderType.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *AccountProviderFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.Address != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" = ?")
		values = append(values, f.Address)
	}

	if f.AddressNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" != ?")
		values = append(values, f.AddressNe)
	}

	if f.AddressGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" > ?")
		values = append(values, f.AddressGt)
	}

	if f.AddressLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" < ?")
		values = append(values, f.AddressLt)
	}

	if f.AddressGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" >= ?")
		values = append(values, f.AddressGte)
	}

	if f.AddressLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" <= ?")
		values = append(values, f.AddressLte)
	}

	if f.AddressIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" IN (?)")
		values = append(values, f.AddressIn)
	}

	if f.AddressLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AddressLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AddressPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AddressPrefix))
	}

	if f.AddressSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AddressSuffix))
	}

	if f.AddressNull != nil {
		if *f.AddressNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("address")+" IS NOT NULL")
		}
	}

	if f.Phone != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" = ?")
		values = append(values, f.Phone)
	}

	if f.PhoneNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" != ?")
		values = append(values, f.PhoneNe)
	}

	if f.PhoneGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" > ?")
		values = append(values, f.PhoneGt)
	}

	if f.PhoneLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" < ?")
		values = append(values, f.PhoneLt)
	}

	if f.PhoneGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" >= ?")
		values = append(values, f.PhoneGte)
	}

	if f.PhoneLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" <= ?")
		values = append(values, f.PhoneLte)
	}

	if f.PhoneIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" IN (?)")
		values = append(values, f.PhoneIn)
	}

	if f.PhoneLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PhoneLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PhonePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhonePrefix))
	}

	if f.PhoneSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneSuffix))
	}

	if f.PhoneNull != nil {
		if *f.PhoneNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("phone")+" IS NOT NULL")
		}
	}

	if f.AccountProviderTypeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" = ?")
		values = append(values, f.AccountProviderTypeID)
	}

	if f.AccountProviderTypeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" != ?")
		values = append(values, f.AccountProviderTypeIDNe)
	}

	if f.AccountProviderTypeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" > ?")
		values = append(values, f.AccountProviderTypeIDGt)
	}

	if f.AccountProviderTypeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" < ?")
		values = append(values, f.AccountProviderTypeIDLt)
	}

	if f.AccountProviderTypeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" >= ?")
		values = append(values, f.AccountProviderTypeIDGte)
	}

	if f.AccountProviderTypeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" <= ?")
		values = append(values, f.AccountProviderTypeIDLte)
	}

	if f.AccountProviderTypeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" IN (?)")
		values = append(values, f.AccountProviderTypeIDIn)
	}

	if f.AccountProviderTypeIDNull != nil {
		if *f.AccountProviderTypeIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderTypeId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *AccountProviderFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.AddressMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") = ?")
		values = append(values, f.AddressMin)
	}

	if f.AddressMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") = ?")
		values = append(values, f.AddressMax)
	}

	if f.AddressMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") != ?")
		values = append(values, f.AddressMinNe)
	}

	if f.AddressMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") != ?")
		values = append(values, f.AddressMaxNe)
	}

	if f.AddressMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") > ?")
		values = append(values, f.AddressMinGt)
	}

	if f.AddressMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") > ?")
		values = append(values, f.AddressMaxGt)
	}

	if f.AddressMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") < ?")
		values = append(values, f.AddressMinLt)
	}

	if f.AddressMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") < ?")
		values = append(values, f.AddressMaxLt)
	}

	if f.AddressMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") >= ?")
		values = append(values, f.AddressMinGte)
	}

	if f.AddressMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") >= ?")
		values = append(values, f.AddressMaxGte)
	}

	if f.AddressMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") <= ?")
		values = append(values, f.AddressMinLte)
	}

	if f.AddressMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") <= ?")
		values = append(values, f.AddressMaxLte)
	}

	if f.AddressMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") IN (?)")
		values = append(values, f.AddressMinIn)
	}

	if f.AddressMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") IN (?)")
		values = append(values, f.AddressMaxIn)
	}

	if f.AddressMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AddressMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AddressMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AddressMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AddressMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AddressMinPrefix))
	}

	if f.AddressMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AddressMaxPrefix))
	}

	if f.AddressMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("address")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AddressMinSuffix))
	}

	if f.AddressMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("address")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AddressMaxSuffix))
	}

	if f.PhoneMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") = ?")
		values = append(values, f.PhoneMin)
	}

	if f.PhoneMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") = ?")
		values = append(values, f.PhoneMax)
	}

	if f.PhoneMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") != ?")
		values = append(values, f.PhoneMinNe)
	}

	if f.PhoneMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") != ?")
		values = append(values, f.PhoneMaxNe)
	}

	if f.PhoneMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") > ?")
		values = append(values, f.PhoneMinGt)
	}

	if f.PhoneMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") > ?")
		values = append(values, f.PhoneMaxGt)
	}

	if f.PhoneMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") < ?")
		values = append(values, f.PhoneMinLt)
	}

	if f.PhoneMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") < ?")
		values = append(values, f.PhoneMaxLt)
	}

	if f.PhoneMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") >= ?")
		values = append(values, f.PhoneMinGte)
	}

	if f.PhoneMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") >= ?")
		values = append(values, f.PhoneMaxGte)
	}

	if f.PhoneMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") <= ?")
		values = append(values, f.PhoneMinLte)
	}

	if f.PhoneMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") <= ?")
		values = append(values, f.PhoneMaxLte)
	}

	if f.PhoneMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") IN (?)")
		values = append(values, f.PhoneMinIn)
	}

	if f.PhoneMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") IN (?)")
		values = append(values, f.PhoneMaxIn)
	}

	if f.PhoneMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PhoneMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PhoneMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PhoneMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PhoneMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhoneMinPrefix))
	}

	if f.PhoneMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhoneMaxPrefix))
	}

	if f.PhoneMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneMinSuffix))
	}

	if f.PhoneMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("phone")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneMaxSuffix))
	}

	if f.AccountProviderTypeIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") = ?")
		values = append(values, f.AccountProviderTypeIDMin)
	}

	if f.AccountProviderTypeIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") = ?")
		values = append(values, f.AccountProviderTypeIDMax)
	}

	if f.AccountProviderTypeIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") != ?")
		values = append(values, f.AccountProviderTypeIDMinNe)
	}

	if f.AccountProviderTypeIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") != ?")
		values = append(values, f.AccountProviderTypeIDMaxNe)
	}

	if f.AccountProviderTypeIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") > ?")
		values = append(values, f.AccountProviderTypeIDMinGt)
	}

	if f.AccountProviderTypeIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") > ?")
		values = append(values, f.AccountProviderTypeIDMaxGt)
	}

	if f.AccountProviderTypeIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") < ?")
		values = append(values, f.AccountProviderTypeIDMinLt)
	}

	if f.AccountProviderTypeIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") < ?")
		values = append(values, f.AccountProviderTypeIDMaxLt)
	}

	if f.AccountProviderTypeIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") >= ?")
		values = append(values, f.AccountProviderTypeIDMinGte)
	}

	if f.AccountProviderTypeIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") >= ?")
		values = append(values, f.AccountProviderTypeIDMaxGte)
	}

	if f.AccountProviderTypeIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") <= ?")
		values = append(values, f.AccountProviderTypeIDMinLte)
	}

	if f.AccountProviderTypeIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") <= ?")
		values = append(values, f.AccountProviderTypeIDMaxLte)
	}

	if f.AccountProviderTypeIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") IN (?)")
		values = append(values, f.AccountProviderTypeIDMinIn)
	}

	if f.AccountProviderTypeIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderTypeId")+") IN (?)")
		values = append(values, f.AccountProviderTypeIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *AccountProviderFilterType) AndWith(f2 ...*AccountProviderFilterType) *AccountProviderFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AccountProviderFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *AccountProviderFilterType) OrWith(f2 ...*AccountProviderFilterType) *AccountProviderFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AccountProviderFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *AccountFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *AccountFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("accounts"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *AccountFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.AccountProvider != nil {
		_alias := alias + "_accountProvider"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("account_providers"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountProviderId"))
		err := f.AccountProvider.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Wallet != nil {
		_alias := alias + "_wallet"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallets"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletId"))
		err := f.Wallet.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Payments != nil {
		_alias := alias + "_payments"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("accountId")+" = "+dialect.Quote(alias)+".id")
		err := f.Payments.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *AccountFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.AccountNumber != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" = ?")
		values = append(values, f.AccountNumber)
	}

	if f.AccountNumberNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" != ?")
		values = append(values, f.AccountNumberNe)
	}

	if f.AccountNumberGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" > ?")
		values = append(values, f.AccountNumberGt)
	}

	if f.AccountNumberLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" < ?")
		values = append(values, f.AccountNumberLt)
	}

	if f.AccountNumberGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" >= ?")
		values = append(values, f.AccountNumberGte)
	}

	if f.AccountNumberLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" <= ?")
		values = append(values, f.AccountNumberLte)
	}

	if f.AccountNumberIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" IN (?)")
		values = append(values, f.AccountNumberIn)
	}

	if f.AccountNumberLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AccountNumberLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AccountNumberPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AccountNumberPrefix))
	}

	if f.AccountNumberSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AccountNumberSuffix))
	}

	if f.AccountNumberNull != nil {
		if *f.AccountNumberNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountNumber")+" IS NOT NULL")
		}
	}

	if f.Balance != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" = ?")
		values = append(values, f.Balance)
	}

	if f.BalanceNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" != ?")
		values = append(values, f.BalanceNe)
	}

	if f.BalanceGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" > ?")
		values = append(values, f.BalanceGt)
	}

	if f.BalanceLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" < ?")
		values = append(values, f.BalanceLt)
	}

	if f.BalanceGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" >= ?")
		values = append(values, f.BalanceGte)
	}

	if f.BalanceLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" <= ?")
		values = append(values, f.BalanceLte)
	}

	if f.BalanceIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" IN (?)")
		values = append(values, f.BalanceIn)
	}

	if f.BalanceNull != nil {
		if *f.BalanceNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("balance")+" IS NOT NULL")
		}
	}

	if f.AccountProviderID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" = ?")
		values = append(values, f.AccountProviderID)
	}

	if f.AccountProviderIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" != ?")
		values = append(values, f.AccountProviderIDNe)
	}

	if f.AccountProviderIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" > ?")
		values = append(values, f.AccountProviderIDGt)
	}

	if f.AccountProviderIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" < ?")
		values = append(values, f.AccountProviderIDLt)
	}

	if f.AccountProviderIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" >= ?")
		values = append(values, f.AccountProviderIDGte)
	}

	if f.AccountProviderIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" <= ?")
		values = append(values, f.AccountProviderIDLte)
	}

	if f.AccountProviderIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" IN (?)")
		values = append(values, f.AccountProviderIDIn)
	}

	if f.AccountProviderIDNull != nil {
		if *f.AccountProviderIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountProviderId")+" IS NOT NULL")
		}
	}

	if f.WalletID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" = ?")
		values = append(values, f.WalletID)
	}

	if f.WalletIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" != ?")
		values = append(values, f.WalletIDNe)
	}

	if f.WalletIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" > ?")
		values = append(values, f.WalletIDGt)
	}

	if f.WalletIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" < ?")
		values = append(values, f.WalletIDLt)
	}

	if f.WalletIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" >= ?")
		values = append(values, f.WalletIDGte)
	}

	if f.WalletIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" <= ?")
		values = append(values, f.WalletIDLte)
	}

	if f.WalletIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IN (?)")
		values = append(values, f.WalletIDIn)
	}

	if f.WalletIDNull != nil {
		if *f.WalletIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *AccountFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.AccountNumberMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") = ?")
		values = append(values, f.AccountNumberMin)
	}

	if f.AccountNumberMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") = ?")
		values = append(values, f.AccountNumberMax)
	}

	if f.AccountNumberMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") != ?")
		values = append(values, f.AccountNumberMinNe)
	}

	if f.AccountNumberMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") != ?")
		values = append(values, f.AccountNumberMaxNe)
	}

	if f.AccountNumberMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") > ?")
		values = append(values, f.AccountNumberMinGt)
	}

	if f.AccountNumberMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") > ?")
		values = append(values, f.AccountNumberMaxGt)
	}

	if f.AccountNumberMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") < ?")
		values = append(values, f.AccountNumberMinLt)
	}

	if f.AccountNumberMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") < ?")
		values = append(values, f.AccountNumberMaxLt)
	}

	if f.AccountNumberMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") >= ?")
		values = append(values, f.AccountNumberMinGte)
	}

	if f.AccountNumberMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") >= ?")
		values = append(values, f.AccountNumberMaxGte)
	}

	if f.AccountNumberMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") <= ?")
		values = append(values, f.AccountNumberMinLte)
	}

	if f.AccountNumberMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") <= ?")
		values = append(values, f.AccountNumberMaxLte)
	}

	if f.AccountNumberMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") IN (?)")
		values = append(values, f.AccountNumberMinIn)
	}

	if f.AccountNumberMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") IN (?)")
		values = append(values, f.AccountNumberMaxIn)
	}

	if f.AccountNumberMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AccountNumberMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AccountNumberMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AccountNumberMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AccountNumberMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AccountNumberMinPrefix))
	}

	if f.AccountNumberMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AccountNumberMaxPrefix))
	}

	if f.AccountNumberMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountNumber")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AccountNumberMinSuffix))
	}

	if f.AccountNumberMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountNumber")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AccountNumberMaxSuffix))
	}

	if f.BalanceMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") = ?")
		values = append(values, f.BalanceMin)
	}

	if f.BalanceMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") = ?")
		values = append(values, f.BalanceMax)
	}

	if f.BalanceAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") = ?")
		values = append(values, f.BalanceAvg)
	}

	if f.BalanceMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") != ?")
		values = append(values, f.BalanceMinNe)
	}

	if f.BalanceMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") != ?")
		values = append(values, f.BalanceMaxNe)
	}

	if f.BalanceAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") != ?")
		values = append(values, f.BalanceAvgNe)
	}

	if f.BalanceMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") > ?")
		values = append(values, f.BalanceMinGt)
	}

	if f.BalanceMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") > ?")
		values = append(values, f.BalanceMaxGt)
	}

	if f.BalanceAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") > ?")
		values = append(values, f.BalanceAvgGt)
	}

	if f.BalanceMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") < ?")
		values = append(values, f.BalanceMinLt)
	}

	if f.BalanceMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") < ?")
		values = append(values, f.BalanceMaxLt)
	}

	if f.BalanceAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") < ?")
		values = append(values, f.BalanceAvgLt)
	}

	if f.BalanceMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") >= ?")
		values = append(values, f.BalanceMinGte)
	}

	if f.BalanceMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") >= ?")
		values = append(values, f.BalanceMaxGte)
	}

	if f.BalanceAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") >= ?")
		values = append(values, f.BalanceAvgGte)
	}

	if f.BalanceMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") <= ?")
		values = append(values, f.BalanceMinLte)
	}

	if f.BalanceMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") <= ?")
		values = append(values, f.BalanceMaxLte)
	}

	if f.BalanceAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") <= ?")
		values = append(values, f.BalanceAvgLte)
	}

	if f.BalanceMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("balance")+") IN (?)")
		values = append(values, f.BalanceMinIn)
	}

	if f.BalanceMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("balance")+") IN (?)")
		values = append(values, f.BalanceMaxIn)
	}

	if f.BalanceAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("balance")+") IN (?)")
		values = append(values, f.BalanceAvgIn)
	}

	if f.AccountProviderIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") = ?")
		values = append(values, f.AccountProviderIDMin)
	}

	if f.AccountProviderIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") = ?")
		values = append(values, f.AccountProviderIDMax)
	}

	if f.AccountProviderIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") != ?")
		values = append(values, f.AccountProviderIDMinNe)
	}

	if f.AccountProviderIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") != ?")
		values = append(values, f.AccountProviderIDMaxNe)
	}

	if f.AccountProviderIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") > ?")
		values = append(values, f.AccountProviderIDMinGt)
	}

	if f.AccountProviderIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") > ?")
		values = append(values, f.AccountProviderIDMaxGt)
	}

	if f.AccountProviderIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") < ?")
		values = append(values, f.AccountProviderIDMinLt)
	}

	if f.AccountProviderIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") < ?")
		values = append(values, f.AccountProviderIDMaxLt)
	}

	if f.AccountProviderIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") >= ?")
		values = append(values, f.AccountProviderIDMinGte)
	}

	if f.AccountProviderIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") >= ?")
		values = append(values, f.AccountProviderIDMaxGte)
	}

	if f.AccountProviderIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") <= ?")
		values = append(values, f.AccountProviderIDMinLte)
	}

	if f.AccountProviderIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") <= ?")
		values = append(values, f.AccountProviderIDMaxLte)
	}

	if f.AccountProviderIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountProviderId")+") IN (?)")
		values = append(values, f.AccountProviderIDMinIn)
	}

	if f.AccountProviderIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountProviderId")+") IN (?)")
		values = append(values, f.AccountProviderIDMaxIn)
	}

	if f.WalletIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") = ?")
		values = append(values, f.WalletIDMin)
	}

	if f.WalletIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") = ?")
		values = append(values, f.WalletIDMax)
	}

	if f.WalletIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") != ?")
		values = append(values, f.WalletIDMinNe)
	}

	if f.WalletIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") != ?")
		values = append(values, f.WalletIDMaxNe)
	}

	if f.WalletIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") > ?")
		values = append(values, f.WalletIDMinGt)
	}

	if f.WalletIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") > ?")
		values = append(values, f.WalletIDMaxGt)
	}

	if f.WalletIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") < ?")
		values = append(values, f.WalletIDMinLt)
	}

	if f.WalletIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") < ?")
		values = append(values, f.WalletIDMaxLt)
	}

	if f.WalletIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") >= ?")
		values = append(values, f.WalletIDMinGte)
	}

	if f.WalletIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") >= ?")
		values = append(values, f.WalletIDMaxGte)
	}

	if f.WalletIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") <= ?")
		values = append(values, f.WalletIDMinLte)
	}

	if f.WalletIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") <= ?")
		values = append(values, f.WalletIDMaxLte)
	}

	if f.WalletIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") IN (?)")
		values = append(values, f.WalletIDMinIn)
	}

	if f.WalletIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") IN (?)")
		values = append(values, f.WalletIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *AccountFilterType) AndWith(f2 ...*AccountFilterType) *AccountFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AccountFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *AccountFilterType) OrWith(f2 ...*AccountFilterType) *AccountFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AccountFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *PaymentChannelFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *PaymentChannelFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("payment_channels"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *PaymentChannelFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Payment != nil {
		_alias := alias + "_payment"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentId"))
		err := f.Payment.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *PaymentChannelFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.PaymentID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" = ?")
		values = append(values, f.PaymentID)
	}

	if f.PaymentIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" != ?")
		values = append(values, f.PaymentIDNe)
	}

	if f.PaymentIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" > ?")
		values = append(values, f.PaymentIDGt)
	}

	if f.PaymentIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" < ?")
		values = append(values, f.PaymentIDLt)
	}

	if f.PaymentIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" >= ?")
		values = append(values, f.PaymentIDGte)
	}

	if f.PaymentIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" <= ?")
		values = append(values, f.PaymentIDLte)
	}

	if f.PaymentIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IN (?)")
		values = append(values, f.PaymentIDIn)
	}

	if f.PaymentIDNull != nil {
		if *f.PaymentIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *PaymentChannelFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.PaymentIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") = ?")
		values = append(values, f.PaymentIDMin)
	}

	if f.PaymentIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") = ?")
		values = append(values, f.PaymentIDMax)
	}

	if f.PaymentIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") != ?")
		values = append(values, f.PaymentIDMinNe)
	}

	if f.PaymentIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") != ?")
		values = append(values, f.PaymentIDMaxNe)
	}

	if f.PaymentIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") > ?")
		values = append(values, f.PaymentIDMinGt)
	}

	if f.PaymentIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") > ?")
		values = append(values, f.PaymentIDMaxGt)
	}

	if f.PaymentIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") < ?")
		values = append(values, f.PaymentIDMinLt)
	}

	if f.PaymentIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") < ?")
		values = append(values, f.PaymentIDMaxLt)
	}

	if f.PaymentIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") >= ?")
		values = append(values, f.PaymentIDMinGte)
	}

	if f.PaymentIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") >= ?")
		values = append(values, f.PaymentIDMaxGte)
	}

	if f.PaymentIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") <= ?")
		values = append(values, f.PaymentIDMinLte)
	}

	if f.PaymentIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") <= ?")
		values = append(values, f.PaymentIDMaxLte)
	}

	if f.PaymentIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") IN (?)")
		values = append(values, f.PaymentIDMinIn)
	}

	if f.PaymentIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") IN (?)")
		values = append(values, f.PaymentIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *PaymentChannelFilterType) AndWith(f2 ...*PaymentChannelFilterType) *PaymentChannelFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PaymentChannelFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *PaymentChannelFilterType) OrWith(f2 ...*PaymentChannelFilterType) *PaymentChannelFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PaymentChannelFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *PaymentTypeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *PaymentTypeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("payment_types"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *PaymentTypeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Payment != nil {
		_alias := alias + "_payment"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payments"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentId"))
		err := f.Payment.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *PaymentTypeFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Description != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" = ?")
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" != ?")
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" > ?")
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" < ?")
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" >= ?")
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" <= ?")
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IN (?)")
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil {
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("description")+" IS NOT NULL")
		}
	}

	if f.PaymentID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" = ?")
		values = append(values, f.PaymentID)
	}

	if f.PaymentIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" != ?")
		values = append(values, f.PaymentIDNe)
	}

	if f.PaymentIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" > ?")
		values = append(values, f.PaymentIDGt)
	}

	if f.PaymentIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" < ?")
		values = append(values, f.PaymentIDLt)
	}

	if f.PaymentIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" >= ?")
		values = append(values, f.PaymentIDGte)
	}

	if f.PaymentIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" <= ?")
		values = append(values, f.PaymentIDLte)
	}

	if f.PaymentIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IN (?)")
		values = append(values, f.PaymentIDIn)
	}

	if f.PaymentIDNull != nil {
		if *f.PaymentIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *PaymentTypeFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.DescriptionMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMin)
	}

	if f.DescriptionMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") = ?")
		values = append(values, f.DescriptionMax)
	}

	if f.DescriptionMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMinNe)
	}

	if f.DescriptionMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") != ?")
		values = append(values, f.DescriptionMaxNe)
	}

	if f.DescriptionMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMinGt)
	}

	if f.DescriptionMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") > ?")
		values = append(values, f.DescriptionMaxGt)
	}

	if f.DescriptionMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMinLt)
	}

	if f.DescriptionMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") < ?")
		values = append(values, f.DescriptionMaxLt)
	}

	if f.DescriptionMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMinGte)
	}

	if f.DescriptionMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") >= ?")
		values = append(values, f.DescriptionMaxGte)
	}

	if f.DescriptionMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMinLte)
	}

	if f.DescriptionMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") <= ?")
		values = append(values, f.DescriptionMaxLte)
	}

	if f.DescriptionMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMinIn)
	}

	if f.DescriptionMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") IN (?)")
		values = append(values, f.DescriptionMaxIn)
	}

	if f.DescriptionMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.DescriptionMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.DescriptionMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMinPrefix))
	}

	if f.DescriptionMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionMaxPrefix))
	}

	if f.DescriptionMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMinSuffix))
	}

	if f.DescriptionMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("description")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionMaxSuffix))
	}

	if f.PaymentIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") = ?")
		values = append(values, f.PaymentIDMin)
	}

	if f.PaymentIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") = ?")
		values = append(values, f.PaymentIDMax)
	}

	if f.PaymentIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") != ?")
		values = append(values, f.PaymentIDMinNe)
	}

	if f.PaymentIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") != ?")
		values = append(values, f.PaymentIDMaxNe)
	}

	if f.PaymentIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") > ?")
		values = append(values, f.PaymentIDMinGt)
	}

	if f.PaymentIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") > ?")
		values = append(values, f.PaymentIDMaxGt)
	}

	if f.PaymentIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") < ?")
		values = append(values, f.PaymentIDMinLt)
	}

	if f.PaymentIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") < ?")
		values = append(values, f.PaymentIDMaxLt)
	}

	if f.PaymentIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") >= ?")
		values = append(values, f.PaymentIDMinGte)
	}

	if f.PaymentIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") >= ?")
		values = append(values, f.PaymentIDMaxGte)
	}

	if f.PaymentIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") <= ?")
		values = append(values, f.PaymentIDMinLte)
	}

	if f.PaymentIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") <= ?")
		values = append(values, f.PaymentIDMaxLte)
	}

	if f.PaymentIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentId")+") IN (?)")
		values = append(values, f.PaymentIDMinIn)
	}

	if f.PaymentIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentId")+") IN (?)")
		values = append(values, f.PaymentIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *PaymentTypeFilterType) AndWith(f2 ...*PaymentTypeFilterType) *PaymentTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PaymentTypeFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *PaymentTypeFilterType) OrWith(f2 ...*PaymentTypeFilterType) *PaymentTypeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PaymentTypeFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *PaymentFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *PaymentFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("payments"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *PaymentFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Wallet != nil {
		_alias := alias + "_wallet"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("wallets"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("walletId"))
		err := f.Wallet.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Account != nil {
		_alias := alias + "_account"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("accounts"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("accountId"))
		err := f.Account.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.PaymentChannel != nil {
		_alias := alias + "_paymentChannel"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payment_channels"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentChannelId"))
		err := f.PaymentChannel.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.PaymentType != nil {
		_alias := alias + "_paymentType"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("payment_types"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("paymentTypeId"))
		err := f.PaymentType.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *PaymentFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.PaymentRef != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" = ?")
		values = append(values, f.PaymentRef)
	}

	if f.PaymentRefNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" != ?")
		values = append(values, f.PaymentRefNe)
	}

	if f.PaymentRefGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" > ?")
		values = append(values, f.PaymentRefGt)
	}

	if f.PaymentRefLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" < ?")
		values = append(values, f.PaymentRefLt)
	}

	if f.PaymentRefGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" >= ?")
		values = append(values, f.PaymentRefGte)
	}

	if f.PaymentRefLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" <= ?")
		values = append(values, f.PaymentRefLte)
	}

	if f.PaymentRefIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" IN (?)")
		values = append(values, f.PaymentRefIn)
	}

	if f.PaymentRefLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PaymentRefLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PaymentRefPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PaymentRefPrefix))
	}

	if f.PaymentRefSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PaymentRefSuffix))
	}

	if f.PaymentRefNull != nil {
		if *f.PaymentRefNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentRef")+" IS NOT NULL")
		}
	}

	if f.Amount != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" = ?")
		values = append(values, f.Amount)
	}

	if f.AmountNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" != ?")
		values = append(values, f.AmountNe)
	}

	if f.AmountGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" > ?")
		values = append(values, f.AmountGt)
	}

	if f.AmountLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" < ?")
		values = append(values, f.AmountLt)
	}

	if f.AmountGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" >= ?")
		values = append(values, f.AmountGte)
	}

	if f.AmountLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" <= ?")
		values = append(values, f.AmountLte)
	}

	if f.AmountIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" IN (?)")
		values = append(values, f.AmountIn)
	}

	if f.AmountNull != nil {
		if *f.AmountNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("amount")+" IS NOT NULL")
		}
	}

	if f.Concept != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" = ?")
		values = append(values, f.Concept)
	}

	if f.ConceptNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" != ?")
		values = append(values, f.ConceptNe)
	}

	if f.ConceptGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" > ?")
		values = append(values, f.ConceptGt)
	}

	if f.ConceptLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" < ?")
		values = append(values, f.ConceptLt)
	}

	if f.ConceptGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" >= ?")
		values = append(values, f.ConceptGte)
	}

	if f.ConceptLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" <= ?")
		values = append(values, f.ConceptLte)
	}

	if f.ConceptIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" IN (?)")
		values = append(values, f.ConceptIn)
	}

	if f.ConceptLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ConceptLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ConceptPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ConceptPrefix))
	}

	if f.ConceptSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ConceptSuffix))
	}

	if f.ConceptNull != nil {
		if *f.ConceptNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("concept")+" IS NOT NULL")
		}
	}

	if f.WalletID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" = ?")
		values = append(values, f.WalletID)
	}

	if f.WalletIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" != ?")
		values = append(values, f.WalletIDNe)
	}

	if f.WalletIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" > ?")
		values = append(values, f.WalletIDGt)
	}

	if f.WalletIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" < ?")
		values = append(values, f.WalletIDLt)
	}

	if f.WalletIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" >= ?")
		values = append(values, f.WalletIDGte)
	}

	if f.WalletIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" <= ?")
		values = append(values, f.WalletIDLte)
	}

	if f.WalletIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IN (?)")
		values = append(values, f.WalletIDIn)
	}

	if f.WalletIDNull != nil {
		if *f.WalletIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("walletId")+" IS NOT NULL")
		}
	}

	if f.AccountID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" = ?")
		values = append(values, f.AccountID)
	}

	if f.AccountIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" != ?")
		values = append(values, f.AccountIDNe)
	}

	if f.AccountIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" > ?")
		values = append(values, f.AccountIDGt)
	}

	if f.AccountIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" < ?")
		values = append(values, f.AccountIDLt)
	}

	if f.AccountIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" >= ?")
		values = append(values, f.AccountIDGte)
	}

	if f.AccountIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" <= ?")
		values = append(values, f.AccountIDLte)
	}

	if f.AccountIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" IN (?)")
		values = append(values, f.AccountIDIn)
	}

	if f.AccountIDNull != nil {
		if *f.AccountIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("accountId")+" IS NOT NULL")
		}
	}

	if f.PaymentChannelID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" = ?")
		values = append(values, f.PaymentChannelID)
	}

	if f.PaymentChannelIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" != ?")
		values = append(values, f.PaymentChannelIDNe)
	}

	if f.PaymentChannelIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" > ?")
		values = append(values, f.PaymentChannelIDGt)
	}

	if f.PaymentChannelIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" < ?")
		values = append(values, f.PaymentChannelIDLt)
	}

	if f.PaymentChannelIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" >= ?")
		values = append(values, f.PaymentChannelIDGte)
	}

	if f.PaymentChannelIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" <= ?")
		values = append(values, f.PaymentChannelIDLte)
	}

	if f.PaymentChannelIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" IN (?)")
		values = append(values, f.PaymentChannelIDIn)
	}

	if f.PaymentChannelIDNull != nil {
		if *f.PaymentChannelIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentChannelId")+" IS NOT NULL")
		}
	}

	if f.PaymentTypeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" = ?")
		values = append(values, f.PaymentTypeID)
	}

	if f.PaymentTypeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" != ?")
		values = append(values, f.PaymentTypeIDNe)
	}

	if f.PaymentTypeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" > ?")
		values = append(values, f.PaymentTypeIDGt)
	}

	if f.PaymentTypeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" < ?")
		values = append(values, f.PaymentTypeIDLt)
	}

	if f.PaymentTypeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" >= ?")
		values = append(values, f.PaymentTypeIDGte)
	}

	if f.PaymentTypeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" <= ?")
		values = append(values, f.PaymentTypeIDLte)
	}

	if f.PaymentTypeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" IN (?)")
		values = append(values, f.PaymentTypeIDIn)
	}

	if f.PaymentTypeIDNull != nil {
		if *f.PaymentTypeIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("paymentTypeId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *PaymentFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.PaymentRefMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") = ?")
		values = append(values, f.PaymentRefMin)
	}

	if f.PaymentRefMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") = ?")
		values = append(values, f.PaymentRefMax)
	}

	if f.PaymentRefMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") != ?")
		values = append(values, f.PaymentRefMinNe)
	}

	if f.PaymentRefMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") != ?")
		values = append(values, f.PaymentRefMaxNe)
	}

	if f.PaymentRefMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") > ?")
		values = append(values, f.PaymentRefMinGt)
	}

	if f.PaymentRefMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") > ?")
		values = append(values, f.PaymentRefMaxGt)
	}

	if f.PaymentRefMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") < ?")
		values = append(values, f.PaymentRefMinLt)
	}

	if f.PaymentRefMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") < ?")
		values = append(values, f.PaymentRefMaxLt)
	}

	if f.PaymentRefMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") >= ?")
		values = append(values, f.PaymentRefMinGte)
	}

	if f.PaymentRefMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") >= ?")
		values = append(values, f.PaymentRefMaxGte)
	}

	if f.PaymentRefMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") <= ?")
		values = append(values, f.PaymentRefMinLte)
	}

	if f.PaymentRefMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") <= ?")
		values = append(values, f.PaymentRefMaxLte)
	}

	if f.PaymentRefMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") IN (?)")
		values = append(values, f.PaymentRefMinIn)
	}

	if f.PaymentRefMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") IN (?)")
		values = append(values, f.PaymentRefMaxIn)
	}

	if f.PaymentRefMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PaymentRefMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PaymentRefMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PaymentRefMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PaymentRefMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PaymentRefMinPrefix))
	}

	if f.PaymentRefMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PaymentRefMaxPrefix))
	}

	if f.PaymentRefMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentRef")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PaymentRefMinSuffix))
	}

	if f.PaymentRefMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentRef")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PaymentRefMaxSuffix))
	}

	if f.AmountMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("amount")+") = ?")
		values = append(values, f.AmountMin)
	}

	if f.AmountMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("amount")+") = ?")
		values = append(values, f.AmountMax)
	}

	if f.AmountAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("amount")+") = ?")
		values = append(values, f.AmountAvg)
	}

	if f.AmountMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("amount")+") != ?")
		values = append(values, f.AmountMinNe)
	}

	if f.AmountMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("amount")+") != ?")
		values = append(values, f.AmountMaxNe)
	}

	if f.AmountAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("amount")+") != ?")
		values = append(values, f.AmountAvgNe)
	}

	if f.AmountMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("amount")+") > ?")
		values = append(values, f.AmountMinGt)
	}

	if f.AmountMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("amount")+") > ?")
		values = append(values, f.AmountMaxGt)
	}

	if f.AmountAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("amount")+") > ?")
		values = append(values, f.AmountAvgGt)
	}

	if f.AmountMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("amount")+") < ?")
		values = append(values, f.AmountMinLt)
	}

	if f.AmountMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("amount")+") < ?")
		values = append(values, f.AmountMaxLt)
	}

	if f.AmountAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("amount")+") < ?")
		values = append(values, f.AmountAvgLt)
	}

	if f.AmountMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("amount")+") >= ?")
		values = append(values, f.AmountMinGte)
	}

	if f.AmountMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("amount")+") >= ?")
		values = append(values, f.AmountMaxGte)
	}

	if f.AmountAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("amount")+") >= ?")
		values = append(values, f.AmountAvgGte)
	}

	if f.AmountMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("amount")+") <= ?")
		values = append(values, f.AmountMinLte)
	}

	if f.AmountMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("amount")+") <= ?")
		values = append(values, f.AmountMaxLte)
	}

	if f.AmountAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("amount")+") <= ?")
		values = append(values, f.AmountAvgLte)
	}

	if f.AmountMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("amount")+") IN (?)")
		values = append(values, f.AmountMinIn)
	}

	if f.AmountMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("amount")+") IN (?)")
		values = append(values, f.AmountMaxIn)
	}

	if f.AmountAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("amount")+") IN (?)")
		values = append(values, f.AmountAvgIn)
	}

	if f.ConceptMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") = ?")
		values = append(values, f.ConceptMin)
	}

	if f.ConceptMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") = ?")
		values = append(values, f.ConceptMax)
	}

	if f.ConceptMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") != ?")
		values = append(values, f.ConceptMinNe)
	}

	if f.ConceptMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") != ?")
		values = append(values, f.ConceptMaxNe)
	}

	if f.ConceptMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") > ?")
		values = append(values, f.ConceptMinGt)
	}

	if f.ConceptMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") > ?")
		values = append(values, f.ConceptMaxGt)
	}

	if f.ConceptMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") < ?")
		values = append(values, f.ConceptMinLt)
	}

	if f.ConceptMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") < ?")
		values = append(values, f.ConceptMaxLt)
	}

	if f.ConceptMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") >= ?")
		values = append(values, f.ConceptMinGte)
	}

	if f.ConceptMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") >= ?")
		values = append(values, f.ConceptMaxGte)
	}

	if f.ConceptMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") <= ?")
		values = append(values, f.ConceptMinLte)
	}

	if f.ConceptMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") <= ?")
		values = append(values, f.ConceptMaxLte)
	}

	if f.ConceptMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") IN (?)")
		values = append(values, f.ConceptMinIn)
	}

	if f.ConceptMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") IN (?)")
		values = append(values, f.ConceptMaxIn)
	}

	if f.ConceptMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ConceptMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ConceptMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ConceptMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ConceptMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ConceptMinPrefix))
	}

	if f.ConceptMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ConceptMaxPrefix))
	}

	if f.ConceptMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("concept")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ConceptMinSuffix))
	}

	if f.ConceptMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("concept")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ConceptMaxSuffix))
	}

	if f.WalletIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") = ?")
		values = append(values, f.WalletIDMin)
	}

	if f.WalletIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") = ?")
		values = append(values, f.WalletIDMax)
	}

	if f.WalletIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") != ?")
		values = append(values, f.WalletIDMinNe)
	}

	if f.WalletIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") != ?")
		values = append(values, f.WalletIDMaxNe)
	}

	if f.WalletIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") > ?")
		values = append(values, f.WalletIDMinGt)
	}

	if f.WalletIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") > ?")
		values = append(values, f.WalletIDMaxGt)
	}

	if f.WalletIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") < ?")
		values = append(values, f.WalletIDMinLt)
	}

	if f.WalletIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") < ?")
		values = append(values, f.WalletIDMaxLt)
	}

	if f.WalletIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") >= ?")
		values = append(values, f.WalletIDMinGte)
	}

	if f.WalletIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") >= ?")
		values = append(values, f.WalletIDMaxGte)
	}

	if f.WalletIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") <= ?")
		values = append(values, f.WalletIDMinLte)
	}

	if f.WalletIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") <= ?")
		values = append(values, f.WalletIDMaxLte)
	}

	if f.WalletIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("walletId")+") IN (?)")
		values = append(values, f.WalletIDMinIn)
	}

	if f.WalletIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("walletId")+") IN (?)")
		values = append(values, f.WalletIDMaxIn)
	}

	if f.AccountIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountId")+") = ?")
		values = append(values, f.AccountIDMin)
	}

	if f.AccountIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountId")+") = ?")
		values = append(values, f.AccountIDMax)
	}

	if f.AccountIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountId")+") != ?")
		values = append(values, f.AccountIDMinNe)
	}

	if f.AccountIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountId")+") != ?")
		values = append(values, f.AccountIDMaxNe)
	}

	if f.AccountIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountId")+") > ?")
		values = append(values, f.AccountIDMinGt)
	}

	if f.AccountIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountId")+") > ?")
		values = append(values, f.AccountIDMaxGt)
	}

	if f.AccountIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountId")+") < ?")
		values = append(values, f.AccountIDMinLt)
	}

	if f.AccountIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountId")+") < ?")
		values = append(values, f.AccountIDMaxLt)
	}

	if f.AccountIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountId")+") >= ?")
		values = append(values, f.AccountIDMinGte)
	}

	if f.AccountIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountId")+") >= ?")
		values = append(values, f.AccountIDMaxGte)
	}

	if f.AccountIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountId")+") <= ?")
		values = append(values, f.AccountIDMinLte)
	}

	if f.AccountIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountId")+") <= ?")
		values = append(values, f.AccountIDMaxLte)
	}

	if f.AccountIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("accountId")+") IN (?)")
		values = append(values, f.AccountIDMinIn)
	}

	if f.AccountIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("accountId")+") IN (?)")
		values = append(values, f.AccountIDMaxIn)
	}

	if f.PaymentChannelIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentChannelId")+") = ?")
		values = append(values, f.PaymentChannelIDMin)
	}

	if f.PaymentChannelIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentChannelId")+") = ?")
		values = append(values, f.PaymentChannelIDMax)
	}

	if f.PaymentChannelIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentChannelId")+") != ?")
		values = append(values, f.PaymentChannelIDMinNe)
	}

	if f.PaymentChannelIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentChannelId")+") != ?")
		values = append(values, f.PaymentChannelIDMaxNe)
	}

	if f.PaymentChannelIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentChannelId")+") > ?")
		values = append(values, f.PaymentChannelIDMinGt)
	}

	if f.PaymentChannelIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentChannelId")+") > ?")
		values = append(values, f.PaymentChannelIDMaxGt)
	}

	if f.PaymentChannelIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentChannelId")+") < ?")
		values = append(values, f.PaymentChannelIDMinLt)
	}

	if f.PaymentChannelIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentChannelId")+") < ?")
		values = append(values, f.PaymentChannelIDMaxLt)
	}

	if f.PaymentChannelIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentChannelId")+") >= ?")
		values = append(values, f.PaymentChannelIDMinGte)
	}

	if f.PaymentChannelIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentChannelId")+") >= ?")
		values = append(values, f.PaymentChannelIDMaxGte)
	}

	if f.PaymentChannelIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentChannelId")+") <= ?")
		values = append(values, f.PaymentChannelIDMinLte)
	}

	if f.PaymentChannelIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentChannelId")+") <= ?")
		values = append(values, f.PaymentChannelIDMaxLte)
	}

	if f.PaymentChannelIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentChannelId")+") IN (?)")
		values = append(values, f.PaymentChannelIDMinIn)
	}

	if f.PaymentChannelIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentChannelId")+") IN (?)")
		values = append(values, f.PaymentChannelIDMaxIn)
	}

	if f.PaymentTypeIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTypeId")+") = ?")
		values = append(values, f.PaymentTypeIDMin)
	}

	if f.PaymentTypeIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTypeId")+") = ?")
		values = append(values, f.PaymentTypeIDMax)
	}

	if f.PaymentTypeIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTypeId")+") != ?")
		values = append(values, f.PaymentTypeIDMinNe)
	}

	if f.PaymentTypeIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTypeId")+") != ?")
		values = append(values, f.PaymentTypeIDMaxNe)
	}

	if f.PaymentTypeIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTypeId")+") > ?")
		values = append(values, f.PaymentTypeIDMinGt)
	}

	if f.PaymentTypeIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTypeId")+") > ?")
		values = append(values, f.PaymentTypeIDMaxGt)
	}

	if f.PaymentTypeIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTypeId")+") < ?")
		values = append(values, f.PaymentTypeIDMinLt)
	}

	if f.PaymentTypeIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTypeId")+") < ?")
		values = append(values, f.PaymentTypeIDMaxLt)
	}

	if f.PaymentTypeIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTypeId")+") >= ?")
		values = append(values, f.PaymentTypeIDMinGte)
	}

	if f.PaymentTypeIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTypeId")+") >= ?")
		values = append(values, f.PaymentTypeIDMaxGte)
	}

	if f.PaymentTypeIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTypeId")+") <= ?")
		values = append(values, f.PaymentTypeIDMinLte)
	}

	if f.PaymentTypeIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTypeId")+") <= ?")
		values = append(values, f.PaymentTypeIDMaxLte)
	}

	if f.PaymentTypeIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("paymentTypeId")+") IN (?)")
		values = append(values, f.PaymentTypeIDMinIn)
	}

	if f.PaymentTypeIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("paymentTypeId")+") IN (?)")
		values = append(values, f.PaymentTypeIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *PaymentFilterType) AndWith(f2 ...*PaymentFilterType) *PaymentFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PaymentFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *PaymentFilterType) OrWith(f2 ...*PaymentFilterType) *PaymentFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PaymentFilterType{
		Or: append(_f2, f),
	}
}
