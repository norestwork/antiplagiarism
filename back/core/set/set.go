package set

type Set struct {
	set map[string]struct{}
}

func New(stringSlice []string) Set {
	newSet := Set{
		set: make(map[string]struct{}),
	}

	for i := range stringSlice {
		newSet.set[stringSlice[i]] = struct{}{}
	}

	return newSet
}

func (s *Set) Add(x string) {
	s.set[x] = struct{}{}
}

func (s *Set) GetLen() float64 {
	return float64(len(s.set))
}

func (s *Set) Union(bSet Set) Set {
	unitedSet := Set{
		set: make(map[string]struct{}),
	}

	for element := range s.set {
		unitedSet.Add(element)
	}

	for element := range bSet.set {
		unitedSet.Add(element)
	}

	return unitedSet
}

func (s *Set) Intersection(bSet Set) Set {
	intersectionSet := Set{
		set: make(map[string]struct{}),
	}

	for element := range s.set {
		if _, ok := bSet.set[element]; ok {
			intersectionSet.Add(element)
		}
	}

	return intersectionSet
}