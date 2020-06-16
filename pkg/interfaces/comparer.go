package interfaces

// Comparer defines interface for comparing one thing to another
type Comparer interface {
	// Compare intended to compare A to B (A being the instance implementing Comparer).  Return negative number if A less than B.  Return 0 if A equals B.  Return 1 if A greater than B.
	Compare(b interface{}) int
}
