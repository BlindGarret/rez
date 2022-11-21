package rez

// ToPtr simple one liner for translating something to a pointer while being sure we broke any references via a copy.
func toPtr[T any](o T) *T {
	temp := o
	return &temp
}
