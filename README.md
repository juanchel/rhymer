# rhymer [![Build Status](https://travis-ci.org/juanchel/rhymer.svg?branch=master)](https://travis-ci.org/juanchel/rhymer)
A very simple rhyming dictionary written in golang, powered by CMU's pronouncing dictionary.

    
```go
import (
  "rhymer"
)
    
func main() {
  r := rhymer.Rhymer()
  
  // returns [[AE K S EH P T][AH K S EH P T]]
  // See the cmudict README for more information
  r.Pronounce("accept")
  
  // Returns 1 because cat and bat rhyme
  r.Rhymes("cat", "bat") 
  
  // Returns 0 because cat and dog don't rhyme
  r.Rhymes("cat", "dog")
  
  // Returns -1 because golang isn't in cmudict :(
  r.Rhymes("cat", "golang")
  
  // Returns an array of strings that rhyme with "cat"
  r.FindRhymesByWord("cat")
  
  // Return an array of strings that rhyme with the second listed pronounciation of accept
  r.FindRhymes(r.Pronounce("accept")[1])
  
  // Returns an array of strings that rhyme with the "acker" sound
  r.FindRhymes([]string{"AA", "K", "ER"})
}
```
