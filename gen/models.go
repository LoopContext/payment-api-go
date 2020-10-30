package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

// WalletResultType struct
type WalletResultType struct {
	EntityResultType
}

// Wallet struct
type Wallet struct {
	ID           string     `json:"id" gorm:"column:id;primary_key"`
	UserID       *string    `json:"userId" gorm:"column:userId;index:useridx"`
	Balance      float64    `json:"balance" gorm:"column:balance;default:0.0"`
	WalletTypeID *string    `json:"walletTypeId" gorm:"column:walletTypeId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	WalletType *WalletType `json:"walletType"`

	Accounts []*Account `json:"accounts" gorm:"foreignkey:WalletID"`

	Payments []*Payment `json:"payments" gorm:"foreignkey:WalletID"`
}

// IsEntity ...
func (m *Wallet) IsEntity() {}

// WalletChanges struct
type WalletChanges struct {
	ID           string
	UserID       *string
	Balance      float64
	WalletTypeID *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string

	AccountsIDs []*string
	PaymentsIDs []*string
}

// WalletTypeResultType struct
type WalletTypeResultType struct {
	EntityResultType
}

// WalletType struct
type WalletType struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        string     `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description"`
	WalletID    *string    `json:"walletId" gorm:"column:walletId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Wallet *Wallet `json:"wallet"`
}

// IsEntity ...
func (m *WalletType) IsEntity() {}

// WalletTypeChanges struct
type WalletTypeChanges struct {
	ID          string
	Name        string
	Description *string
	WalletID    *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// AccountProviderTypeResultType struct
type AccountProviderTypeResultType struct {
	EntityResultType
}

// AccountProviderType struct
type AccountProviderType struct {
	ID                string     `json:"id" gorm:"column:id;primary_key"`
	Name              string     `json:"name" gorm:"column:name"`
	Description       *string    `json:"description" gorm:"column:description"`
	AccountProviderID *string    `json:"accountProviderId" gorm:"column:accountProviderId"`
	UpdatedAt         *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt         time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy         *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy         *string    `json:"createdBy" gorm:"column:createdBy"`

	AccountProvider *AccountProvider `json:"accountProvider"`
}

// IsEntity ...
func (m *AccountProviderType) IsEntity() {}

// AccountProviderTypeChanges struct
type AccountProviderTypeChanges struct {
	ID                string
	Name              string
	Description       *string
	AccountProviderID *string
	UpdatedAt         *time.Time
	CreatedAt         time.Time
	UpdatedBy         *string
	CreatedBy         *string
}

// AccountProviderResultType struct
type AccountProviderResultType struct {
	EntityResultType
}

// AccountProvider struct
type AccountProvider struct {
	ID                    string     `json:"id" gorm:"column:id;primary_key"`
	Name                  string     `json:"name" gorm:"column:name"`
	Description           *string    `json:"description" gorm:"column:description"`
	Address               *string    `json:"address" gorm:"column:address"`
	Phone                 *string    `json:"phone" gorm:"column:phone"`
	AccountProviderTypeID *string    `json:"accountProviderTypeId" gorm:"column:accountProviderTypeId"`
	UpdatedAt             *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt             time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy             *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy             *string    `json:"createdBy" gorm:"column:createdBy"`

	Accounts []*Account `json:"accounts" gorm:"foreignkey:AccountProviderID"`

	AccountProviderType *AccountProviderType `json:"accountProviderType"`
}

// IsEntity ...
func (m *AccountProvider) IsEntity() {}

// AccountProviderChanges struct
type AccountProviderChanges struct {
	ID                    string
	Name                  string
	Description           *string
	Address               *string
	Phone                 *string
	AccountProviderTypeID *string
	UpdatedAt             *time.Time
	CreatedAt             time.Time
	UpdatedBy             *string
	CreatedBy             *string

	AccountsIDs []*string
}

// AccountResultType struct
type AccountResultType struct {
	EntityResultType
}

// Account struct
type Account struct {
	ID                string     `json:"id" gorm:"column:id;primary_key"`
	AccountNumber     string     `json:"accountNumber" gorm:"column:accountNumber"`
	Balance           float64    `json:"balance" gorm:"column:balance;default:0.0"`
	AccountProviderID *string    `json:"accountProviderId" gorm:"column:accountProviderId"`
	WalletID          *string    `json:"walletId" gorm:"column:walletId"`
	UpdatedAt         *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt         time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy         *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy         *string    `json:"createdBy" gorm:"column:createdBy"`

	AccountProvider *AccountProvider `json:"accountProvider"`

	Wallet *Wallet `json:"wallet"`

	Payments []*Payment `json:"payments" gorm:"foreignkey:AccountID"`
}

// IsEntity ...
func (m *Account) IsEntity() {}

// AccountChanges struct
type AccountChanges struct {
	ID                string
	AccountNumber     string
	Balance           float64
	AccountProviderID *string
	WalletID          *string
	UpdatedAt         *time.Time
	CreatedAt         time.Time
	UpdatedBy         *string
	CreatedBy         *string

	PaymentsIDs []*string
}

// PaymentChannelResultType struct
type PaymentChannelResultType struct {
	EntityResultType
}

// PaymentChannel struct
type PaymentChannel struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        string     `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description"`
	PaymentID   *string    `json:"paymentId" gorm:"column:paymentId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Payment *Payment `json:"payment"`
}

// IsEntity ...
func (m *PaymentChannel) IsEntity() {}

// PaymentChannelChanges struct
type PaymentChannelChanges struct {
	ID          string
	Name        string
	Description *string
	PaymentID   *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// PaymentTypeResultType struct
type PaymentTypeResultType struct {
	EntityResultType
}

// PaymentType struct
type PaymentType struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Name        string     `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description"`
	PaymentID   *string    `json:"paymentId" gorm:"column:paymentId"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`

	Payment *Payment `json:"payment"`
}

