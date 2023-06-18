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

package length_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.pitz.tech/units/length"
)

func TestLength(t *testing.T) {
	require.Equal(t, 10.0, length.Kilometer.As(length.Hectometer))
	require.Equal(t, 10.0, length.Hectometer.As(length.Decameter))
	require.Equal(t, 10.0, length.Decameter.As(length.Meter))
	require.Equal(t, 10.0, length.Meter.As(length.Decimeter))
	require.Equal(t, 10.0, length.Decimeter.As(length.Centimeter))
	require.Equal(t, 10.0, length.Centimeter.As(length.Millimeter))
	require.Equal(t, 1000.0, length.Millimeter.As(length.Micrometer))
	require.Equal(t, 1000.0, length.Micrometer.As(length.Nanometer))

	require.Equal(t, 3.0, length.League.As(length.Mile))
	require.Equal(t, 1760.0, length.Mile.As(length.Yard))
	require.Equal(t, 3.0, length.Yard.As(length.Foot))
	require.Equal(t, 12.0, length.Foot.As(length.Inch))
	require.Equal(t, 1000.0, length.Inch.As(length.Thou))

	require.Equal(t, "1km", length.Kilometer.String())
	require.Equal(t, "1hm", length.Hectometer.String())
	require.Equal(t, "1dam", length.Decameter.String())
	require.Equal(t, "1m", length.Meter.String())
	require.Equal(t, "1dm", length.Decimeter.String())
	require.Equal(t, "1cm", length.Centimeter.String())
	require.Equal(t, "1mm", length.Millimeter.String())
	require.Equal(t, "1μm", length.Micrometer.String())
	require.Equal(t, "1nm", length.Nanometer.String())

	require.Equal(t, "1lea", length.Imperial.Format(length.League))
	require.Equal(t, "1mi", length.Imperial.Format(length.Mile))
	require.Equal(t, "1yd", length.Imperial.Format(length.Yard))
	require.Equal(t, "1ft", length.Imperial.Format(length.Foot))
	require.Equal(t, "1in", length.Imperial.Format(length.Inch))

	basic := 100 * length.Meter

	testCases := []struct {
		set      string
		err      bool
		expected length.Length
	}{
		{"", false, 0},
		{"10km", false, 10 * length.Kilometer},
		{"100DNE", true, 0},
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
