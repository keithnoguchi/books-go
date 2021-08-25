package ch07

import (
	"flag"
	"fmt"
	"strings"
)

type Celcious float64

func (c Celcious) String() string {
	return fmt.Sprintf("%.05fC", c)
}

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.05F", f)
}

func ftoc(f Fahrenheit) Celcious {
	return Celcious((f - 32) * 5 / 9)
}

func CelciousFlag(name string, val Celcious, usage string) *Celcious {
	var f celciousFlag
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celcious
}

type celciousFlag struct{ Celcious }

func (f *celciousFlag) Set(s string) error {
	var unit string
	var val float64
	fmt.Sscanf(s, "%f%s", &val, &unit)
	switch strings.ToUpper(unit) {
	case "C":
		f.Celcious = Celcious(val)
		return nil
	case "F":
		f.Celcious = ftoc(Fahrenheit(val))
		return nil
	}
	return fmt.Errorf("invalid temperature: %q", s)
}
