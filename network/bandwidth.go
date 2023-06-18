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
	v, err := all.Parse(val)
	if err != nil {
		return err
	}

	*u = v
	return nil
}

func (u Bandwidth) String() string {
	return BinaryIEC.Format(u)
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
		{Bit, []string{"bps"}},
		{Kilobit, []string{"kbps"}},
		{Megabit, []string{"Mbps"}},
		{Gigabit, []string{"Gbps"}},
		{Terabit, []string{"Tbps"}},
		{Petabit, []string{"Pbps"}},
	}

	BinaryIEC = units.Unit[Bandwidth]{
		{Bit, []string{"bps"}},
		{Kibibit, []string{"Kibps"}},
		{Mebibit, []string{"Mibps"}},
		{Gibibit, []string{"Gibps"}},
		{Tebibit, []string{"Tibps"}},
		{Pebibit, []string{"Pibps"}},
	}

	all = units.Unit[Bandwidth]{
		{Bit, []string{"bps"}},
		{Kilobit, []string{"kbps"}},
		{Kibibit, []string{"Kibps"}},
		{Megabit, []string{"Mbps"}},
		{Mebibit, []string{"Mibps"}},
		{Gigabit, []string{"Gbps"}},
		{Gibibit, []string{"Gibps"}},
		{Terabit, []string{"Tbps"}},
		{Tebibit, []string{"Tibps"}},
		{Petabit, []string{"Pbps"}},
		{Pebibit, []string{"Pibps"}},
	}
)

func init() {
	// ensure all is sorted
	// this is easy to do on a unit by unit basis, but is much harder when intermixing units
	sort.Slice(all, func(i, j int) bool {
		return all[i].Size < all[j].Size
	})
}
