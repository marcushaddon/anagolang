package anagrams

// WordRepository can lookup words
type WordRepository interface {
	Search(pattern string, limit int) ([]string, error)
}

// AnagramFinder finds all anagrams
type AnagramFinder struct {
	wr WordRepository
}

// TODO: handle errors!
func (af AnagramFinder) getEnglishPermutations(b []byte, left, right int, perms *[]string) {
	if left == right {
		perm := string(b)
		*perms = append(*perms, perm)
	} else {

		for i := left; i <= right; i++ {
			b[i], b[left] = b[left], b[i]
			prefix := string(b[0 : i+1])
			result, _ := af.wr.Search(prefix, 1)
			if len(result) > 0 {
				af.getEnglishPermutations(b, left+1, right, perms)
			}
			b[i], b[left] = b[left], b[i]
		}
	}
}

// GetAnagrams prints all
func (af AnagramFinder) GetAnagrams(s string) []string {
	bytes, left, right := []byte(s), 0, len(s)-1
	p := &[]string{}
	af.getEnglishPermutations(bytes, left, right, p)
	return *p
}
