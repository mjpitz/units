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

package length

import (
	"sort"

	"go.pitz.tech/units"
)

// Length is a fundamental physical quantity that measures the extent of an object or distance between two points. It is
// typically defined as the longest dimension of an object or the spatial separation between two locations. Length is
// commonly measured in units such as meters, feet, or inches, and plays a crucial role in various scientific,
// engineering, and everyday applications.
type Length int64

func (u Length) As(other Length) float64 {
	return float64(u) / float64(other)
}

func (u *Length) Set(val string) error {
	v, err := all.Parse(val)
	if err != nil {
		return err
	}

	*u = v
	return nil
}

func (u Length) String() string {
	return SI.Format(u)
}

const (
	Nanometer  Length = 1
	Micrometer        = 1000 * Nanometer
	Millimeter        = 1000 * Micrometer
	Centimeter        = 10 * Millimeter
	Decimeter         = 10 * Centimeter
	Meter             = 10 * Decimeter
	Decameter         = 10 * Meter
	Hectometer        = 10 * Decameter
	Kilometer         = 10 * Hectometer

	Thou   = 25400 * Nanometer
	Inch   = 1000 * Thou
	Foot   = 12 * Inch
	Yard   = 3 * Foot
	Mile   = 1760 * Yard
	League = 3 * Mile
)

var (
	SI = units.Unit[Length]{
		{Nanometer, []string{"nm"}},
		{Micrometer, []string{"μm", "um"}},
		{Millimeter, []string{"mm"}},
		{Centimeter, []string{"cm"}},
		{Decimeter, []string{"dm"}},
		{Meter, []string{"m"}},
		{Decameter, []string{"dam"}},
		{Hectometer, []string{"hm"}},
		{Kilometer, []string{"km"}},
	}

	Imperial = units.Unit[Length]{
		{Inch, []string{"in", "\""}},
		{Foot, []string{"ft", "'"}},
		{Yard, []string{"yd"}},
		{Mile, []string{"mi"}},
		{League, []string{"lea"}},
	}

	all units.Unit[Length]
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
