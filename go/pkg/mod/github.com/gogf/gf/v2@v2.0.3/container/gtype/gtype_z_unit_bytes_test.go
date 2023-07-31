// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtype_test

import (
	"testing"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Bytes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewBytes([]byte("abc"))
		iClone := i.Clone()
		t.AssertEQ(iClone.Set([]byte("123")), []byte("abc"))
		t.AssertEQ(iClone.Val(), []byte("123"))

		// 空参测试
		i1 := gtype.NewBytes()
		t.AssertEQ(i1.Val(), nil)
	})
}

func Test_Bytes_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		b := []byte("i love gf")
		i := gtype.NewBytes(b)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewBytes()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.Assert(err, nil)
		t.Assert(i2.Val(), b)
	})
}

func Test_Bytes_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Bytes
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