// IsEntity ...
func (m *PaymentType) IsEntity() {}

// PaymentTypeChanges struct
type PaymentTypeChanges struct {
	ID          string
	Name        string
	Description *string
	PaymentID   *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// PaymentResultType struct
type PaymentResultType struct {
	EntityResultType
}

// Payment struct
type Payment struct {
	ID               string     `json:"id" gorm:"column:id;primary_key"`
	PaymentRef       *string    `json:"paymentRef" gorm:"column:paymentRef;index:prefidx"`
	Amount           float64    `json:"amount" gorm:"column:amount;default:0.0"`
	Concept          *string    `json:"concept" gorm:"column:concept"`
	WalletID         *string    `json:"walletId" gorm:"column:walletId"`
	AccountID        *string    `json:"accountId" gorm:"column:accountId"`
	PaymentChannelID *string    `json:"paymentChannelId" gorm:"column:paymentChannelId"`
	PaymentTypeID    *string    `json:"paymentTypeId" gorm:"column:paymentTypeId"`
	UpdatedAt        *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt        time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy        *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy        *string    `json:"createdBy" gorm:"column:createdBy"`

	Wallet *Wallet `json:"wallet"`

	Account *Account `json:"account"`

	PaymentChannel *PaymentChannel `json:"paymentChannel"`

	PaymentType *PaymentType `json:"paymentType"`
}

// IsEntity ...
func (m *Payment) IsEntity() {}

// PaymentChanges struct
type PaymentChanges struct {
	ID               string
	PaymentRef       *string
	Amount           float64
	Concept          *string
	WalletID         *string
	AccountID        *string
	PaymentChannelID *string
	PaymentTypeID    *string
	UpdatedAt        *time.Time
	CreatedAt        time.Time
	UpdatedBy        *string
	CreatedBy        *string
}

// ApplyChanges used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
