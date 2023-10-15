package leanauthz

import "strings"

// Resource is a string representing a resource identifier or locator. For example, a URL, a file path, or a namespaced id.
// It is usually in the form of <type>/<id> but this is up to your application.
// e.g., "documents/doc_abcdef", "articles/ABdj24aFp_aD/comments/65F1sd46", "s3://my-bucket/some-folder/my-object"
type Resource string

func (r Resource) String() string {
	return string(r)
}

// SplitBy splits a Resource string by the given separator and returns its parts.
func (r Resource) SplitBy(sep string) []string {
	return strings.Split(string(r), sep)
}

// SplitBySlash splits a Resource string by the slash separator and returns its parts.
func (r Resource) SplitBySlash() []string {
	return r.SplitBy(`/`)
}

// NewResource returns a new Resource. It cannot be empty.
func NewResource(str string) Resource {
	if str == "" {
		panic("authz: resource string cannot be empty")
	}
	return Resource(str)
}

// NewNamespacedResource returns a new Resource with the given namespace.
func NewNamespacedResource(ns, id string) Resource {
	if ns == "" {
		panic("authz: resource namespace string cannot be empty")
	}
	if id == "" {
		panic("authz: resource id string cannot be empty")
	}
	return Resource(strings.Join([]string{ns, id}, `/`))
}
