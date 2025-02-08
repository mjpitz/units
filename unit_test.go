package units_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.pitz.tech/units"
)

func TestOptions(t *testing.T) {
	options := units.Options{'f', -1}
	original := options

	units.Precision(2).Apply(&options)
	require.Equal(t, 2, options.Precision, "precision was not set properly")

	units.Format('b').Apply(&options)
	require.Equal(t, byte('b'), options.Format, "format was not set properly")

	original.Apply(&options)
	require.Equal(t, -1, options.Precision, "precision was not reset properly")
	require.Equal(t, byte('f'), options.Format, "format was not reset properly")
}
