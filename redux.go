package rhymer

import (
    "os"
    "bufio"
    "strings"
    "runtime"
    "path"
)

func Redux() {
    r := new(rhymer)

    // Read the file
    _, filename, _, _ := runtime.Caller(0)
    file, err := os.Open(path.Join(path.Dir(filename), "data", "cmudict"))
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

    f, err := os.Create(path.Join(path.Dir(filename), "data", "reduxdict"))
    check(err)
    defer f.Close()

    for k, v := range r.dictionary {
        unique := make([][]string, 0)
        for _, pronounce := range v {
            skip := false
            for _, existing := range unique {
                if samePhonemes(pronounce, existing) {
                    skip = true
                }
            }
            if !skip {
                f.WriteString(k + " " + strings.Join(pronounce, " ") + "\n")
                unique = append(unique, pronounce)
            }
        }
    }

    f.Sync()
}

func samePhonemes(n, m []string) bool {
    if len(n) != len(m) {
        return false
    } else {
        for i, _ := range n {
            if n[i] != m[i] {
                return false
            }
        }
        return true
    }
}