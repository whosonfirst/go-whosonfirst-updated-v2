package chicken

import (
	"errors"
	"fmt"
	"github.com/cooperhewitt/go-ucd"
	emoji "github.com/thisisaaronland/go-chicken/emoji"
	"github.com/thisisaaronland/go-chicken/strings"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const emoji_pattern = "[\\x{2712}\\x{2714}\\x{2716}\\x{271d}\\x{2721}\\x{2728}\\x{2733}\\x{2734}\\x{2744}\\x{2747}\\x{274c}\\x{274e}\\x{2753}-\\x{2755}\\x{2757}\\x{2763}\\x{2764}\\x{2795}-\\x{2797}\\x{27a1}\\x{27b0}\\x{27bf}\\x{2934}\\x{2935}\\x{2b05}-\\x{2b07}\\x{2b1b}\\x{2b1c}\\x{2b50}\\x{2b55}\\x{3030}\\x{303d}\\x{1f004}\\x{1f0cf}\\x{1f170}\\x{1f171}\\x{1f17e}\\x{1f17f}\\x{1f18e}\\x{1f191}-\\x{1f19a}\\x{1f201}\\x{1f202}\\x{1f21a}\\x{1f22f}\\x{1f232}-\\x{1f23a}\\x{1f250}\\x{1f251}\\x{1f300}-\\x{1f321}\\x{1f324}-\\x{1f393}\\x{1f396}\\x{1f397}\\x{1f399}-\\x{1f39b}\\x{1f39e}-\\x{1f3f0}\\x{1f3f3}-\\x{1f3f5}\\x{1f3f7}-\\x{1f4fd}\\x{1f4ff}-\\x{1f53d}\\x{1f549}-\\x{1f54e}\\x{1f550}-\\x{1f567}\\x{1f56f}\\x{1f570}\\x{1f573}-\\x{1f579}\\x{1f587}\\x{1f58a}-\\x{1f58d}\\x{1f590}\\x{1f595}\\x{1f596}\\x{1f5a5}\\x{1f5a8}\\x{1f5b1}\\x{1f5b2}\\x{1f5bc}\\x{1f5c2}-\\x{1f5c4}\\x{1f5d1}-\\x{1f5d3}\\x{1f5dc}-\\x{1f5de}\\x{1f5e1}\\x{1f5e3}\\x{1f5ef}\\x{1f5f3}\\x{1f5fa}-\\x{1f64f}\\x{1f680}-\\x{1f6c5}\\x{1f6cb}-\\x{1f6d0}\\x{1f6e0}-\\x{1f6e5}\\x{1f6e9}\\x{1f6eb}\\x{1f6ec}\\x{1f6f0}\\x{1f6f3}\\x{1f910}-\\x{1f918}\\x{1f980}-\\x{1f984}\\x{1f9c0}\\x{3297}\\x{3299}\\x{a9}\\x{ae}\\x{203c}\\x{2049}\\x{2122}\\x{2139}\\x{2194}-\\x{2199}\\x{21a9}\\x{21aa}\\x{231a}\\x{231b}\\x{2328}\\x{2388}\\x{23cf}\\x{23e9}-\\x{23f3}\\x{23f8}-\\x{23fa}\\x{24c2}\\x{25aa}\\x{25ab}\\x{25b6}\\x{25c0}\\x{25fb}-\\x{25fe}\\x{2600}-\\x{2604}\\x{260e}\\x{2611}\\x{2614}\\x{2615}\\x{2618}\\x{261d}\\x{2620}\\x{2622}\\x{2623}\\x{2626}\\x{262a}\\x{262e}\\x{262f}\\x{2638}-\\x{263a}\\x{2648}-\\x{2653}\\x{2660}\\x{2663}\\x{2665}\\x{2666}\\x{2668}\\x{267b}\\x{267f}\\x{2692}-\\x{2694}\\x{2696}\\x{2697}\\x{2699}\\x{269b}\\x{269c}\\x{26a0}\\x{26a1}\\x{26aa}\\x{26ab}\\x{26b0}\\x{26b1}\\x{26bd}\\x{26be}\\x{26c4}\\x{26c5}\\x{26c8}\\x{26ce}\\x{26cf}\\x{26d1}\\x{26d3}\\x{26d4}\\x{26e9}\\x{26ea}\\x{26f0}-\\x{26f5}\\x{26f7}-\\x{26fa}\\x{26fd}\\x{2702}\\x{2705}\\x{2708}-\\x{270d}\\x{270f}]|\\x{23}\\x{20e3}|\\x{2a}\\x{20e3}|\\x{30}\\x{20e3}|\\x{31}\\x{20e3}|\\x{32}\\x{20e3}|\\x{33}\\x{20e3}|\\x{34}\\x{20e3}|\\x{35}\\x{20e3}|\\x{36}\\x{20e3}|\\x{37}\\x{20e3}|\\x{38}\\x{20e3}|\\x{39}\\x{20e3}|\\x{1f1e6}[\\x{1f1e8}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f2}\\x{1f1f4}\\x{1f1f6}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fd}\\x{1f1ff}]|\\x{1f1e7}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ef}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1e8}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ee}\\x{1f1f0}-\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}-\\x{1f1ff}]|\\x{1f1e9}[\\x{1f1ea}\\x{1f1ec}\\x{1f1ef}\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1ff}]|\\x{1f1ea}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ed}\\x{1f1f7}-\\x{1f1fa}]|\\x{1f1eb}[\\x{1f1ee}-\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1f7}]|\\x{1f1ec}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ee}\\x{1f1f1}-\\x{1f1f3}\\x{1f1f5}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1ed}[\\x{1f1f0}\\x{1f1f2}\\x{1f1f3}\\x{1f1f7}\\x{1f1f9}\\x{1f1fa}]|\\x{1f1ee}[\\x{1f1e8}-\\x{1f1ea}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}]|\\x{1f1ef}[\\x{1f1ea}\\x{1f1f2}\\x{1f1f4}\\x{1f1f5}]|\\x{1f1f0}[\\x{1f1ea}\\x{1f1ec}-\\x{1f1ee}\\x{1f1f2}\\x{1f1f3}\\x{1f1f5}\\x{1f1f7}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1f1}[\\x{1f1e6}-\\x{1f1e8}\\x{1f1ee}\\x{1f1f0}\\x{1f1f7}-\\x{1f1fb}\\x{1f1fe}]|\\x{1f1f2}[\\x{1f1e6}\\x{1f1e8}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1ff}]|\\x{1f1f3}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f4}\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}\\x{1f1ff}]|\\x{1f1f4}\\x{1f1f2}|\\x{1f1f5}[\\x{1f1e6}\\x{1f1ea}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1f3}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1f6}\\x{1f1e6}|\\x{1f1f7}[\\x{1f1ea}\\x{1f1f4}\\x{1f1f8}\\x{1f1fa}\\x{1f1fc}]|\\x{1f1f8}[\\x{1f1e6}-\\x{1f1ea}\\x{1f1ec}-\\x{1f1f4}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fd}-\\x{1f1ff}]|\\x{1f1f9}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ed}\\x{1f1ef}-\\x{1f1f4}\\x{1f1f7}\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1ff}]|\\x{1f1fa}[\\x{1f1e6}\\x{1f1ec}\\x{1f1f2}\\x{1f1f8}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1fb}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ee}\\x{1f1f3}\\x{1f1fa}]|\\x{1f1fc}[\\x{1f1eb}\\x{1f1f8}]|\\x{1f1fd}\\x{1f1f0}|\\x{1f1fe}[\\x{1f1ea}\\x{1f1f9}]|\\x{1f1ff}[\\x{1f1e6}\\x{1f1f2}\\x{1f1fc}]|[\\x{1f910}-\\x{1f9c0}]"

var emoji_re = regexp.MustCompile(emoji_pattern)

func GetChickenForLanguageTag(tag string, clucking bool) (*Chicken, error) {

	// see also: https://github.com/thisisaaronland/go-chicken/issues/1

	c, ok := chicken.CHICKENS[tag]

	if !ok {
		return nil, errors.New("unknown (or untranslated) chicken")
	}

	debug := false

	ch := Chicken{tag, c, clucking, debug}
	return &ch, nil
}

type Chicken struct {
	tag      string // sudo make me a https://godoc.org/golang.org/x/text/language thingy
	chicken  string
	clucking bool
	debug    bool
}

func (ch *Chicken) String() string {

	if ch.clucking {
		return ch.Cluck()
	}

	return ch.Chicken()
}

func (ch *Chicken) Chicken() string {

	return ch.chicken
}

func (ch *Chicken) Cluck() string {

	tag := ch.tag
	clucks, ok := chicken.CLUCKING[tag]

	if !ok {
		return ch.Chicken()
	}

	count := len(clucks)
	idx := rand.Intn(count)

	return clucks[idx]
}

func (ch *Chicken) ExpandAlphaCodes(txt string) string {

	cb := func(in []byte) []byte {

		str_in := string(in)

		code, ok := emoji.EmojiAlphaCodes[str_in]

		if !ok {
			return in
		}

		num, _ := strconv.ParseInt(code, 0, 32)
		rune := fmt.Sprintf("%c", num)

		return []byte(rune)
	}

	re := regexp.MustCompile(`:(\w+):`)

	bytes_in := []byte(txt)
	bytes_out := re.ReplaceAllFunc(bytes_in, cb)

	return string(bytes_out)
}

func (ch *Chicken) TextToChicken(txt string) string {

	re := regexp.MustCompile(`:(\w+):`)

	if re.MatchString(txt) {

		if ch.debug {
			log.Printf("TEXT is a string")
		}

		if ch.debug {
			log.Printf("BEFORE %s\n", txt)
		}

		txt = ch.ExpandAlphaCodes(txt)

		if ch.debug {
			fmt.Printf("AFTER %s\n", txt)
		}
	}

	buf := make([]string, 0)
	word := false

	for _, char := range txt {

		// https://golang.org/pkg/unicode/#IsLetter
		// fmt.Println(idx, char, unicode.IsLetter(char), unicode.IsSpace(char))

		if unicode.IsSpace(char) {

			if ch.debug {
				log.Printf("CHAR '%d' is a space\n", char)
			}

			buf = append(buf, string(char))
			word = false

		} else if unicode.IsSymbol(char) {

			if ch.debug {
				log.Printf("CHAR '%d' is a symbol\n", char)
			}

			if ch.debug {
				log.Printf("BEFORE UCD '%d'\n", char)
			}

			n := ucd.Name(string(char))

			if ch.debug {
				log.Printf("AFTER UCD '%s'\n", n.Name)
			}

			b := ch.TextToChicken(n.Name)
			buf = append(buf, b)

		} else if unicode.IsLetter(char) {

			if ch.debug {
				log.Printf("CHAR '%d' is a letter\n", char)
			}

			if !word {
				buf = append(buf, ch.String())
				word = true
			}

			// For example like when Unicode 9.0 was released before Go itself was
			// updated (20160623/thisisaaronland)

		} else if ch.IsCharEmoji(char) {

			if ch.debug {
				log.Printf("CHAR '%d' is emoji\n", char)
			}

			if ch.debug {
				log.Printf("BEFORE UCD '%d'\n", char)
			}

			n := ucd.Name(string(char))

			if ch.debug {
				log.Printf("AFTER UCD '%s'\n", n.Name)
			}

			b := ch.TextToChicken(n.Name)
			buf = append(buf, b)

		} else {

			if ch.debug {
				log.Printf("CHAR '%d' is a char\n", char)
			}

			buf = append(buf, string(char))
			word = false
		}
	}

	return strings.Join(buf, "")
}

func (ch *Chicken) IsCharEmoji(char rune) bool {

	bs := []byte(string(char))
	return emoji_re.Match(bs)
}
