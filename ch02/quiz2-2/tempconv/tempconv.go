package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC             = 0
	BoilingC              = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func (pd Pound) String() string    { return fmt.Sprintf("%gpound", pd) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
