// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalTranslateLangDao is internal type for wrapping internal DAO implements.
type internalTranslateLangDao = *internal.TranslateLangDao

// translateLangDao is the data access object for table translate_lang.
// You can define custom methods on it to extend its functionality as you wish.
type translateLangDao struct {
	internalTranslateLangDao
}

var (
	// TranslateLang is globally public accessible object for table translate_lang operations.
	TranslateLang = translateLangDao{
		internal.NewTranslateLangDao(),
	}
)

// Fill with you ideas below.
