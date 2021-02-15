package ch04

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type UTF8ByteSize int

const (
	Invalid UTF8ByteSize = iota
	Single
	Double
	Tripple
	Quad
)

type Counter struct {
	in      io.Reader
	counts  map[rune]int
	utflen  [Quad]int
	invalid int
}

func NewCounter(in io.Reader) Counter {
	return Counter{
		in:     in,
		counts: map[rune]int{},
	}
}

func (c *Counter) Scan() error {
	in := bufio.NewReader(c.in)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if r == unicode.ReplacementChar && n == 1 {
			c.invalid++
			continue
		}
		c.counts[r]++
		c.utflen[n]++
	}
	return nil
}

func (c *Counter) Count(r rune) int {
	return c.counts[r]
}

func (c *Counter) CountByUTF8ByteSize(size UTF8ByteSize) (int, error) {
	if size == Invalid {
		return 0, fmt.Errorf("invalid utf8 byte size")
	}
	return c.utflen[size], nil
}

func (c *Counter) InvalidCount() int {
	return c.invalid
}
