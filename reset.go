package reset

const reuseThreshold = 100_0000

// ClearMap clears all entries in the given map.
// For small maps (length ≤ reuseThreshold),
// it deletes keys to reuse the existing map
// and keep its allocated capacity.
// For large maps (length > reuseThreshold),
// it creates a new map to free memory and
// avoid holding a large internal structure.This
// approach balances memory usage and performance,
// since Go maps do not expose their internal
// capacity like slices.
func ClearMap[K comparable, V any](m *map[K]V) {
	if len(*m) > reuseThreshold {
		*m = make(map[K]V)
	} else {
		for k := range *m {
			delete(*m, k)
		}
	}
}
