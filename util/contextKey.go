package util

//ContextKey string key for context value lookups
type ContextKey string

//NewKey creates a new context string key using a package and name designator
func NewKey(pkg string, name string) ContextKey {
	return ContextKey(pkg + "." + name)
}
