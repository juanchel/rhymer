package rhymer

import (
    "os"
    "bufio"
    "strings"
    "runtime"
    "path"
)

type rhymer struct {
    dictionary map[string][][]string
}

func (r *rhymer) Pronounce(s string) [][]string {
    return r.dictionary[strings.ToUpper(s)]
}

func (r *rhymer) Rhymes(s1, s2 string) int {
    s1 = strings.ToUpper(s1)
    s2 = strings.ToUpper(s2)
    if len(r.Pronounce(s1)) == 0 || len(r.Pronounce(s2)) == 0 {
        return -1
    }

    rhymes := 0

    for _, v := range r.Pronounce(s1) {
        for _, w := range r.Pronounce(s2) {
            if pronounceRhymes(v, w) {
                rhymes = 1
            }
        }
    }
    return rhymes
}

func vowelSound(s string) bool {
    switch s[0] {
    case 'A', 'E', 'I', 'O', 'U':
        return true
    default:
        return false
    }
}

func pronounceRhymes(a1, a2 []string) bool {
    // Find the shorter word
    var longer []string
    var shorter []string
    if len(a1) > len(a2) {
        longer = a1
        shorter = a2
    } else {
        longer = a2
        shorter = a1
    }

    diff := len(longer) - len(shorter)
    full := true

    offset := 1

    // We need to match the entire shorter word if it starts with a vowel
    if vowelSound(shorter[0]) {
        offset = 0
    }

    // Check if the words sound the same, ignoring the first constanant sound of the shorter word
    for i, v := range shorter[offset:] {
        if longer[diff+i+offset] != v {
            full = false
        }
    }

    return full
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func NewRhymer() *rhymer {
    r := new(rhymer)

    // Read the file
    _, filename, _, _ := runtime.Caller(0)
    file, err := os.Open(path.Join(path.Dir(filename), "cmudict", "cmudict"))
    check(err)
    scanner := bufio.NewScanner(file)
    defer file.Close()
    check(err)

    r.dictionary = make(map[string][][]string)

    // Scan the file line by line
    for scanner.Scan() {
        // Split the line by whitespace
        f := strings.Fields(scanner.Text())
        // Trim the stress numbers
        for i, v := range f[2:] {
            f[i+2] = strings.TrimRight(v, "012")
        }
        // f[0] is the string, and f[2:] is the pronounciation
        r.dictionary[f[0]] = append(r.dictionary[f[0]], f[2:])
    }

    return r
}