package swiss

func (m *Map[K, V]) Keys() []K {
	var keys = make([]K, m.Count())
	// take a consistent view of the table in case
	// we rehash during iteration
	ctrl, groups := m.ctrl, m.groups
	// pick a random starting group
	g := randIntN(len(groups))
	keysIndex := 0
	for n := 0; n < len(groups); n++ {
		for s, c := range ctrl[g] {
			if c == empty || c == tombstone {
				continue
			}
			k := groups[g].keys[s]
			keys[keysIndex] = k
			keysIndex++
		}
		g++
		if g >= uint32(len(groups)) {
			g = 0
		}
	}
	return keys
}
