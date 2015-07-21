package rhymer

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

type Pronounce struct {
    phonemes []string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Read the file
    file, err := os.Open("cmudict/cmudict")
    check(err)
    scanner := bufio.NewScanner(file)
    defer file.Close()
    check(err)

    dictionary := make(map[string][]Pronounce)

    // Scan the file line by line
    for scanner.Scan() {
        // Split the line by whitespace
        f := strings.Fields(scanner.Text())
        // f[0] is the string, and f[2:] is the pronounciation

        dictionary[f[0]] = []Pronounce{Pronounce{f[2:]}}
    }

    // fmt.Println(dictionary)
    fmt.Println(dictionary["HELLO"][0].phonemes[2])
}