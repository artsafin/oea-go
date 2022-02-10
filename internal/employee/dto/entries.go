package dto

type Entries []*Entry

func (e Entries) Len() int {
	return len(e)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (e Entries) Less(i, j int) bool {
	return e[i].Sort < e[j].Sort
}

// Swap swaps the elements with indexes i and j.
func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
