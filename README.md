# rhymer [![Build Status](https://travis-ci.org/juanchel/rhymer.svg?branch=master)](https://travis-ci.org/juanchel/rhymer)
A very simple rhyming dictionary written in golang, powered by CMU's pronouncing dictionary.
t
### Installation

    go get github.com/juanchel/rhymer
    
### Usage

Import the project and get a new rhymer

```go
import (
  "github.com/juanchel/rhymer"
)
    
func main() {
  r := rhymer.NewRhymer()
  // do stuff with r
}
```

You can recieve a slice of ways pronounce a word with `Pronounce`. Each member of the slice contains another slice of strings which contains the phonemes of each pronounciation.

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

You can also find all words that rhyme with a certain word with `FindRhymesByWord`, or with a phoneme slice `FindRhymes`.

```go
  // Returns a slice of strings that rhyme with "cat"
  r.FindRhymesByWord("cat")

  // Returns an array of strings that rhyme with the "acker" sound
  r.FindRhymes([]string{"AA", "K", "ER"})
}
```
