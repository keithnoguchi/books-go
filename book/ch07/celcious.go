package ch07

import (
	"flag"
	"fmt"
)

type Celcious float64

func (c Celcious) String() string {
	return fmt.Sprintf("%.05fC", c)
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
	switch unit {
	case "C", "c":
		f.Celcious = Celcious(val)
		return nil
	}
	return fmt.Errorf("invalid temperature: %q", s)
}
