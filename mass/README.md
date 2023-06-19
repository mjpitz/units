# mass

```go
import "go.pitz.tech/units/mass"
```

## Usage

```go
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

	Slug = 14593902937205 * Nanogram
)
```

```go
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
		{TroyPennyweight, []string{"dw t"}},
		{TroyOunce, []string{"oz t"}},
		{TroyPound, []string{"lb t"}},
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
)
```

#### type Mass

```go
type Mass int64
```

Mass is a fundamental property of matter that quantifies the amount of substance
contained in an object. It represents the resistance of an object to changes in
its motion and is a measure of the total number and type of particles present.
Mass is commonly measured in units such as kilograms or pounds and is a key
factor in determining the gravitational force acting on an object.

#### func (Mass) As

```go
func (u Mass) As(other Mass) float64
```

#### func (\*Mass) Set

```go
func (u *Mass) Set(val string) error
```

#### func (Mass) String

```go
func (u Mass) String() string
```
