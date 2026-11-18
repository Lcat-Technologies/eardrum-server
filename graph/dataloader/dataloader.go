package dataloader

import (
	"context"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/Lcat-Technologies/eardrum-postgres/device"
	"github.com/Lcat-Technologies/eardrum-postgres/transaction"
	"github.com/Lcat-Technologies/eardrum-postgres/user"
	"github.com/Lcat-Technologies/eardrum-server/graph/model"
	"github.com/vikstrous/dataloadgen"
	"gorm.io/gorm"
)

// Loaders holds all individual batch loaders for a single request
type Loaders struct {
	UserLoader   *dataloadgen.Loader[string, *model.User]
	DeviceLoader *dataloadgen.Loader[string, *string]
}

// NewLoaders creates a fresh set of loaders
func NewLoaders(db *gorm.DB) *Loaders {
	// 1. User Batch Loader Function
	userBatchFn := func(ctx context.Context, usernames []string) ([]*model.User, []error) {
		// Fetch all requested users in batch from database
		dbUsers, err := user.GetUsersWithUsernames(db, usernames)
		if err != nil {
			return nil, []error{err}
		}

		// Map database results by username for quick lookup
		userMap := make(map[string]*model.User)
		for _, dbuser := range dbUsers {
			uuidStr := dbuser.GetUUID()
			userMap[dbuser.GetUserName()] = &model.User{
				Username:              dbuser.GetUserName(),
				PhoneNumber:           dbuser.GetPhoneNumber(),
				UUID:                  &uuidStr,
				FacialEmbeddings:      pointer.Get(dbuser.GetFacialEmbeddings()),
				AccountBalanceInCents: int(dbuser.GetAccountBalanceInCents()),
				FaceEnrollmentStatus:  dbuser.GetFaceEnrollmentStatus(),
				PinEnrollmentStatus:   dbuser.GetPinStatus(),
			}
		}

		// Return results in the EXACT same order as requested usernames
		output := make([]*model.User, len(usernames))
		for i, uname := range usernames {
			if u, ok := userMap[uname]; ok {
				output[i] = u
			}
		}
		return output, nil
	}

	// 2. Corrected Device Batch Loader Function (keyed by Transaction ID)
	deviceBatchFn := func(ctx context.Context, txIDs []string) ([]*string, []error) {
		// 1. Fetch ALL transactions in a SINGLE query using an IN clause
		// e.g., transaction.GetTransactionsByReferences(db, txIDs)
		transactions, err := transaction.GetTransactionsByReferences(db, txIDs)
		if err != nil {
			return nil, []error{err}
		}

		// 2. Collect all unique Device IDs from those transactions
		deviceIDs := make([]string, 0, len(transactions))
		txToDeviceMap := make(map[string]string) // maps txID -> deviceID
		for _, ts := range transactions {
			txID := ts.GetTransactionID() // or reference
			deviceID, err := ts.GetTransactionDeviceID()
			if err == nil && deviceID != "" {
				deviceIDs = append(deviceIDs, deviceID)
				txToDeviceMap[txID] = deviceID
			}
		}

		// 3. Fetch ALL corresponding devices in a SINGLE query using an IN clause
		// e.g., device.GetDevicesByIDs(db, deviceIDs)
		devices, err := device.GetDevicesByIDs(deviceIDs, db)
		if err != nil {
			return nil, []error{err}
		}

		// Map deviceID -> device model string for quick lookup
		deviceModelMap := make(map[string]string)
		for _, ds := range devices {
			deviceModelMap[ds.GetDeviceID()] = pointer.GetString(ds.GetDeviceModel())
		}

		// 4. Build the final output slice matching the EXACT order of the requested txIDs
		output := make([]*string, len(txIDs))

		for i, txID := range txIDs {
			deviceID, hasDevice := txToDeviceMap[txID]
			if !hasDevice {
				output[i] = nil
				continue
			}
			modelStr, hasModel := deviceModelMap[deviceID]
			if !hasModel {
				output[i] = nil
				continue
			}
			output[i] = &modelStr
		}

		return output, nil
	}

	return &Loaders{
		UserLoader:   dataloadgen.NewLoader(userBatchFn, dataloadgen.WithWait(2*time.Millisecond), dataloadgen.WithBatchCapacity(100)),
		DeviceLoader: dataloadgen.NewLoader(deviceBatchFn, dataloadgen.WithWait(2*time.Millisecond), dataloadgen.WithBatchCapacity(100)),
	}
}
