// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package testutil

import (
	"reflect"

	. "github.com/onsi/gomega"
)

func ExpectFuncsEqual(f1, f2 interface{}) {
	val1 := reflect.ValueOf(f1)
	val2 := reflect.ValueOf(f2)
	Expect(val1.Pointer()).To(Equal(val2.Pointer()))
}
