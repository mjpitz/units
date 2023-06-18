# data

```go
import go.pitz.tech/units/data
```

## Usage

```go
const (
	Byte Size = 1

	Kibibyte = Byte << 10
	Mebibyte = Kibibyte << 10
	Gibibyte = Mebibyte << 10
	Tebibyte = Gibibyte << 10
	Pebibyte = Tebibyte << 10

	Kilobyte = Byte * 1000
	Megabyte = Kilobyte * 1000
	Gigabyte = Megabyte * 1000
	Terabyte = Gigabyte * 1000
	Petabyte = Terabyte * 1000
)
```

```go
var (
	Decimal = units.Unit[Size]{
		{Byte, []string{"B"}},
		{Kilobyte, []string{"kB"}},
		{Megabyte, []string{"MB"}},
		{Gigabyte, []string{"GB"}},
		{Terabyte, []string{"TB"}},
		{Petabyte, []string{"PB"}},
	}

	BinaryIEC = units.Unit[Size]{
		{Byte, []string{"B"}},
		{Kibibyte, []string{"KiB"}},
		{Mebibyte, []string{"MiB"}},
		{Gibibyte, []string{"GiB"}},
		{Tebibyte, []string{"TiB"}},
		{Pebibyte, []string{"PiB"}},
	}

	BinaryMemory = units.Unit[Size]{
		{Byte, []string{"B"}},
		{Kibibyte, []string{"KB"}},
		{Mebibyte, []string{"MB"}},
		{Gibibyte, []string{"GB"}},
		{Tebibyte, []string{"TB"}},
	}
)
```

#### type Size

```go
type Size int64
```

Size defines how we measure digital information (typically in the number of
bytes). Digital information refers to any data or content that is represented,
stored, or transmitted in a digital format. It is composed of discrete units
known as bits, which can be either 0 or 1. Digital information can take various
forms, such as text, images, audio, video, or any other type of data that can be
encoded and processed electronically.

#### func (Size) As

```go
func (u Size) As(other Size) float64
```

#### func (\*Size) Set

```go
func (u *Size) Set(val string) error
```

#### func (Size) String

```go
func (u Size) String() string
```
