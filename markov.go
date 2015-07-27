// Generates markov chains
// Based off of https://golang.org/doc/codewalk/markov.go

package rhymer

const prefixLen int = 2

type Prefix [prefixLen]string

func (p Prefix) Shift(word string) {
    copy(p, p[1:])
    p[prefixLen-1] = word
}

type Chain map[string][]string

