// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalHtmlFormatDao is internal type for wrapping internal DAO implements.
type internalHtmlFormatDao = *internal.HtmlFormatDao

// htmlFormatDao is the data access object for table html_format.
// You can define custom methods on it to extend its functionality as you wish.
type htmlFormatDao struct {
	internalHtmlFormatDao
}

var (
	// HtmlFormat is globally public accessible object for table html_format operations.
	HtmlFormat = htmlFormatDao{
		internal.NewHtmlFormatDao(),
	}
)

// Fill with you ideas below.