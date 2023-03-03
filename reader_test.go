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

	It("reads uint8s", func() {
		r := NewReader([]byte{0x42, 77})
		Expect(r.Uint8()).To(Equal(uint8(0x42)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint8()).To(Equal(uint8(77)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint8()).To(BeZero())
		Expect(r.Error()).To(HaveOccurred())
		Expect(r.Uint8()).To(BeZero())
	})

	It("reads uint16s", func() {
		data := make([]byte, 2*2)
		native.PutUint16(data[0:2], 0x42)
		native.PutUint16(data[2:4], 666)
		r := NewReader(data)
		Expect(r.Uint16()).To(Equal(uint16(0x42)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint16()).To(Equal(uint16(666)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint16()).To(BeZero())
		Expect(r.Error()).To(HaveOccurred())
		Expect(r.Uint16()).To(BeZero())
	})

	It("reads uint32s", func() {
		data := make([]byte, 2*4)
		native.PutUint32(data[0:4], 0x42)
		native.PutUint32(data[4:8], 0x60066)
		r := NewReader(data)
		Expect(r.Uint32()).To(Equal(uint32(0x42)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint32()).To(Equal(uint32(0x60066)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint32()).To(BeZero())
		Expect(r.Error()).To(HaveOccurred())
		Expect(r.Uint32()).To(BeZero())
	})

	It("reads uint64s", func() {
		data := make([]byte, 2*8)
		native.PutUint64(data[0:8], 0x42)
		native.PutUint64(data[8:16], 0x600000066)
		r := NewReader(data)
		Expect(r.Uint64()).To(Equal(uint64(0x42)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint64()).To(Equal(uint64(0x600000066)))
		Expect(r.Error()).NotTo(HaveOccurred())
		Expect(r.Uint64()).To(BeZero())
		Expect(r.Error()).To(HaveOccurred())
		Expect(r.Uint64()).To(BeZero())
	})

})
