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

package mass_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.pitz.tech/units/mass"
)

func TestMass(t *testing.T) {
	require.Equal(t, 10.0, mass.Kilogram.As(mass.Hectogram))
	require.Equal(t, 10.0, mass.Hectogram.As(mass.Decagram))
	require.Equal(t, 10.0, mass.Decagram.As(mass.Gram))
	require.Equal(t, 10.0, mass.Gram.As(mass.Decigram))
	require.Equal(t, 10.0, mass.Decigram.As(mass.Centigram))
	require.Equal(t, 10.0, mass.Centigram.As(mass.Milligram))
	require.Equal(t, 1000.0, mass.Milligram.As(mass.Microgram))
	require.Equal(t, 1000.0, mass.Microgram.As(mass.Nanogram))

	require.Equal(t, 16.0, mass.Pound.As(mass.Ounce))
	require.Equal(t, 7000.0, mass.Pound.As(mass.Grain))
	require.Equal(t, 14.0, mass.Stone.As(mass.Pound))
	require.Equal(t, 2.0, mass.Quarter.As(mass.Stone))
	require.Equal(t, 4.0, mass.Hundredweight.As(mass.Quarter))
	require.Equal(t, 20.0, mass.Ton.As(mass.Hundredweight))

	require.Equal(t, "1kg", mass.Kilogram.String())
	require.Equal(t, "1hg", mass.Hectogram.String())
	require.Equal(t, "1dag", mass.Decagram.String())
	require.Equal(t, "1g", mass.Gram.String())
	require.Equal(t, "1dg", mass.Decigram.String())
	require.Equal(t, "1cg", mass.Centigram.String())
	require.Equal(t, "1mg", mass.Milligram.String())
	require.Equal(t, "1μg", mass.Microgram.String())
	require.Equal(t, "1ng", mass.Nanogram.String())

	require.Equal(t, "", mass.Imperial.Format(0))
	require.Equal(t, "1ton", mass.Imperial.Format(mass.Ton))
	require.Equal(t, "1cwt", mass.Imperial.Format(mass.Hundredweight))
	require.Equal(t, "1qr", mass.Imperial.Format(mass.Quarter))
	require.Equal(t, "1st", mass.Imperial.Format(mass.Stone))
	require.Equal(t, "1lb", mass.Imperial.Format(mass.Pound))
	require.Equal(t, "1oz", mass.Imperial.Format(mass.Ounce))
	require.Equal(t, "1gr", mass.Imperial.Format(mass.Grain))

	basic := 100 * mass.Gram

	testCases := []struct {
		set      string
		err      bool
		expected mass.Mass
	}{
		{"", false, 0},
		{"10kg", false, 10 * mass.Kilogram},
		{"1kg1hg1dag", false, mass.Kilogram + mass.Hectogram + mass.Decagram},
		{"100DNE", true, 0},
		{"BAD", true, 0},
	}

	for _, testCase := range testCases {
		//set empty
		err := (&basic).Set(testCase.set)
		if testCase.err {
			require.Error(t, err)
			continue
		}

		require.NoError(t, err)
		require.Equal(t, testCase.expected, basic)
	}
}
