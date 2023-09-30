package swiss

func (m *Map[K, V]) Keys() []K {
	var allKeys = make([]K, 0, m.Count())
	// take a consistent view of the table in case
	// we rehash during iteration
	ctrl, groups := m.ctrl, m.groups
	// pick a random starting group
	g := randIntN(len(groups))
	allKeysIndex := 0
	for n := 0; n < len(groups); n++ {
		for s, c := range ctrl[g] {
			if c == empty || c == tombstone {
				continue
			}
			k := groups[g].keys[s]
			allKeys[allKeysIndex] = k
			allKeysIndex++
			g++
			if g >= uint32(len(groups)) {
				g = 0
			}
		}
	}
	return allKeys
}
