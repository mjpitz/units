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

package units

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

var (
	// ErrValueDoesNotMatchPattern notifies the caller that the provided string text did not match our expected format.
	ErrValueDoesNotMatchPattern = fmt.Errorf("value does not match pattern")

	format = regexp.MustCompile("(?P<measure>[0-9.]+)(?P<symbol>[a-zA-Z ]{1,6})")
)

// Number defines a constraint to ensure the values provided to units are integer based (i.e. we're working with whole
// numbers). This makes sure we're working with whole numbers and not handling fractions internally. This forces the
// programmer to handle all rounding and truncation.
type Number interface {
	constraints.Signed
}

// Symbol defines how various sizes should be labeled. Some values may contain multiple labels, but the preferred label
// that will be used when printing should come first in the list.
type Symbol[T Number] struct {
	Size  T
	Label []string
}

// A Unit of measure is a standardized quantity used to quantify and express the magnitude or value of a physical
// quantity. It establishes a reference point or a standard against which measurements can be made and compared. Units
// of measure provide a consistent and universally understood way to communicate and exchange information about
// quantities in various domains, such as length, mass, time, temperature, volume, and many others.
type Unit[T Number] []Symbol[T]

// Format uses the underlying Unit to convert the provided value to a human-readable string. The benefit to this
// abstraction is that so long as a unit shares a common base unit, multiple formats can be used to represent the
// underlying value (for example, metric vs imperial).
func (u Unit[T]) Format(value T, opts ...Option) (str string) {
	if len(u) == 0 || value == 0 {
		return ""
	}

	options := Options{'f', -1}
	for _, opt := range opts {
		opt.Apply(&options)
	}

	for i := len(u); value > 0 && i > 1; i-- {
		if value >= u[i-1].Size {
			str += strconv.FormatInt(int64(value/u[i-1].Size), 10) + u[i-1].Label[0]
			value = value % u[i-1].Size
		}
	}

	if value > 0 {
		rem := float64(value) / float64(u[0].Size)
		str += strconv.FormatFloat(rem, options.Format, options.Precision, 64) + u[0].Label[0]
	}

	return str
}

// Parse attempts to convert the provided string value to its equivalent numeric representation.
func (u Unit[T]) Parse(val string) (size T, err error) {
	val = strings.TrimSpace(val)
	if val == "" || val == "0" {
		return 0, nil
	}

	var factor T = 1
	switch val[0] {
	case '-':
		factor = T(-1)
		val = val[1:]
	case '+':
		val = val[1:]
	}

	// todo: improve this so we don't need regex
	if !format.MatchString(val) {
		return size, ErrValueDoesNotMatchPattern
	}

	// todo: memoize this
	idx := make(map[string]T)
	for i := len(u); i > 0; i-- {
		for _, label := range u[i-1].Label {
			idx[label] = u[i-1].Size
		}
	}

	for _, match := range format.FindAllStringSubmatch(val, -1) {
		label := strings.TrimSpace(match[2])
		unit, ok := idx[label]
		if !ok {
			return 0, fmt.Errorf("unrecognized symbol: %s", label)
		}

		parsed, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			return 0, err
		}

		size += T(parsed * float64(unit))
	}

	return factor * size, nil
}

// Options defines various formatting options that can be used to tailor a given Unit.Format call including which number
// format is used to render a floating point number and its associated precision.
type Options struct {
	Format    byte
	Precision int
}

// Apply this Options configured values to the destination Options configured values.
func (o Options) Apply(dst *Options) {
	if o.Format > 0 {
		dst.Format = o.Format
	}

	if o.Precision != 0 {
		dst.Precision = o.Precision
	}
}

// Option provides a mechanism to tune specific behaviors of the system such as formatting and numeric precision.
type Option interface {
	Apply(*Options)
}

// OptionFunc provides a functional way to tune specific options of the underlying behavior.
type OptionFunc func(*Options)

func (f OptionFunc) Apply(opts *Options) {
	if f != nil {
		f(opts)
	}
}

// Format configures how floating point numbers are rendered. The section below can be found in strconv.FormatFloat.
//
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise),
// 'G' ('E' for large exponents, 'f' otherwise),
// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
//
// For further information, please consult the strconv documents.
func Format(fmt byte) OptionFunc {
	return func(opts *Options) {
		opts.Format = fmt
	}
}

// Precision configures how many decimal places are used when rendering floating point numbers.
func Precision(precision int) OptionFunc {
	return func(opts *Options) {
		opts.Precision = precision
	}
}
