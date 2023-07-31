// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalFontDao is internal type for wrapping internal DAO implements.
type internalFontDao = *internal.FontDao

// fontDao is the data access object for table font.
// You can define custom methods on it to extend its functionality as you wish.
type fontDao struct {
	internalFontDao
}

var (
	// Font is globally public accessible object for table font operations.
	Font = fontDao{
		internal.NewFontDao(),
	}
)

// Fill with you ideas below.