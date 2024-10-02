package shape

type bound struct {
	min, max int32
}

type bounds struct {
	start        int
	upper, lower []bound
}

func (b *bounds) getBounds() (int, int) {
	return b.start - len(b.upper), b.start + len(b.lower)
}

func (b *bounds) getBound(i int) (int, int) {
	i -= b.start

	rows := b.upper

	if i < 0 {
		i = ^i
		rows = b.lower
	}

	if i >= len(rows) {
		return 0, 0
	}

	return int(rows[i].min), int(rows[i].max)
}

func (b *bounds) addPoint(i, j int) {
	if b.upper == nil && b.lower == nil {
		b.start = i
	}

	i -= b.start
	rows := &b.upper
	if i < 0 {
		rows = &b.lower
		i = ^i
	}

	for i >= len(*rows) {
		*rows = append(*rows, bound{})
	}

	r := &(*rows)[i]
	if r.min == r.max {
		*r = bound{int32(j), int32(j + 1)}
	} else {
		*r = bound{min(r.min, int32(j)), max(r.max, int32(j+1))}
	}
}
