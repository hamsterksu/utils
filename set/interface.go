// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package set

import (
	"github.com/juju/names"
)

type baseSet interface {
	Size() int
	IsEmpty() bool
}

type tagSet interface {
	Add(names.Tag)
	Remove(names.Tag)
	Contains(names.Tag) bool
	Values() []names.Tag
	SortedValues() []names.Tag

	Union(Tags) Tags
	Intersection(Tags) Tags
	Difference(Tags) Tags

	baseSet
}