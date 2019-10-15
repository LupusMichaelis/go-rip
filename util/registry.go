package util

type Registry map[string][]string

func MakeRegistry() (r Registry) {
	r = make(map[string][]string, 0)
	return
}

func (r *Registry) Add(key string, text string) {

	if 0 < len((*r)[key]) {

		(*r)[key] = make([]string, 1)
	}

	(*r)[key] = append((*r)[key], text)
	return
}
