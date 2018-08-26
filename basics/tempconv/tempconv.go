// tempconv package which does temperature conversion
// computation from celsius to fahrenheit and vice versa

package tempconv

import (
    "fmt"
)

type Fahrenheit float64
type Celsius float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func (c Celsius) String() string {
    return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g °F", f)
}

func C2F (c Celsius) Fahrenheit {
    return Fahrenheit(c * 9 / 5 + 32)
}

func F2C (f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}
