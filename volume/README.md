# volume

```go
import "go.pitz.tech/units/volume"
```

## Usage

```go
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
```

```go
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
)
```

#### type Volume

```go
type Volume int64
```

Volume is a physical quantity that measures the amount of space occupied by an
object or a substance. It is often described as the three-dimensional extent of
an object or the capacity of a container to hold a substance. Volume is
typically measured in cubic units, such as cubic meters or cubic feet, and plays
a crucial role in various fields such as physics, engineering, and fluid
dynamics.

#### func (Volume) As

```go
func (u Volume) As(other Volume) float64
```

#### func (\*Volume) Set

```go
func (u *Volume) Set(val string) error
```

#### func (Volume) String

```go
func (u Volume) String() string
```
