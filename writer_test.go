// Copyright 2023 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package native

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("reading natives", func() {

	It("writes uint8s", func() {
		w := NewWriter(0)
		w.Uint8(0x42)
		w.Uint8(66)
		Expect(w.Bytes()).To(Equal([]byte{0x42, 66}))
	})

	It("writes uint16s", func() {
		w := NewWriter(2 * 2)
		w.Uint16(0x1042)
		w.Uint16(666)
		r := NewReader(w.Bytes())
		Expect(r.Uint16()).To(Equal(uint16(0x1042)))
		Expect(r.Uint16()).To(Equal(uint16(666)))
	})

	It("writes uint32s", func() {
		w := NewWriter(0)
		w.Uint32(0x11220042)
		w.Uint32(666)
		r := NewReader(w.Bytes())
		Expect(r.Uint32()).To(Equal(uint32(0x11220042)))
		Expect(r.Uint32()).To(Equal(uint32(666)))
	})

	It("writes uint64s", func() {
		w := NewWriter(0)
		w.Uint64(0x1122334455660042)
		w.Uint64(666)
		r := NewReader(w.Bytes())
		Expect(r.Uint64()).To(Equal(uint64(0x1122334455660042)))
		Expect(r.Uint64()).To(Equal(uint64(666)))
	})

})
