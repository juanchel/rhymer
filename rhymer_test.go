package rhymer_test

import (
    "testing"
    "os"
    "github.com/juanchel/rhymer"
)

var r = rhymer.NewRhymer()

var rhymeTests = []struct {
    a string     // first input
    b string     // second input
    expected int // expected result
} {
    {"cat", "cat", 1},
    {"cat", "bat", 1},
    {"cat", "acrobat", 1},
    {"over", "clover", 1},
    {"master", "raster", 1},
    {"cat", "dog", 0},
    {"over", "ever", 0},
    {"ever", "clover", 0},
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}

func TestRhymes(m *testing.T) {
    for _, i := range rhymeTests {
        actual := r.Rhymes(i.a, i.b)
        if actual != i.expected {
            m.Errorf("Rhymes(%s, %s): expected %d but got %d", i.a, i.b, i.expected, actual)
        }
    }
}