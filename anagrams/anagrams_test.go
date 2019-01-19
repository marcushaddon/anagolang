package anagrams

import (
	"fmt"
	"strings"
	"testing"
)

// MockWordRepo sometimes return a result
type OneResultWordRepo struct {
	result string
}

// Search finds results words beginning with 'ar'
func (orr OneResultWordRepo) Search(pattern string, limit int) ([]string, error) {
	if strings.HasPrefix(orr.result, pattern) {
		return []string{"aardvark"}, nil
	}

	return nil, nil
}

// FakeWordRepo always returns a result
type FakeWordRepo struct{}

// Search alwyas returns a result
func (fwr FakeWordRepo) Search(pattern string, limit int) ([]string, error) {
	return []string{"dog"}, nil
}

/**
 * TestAllPermutationsAreAnagrams
 * checks for correct number (n!)
 * of permutations
 */
func TestPermutationsCount(t *testing.T) {
	af := AnagramFinder{
		wr: FakeWordRepo{},
	}

	word := "adongfish"

	anagrams := af.GetAnagrams(word)

	anagramsBang, factor := len(word), len(word)-1

	for factor > 0 {
		anagramsBang *= factor
		factor--
	}

	if len(anagrams) != anagramsBang {
		fmt.Printf("Expected %v but got %v\n", anagramsBang, len(anagrams))
		t.Fail()
	}
}

/**
 * TestPrefixes asserts that if a word only
 * has one anagram, one and only one will be found.
 */
func TestPrefixes(t *testing.T) {
	result := "arcm"

	wr := OneResultWordRepo{
		result: result,
	}

	af := AnagramFinder{
		wr: wr,
	}

	anagrams := af.GetAnagrams("cram")
	fmt.Println(anagrams)
	if len(anagrams) != 1 {
		fmt.Printf("Expected one anagram, found %v", len(anagrams))
		t.Fail()
	}
}

// TODO: Test to make sure all anagrams are unique.
