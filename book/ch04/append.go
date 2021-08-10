package ch04

func Append(a []interface{}, b ...interface{}) []interface{} {
	clen := len(a) + len(b)
	var c []interface{}
	if clen <= cap(a) {
		c = a[:clen]
	} else {
		c = make([]interface{}, clen, clen*2)
		copy(c, a)
	}
	for i, v := range b {
		c[len(a)+i] = v
	}
	return c
}
