package main

func HexLiteral() bool {
	x := 0xFff
	y := 0xFFF
	z := 0xfff
	a := 0xffe

	return (x == y) && (y == z)
}
