package rhymer

import (
    "testing"
)

func TestRhymer(t *testing.T) {
    r := Rhymer()
    r.Rhymes("cat", "bat")
}