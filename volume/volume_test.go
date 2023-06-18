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

package volume_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.pitz.tech/units/volume"
)

func TestVolume(t *testing.T) {
	require.Equal(t, 10.0, volume.Kiloliter.As(volume.Hectoliter))
	require.Equal(t, 10.0, volume.Hectoliter.As(volume.Decaliter))
	require.Equal(t, 10.0, volume.Decaliter.As(volume.Liter))
	require.Equal(t, 10.0, volume.Liter.As(volume.Deciliter))
	require.Equal(t, 10.0, volume.Deciliter.As(volume.Centiliter))
	require.Equal(t, 10.0, volume.Centiliter.As(volume.Milliliter))
	require.Equal(t, 1000.0, volume.Milliliter.As(volume.Microliter))
	require.Equal(t, 1000.0, volume.Microliter.As(volume.Nanoliter))

	require.Equal(t, 4.0, volume.Gallon.As(volume.Quart))
	require.Equal(t, 2.0, volume.Quart.As(volume.Pint))
	require.Equal(t, 4.0, volume.Pint.As(volume.Gill))
	require.Equal(t, 5.0, volume.Gill.As(volume.FluidOunce))

	require.Equal(t, "1kL", volume.Kiloliter.String())
	require.Equal(t, "1hL", volume.Hectoliter.String())
	require.Equal(t, "1daL", volume.Decaliter.String())
	require.Equal(t, "1L", volume.Liter.String())
	require.Equal(t, "1dL", volume.Deciliter.String())
	require.Equal(t, "1cL", volume.Centiliter.String())
	require.Equal(t, "1mL", volume.Milliliter.String())
	require.Equal(t, "1μL", volume.Microliter.String())
	require.Equal(t, "1nL", volume.Nanoliter.String())

	require.Equal(t, "1gal", volume.Imperial.Format(volume.Gallon))
	require.Equal(t, "1qt", volume.Imperial.Format(volume.Quart))
	require.Equal(t, "1pt", volume.Imperial.Format(volume.Pint))
	require.Equal(t, "1gi", volume.Imperial.Format(volume.Gill))
	require.Equal(t, "1fl oz", volume.Imperial.Format(volume.FluidOunce))

	basic := 100 * volume.Liter

	testCases := []struct {
		set      string
		err      bool
		expected volume.Volume
	}{
		{"", false, 0},
		{"10kL", false, 10 * volume.Kiloliter},
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
