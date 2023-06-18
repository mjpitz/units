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

package data_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.pitz.tech/units/data"
)

func TestSize(t *testing.T) {
	require.Equal(t, 1024.0, data.Pebibyte.As(data.Tebibyte))
	require.Equal(t, 1024.0, data.Tebibyte.As(data.Gibibyte))
	require.Equal(t, 1024.0, data.Gibibyte.As(data.Mebibyte))
	require.Equal(t, 1024.0, data.Mebibyte.As(data.Kibibyte))
	require.Equal(t, 1024.0, data.Kibibyte.As(data.Byte))

	require.Equal(t, 1000.0, data.Petabyte.As(data.Terabyte))
	require.Equal(t, 1000.0, data.Terabyte.As(data.Gigabyte))
	require.Equal(t, 1000.0, data.Gigabyte.As(data.Megabyte))
	require.Equal(t, 1000.0, data.Megabyte.As(data.Kilobyte))
	require.Equal(t, 1000.0, data.Kilobyte.As(data.Byte))

	require.Equal(t, "1PiB", data.BinaryIEC.Format(data.Pebibyte))
	require.Equal(t, "1TiB", data.BinaryIEC.Format(data.Tebibyte))
	require.Equal(t, "1GiB", data.BinaryIEC.Format(data.Gibibyte))
	require.Equal(t, "1MiB", data.BinaryIEC.Format(data.Mebibyte))
	require.Equal(t, "1KiB", data.BinaryIEC.Format(data.Kibibyte))

	require.Equal(t, "1PB", data.Petabyte.String())
	require.Equal(t, "1TB", data.Terabyte.String())
	require.Equal(t, "1GB", data.Gigabyte.String())
	require.Equal(t, "1MB", data.Megabyte.String())
	require.Equal(t, "1kB", data.Kilobyte.String())
	require.Equal(t, "1B", data.Byte.String())

	basic := 100 * data.Gibibyte

	testCases := []struct {
		set      string
		err      bool
		expected data.Size
	}{
		{"", false, 0},
		{"10GiB", false, 10 * data.Gibibyte},
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
