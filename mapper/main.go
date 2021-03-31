/*
	1 . Write a function with this signature:

	func CapitalizeEveryThirdAlphanumericChar(s string) string {
		// your code goes here
	}

	that capitalizes *only* every third alphanumeric character of a string, so that if I hand you

	Aspiration.com

	You hand me back

	asPirAtiOn.cOm

	Please note:

	- Characters other than each third should be downcased
	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9].

	2. Do the same problem, but this time create a "mapper" package that looks like this:
*/

package mapper

import (
	"unicode"
)

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// your code goes here
	var index int
	for i, v := range s {
		if unicode.IsLetter(v) {
			index++
		} else {
			continue
		}
		if index%3 == 0 {
			if unicode.IsLower(v) {
				s = replaceAtIndex(s, unicode.ToUpper(v), i)
			}
		} else {
			if unicode.IsUpper(v) {
				s = replaceAtIndex(s, unicode.ToLower(v), i)
			}
		}
	}
	return s

}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos, _ := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

/*
  Make sure your fmt.Println(s) output looks nice and clean, ie implement the Stringer interface.
  Zip up your file(s) and email them to the recruiter
*/

type StringTransformer struct {
	// StepCount counter to replace values at
	StepCount int
	// Data is the string on which actions are to be performed on
	Data string
	// Counter is the real incremental counter for replacing index after stepcount
	counter int
}

// NewStringTransformer returns a new StringTransformer instance
func NewStringTransformer(stepCount int, data string) StringTransformer {
	return StringTransformer{StepCount: stepCount, Data: data}
}

// TransformRune transforms rune at position
func (s *StringTransformer) TransformRune(pos int) {
	v := rune(s.Data[pos])
	if unicode.IsLetter(v) {
		s.counter++
	} else {
		return
	}
	if s.counter%s.StepCount == 0 {
		if unicode.IsLower(v) {
			s.Data = replaceAtIndex(s.Data, unicode.ToUpper(v), pos)
		}
	} else {
		if unicode.IsUpper(v) {
			s.Data = replaceAtIndex(s.Data, unicode.ToLower(v), pos)
		}
	}
}

//GetValueAsRuneSlice returns the value of rune slice
func (s *StringTransformer) GetValueAsRuneSlice() []rune {
	return []rune(s.Data)
}

func (s StringTransformer) String() string {
	return s.Data
}

func NewSkipString(skipper int, s string) StringTransformer {
	return NewStringTransformer(skipper, s)
}
