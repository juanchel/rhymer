# rhymer

Simple rhyming dictionary written in golang, powered by CMU's pronouncing dictionary.

    import (
      "rhymer"
    )
    
    func main() {
      r = rhymer.Rhymer()
      
      // Returns an array of arrays of strings that represents the phonemes of possible pronounciations of "accept"
      // For "accept", it will returns [[AE K S EH P T][AH K S EH P T]]
      // See the README for the cmudict for more information about the phonemes
      r.Pronounce("accept")
      
      // Returns 1 when the words rhyme, 0 when they dont, and -1 when the rhymer doesn't know one of the words
      r.Rhymes("cat", "bat") 
      
      // Returns an array of strings that rhyme with "cat"
      // It uses an arbitrary pronounciation of the word when there are multiple
      r.FindRhymesByWord("cat")
      
      // Return an array that rhymes with the second listed pronounciation of accept
      r.FindRhymes(r.Pronounce("accept")[1])
      
      // Returns an array that rhymes with the "acker" sound
      r.FindRhymes([]string{"AA", "K", "ER"})
    }
