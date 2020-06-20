package datastructure

// use map impl simple set struct

type Set map[string]struct{}

func (set Set) Add(val string) {
	set[val] = struct{}{}
}

func (set Set) Discard(val string) {
	delete(set, val)
}

// return a, b set intersection
func (set Set) Intersection(other Set)  Set {
	s := Set{}
	for k := range set {
		if _, ok := other[k]; ok {
			s[k] = struct{}{}
		}
	}
	return s
}

func (set Set) Union(other Set) Set {
	s := Set{}
	for k := range set {
		s[k] = struct{}{}
	}
	for k := range other {
		s[k] = struct{}{}
	}
	return s
}

func (set Set) Difference(other Set) Set {
	s := Set{}
	for k := range set {
		if _, ok := other[k]; !ok {
			s[k] = struct{}{}
		}
	}
	return s
}

// set and other has no inter section.
func (set Set) IsDisjoint(other Set) bool {
	for k := range set {
		if _, ok := other[k]; ok {
			return false
		}
	}
	return true
}

// set is sub set of other.
func (set Set) IsSubset(other Set) bool {
	for k := range set {
		if _, ok := other[k]; !ok {
			return false
		}
	}
	return true
}

func (set Set) IsSuperSet(other Set) bool {
	for k := range set {
		if _, ok := other[k]; !ok {
			return false
		}
	}
	return true
}

func (set Set) SymmetricDifference(other Set) Set {
	interSection := set.Intersection(other)
	return set.Difference(interSection).Union(other.Difference(interSection))
}
