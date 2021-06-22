package registry

type Map map[string][]string

func New() (r Map) {
	r = make(map[string][]string, 0)
	return
}

func (r *Map) Add(key string, text string) {

	if 0 < len((*r)[key]) {

		(*r)[key] = make([]string, 1)
	}

	(*r)[key] = append((*r)[key], text)
	return
}
