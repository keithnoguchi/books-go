package ch04

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

type CharCounter struct {
	Counts  map[rune]int
	invalid int
	utflen  [utf8.UTFMax + 1]int
}

func NewCharCounter() CharCounter {
	return CharCounter{
		Counts: make(map[rune]int),
	}
}

func (c *CharCounter) Count(in io.Reader) error {
	reader := bufio.NewReader(in)
	for {
		r, n, err := reader.ReadRune()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if r == unicode.ReplacementChar && n == 1 {
			c.invalid++
			continue
		}
		c.Counts[r]++
		c.utflen[n]++
	}
	return nil
}

func (c *CharCounter) UTFLen() [utf8.UTFMax + 1]int {
	return c.utflen
}

func (c *CharCounter) InvalidCount() int {
	return c.invalid
}
