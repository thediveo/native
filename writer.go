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
	"bytes"
	"unsafe"
)

// Writer writes numbers in various fixed sizes and always in native endianess
// to byte sequences. It is a much more convenient combination of [bytes.Buffer]
// with [encoding.Binary], with better performance on top of it.
type Writer struct {
	buff *bytes.Buffer
}

// NewWriter returns a Writer to write numbers in various fixed sizes to a
// buffer. If the capacity is non-zero, then the Writer returned will already
// have the specified capacity allocated for its write buffer.
func NewWriter(cap int) *Writer {
	w := &Writer{
		buff: &bytes.Buffer{},
	}
	if cap > 0 {
		w.buff.Grow(cap)
	}
	return w
}

// Bytes returns the data written. Any later write after calling Bytes will
// invalidate the previously returned result.
func (w *Writer) Bytes() []byte {
	return w.buff.Bytes()
}

// Uint8 writes the specified uint8 value to the buffer.
func (w *Writer) Uint8(v uint8) {
	w.buff.WriteByte(v)
}

// Uint16 writes the specified uint16 value to the buffer.
func (w *Writer) Uint16(v uint16) {
	w.buff.Write((*[2]byte)(unsafe.Pointer(&v))[:])
}

// Uint32 writes the specified uint32 value to the buffer.
func (w *Writer) Uint32(v uint32) {
	w.buff.Write((*[4]byte)(unsafe.Pointer(&v))[:])
}

// Uint64 writes the specified uint64 value to the buffer.
func (w *Writer) Uint64(v uint64) {
	w.buff.Write((*[8]byte)(unsafe.Pointer(&v))[:])
}
