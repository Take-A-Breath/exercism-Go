// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // Not a triangle
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	invalid := !valid(a, b, c)
	equilateral := a == b && b == c && c == a
	isosceles := a == b || b == c || c == a

	switch {
	case invalid:
		return NaT
	case equilateral:
		return Equ
	case isosceles:
		return Iso
	default:
		return Sca
	}

}

func valid(a, b, c float64) bool {
	ab := a+b > c
	bc := b+c > a
	ac := a+c > b

	inequalityPass := ab && bc && ac
	notZero := a > 0 && b > 0 && c > 0

	return inequalityPass && notZero
}
