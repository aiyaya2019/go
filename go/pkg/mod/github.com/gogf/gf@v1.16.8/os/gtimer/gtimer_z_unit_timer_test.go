// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Timer Operations

package gtimer_test

import (
	"testing"
	"time"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/test/gtest"
)

func New() *gtimer.Timer {
	return gtimer.New()
}

func TestTimer_Add_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		//fmt.Println("start", time.Now())
		timer.Add(200*time.Millisecond, func() {
			//fmt.Println("job1", time.Now())
			array.Append(1)
		})
		timer.Add(200*time.Millisecond, func() {
			//fmt.Println("job2", time.Now())
			array.Append(1)
		})
		timer.Add(400*time.Millisecond, func() {
			//fmt.Println("job3", time.Now())
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 5)
		timer.Close()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_Start_Stop_Close(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.Add(1000*time.Millisecond, func() {
			array.Append(1)
		})
		t.Assert(array.Len(), 0)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		timer.Stop()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		timer.Start()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
		timer.Close()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestJob_Reset(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		job := timer.AddSingleton(500*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		job.Reset()
		time.Sleep(300 * time.Millisecond)
		job.Reset()
		time.Sleep(300 * time.Millisecond)
		job.Reset()
		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_AddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.AddSingleton(200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_AddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.AddOnce(200*time.Millisecond, func() {
			array.Append(1)
		})
		timer.AddOnce(200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		timer.Close()
		time.Sleep(250 * time.Millisecond)
		fixedLength := array.Len()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), fixedLength)
	})
}

func TestTimer_AddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.AddTimes(200*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestTimer_DelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.DelayAdd(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddJob(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.DelayAddEntry(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		}, false, 100, gtimer.StatusReady)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.DelayAddSingleton(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.DelayAddOnce(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_DelayAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.DelayAddTimes(200*time.Millisecond, 500*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(600 * time.Millisecond)
		t.Assert(array.Len(), 2)

		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestTimer_AddLessThanInterval(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := gtimer.New(gtimer.TimerOptions{
			Interval: 100 * time.Millisecond,
		})
		array := garray.New(true)
		timer.Add(20*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(50 * time.Millisecond)
		t.Assert(array.Len(), 0)

		time.Sleep(110 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(110 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestTimer_AddLeveledJob1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		//glog.Println("start")
		timer.DelayAdd(1000*time.Millisecond, 1000*time.Millisecond, func() {
			//glog.Println("add")
			array.Append(1)
		})
		time.Sleep(1500 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1300 * time.Millisecond)
		//glog.Println("check")
		t.Assert(array.Len(), 1)
	})
}

func TestTimer_Exit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		timer := New()
		array := garray.New(true)
		timer.Add(200*time.Millisecond, func() {
			array.Append(1)
			gtimer.Exit()
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}
