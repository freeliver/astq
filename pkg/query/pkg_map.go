package query

// Code generated by astgen v1.0 with go1.11.2 DO NOT EDIT

import (
	"strings"
)

// Keys returns a new slice containing the set of map keys. The order is unspecified.
func (m Packages) Keys() (keys []string) {
	for name := range m {
		keys = append(keys, name)
	}

	return
}

// Values returns a new slice containing the set of map values. The order is unspecified.
func (m Packages) Values() (values []*Package) {
	for _, value := range m {
		values = append(values, value)
	}

	return
}

// Contains reports whether key is within map.
func (m Packages) Contains(key string) bool {
	_, found := m[key]

	return found
}

// Clone returns a shadow copy of map.
func (m Packages) Clone() Packages {
	cloned := make(Packages)

	for key, value := range m {
		cloned[key] = value
	}

	return cloned
}

// Filter filters the map to only include elements for which filter returns true.
func (m Packages) Filter(filter func(key string, value *Package) bool) Packages {
	filtered := make(Packages)

	for key, value := range m {
		if filter(key, value) {
			filtered[key] = value
		}
	}

	return filtered
}

// WithPrefix filters the map to only include elements for which contains prefix.
func (m Packages) WithPrefix(prefix string) Packages {
	return m.Filter(func(key string, value *Package) bool {
		return strings.HasPrefix(key, prefix)
	})
}

// WithSuffix filters the map to only include elements for which contains suffix.
func (m Packages) WithSuffix(suffix string) Packages {
	return m.Filter(func(key string, value *Package) bool {
		return strings.HasSuffix(key, suffix)
	})
}
