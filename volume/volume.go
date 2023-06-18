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

package volume

import (
	"sort"

	"go.pitz.tech/units"
)

// Volume is a physical quantity that measures the amount of space occupied by an object or a substance. It is often
// described as the three-dimensional extent of an object or the capacity of a container to hold a substance. Volume is
// typically measured in cubic units, such as cubic meters or cubic feet, and plays a crucial role in various fields
// such as physics, engineering, and fluid dynamics.
type Volume int64

func (u Volume) As(other Volume) float64 {
	return float64(u) / float64(other)
}

func (u *Volume) Set(val string) error {
	v, err := all.Parse(val)
	if err != nil {
		return err
	}

	*u = v
	return nil
}

func (u Volume) String() string {
	return SI.Format(u)
}

const (
	Nanoliter  Volume = 1
	Microliter        = 1000 * Nanoliter
	Milliliter        = 1000 * Microliter
	Centiliter        = 10 * Milliliter
	Deciliter         = 10 * Centiliter
	Liter             = 10 * Deciliter
	Decaliter         = 10 * Liter
	Hectoliter        = 10 * Decaliter
	Kiloliter         = 10 * Hectoliter

	FluidOunce = 29573529 * Nanoliter // floz
	Gill       = 5 * FluidOunce       // gi
	Pint       = 4 * Gill             // pt
	Quart      = 2 * Pint             // qt
	Gallon     = 4 * Quart            // gal
)

var (
	SI = units.Unit[Volume]{
		{Nanoliter, []string{"nL"}},
		{Microliter, []string{"μL", "uL"}},
		{Milliliter, []string{"mL"}},
		{Centiliter, []string{"cL"}},
		{Deciliter, []string{"dL"}},
		{Liter, []string{"L"}},
		{Decaliter, []string{"daL"}},
		{Hectoliter, []string{"hL"}},
		{Kiloliter, []string{"kL"}},
	}

	Imperial = units.Unit[Volume]{
		{FluidOunce, []string{"fl oz"}},
		{Gill, []string{"gi"}},
		{Pint, []string{"pt"}},
		{Quart, []string{"qt"}},
		{Gallon, []string{"gal"}},
	}

	all units.Unit[Volume]
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
