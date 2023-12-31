// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalDeviceCustomDao is internal type for wrapping internal DAO implements.
type internalDeviceCustomDao = *internal.DeviceCustomDao

// deviceCustomDao is the data access object for table device_custom.
// You can define custom methods on it to extend its functionality as you wish.
type deviceCustomDao struct {
	internalDeviceCustomDao
}

var (
	// DeviceCustom is globally public accessible object for table device_custom operations.
	DeviceCustom = deviceCustomDao{
		internal.NewDeviceCustomDao(),
	}
)

// Fill with you ideas below.
