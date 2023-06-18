# network

```go
import go.pitz.tech/units/network
```

## Usage

```go
const (
	Bit Bandwidth = 1

	Kibibit = Bit << 10
	Mebibit = Kibibit << 10
	Gibibit = Mebibit << 10
	Tebibit = Gibibit << 10
	Pebibit = Tebibit << 10

	Kilobit = Bit * 1000
	Megabit = Kilobit * 1000
	Gigabit = Megabit * 1000
	Terabit = Gigabit * 1000
	Petabit = Terabit * 1000
)
```

```go
var (
	Decimal = units.Unit[Bandwidth]{
		{Bit, []string{"b"}},
		{Kilobit, []string{"kb"}},
		{Megabit, []string{"Mb"}},
		{Gigabit, []string{"Gb"}},
		{Terabit, []string{"Tb"}},
		{Petabit, []string{"Pb"}},
	}

	BinaryIEC = units.Unit[Bandwidth]{
		{Bit, []string{"b"}},
		{Kibibit, []string{"Kib"}},
		{Mebibit, []string{"Mib"}},
		{Gibibit, []string{"Gib"}},
		{Tebibit, []string{"Tib"}},
		{Pebibit, []string{"Pib"}},
	}
)
```

#### type Bandwidth

```go
type Bandwidth int64
```

Bandwidth refers to the maximum amount of data that can be transmitted over a
network connection within a given timeframe. It is typically measured in bits
per second (bps) and represents the capacity of the network to transfer data. A
higher bandwidth allows for faster data transmission, while a lower bandwidth
may result in slower or delayed data transfer.

#### func (Bandwidth) As

```go
func (u Bandwidth) As(other Bandwidth) float64
```

#### func (\*Bandwidth) Set

```go
func (u *Bandwidth) Set(val string) error
```

#### func (Bandwidth) String

```go
func (u Bandwidth) String() string
```
