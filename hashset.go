package hashset

type CustomHashSetMap map[string]struct{}

type HashSet struct {
	set CustomHashSetMap
}

func (h *HashSet) Size() int {
	return len(h.set)
}

func (h *HashSet) Add(value string) {
	h.set[value] = struct{}{}
}

func (h *HashSet) ToArray() []string {
	var result []string
	for key := range h.set {
		result = append(result, key)
	}
	return result
}

func New(values ...string) *HashSet {
	set := HashSet{make(CustomHashSetMap)}
	for _, value := range values {
		set.set[value] = struct{}{}
	}
	return &set
}
