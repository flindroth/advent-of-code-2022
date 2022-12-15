package util

func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	chunks := [][]T{}
	tmp := []T{}
	for i := 0; i < len(slice); i++ {
		tmp = append(tmp, slice[i])
		if len(tmp) == chunkSize {
			chunks = append(chunks, tmp)
			tmp = []T{}
		}
	}
	chunks = append(chunks, tmp)
	return chunks
}
