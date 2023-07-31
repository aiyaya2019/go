// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalTaskQueueDao is internal type for wrapping internal DAO implements.
type internalTaskQueueDao = *internal.TaskQueueDao

// taskQueueDao is the data access object for table task_queue.
// You can define custom methods on it to extend its functionality as you wish.
type taskQueueDao struct {
	internalTaskQueueDao
}

var (
	// TaskQueue is globally public accessible object for table task_queue operations.
	TaskQueue = taskQueueDao{
		internal.NewTaskQueueDao(),
	}
)

// Fill with you ideas below.