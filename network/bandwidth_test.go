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

package network_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.pitz.tech/units/network"
)

func TestBandwidth(t *testing.T) {
	require.Equal(t, 1024.0, network.Pebibit.As(network.Tebibit))
	require.Equal(t, 1024.0, network.Tebibit.As(network.Gibibit))
	require.Equal(t, 1024.0, network.Gibibit.As(network.Mebibit))
	require.Equal(t, 1024.0, network.Mebibit.As(network.Kibibit))
	require.Equal(t, 1024.0, network.Kibibit.As(network.Bit))

	require.Equal(t, 1000.0, network.Petabit.As(network.Terabit))
	require.Equal(t, 1000.0, network.Terabit.As(network.Gigabit))
	require.Equal(t, 1000.0, network.Gigabit.As(network.Megabit))
	require.Equal(t, 1000.0, network.Megabit.As(network.Kilobit))
	require.Equal(t, 1000.0, network.Kilobit.As(network.Bit))

	require.Equal(t, "1Pibps", network.Pebibit.String())
	require.Equal(t, "1Tibps", network.Tebibit.String())
	require.Equal(t, "1Gibps", network.Gibibit.String())
	require.Equal(t, "1Mibps", network.Mebibit.String())
	require.Equal(t, "1Kibps", network.Kibibit.String())
	require.Equal(t, "1bps", network.Bit.String())

	require.Equal(t, "1Pbps", network.Decimal.Format(network.Petabit))
	require.Equal(t, "1Tbps", network.Decimal.Format(network.Terabit))
	require.Equal(t, "1Gbps", network.Decimal.Format(network.Gigabit))
	require.Equal(t, "1Mbps", network.Decimal.Format(network.Megabit))
	require.Equal(t, "1kbps", network.Decimal.Format(network.Kilobit))
	require.Equal(t, "1bps", network.Decimal.Format(network.Bit))

	basic := 100 * network.Gibibit

	testCases := []struct {
		set      string
		err      bool
		expected network.Bandwidth
	}{
		{"", false, 0},
		{"10Gibps", false, 10 * network.Gibibit},
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
