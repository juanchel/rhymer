# rhymer [![Build Status](https://travis-ci.org/juanchel/rhymer.svg?branch=master)](https://travis-ci.org/juanchel/rhymer)
A simple pronunciation and rhyming library written in golang, powered by a streamlined version of CMU's pronoucniation dictionary.

### Installation

    go get github.com/juanchel/rhymer
    
### Usage

You can view the godocs [here](https://godoc.org/github.com/juanchel/rhymer).

Import the project and get a new rhymer.

```go
import (
  "github.com/juanchel/rhymer"
)
    
func main() {
  r := rhymer.New()
  // do stuff with r
}
```

You can recieve a slice of ways pronounce a word with `Pronounce`. Each member of the slice contains another slice of strings which contains the phonemes of each pronunciation.

```go
// Returns [[AE N T][AO N T]]
r.Pronounce("aunt")
```

You can check if two words rhyme with `Rhymes`. It will return 1 if they do, 0 if they don't, and -1 when one of the words is unknown.
  
```go
// Returns 1
r.Rhymes("bat", "cat") 

// Returns 0
r.Rhymes("cat", "dog")

// Returns -1
r.Rhymes("dog", "tog")
```

Even if there are words that aren't in the dictionary, if you can sound out their phonemes, you can still check if their sounds rhyme.

```go  
// Returns 1 because dog rhymes with the "tog" sound
r.RhymesToPhonetic("dog", {}string["T", "AW", "G"])

// Returns 1 because the "dog" sound rhymes with the "tog" sound
r.RhymesFullPhonetic("{}string["D", "AW", "G"], {}string["T", "AW", "G"])
```

You can also find all words that rhyme with a certain word with `FindRhymesByWord`, or with a phoneme slice `FindRhymes`. `FindRhymesByWord` will return an empty slice if the word is not recognized.

```go
// Returns a slice of strings that rhyme with "cat"
r.FindRhymesByWord("cat")

// Returns a slice of strings that rhyme with the "acker" sound
r.FindRhymes([]string{"AA", "K", "ER"})
```

The library also provides a few simple methods for manipulating slices of phonemes. These are called from the rhymer package, not the object:
 - If you want to reduce a slice down to the rhyming portion, you can call `RhymerReduce`, which would reduce the phonemes of "wombat" to "ombat".
 - If you want to reduce a slice down to the very last rhymable syllable, you can call `SyllabicReduce`, which would reduce the phoneme of "wombat" to "at".

```go
// Returns [AA M B AE T]
rhymer.RhymerReduce(r.Pronounce("wombat")[0])

// Returns [AE T]
rhymer.RhymerReduce(r.Pronounce("wombat")[0])
```

### Usage Notes

Note that `Rhymes` is not necessarily transitive because some words can be pronounced in different ways, e.g. "read" rhymes with "creed", and with "bed". This also means that the results of certain inputs to `FindRhymesByWord` will not all rhyme with each other. If you only want to use for a certain pronunciation of a word, you can use the functions that take in a slice of phonemes. You can easily access these slices through `Pronounce`.

### License

This software is released under the MIT License.
