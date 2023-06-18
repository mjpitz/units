# network

```go
import "go.pitz.tech/units/network"
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
		{Bit, []string{"bps"}},
		{Kilobit, []string{"kbps"}},
		{Megabit, []string{"Mbps"}},
		{Gigabit, []string{"Gbps"}},
		{Terabit, []string{"Tbps"}},
		{Petabit, []string{"Pbps"}},
	}

	BinaryIEC = units.Unit[Bandwidth]{
		{Bit, []string{"bps"}},
		{Kibibit, []string{"Kibps"}},
		{Mebibit, []string{"Mibps"}},
		{Gibibit, []string{"Gibps"}},
		{Tebibit, []string{"Tibps"}},
		{Pebibit, []string{"Pibps"}},
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
