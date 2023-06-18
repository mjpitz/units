# units

Package units offers shared code for handling standard units of measurement. It
takes inspiration from Golang's time library and applies similar ideas to other
units of measure. This is achieved through a common declaration called a Unit,
which organizes measurements in increasing order, often derived from a base
unit. A Unit also facilitates parsing and formatting of measurements to improve
computer use and human readability (respectively). To get a better understanding
of its functioning, you can explore the provided packages for more specific
information.

```go
import "go.pitz.tech/units"
```

## Usage

```go
var (
	// ErrValueDoesNotMatchPattern notifies the caller that the provided string text did not match our expected format.
	ErrValueDoesNotMatchPattern = fmt.Errorf("value does not match pattern")
)
```

#### type Number

```go
type Number interface {
	constraints.Integer
}
```

Number defines a constraint to ensure the values provided to units are integer
based (i.e. we're working with whole numbers). This makes sure we're working
with whole numbers and not handling fractions internally. This forces the
programmer to handle all rounding and truncation.

#### type Symbol

```go
type Symbol[T Number] struct {
	Size  T
	Label []string
}
```

Symbol defines how various sizes should be labeled. Some values may contain
multiple labels, but the preferred label that will be used when printing should
come first in the list.

#### type Unit

```go
type Unit[T Number] []Symbol[T]
```

A Unit of measure is a standardized quantity used to quantify and express the
magnitude or value of a physical quantity. It establishes a reference point or a
standard against which measurements can be made and compared. Units of measure
provide a consistent and universally understood way to communicate and exchange
information about quantities in various domains, such as length, mass, time,
temperature, volume, and many others.

#### func (Unit[T]) Format

```go
func (u Unit[T]) Format(value T) (str string)
```

Format uses the underlying Unit to convert the provided value to a
human-readable string. The benefit to this abstraction is that so long as a unit
shares a common base unit, multiple formats can be used to represent the
underlying value (for example, metric vs imperial).

#### func (Unit[T]) Parse

```go
func (u Unit[T]) Parse(val string) (size T, err error)
```

Parse attempts to convert the provided string value to its equivalent numeric
representation.
