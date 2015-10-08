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
    {"cat",    "cat",     1},
    {"cat",    "bat",     1},
    {"cat",    "acrobat", 1},
    {"over",   "clover",  1},
    {"master", "raster",  1},
    {"masTER", "RaStEr",  1},
    {"aunt",   "rant",    1},
    {"aunt",   "want",    1},
    {"rant",   "want",    0},
    {"cat",    "dog",     0},
    {"over",   "ever",    0},
    {"ever",   "clover",  0},
    {"kanye",  "cat",     -1},
    {"kanye",  "yeezy",   -1},
    {"",       "cat",     -1},
    {"",       "",        -1},
    {"^cat",   "&bat",    -1},
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

func TestPronounceSimple(m *testing.T) {
    expected := [3]string{"K", "AE", "T"}
    actual := r.Pronounce("cat")
    if len(actual) != 1 {
        m.Errorf("Pronounce(cat) returned the wrong number of results: %d", len(actual))
    } else if len(actual[0]) != 3 {
        m.Errorf("Pronounce(cat) returned the wrong number of phonemes: %d")
    } else {
        for i, v := range actual[0] {
            if v != expected[i] {
                m.Errorf("Pronounce(cat) returned the wrong phoneme at index %d: got %v, expected %v", i, actual[0], expected)
            }
        }
    }
}

func TestPronounceMultiple(m *testing.T) {
    expectedA := [3]string{"AE", "N", "T"}
    expectedB := [3]string{"AO", "N", "T"}
    actual := r.Pronounce("aunt")
    if len(actual) != 2 {
        m.Errorf("Pronounce(aunt) returned the wrong number of results: %d", len(actual))
    } else if len(actual[0]) != 3 || len(actual[1]) != 3 {
        m.Errorf("Pronounce(aunt) returned the wrong number of phonemes")
    } else {
        for i := range actual[0] {
            if actual[0][i] != expectedA[i] && actual[0][i] != expectedB[i] {
                m.Errorf("Pronounce(aunt) returned the wrong phoneme: got %v, expected [%v %v]", actual, expectedA, expectedB)
            }
            if actual[1][i] != expectedA[i] && actual[1][i] != expectedB[i] {
                m.Errorf("Pronounce(aunt) returned the wrong phoneme: got %v, expected [%v %v]", actual, expectedA, expectedB)
            }
        }
    }
}

func TestPronounceNotFound(m *testing.T) {
    actual := r.Pronounce("naenae")
    if len(actual) != 0 {
        m.Errorf("Pronounce(naenae) should have returned nothing but returned: %v", actual)
    }
    actual = r.Pronounce("!@#$^&")
    if len(actual) != 0 {
        m.Errorf("Pronounce(!@#$^&) should have returned nothing but returned: %v", actual)
    }
}

func TestFindRhymes(m *testing.T) {
    actualWord := r.FindRhymesByWord("crunk")
    actualPhon := r.FindRhymes([]string{"AH", "NG", "K"})
    wordSet := make(map[string]bool)

    for _, v := range actualWord {
        wordSet[v] = true
    }
    for _, v := range actualPhon {
        if !wordSet[v] {
            m.Errorf("Mismatch in FindRhymesByWord(crunk) and FindRhymes([AH NG K])")
        }
    }

    if len(actualWord) != 54 {
        m.Errorf("FindRhymesByWord(crunk) returned %d results, expected 54", len(actualWord))
    }
    if len(actualPhon) != 54 {
        m.Errorf("FindRhymes([AH NG K]) returned %d results, expected 54", len(actualPhon))
    }
}

func TestFindRhymesNotFound(m *testing.T) {
    actualWord := r.FindRhymesByWord("abcd")
    actualPhon := r.FindRhymes([]string{"T", "K", "O"})

    if len(actualWord) != 0 {
        m.Errorf("FindRhymesByWord(abcd) returned %d results, expected 0", len(actualWord))
    }
    if len(actualPhon) != 0 {
        m.Errorf("FindRhymes([AB CD]) returned %d results, expected 0", len(actualPhon))
    }
}