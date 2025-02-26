package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Contains("seafood", "")) // true
	fmt.Println(strings.Contains("", ""))        // true

	fmt.Println(strings.ContainsAny("foo", "")) // false
	fmt.Println(strings.ContainsAny("", ""))    // false

	f := func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	fmt.Println(strings.ContainsFunc("hello", f)) // true

	fmt.Println(strings.ContainsRune("timeout", 97)) // false : 97->'a'

	fmt.Println(strings.Count("cheese", "e")) // 3

	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "ph")
	show("Gopher", "oa")

	show = func(s, sep string) {
		after, found := strings.CutPrefix(s, sep)
		fmt.Printf("CutPrefix(%q, %q) = %q, %v\n", s, sep, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")

	show = func(s, sep string) {
		before, found := strings.CutSuffix(s, sep)
		fmt.Printf("CutSuffix(%q, %q) = %q, %v\n", s, sep, before, found)
	}
	show("Gopher", "Go")
	show("Gopher", "er")

	// EqualFold case-insensitive
	fmt.Println(strings.EqualFold("Go", "gos")) // false
	fmt.Println(strings.EqualFold("AB", "aB"))  // true

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	fmt.Printf("Fields are: %q\n", strings.Fields("      "))

	f = func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,   ,baz3...", f))
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,   ,baz3...", f))

	fmt.Println(strings.HasPrefix("Gopher", "G")) // true
	fmt.Println(strings.HasPrefix("Gopher", ""))  // true

	fmt.Println(strings.HasSuffix("Amigo", "ig")) // false
	fmt.Println(strings.HasSuffix("Amigo", ""))   // true

	fmt.Println(strings.Index("chicken", "ck"))  // 3
	fmt.Println(strings.Index("chicken", "dmr")) // -1
	fmt.Println(strings.Index("chicken", ""))    // 0

	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1

	fmt.Println(strings.IndexByte("golaang", 'a')) // 3
	fmt.Println(strings.IndexByte("golang", '^'))  // -1

	f = func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("abcd世界", f)) // 4
	fmt.Println(strings.IndexFunc("abc", f))    // -1

	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, "...")) // foo...bar...baz

	fmt.Println(strings.LastIndex("go gopher", "go")) // 3
	fmt.Println(strings.LastIndex("go gopher", "re")) // -1

	fmt.Println(strings.LastIndexAny("go gopher", "go")) // 4
	fmt.Println(strings.LastIndexAny("go gopher", "re")) // 8
	fmt.Println(strings.LastIndexAny("go gopher", "x"))  // -1

	fmt.Println(strings.LastIndexByte("Hello, world", 'l')) // 10
	fmt.Println(strings.LastIndexByte("Hello, world", 'x')) // -1

	fmt.Println(strings.IndexFunc("go 123", unicode.IsNumber))     // 3 : IndexFunc
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber)) // 2 : LastIndexFunc
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))     // -1

	rot1 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+1)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+1)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot1, "'Twas brillig and the slithy gopher..."))

	//It panics if count is negative or if the result of (len(s) * count) overflows.
	fmt.Println("ba" + strings.Repeat("na", 3)) // bananana

	//Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	fmt.Println(strings.Replace("oink oink oink", "k", "rrr", 2))
	fmt.Println(strings.Replace("oink oink oink", "", "rrr", 2)) // f old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a,b,c", ".."))                       // ["a,b,c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "abc"))                           // [""]
	fmt.Printf("%q\n", strings.Split("", ""))                              // []

	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]

	//n > 0: at most n substrings; the last substring will be the unsplit remainder;
	//n == 0: the result is nil (zero substrings);
	//n < 0: all substrings.
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", ",", 4))    // ["a," "b," "c," "d,e,f"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", ",", 0))    // []
	fmt.Printf("%q\n", strings.SplitAfterN(",a,b,c,d,e,f,", ",", -1)) // ["," "a," "b," "c," "d," "e," "f", ""]

	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)       // ["a" "b,c"]
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)

	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Örnek İş"))
	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

	//Note "ǳ" here is a single character, not a "d" followed by a "z": ǳ vs dz
	str := "ǳ"
	fmt.Println(strings.ToTitle(str))
	fmt.Println(strings.ToUpper(str))

	fmt.Println(strings.ToUpper("Gopher"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))

	// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences replaced by the replacement string, which may be empty.
	fmt.Printf("%s\n", strings.ToValidUTF8("abc", "\uFFFD"))
	fmt.Printf("%s\n", strings.ToValidUTF8("a\xffb\xC0\xAFc\xff", ""))
	fmt.Printf("%s\n", strings.ToValidUTF8("a\xffb\xC0\xAFc\xff", "z"))
	fmt.Printf("%s\n", strings.ToValidUTF8("\xed\xa0\x80\xe0", "abc"))

	fmt.Println(strings.Trim("¡¡¡Hello, !!Gophers!!!", "!¡"))
	fmt.Println(strings.TrimFunc("¡¡¡Hello, !!Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	fmt.Println(strings.TrimPrefix("¡¡¡Hello, Gophers!!!", "¡¡¡Hello, "))
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println(strings.TrimSuffix("¡¡¡Hello, Gophers!!!", "!!!"))
}
