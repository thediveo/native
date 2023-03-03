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

// Reader reads numbers in various fixed sizes and always in native endianess
// from byte sequences. It is a much more convenient combination of
// [bytes.Buffer] with [encoding.Binary], with better error handling on top of
// it.
//
// In particular, Reader enters an error state when trying to read past the end
// of the supplied data, returning only 0 for the numbers to read. A typical
// pattern thus is to read (or trying to read) all numbers from the data and
// only at the end checking if all was for nothing.
type Reader struct {
	buff *bytes.Buffer
	err  error
}

// NewReader returns a Reader to read numbers in various fixed sizes from the
// specified data.
func NewReader(data []byte) *Reader {
	return &Reader{
		buff: bytes.NewBuffer(data),
	}
}

// Error returns the first error encountered while reading from the data, or nil
// if there was no error.
func (r *Reader) Error() error { return r.err }

// Uint8 reads and returns an uint8. If it encounters an error or if the Reader
// is already in error state, 0 is returned instead.
func (r *Reader) Uint8() uint8 {
	if r.err != nil {
		return 0
	}
	var b uint8
	b, r.err = r.buff.ReadByte()
	return b
}

// Uint16 reads and returns an uint16. If it encounters an error or if the
// Reader is already in error state, 0 is returned instead.
func (r *Reader) Uint16() uint16 {
	if r.err != nil {
		return 0
	}
	// Benchmarking has shown that using unsafe.Pointer is 3½ times as fast as
	// using encoding/binary and without any heap allocations (whereas
	// encoding/binary causes a single heap allocation).
	var v uint16
	_, r.err = r.buff.Read((*[2]byte)(unsafe.Pointer(&v))[:])
	if r.err != nil {
		return 0
	}
	return v
}

// Uint32 reads and returns an uint32. If it encounters an error or if the
// Reader is already in error state, 0 is returned instead.
func (r *Reader) Uint32() uint32 {
	if r.err != nil {
		return 0
	}
	var v uint32
	_, r.err = r.buff.Read((*[4]byte)(unsafe.Pointer(&v))[:])
	if r.err != nil {
		return 0
	}
	return v
}

// Uint64 reads and returns an uint64. If it encounters an error or if the
// Reader is already in error state, 0 is returned instead.
func (r *Reader) Uint64() uint64 {
	if r.err != nil {
		return 0
	}
	var v uint64
	_, r.err = r.buff.Read((*[8]byte)(unsafe.Pointer(&v))[:])
	if r.err != nil {
		return 0
	}
	return v
}
