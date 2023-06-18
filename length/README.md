# length

```go
import "go.pitz.tech/units/length"
```

## Usage

```go
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
```

```go
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
)
```

#### type Length

```go
type Length int64
```

Length is a fundamental physical quantity that measures the extent of an object
or distance between two points. It is typically defined as the longest dimension
of an object or the spatial separation between two locations. Length is commonly
measured in units such as meters, feet, or inches, and plays a crucial role in
various scientific, engineering, and everyday applications.

#### func (Length) As

```go
func (u Length) As(other Length) float64
```

#### func (\*Length) Set

```go
func (u *Length) Set(val string) error
```

#### func (Length) String

```go
func (u Length) String() string
```
