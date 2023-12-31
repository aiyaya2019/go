// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMicrophoneServerDao is internal type for wrapping internal DAO implements.
type internalMicrophoneServerDao = *internal.MicrophoneServerDao

// microphoneServerDao is the data access object for table microphone_server.
// You can define custom methods on it to extend its functionality as you wish.
type microphoneServerDao struct {
	internalMicrophoneServerDao
}

var (
	// MicrophoneServer is globally public accessible object for table microphone_server operations.
	MicrophoneServer = microphoneServerDao{
		internal.NewMicrophoneServerDao(),
	}
)

// Fill with you ideas below.
