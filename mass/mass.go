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

package mass

import (
	"sort"

	"go.pitz.tech/units"
)

// Mass is a fundamental property of matter that quantifies the amount of substance contained in an object. It
// represents the resistance of an object to changes in its motion and is a measure of the total number and type of
// particles present. Mass is commonly measured in units such as kilograms or pounds and is a key factor in determining
// the gravitational force acting on an object.
type Mass int64

func (u Mass) As(other Mass) float64 {
	return float64(u) / float64(other)
}

func (u *Mass) Set(val string) error {
	v, err := all.Parse(val)
	if err != nil {
		return err
	}

	*u = v
	return nil
}

func (u Mass) String() string {
	return SI.Format(u)
}

const (
	Nanogram  Mass = 1
	Microgram      = 1000 * Nanogram
	Milligram      = 1000 * Microgram
	Centigram      = 10 * Milligram
	Decigram       = 10 * Centigram
	Gram           = 10 * Decigram
	Decagram       = 10 * Gram
	Hectogram      = 10 * Decagram
	Kilogram       = 10 * Hectogram

	Grain         = 64798910 * Nanogram
	Dram          = Ounce / 16
	Ounce         = Pound / 16
	Pound         = 7000 * Grain
	Stone         = 14 * Pound
	Quarter       = 2 * Stone
	Hundredweight = 4 * Quarter
	Ton           = 20 * Hundredweight

	USCanadaHundredweight = 100 * Pound
	USCanadaTon           = 20 * USCanadaHundredweight

	TroyPennyweight = 24 * Grain
	TroyOunce       = 20 * TroyPennyweight
	TroyPound       = 12 * TroyOunce
)

var (
	SI = units.Unit[Mass]{
		{Nanogram, []string{"ng"}},
		{Microgram, []string{"μg", "ug"}},
		{Milligram, []string{"mg"}},
		{Centigram, []string{"cg"}},
		{Decigram, []string{"dg"}},
		{Gram, []string{"g"}},
		{Decagram, []string{"dag"}},
		{Hectogram, []string{"hg"}},
		{Kilogram, []string{"kg"}},
	}

	Imperial = units.Unit[Mass]{
		{Grain, []string{"gr"}},
		{Dram, []string{"dr"}},
		{Ounce, []string{"oz"}},
		{Pound, []string{"lb"}},
		{Stone, []string{"st"}},
		{Quarter, []string{"qr"}},
		{Hundredweight, []string{"cwt"}},
		{Ton, []string{"ton"}},
	}

	// The Troy unit of measure is frequently used when dealing with precious metals.
	Troy = units.Unit[Mass]{
		{Grain, []string{"gr"}},
		{TroyPennyweight, []string{"dwt"}},
		{TroyOunce, []string{"ozt"}},
		{TroyPound, []string{"lbt"}},
	}

	// USCanada is a special format that uses smaller values for Hundredweight and Ton.
	USCanada = units.Unit[Mass]{
		{Grain, []string{"gr"}},
		{Dram, []string{"dr"}},
		{Ounce, []string{"oz"}},
		{Pound, []string{"lb"}},
		{USCanadaHundredweight, []string{"cwt"}},
		{USCanadaTon, []string{"ton"}},
	}

	all units.Unit[Mass]
)

func init() {
	all = append(all, SI...)
	all = append(all, Imperial...)

	// ensure all is sorted
	// this is easy to do on a unit by unit basis, but is much harder when intermixing units
	sort.Slice(all, func(i, j int) bool {
		return all[i].Size < all[j].Size
	})
}
