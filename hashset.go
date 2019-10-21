package hashset

type CustomHashSetMap map[interface{}]struct{}

type HashSet struct {
	set CustomHashSetMap
}

func New(values ...interface{}) HashSet {
	hs := HashSet{set: make(CustomHashSetMap)}
	for _, value := range values {
		hs.set[value] = struct{}{}
	}
	return hs
}

func (hs HashSet) Size() int {
	return len(hs.set)
}

func (h HashSet) Add(value interface{}) {
	h.set[value] = struct{}{}

	//if _, ok := h.set[value]; ok {
	//	return errors.New("Cannot add duplicate values.")
	//}
}

func (h *HashSet) ToArray() []interface{} {
	var result []interface{}
	for key := range h.set {
		result = append(result, key)
	}
	return result
}
