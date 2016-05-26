package main

import(
	"bytes"
	"time"
	"strconv"
	"strings"
)

func padder (a string, i int, fill string, right bool) string{
	b := make([]byte,(i - len(a)))
	for k := range b{
		b[k] = fill[0]
	}
	if right {
		return string_concat(a, string(b))
	}
	return string_concat(string(b), a)
}

func pad_space_left(a string, i int) string{
	return padder(a,i," ",false)
}

func pad_space_right(a string, i int) string{
	return padder(a,i," ",true)
}

func pad_zero_left(a string, i int) string{
	return padder(a,i,"0",false)
}

func pad_zero_right(a string, i int) string{
	return padder(a,i,"0",true)
}

func left_trim(a string, t string) string{
	for k, v := range a{
		if string(v) == t{
			continue
		}
		return a[k:]
	}
	return a
}

func to_decimal(a string) string{
	l := len(a)
	return string_concat(a[:l-2],".", a[l-2:])
}

func push_slice(s *[]string, t string) int{
	if t == ""{
		return 0
	}
	*s = append((*s), t)
	return 1
}

func string_concat(s ...string) string{
	var b bytes.Buffer
	for _, v := range s{
		b.WriteString(v)
	}
	return b.String()
}

func keys_join(m map[string]bool) string{
	pi := 0
	a := make([]string, len(m))
	for k, _ := range m{
		a[pi] = k
		pi++
	}
	return strings.Join(a,",")
}
