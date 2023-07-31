// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMicrophoneServerTerminalDao is internal type for wrapping internal DAO implements.
type internalMicrophoneServerTerminalDao = *internal.MicrophoneServerTerminalDao

// microphoneServerTerminalDao is the data access object for table microphone_server_terminal.
// You can define custom methods on it to extend its functionality as you wish.
type microphoneServerTerminalDao struct {
	internalMicrophoneServerTerminalDao
}

var (
	// MicrophoneServerTerminal is globally public accessible object for table microphone_server_terminal operations.
	MicrophoneServerTerminal = microphoneServerTerminalDao{
		internal.NewMicrophoneServerTerminalDao(),
	}
)

// Fill with you ideas below.