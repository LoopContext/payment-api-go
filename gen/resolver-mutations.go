package gen

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/loopcontext/go-graphql-orm/events"
)

// GeneratedMutationResolver struct
type GeneratedMutationResolver struct{ *GeneratedResolver }

// MutationEvents struct
type MutationEvents struct {
	Events []events.Event
}

// EnrichContextWithMutations method
func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.GetDB(ctx).Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}

// FinishMutationContext method
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}

// RollbackMutationContext method
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}

// GetMutationEventStore method
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}

// AddMutationEvent method
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

// CreateWallet method
func (r *GeneratedMutationResolver) CreateWallet(ctx context.Context, input map[string]interface{}) (item *Wallet, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateWallet(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateWalletHandler handler
func CreateWalletHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Wallet, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Wallet{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Wallet",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes WalletChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {
		item.UserID = changes.UserID

		event.AddNewValue("userId", changes.UserID)
	}

	if _, ok := input["balance"]; ok && (item.Balance != changes.Balance) {
		item.Balance = changes.Balance

		event.AddNewValue("balance", changes.Balance)
	}

	if _, ok := input["walletTypeId"]; ok && (item.WalletTypeID != changes.WalletTypeID) && (item.WalletTypeID == nil || changes.WalletTypeID == nil || *item.WalletTypeID != *changes.WalletTypeID) {
		item.WalletTypeID = changes.WalletTypeID

		event.AddNewValue("walletTypeId", changes.WalletTypeID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["accountsIds"]; exists {
		items := []Account{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Accounts")
		association.Replace(items)
	}

	if ids, exists := input["paymentsIds"]; exists {
		items := []Payment{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Payments")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateWallet method
func (r *GeneratedMutationResolver) UpdateWallet(ctx context.Context, id string, input map[string]interface{}) (item *Wallet, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateWallet(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateWalletHandler handler
func UpdateWalletHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Wallet, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Wallet{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Wallet",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes WalletChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["userId"]; ok && (item.UserID != changes.UserID) && (item.UserID == nil || changes.UserID == nil || *item.UserID != *changes.UserID) {
		event.AddOldValue("userId", item.UserID)
		event.AddNewValue("userId", changes.UserID)
		item.UserID = changes.UserID
	}

	if _, ok := input["balance"]; ok && (item.Balance != changes.Balance) {
		event.AddOldValue("balance", item.Balance)
		event.AddNewValue("balance", changes.Balance)
		item.Balance = changes.Balance
	}

	if _, ok := input["walletTypeId"]; ok && (item.WalletTypeID != changes.WalletTypeID) && (item.WalletTypeID == nil || changes.WalletTypeID == nil || *item.WalletTypeID != *changes.WalletTypeID) {
		event.AddOldValue("walletTypeId", item.WalletTypeID)
		event.AddNewValue("walletTypeId", changes.WalletTypeID)
		item.WalletTypeID = changes.WalletTypeID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["accountsIds"]; exists {
		items := []Account{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Accounts")
		association.Replace(items)
	}

	if ids, exists := input["paymentsIds"]; exists {
		items := []Payment{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Payments")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteWallet method
func (r *GeneratedMutationResolver) DeleteWallet(ctx context.Context, id string) (item *Wallet, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteWallet(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteWalletHandler handler
func DeleteWalletHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Wallet, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Wallet{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Wallet",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("wallets")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllWallets method
func (r *GeneratedMutationResolver) DeleteAllWallets(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllWallets(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllWalletsHandler handler
func DeleteAllWalletsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&Wallet{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateWalletType method
func (r *GeneratedMutationResolver) CreateWalletType(ctx context.Context, input map[string]interface{}) (item *WalletType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateWalletType(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateWalletTypeHandler handler
func CreateWalletTypeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *WalletType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &WalletType{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "WalletType",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes WalletTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["walletId"]; ok && (item.WalletID != changes.WalletID) && (item.WalletID == nil || changes.WalletID == nil || *item.WalletID != *changes.WalletID) {
		item.WalletID = changes.WalletID

		event.AddNewValue("walletId", changes.WalletID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateWalletType method
func (r *GeneratedMutationResolver) UpdateWalletType(ctx context.Context, id string, input map[string]interface{}) (item *WalletType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateWalletType(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateWalletTypeHandler handler
func UpdateWalletTypeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *WalletType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &WalletType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "WalletType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes WalletTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["walletId"]; ok && (item.WalletID != changes.WalletID) && (item.WalletID == nil || changes.WalletID == nil || *item.WalletID != *changes.WalletID) {
		event.AddOldValue("walletId", item.WalletID)
		event.AddNewValue("walletId", changes.WalletID)
		item.WalletID = changes.WalletID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteWalletType method
func (r *GeneratedMutationResolver) DeleteWalletType(ctx context.Context, id string) (item *WalletType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteWalletType(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteWalletTypeHandler handler
func DeleteWalletTypeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *WalletType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &WalletType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "WalletType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("wallet_types")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllWalletTypes method
func (r *GeneratedMutationResolver) DeleteAllWalletTypes(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllWalletTypes(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllWalletTypesHandler handler
func DeleteAllWalletTypesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&WalletType{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateAccountProviderType method
func (r *GeneratedMutationResolver) CreateAccountProviderType(ctx context.Context, input map[string]interface{}) (item *AccountProviderType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateAccountProviderType(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateAccountProviderTypeHandler handler
func CreateAccountProviderTypeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *AccountProviderType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &AccountProviderType{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "AccountProviderType",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AccountProviderTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["accountProviderId"]; ok && (item.AccountProviderID != changes.AccountProviderID) && (item.AccountProviderID == nil || changes.AccountProviderID == nil || *item.AccountProviderID != *changes.AccountProviderID) {
		item.AccountProviderID = changes.AccountProviderID

		event.AddNewValue("accountProviderId", changes.AccountProviderID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateAccountProviderType method
func (r *GeneratedMutationResolver) UpdateAccountProviderType(ctx context.Context, id string, input map[string]interface{}) (item *AccountProviderType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateAccountProviderType(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateAccountProviderTypeHandler handler
func UpdateAccountProviderTypeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *AccountProviderType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &AccountProviderType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "AccountProviderType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AccountProviderTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["accountProviderId"]; ok && (item.AccountProviderID != changes.AccountProviderID) && (item.AccountProviderID == nil || changes.AccountProviderID == nil || *item.AccountProviderID != *changes.AccountProviderID) {
		event.AddOldValue("accountProviderId", item.AccountProviderID)
		event.AddNewValue("accountProviderId", changes.AccountProviderID)
		item.AccountProviderID = changes.AccountProviderID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteAccountProviderType method
func (r *GeneratedMutationResolver) DeleteAccountProviderType(ctx context.Context, id string) (item *AccountProviderType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteAccountProviderType(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteAccountProviderTypeHandler handler
func DeleteAccountProviderTypeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *AccountProviderType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &AccountProviderType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "AccountProviderType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("account_provider_types")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllAccountProviderTypes method
func (r *GeneratedMutationResolver) DeleteAllAccountProviderTypes(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllAccountProviderTypes(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllAccountProviderTypesHandler handler
func DeleteAllAccountProviderTypesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&AccountProviderType{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateAccountProvider method
func (r *GeneratedMutationResolver) CreateAccountProvider(ctx context.Context, input map[string]interface{}) (item *AccountProvider, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateAccountProvider(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateAccountProviderHandler handler
func CreateAccountProviderHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *AccountProvider, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &AccountProvider{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "AccountProvider",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AccountProviderChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["address"]; ok && (item.Address != changes.Address) && (item.Address == nil || changes.Address == nil || *item.Address != *changes.Address) {
		item.Address = changes.Address

		event.AddNewValue("address", changes.Address)
	}

	if _, ok := input["phone"]; ok && (item.Phone != changes.Phone) && (item.Phone == nil || changes.Phone == nil || *item.Phone != *changes.Phone) {
		item.Phone = changes.Phone

		event.AddNewValue("phone", changes.Phone)
	}

	if _, ok := input["accountProviderTypeId"]; ok && (item.AccountProviderTypeID != changes.AccountProviderTypeID) && (item.AccountProviderTypeID == nil || changes.AccountProviderTypeID == nil || *item.AccountProviderTypeID != *changes.AccountProviderTypeID) {
		item.AccountProviderTypeID = changes.AccountProviderTypeID

		event.AddNewValue("accountProviderTypeId", changes.AccountProviderTypeID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["accountsIds"]; exists {
		items := []Account{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Accounts")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateAccountProvider method
func (r *GeneratedMutationResolver) UpdateAccountProvider(ctx context.Context, id string, input map[string]interface{}) (item *AccountProvider, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateAccountProvider(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateAccountProviderHandler handler
func UpdateAccountProviderHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *AccountProvider, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &AccountProvider{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "AccountProvider",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AccountProviderChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["address"]; ok && (item.Address != changes.Address) && (item.Address == nil || changes.Address == nil || *item.Address != *changes.Address) {
		event.AddOldValue("address", item.Address)
		event.AddNewValue("address", changes.Address)
		item.Address = changes.Address
	}

	if _, ok := input["phone"]; ok && (item.Phone != changes.Phone) && (item.Phone == nil || changes.Phone == nil || *item.Phone != *changes.Phone) {
		event.AddOldValue("phone", item.Phone)
		event.AddNewValue("phone", changes.Phone)
		item.Phone = changes.Phone
	}

	if _, ok := input["accountProviderTypeId"]; ok && (item.AccountProviderTypeID != changes.AccountProviderTypeID) && (item.AccountProviderTypeID == nil || changes.AccountProviderTypeID == nil || *item.AccountProviderTypeID != *changes.AccountProviderTypeID) {
		event.AddOldValue("accountProviderTypeId", item.AccountProviderTypeID)
		event.AddNewValue("accountProviderTypeId", changes.AccountProviderTypeID)
		item.AccountProviderTypeID = changes.AccountProviderTypeID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["accountsIds"]; exists {
		items := []Account{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Accounts")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteAccountProvider method
func (r *GeneratedMutationResolver) DeleteAccountProvider(ctx context.Context, id string) (item *AccountProvider, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteAccountProvider(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteAccountProviderHandler handler
func DeleteAccountProviderHandler(ctx context.Context, r *GeneratedResolver, id string) (item *AccountProvider, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &AccountProvider{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "AccountProvider",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("account_providers")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllAccountProviders method
func (r *GeneratedMutationResolver) DeleteAllAccountProviders(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllAccountProviders(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllAccountProvidersHandler handler
func DeleteAllAccountProvidersHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&AccountProvider{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateAccount method
func (r *GeneratedMutationResolver) CreateAccount(ctx context.Context, input map[string]interface{}) (item *Account, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateAccount(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateAccountHandler handler
func CreateAccountHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Account, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Account{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Account",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AccountChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["accountNumber"]; ok && (item.AccountNumber != changes.AccountNumber) {
		item.AccountNumber = changes.AccountNumber

		event.AddNewValue("accountNumber", changes.AccountNumber)
	}

	if _, ok := input["balance"]; ok && (item.Balance != changes.Balance) {
		item.Balance = changes.Balance

		event.AddNewValue("balance", changes.Balance)
	}

	if _, ok := input["accountProviderId"]; ok && (item.AccountProviderID != changes.AccountProviderID) && (item.AccountProviderID == nil || changes.AccountProviderID == nil || *item.AccountProviderID != *changes.AccountProviderID) {
		item.AccountProviderID = changes.AccountProviderID

		event.AddNewValue("accountProviderId", changes.AccountProviderID)
	}

	if _, ok := input["walletId"]; ok && (item.WalletID != changes.WalletID) && (item.WalletID == nil || changes.WalletID == nil || *item.WalletID != *changes.WalletID) {
		item.WalletID = changes.WalletID

		event.AddNewValue("walletId", changes.WalletID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["paymentsIds"]; exists {
		items := []Payment{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Payments")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateAccount method
func (r *GeneratedMutationResolver) UpdateAccount(ctx context.Context, id string, input map[string]interface{}) (item *Account, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateAccount(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateAccountHandler handler
func UpdateAccountHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Account, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Account{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Account",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes AccountChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["accountNumber"]; ok && (item.AccountNumber != changes.AccountNumber) {
		event.AddOldValue("accountNumber", item.AccountNumber)
		event.AddNewValue("accountNumber", changes.AccountNumber)
		item.AccountNumber = changes.AccountNumber
	}

	if _, ok := input["balance"]; ok && (item.Balance != changes.Balance) {
		event.AddOldValue("balance", item.Balance)
		event.AddNewValue("balance", changes.Balance)
		item.Balance = changes.Balance
	}

	if _, ok := input["accountProviderId"]; ok && (item.AccountProviderID != changes.AccountProviderID) && (item.AccountProviderID == nil || changes.AccountProviderID == nil || *item.AccountProviderID != *changes.AccountProviderID) {
		event.AddOldValue("accountProviderId", item.AccountProviderID)
		event.AddNewValue("accountProviderId", changes.AccountProviderID)
		item.AccountProviderID = changes.AccountProviderID
	}

	if _, ok := input["walletId"]; ok && (item.WalletID != changes.WalletID) && (item.WalletID == nil || changes.WalletID == nil || *item.WalletID != *changes.WalletID) {
		event.AddOldValue("walletId", item.WalletID)
		event.AddNewValue("walletId", changes.WalletID)
		item.WalletID = changes.WalletID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["paymentsIds"]; exists {
		items := []Payment{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Payments")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteAccount method
func (r *GeneratedMutationResolver) DeleteAccount(ctx context.Context, id string) (item *Account, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteAccount(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteAccountHandler handler
func DeleteAccountHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Account, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Account{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Account",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("accounts")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllAccounts method
func (r *GeneratedMutationResolver) DeleteAllAccounts(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllAccounts(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllAccountsHandler handler
func DeleteAllAccountsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&Account{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreatePaymentChannel method
func (r *GeneratedMutationResolver) CreatePaymentChannel(ctx context.Context, input map[string]interface{}) (item *PaymentChannel, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreatePaymentChannel(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreatePaymentChannelHandler handler
func CreatePaymentChannelHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentChannel, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &PaymentChannel{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "PaymentChannel",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PaymentChannelChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["paymentId"]; ok && (item.PaymentID != changes.PaymentID) && (item.PaymentID == nil || changes.PaymentID == nil || *item.PaymentID != *changes.PaymentID) {
		item.PaymentID = changes.PaymentID

		event.AddNewValue("paymentId", changes.PaymentID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdatePaymentChannel method
func (r *GeneratedMutationResolver) UpdatePaymentChannel(ctx context.Context, id string, input map[string]interface{}) (item *PaymentChannel, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdatePaymentChannel(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdatePaymentChannelHandler handler
func UpdatePaymentChannelHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentChannel, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &PaymentChannel{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "PaymentChannel",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PaymentChannelChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["paymentId"]; ok && (item.PaymentID != changes.PaymentID) && (item.PaymentID == nil || changes.PaymentID == nil || *item.PaymentID != *changes.PaymentID) {
		event.AddOldValue("paymentId", item.PaymentID)
		event.AddNewValue("paymentId", changes.PaymentID)
		item.PaymentID = changes.PaymentID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeletePaymentChannel method
func (r *GeneratedMutationResolver) DeletePaymentChannel(ctx context.Context, id string) (item *PaymentChannel, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeletePaymentChannel(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeletePaymentChannelHandler handler
func DeletePaymentChannelHandler(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentChannel, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &PaymentChannel{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "PaymentChannel",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("payment_channels")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllPaymentChannels method
func (r *GeneratedMutationResolver) DeleteAllPaymentChannels(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllPaymentChannels(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllPaymentChannelsHandler handler
func DeleteAllPaymentChannelsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&PaymentChannel{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreatePaymentType method
func (r *GeneratedMutationResolver) CreatePaymentType(ctx context.Context, input map[string]interface{}) (item *PaymentType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreatePaymentType(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreatePaymentTypeHandler handler
func CreatePaymentTypeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PaymentType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &PaymentType{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "PaymentType",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PaymentTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		item.Description = changes.Description

		event.AddNewValue("description", changes.Description)
	}

	if _, ok := input["paymentId"]; ok && (item.PaymentID != changes.PaymentID) && (item.PaymentID == nil || changes.PaymentID == nil || *item.PaymentID != *changes.PaymentID) {
		item.PaymentID = changes.PaymentID

		event.AddNewValue("paymentId", changes.PaymentID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdatePaymentType method
func (r *GeneratedMutationResolver) UpdatePaymentType(ctx context.Context, id string, input map[string]interface{}) (item *PaymentType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdatePaymentType(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdatePaymentTypeHandler handler
func UpdatePaymentTypeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PaymentType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &PaymentType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "PaymentType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PaymentTypeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["description"]; ok && (item.Description != changes.Description) && (item.Description == nil || changes.Description == nil || *item.Description != *changes.Description) {
		event.AddOldValue("description", item.Description)
		event.AddNewValue("description", changes.Description)
		item.Description = changes.Description
	}

	if _, ok := input["paymentId"]; ok && (item.PaymentID != changes.PaymentID) && (item.PaymentID == nil || changes.PaymentID == nil || *item.PaymentID != *changes.PaymentID) {
		event.AddOldValue("paymentId", item.PaymentID)
		event.AddNewValue("paymentId", changes.PaymentID)
		item.PaymentID = changes.PaymentID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeletePaymentType method
func (r *GeneratedMutationResolver) DeletePaymentType(ctx context.Context, id string) (item *PaymentType, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeletePaymentType(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeletePaymentTypeHandler handler
func DeletePaymentTypeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *PaymentType, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &PaymentType{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "PaymentType",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("payment_types")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllPaymentTypes method
func (r *GeneratedMutationResolver) DeleteAllPaymentTypes(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllPaymentTypes(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllPaymentTypesHandler handler
func DeleteAllPaymentTypesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&PaymentType{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreatePayment method
func (r *GeneratedMutationResolver) CreatePayment(ctx context.Context, input map[string]interface{}) (item *Payment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreatePayment(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreatePaymentHandler handler
func CreatePaymentHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Payment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Payment{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Payment",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PaymentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["paymentRef"]; ok && (item.PaymentRef != changes.PaymentRef) && (item.PaymentRef == nil || changes.PaymentRef == nil || *item.PaymentRef != *changes.PaymentRef) {
		item.PaymentRef = changes.PaymentRef

		event.AddNewValue("paymentRef", changes.PaymentRef)
	}

	if _, ok := input["amount"]; ok && (item.Amount != changes.Amount) {
		item.Amount = changes.Amount

		event.AddNewValue("amount", changes.Amount)
	}

	if _, ok := input["concept"]; ok && (item.Concept != changes.Concept) && (item.Concept == nil || changes.Concept == nil || *item.Concept != *changes.Concept) {
		item.Concept = changes.Concept

		event.AddNewValue("concept", changes.Concept)
	}

	if _, ok := input["walletId"]; ok && (item.WalletID != changes.WalletID) && (item.WalletID == nil || changes.WalletID == nil || *item.WalletID != *changes.WalletID) {
		item.WalletID = changes.WalletID

		event.AddNewValue("walletId", changes.WalletID)
	}

	if _, ok := input["accountId"]; ok && (item.AccountID != changes.AccountID) && (item.AccountID == nil || changes.AccountID == nil || *item.AccountID != *changes.AccountID) {
		item.AccountID = changes.AccountID

		event.AddNewValue("accountId", changes.AccountID)
	}

	if _, ok := input["paymentChannelId"]; ok && (item.PaymentChannelID != changes.PaymentChannelID) && (item.PaymentChannelID == nil || changes.PaymentChannelID == nil || *item.PaymentChannelID != *changes.PaymentChannelID) {
		item.PaymentChannelID = changes.PaymentChannelID

		event.AddNewValue("paymentChannelId", changes.PaymentChannelID)
	}

	if _, ok := input["paymentTypeId"]; ok && (item.PaymentTypeID != changes.PaymentTypeID) && (item.PaymentTypeID == nil || changes.PaymentTypeID == nil || *item.PaymentTypeID != *changes.PaymentTypeID) {
		item.PaymentTypeID = changes.PaymentTypeID

		event.AddNewValue("paymentTypeId", changes.PaymentTypeID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdatePayment method
func (r *GeneratedMutationResolver) UpdatePayment(ctx context.Context, id string, input map[string]interface{}) (item *Payment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdatePayment(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdatePaymentHandler handler
func UpdatePaymentHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Payment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Payment{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Payment",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PaymentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["paymentRef"]; ok && (item.PaymentRef != changes.PaymentRef) && (item.PaymentRef == nil || changes.PaymentRef == nil || *item.PaymentRef != *changes.PaymentRef) {
		event.AddOldValue("paymentRef", item.PaymentRef)
		event.AddNewValue("paymentRef", changes.PaymentRef)
		item.PaymentRef = changes.PaymentRef
	}

	if _, ok := input["amount"]; ok && (item.Amount != changes.Amount) {
		event.AddOldValue("amount", item.Amount)
		event.AddNewValue("amount", changes.Amount)
		item.Amount = changes.Amount
	}

	if _, ok := input["concept"]; ok && (item.Concept != changes.Concept) && (item.Concept == nil || changes.Concept == nil || *item.Concept != *changes.Concept) {
		event.AddOldValue("concept", item.Concept)
		event.AddNewValue("concept", changes.Concept)
		item.Concept = changes.Concept
	}

	if _, ok := input["walletId"]; ok && (item.WalletID != changes.WalletID) && (item.WalletID == nil || changes.WalletID == nil || *item.WalletID != *changes.WalletID) {
		event.AddOldValue("walletId", item.WalletID)
		event.AddNewValue("walletId", changes.WalletID)
		item.WalletID = changes.WalletID
	}

	if _, ok := input["accountId"]; ok && (item.AccountID != changes.AccountID) && (item.AccountID == nil || changes.AccountID == nil || *item.AccountID != *changes.AccountID) {
		event.AddOldValue("accountId", item.AccountID)
		event.AddNewValue("accountId", changes.AccountID)
		item.AccountID = changes.AccountID
	}

	if _, ok := input["paymentChannelId"]; ok && (item.PaymentChannelID != changes.PaymentChannelID) && (item.PaymentChannelID == nil || changes.PaymentChannelID == nil || *item.PaymentChannelID != *changes.PaymentChannelID) {
		event.AddOldValue("paymentChannelId", item.PaymentChannelID)
		event.AddNewValue("paymentChannelId", changes.PaymentChannelID)
		item.PaymentChannelID = changes.PaymentChannelID
	}

	if _, ok := input["paymentTypeId"]; ok && (item.PaymentTypeID != changes.PaymentTypeID) && (item.PaymentTypeID == nil || changes.PaymentTypeID == nil || *item.PaymentTypeID != *changes.PaymentTypeID) {
		event.AddOldValue("paymentTypeId", item.PaymentTypeID)
		event.AddNewValue("paymentTypeId", changes.PaymentTypeID)
		item.PaymentTypeID = changes.PaymentTypeID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeletePayment method
func (r *GeneratedMutationResolver) DeletePayment(ctx context.Context, id string) (item *Payment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeletePayment(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeletePaymentHandler handler
func DeletePaymentHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Payment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Payment{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Payment",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("payments")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllPayments method
func (r *GeneratedMutationResolver) DeleteAllPayments(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllPayments(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllPaymentsHandler handler
func DeleteAllPaymentsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&Payment{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
