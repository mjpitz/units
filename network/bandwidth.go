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

package network

import (
	"sort"
	"strings"

	"go.pitz.tech/units"
)

// Bandwidth refers to the maximum amount of data that can be transmitted over a network connection within a given
// timeframe. It is typically measured in bits per second (bps) and represents the capacity of the network to transfer
// data. A higher bandwidth allows for faster data transmission, while a lower bandwidth may result in slower or delayed
// data transfer.
type Bandwidth int64

func (u Bandwidth) As(other Bandwidth) float64 {
	return float64(u) / float64(other)
}

func (u *Bandwidth) Set(val string) error {
	val = strings.TrimSuffix(val, "/s")
	val = strings.TrimSuffix(val, "ps")

	v, err := all.Parse(val)
	if err != nil {
		return err
	}

	*u = v
	return nil
}

func (u Bandwidth) String() string {
	return BinaryIEC.Format(u) + "/s"
}

const (
	Bit Bandwidth = 1

	Kibibit = Bit << 10
	Mebibit = Kibibit << 10
	Gibibit = Mebibit << 10
	Tebibit = Gibibit << 10
	Pebibit = Tebibit << 10

	Kilobit = Bit * 1000
	Megabit = Kilobit * 1000
	Gigabit = Megabit * 1000
	Terabit = Gigabit * 1000
	Petabit = Terabit * 1000
)

var (
	Decimal = units.Unit[Bandwidth]{
		{Bit, []string{"b"}},
		{Kilobit, []string{"kb"}},
		{Megabit, []string{"Mb"}},
		{Gigabit, []string{"Gb"}},
		{Terabit, []string{"Tb"}},
		{Petabit, []string{"Pb"}},
	}

	BinaryIEC = units.Unit[Bandwidth]{
		{Bit, []string{"b"}},
		{Kibibit, []string{"Kib"}},
		{Mebibit, []string{"Mib"}},
		{Gibibit, []string{"Gib"}},
		{Tebibit, []string{"Tib"}},
		{Pebibit, []string{"Pib"}},
	}

	all = units.Unit[Bandwidth]{
		{Bit, []string{"b"}},
		{Kilobit, []string{"kb"}},
		{Kibibit, []string{"Kib"}},
		{Megabit, []string{"Mb"}},
		{Mebibit, []string{"Mib"}},
		{Gigabit, []string{"Gb"}},
		{Gibibit, []string{"Gib"}},
		{Terabit, []string{"Tb"}},
		{Tebibit, []string{"Tib"}},
		{Petabit, []string{"Pb"}},
		{Pebibit, []string{"Pib"}},
	}
)

func init() {
	// ensure all is sorted
	// this is easy to do on a unit by unit basis, but is much harder when intermixing units
	sort.Slice(all, func(i, j int) bool {
		return all[i].Size < all[j].Size
	})
}
