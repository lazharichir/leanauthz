package leanauthz

// Pattern is a glob-like pattern. It is used to match actiosn, resources, and principals.
type Pattern string

// NewPattern returns a new Pattern.
func NewPattern(str string) Pattern {
	return Pattern(str)
}

// String returns the string representation of a Pattern.
func (p Pattern) String() string {
	return string(p)
}
