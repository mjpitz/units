// Copyright (c) 2023 Mya Pitzeruse
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
// OR OTHER DEALINGS IN THE SOFTWARE.

package data

import (
	"sort"

	"go.pitz.tech/units"
)

// Size defines how we measure digital information (typically in the number of bytes). Digital information refers to any
// data or content that is represented, stored, or transmitted in a digital format. It is composed of discrete units
// known as bits, which can be either 0 or 1. Digital information can take various forms, such as text, images, audio,
// video, or any other type of data that can be encoded and processed electronically.
type Size int64

func (u Size) As(other Size) float64 {
	return float64(u) / float64(other)
}

func (u *Size) Set(val string) error {
	v, err := all.Parse(val)
	if err != nil {
		return err
	}

	*u = v
	return nil
}

func (u Size) String() string {
	return Decimal.Format(u)
}

const (
	Byte Size = 1

	Kibibyte = Byte << 10
	Mebibyte = Kibibyte << 10
	Gibibyte = Mebibyte << 10
	Tebibyte = Gibibyte << 10
	Pebibyte = Tebibyte << 10

	Kilobyte = Byte * 1000
	Megabyte = Kilobyte * 1000
	Gigabyte = Megabyte * 1000
	Terabyte = Gigabyte * 1000
	Petabyte = Terabyte * 1000
)

var (
	Decimal = units.Unit[Size]{
		{Byte, []string{"B"}},
		{Kilobyte, []string{"kB"}},
		{Megabyte, []string{"MB"}},
		{Gigabyte, []string{"GB"}},
		{Terabyte, []string{"TB"}},
		{Petabyte, []string{"PB"}},
	}

	BinaryIEC = units.Unit[Size]{
		{Byte, []string{"B"}},
		{Kibibyte, []string{"KiB"}},
		{Mebibyte, []string{"MiB"}},
		{Gibibyte, []string{"GiB"}},
		{Tebibyte, []string{"TiB"}},
		{Pebibyte, []string{"PiB"}},
	}

	BinaryMemory = units.Unit[Size]{
		{Byte, []string{"B"}},
		{Kibibyte, []string{"KB"}}, // "kilobyte"
		{Mebibyte, []string{"MB"}}, // "megabyte"
		{Gibibyte, []string{"GB"}}, // "gigabyte"
		{Tebibyte, []string{"TB"}}, // "terabyte"
	}

	all = units.Unit[Size]{
		{Byte, []string{"B"}},
		{Kilobyte, []string{"kB"}},
		{Kibibyte, []string{"KiB"}},
		{Megabyte, []string{"MB"}},
		{Mebibyte, []string{"MiB"}},
		{Gigabyte, []string{"GB"}},
		{Gibibyte, []string{"GiB"}},
		{Terabyte, []string{"TB"}},
		{Tebibyte, []string{"TiB"}},
		{Petabyte, []string{"PB"}},
		{Pebibyte, []string{"PiB"}},
	}
)

func init() {
	// ensure all is sorted
	// this is easy to do on a unit by unit basis, but is much harder when intermixing units
	sort.Slice(all, func(i, j int) bool {
		return all[i].Size < all[j].Size
	})
}
