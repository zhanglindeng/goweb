package util

import "testing"

func TestSizeFormat(t *testing.T) {

	t.Log(SizeFormat(1))
	t.Log(SizeFormat(12))
	t.Log(SizeFormat(123))
	t.Log(SizeFormat(1234))
	t.Log(SizeFormat(12345))
	t.Log(SizeFormat(123456))
	t.Log(SizeFormat(1234567))
	t.Log(SizeFormat(12345678))
	t.Log(SizeFormat(123456789))
	t.Log(SizeFormat(100<<20))

}

