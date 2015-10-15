// Rhyming dictionary in go

package rhymer

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strings"
)

type phonTrie struct {
	leaves map[string]*phonTrie
	words  map[string]bool
}

type Rhymer struct {
	dictionary map[string][][]string
	trie       phonTrie
}

// Create a new Rhymer by reading the pronunciation dictionary.
func NewRhymer() *Rhymer {
	r := new(Rhymer)

	// Read the file.
	_, filename, _, _ := runtime.Caller(0)
	file, err := os.Open(path.Join(path.Dir(filename), "data", "reduxdict"))
	check(err)
	scanner := bufio.NewScanner(file)
	defer file.Close()
	check(err)

	r.dictionary = make(map[string][][]string)
	r.trie.leaves = make(map[string]*phonTrie)

	// Scan the file line by line.
	for scanner.Scan() {
		// Split the line by whitespace.
		f := strings.Fields(scanner.Text())

		// Build the dictionary for quick lookups.
		// f[0] is the string, and f[1:] is the pronunciation.
		r.dictionary[f[0]] = append(r.dictionary[f[0]], f[1:])

		// Reduce the word to the rhyming part for trie insertion.
		rhymeSound := Rhymereduce(f[1:])

		// Insert the word into the trie.
		cur := &r.trie
		for i := len(rhymeSound) - 1; i >= 0; i-- {
			if cur.leaves[rhymeSound[i]] == nil {
				cur.leaves[rhymeSound[i]] = &phonTrie{make(map[string]*phonTrie), make(map[string]bool)}
			}
			cur = cur.leaves[rhymeSound[i]]
			if i == 0 {
				cur.words[f[0]] = true
			}
		}
	}
	return r
}

// Cheks whether or now two slices of phonemes rhyme. Returns 1 if they do and 0 if they don't.
func RhymesFullPhonetic(a1, a2 []string) bool {
	// Find the word with less rhymable phonemes.
	var longer []string
	var shorter []string

	if vowelOffset(a1) == -1 || vowelOffset(a2) == -1 {
		return false
	}

	if len(a1)-vowelOffset(a1) > len(a2)-vowelOffset(a2) {
		longer = a1
		shorter = a2
	} else {
		longer = a2
		shorter = a1
	}

	return rhymeTo(longer, shorter)
}

// Reduce a slice of phonemes to the section that would be required to match when checking for rhymes.
func Rhymereduce(phon []string) []string {
	var res []string
	vowelFound := false
	for _, v := range phon {
		if len(v) == 0 {
			return res
		}
		if !vowelFound {
			switch v[0] {
			case 'A', 'E', 'I', 'O', 'U':
				vowelFound = true
				res = []string{v}
			}
		} else {
			res = append(res, v)
		}
	}
	return res
}

// Reduces a slice of phonemes to the very last rhymable chain.
func SyllabicReduce(phon []string) []string {
	if len(phon) == 0 {
		return []string{}
	}
	var res []string
	vowelFound := false
	for i := len(phon) - 1; i >= 0; i-- {
		if len(phon[i]) == 0 {
			return []string{}
		}
		switch phon[i][0] {
		case 'A', 'E', 'I', 'O', 'U':
			vowelFound = true
			res = append([]string{phon[i]}, res...)
		default:
			if vowelFound {
				return res
			}
			res = append([]string{phon[i]}, res...)
		}
	}
	if vowelFound {
		return res
	} else {
		return []string{}
	}
}

// Returns a slice of strings that contain all known words that rhyme with a phoneme slice s.
func (r *Rhymer) FindRhymes(s []string) []string {
	rhymeSound := Rhymereduce(s)
	cur := &r.trie
	for i := len(rhymeSound) - 1; i >= 0; i-- {
		if cur.leaves[rhymeSound[i]] != nil {
			cur = cur.leaves[rhymeSound[i]]
		} else {
			return []string{}
		}
	}
	return cur.getFullSet()
}

// Returns a slice that contains the various possible pronunciations of string s. Each slice contains a slice of strings that represent the phonemes that make up the pronunciation.
func (r *Rhymer) Pronounce(s string) [][]string {
	return r.dictionary[strings.ToUpper(s)]
}

// Returns a slice of strings that contain all known words that rhyme with any pronunciation of string s.
func (r *Rhymer) FindRhymesByWord(s string) []string {
	s = strings.ToUpper(s)
	if _, ok := r.dictionary[s]; !ok {
		return []string{}
	}
	return r.FindRhymes(r.dictionary[s][0])
}

// Checks whether or not string s rhymes with a phoneme slice p1. Returns 1 if they do, 0 if they don't, and -1 if one or more of the words are unknown.
func (r *Rhymer) RhymesPhonetic(s string, p1 []string) int {
	s = strings.ToUpper(s)
	p2 := r.Pronounce(s)

	if len(p2) == 0 {
		return -1
	}

	for _, v := range p2 {
		if RhymesFullPhonetic(v, p1) {
			return 1
		}
	}
	return 0
}

// Checks whether or not strings s1 and s2 rhyme. Returns 1 if they do, 0 if they don't, and -1 if one or more of the words are unknown.
func (r *Rhymer) Rhymes(s1, s2 string) int {
	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)

	p1 := r.Pronounce(s1)
	p2 := r.Pronounce(s2)

	// Return -1 if one of the words is unknown.
	if len(p1) == 0 || len(p2) == 0 {
		return -1
	}

	// Return 1 if any of the prounounciations rhyme.
	for _, v := range p1 {
		for _, w := range p2 {
			if RhymesFullPhonetic(v, w) {
				return 1
			}
		}
	}
	return 0
}

// Returns a slice of all the words at the current level of a slice and recurses on its children.
func (p *phonTrie) getFullSet() []string {
	res := []string{}
	for k, _ := range p.words {
		res = append(res, k)
	}
	for _, v := range p.leaves {
		res = append(res, v.getFullSet()...)
	}
	return res
}

// Finds how many phonemes the vowel is offset so we know where to start rhyming. Returns -1 if there are none.
func vowelOffset(s []string) int {
	for i, v := range s {
		if len(v) == 0 {
			return -1
		}
		switch v[0] {
		case 'A', 'E', 'I', 'O', 'U':
			return i
		}
	}
	return -1
}

func rhymeTo(l, s []string) bool {
	diff := len(l) - len(s)
	ret := true

	offset := vowelOffset(s)
	if offset == -1 {
		return false
	}

	// Check if the words sound the same, ignoring the first constanant sounds of the shorter word.
	for i, v := range s[offset:] {
		if l[diff+i+offset] != v {
			ret = false
		}
	}

	return ret
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
