package tokenizer

import (
	"fmt"
	"log"
	"unicode"
)

type Tokenizer struct {
	tokens []Token
	read rune
	lookahead rune
	finished bool
	text string
	position int
}

func ComposesWords(ch rune) bool {
	if unicode.IsLetter(ch) {
		return true
	}

	if ch == '_' {
		return true
	}

	if unicode.IsDigit(ch) {
		return true
	}

	return false
}

func (tokenizer *Tokenizer) StoreToken(start int, end int, kind TokenKind) {
	tokenizer.tokens = append(tokenizer.tokens, Token {
		Start: start,
		End: end,
		Kind: kind,
		Piece: tokenizer.text[start:end],
	})
}

func (tokenizer *Tokenizer) Read() (ch rune) {
	if tokenizer.read != 0 {
		tokenizer.position += 1
	} 

	if tokenizer.position >= len(tokenizer.text) - 1 {
		tokenizer.finished = true
		return 0
	}

	tokenizer.read = rune(tokenizer.text[tokenizer.position])
	if tokenizer.position < len(tokenizer.text) {
		tokenizer.lookahead = rune(tokenizer.text[tokenizer.position + 1])
	} else {
		tokenizer.lookahead = 0
	}

	tokenizer.finished = false
	return tokenizer.read
}

func (tokenizer *Tokenizer) Tokenize(script string) {
	tokenizer.read = 0
	tokenizer.lookahead = 0

	tokenizer.text = script

	log.Printf("reading first token\n")
	for !tokenizer.finished {
		// Single characters
		if kind, ok := token_runes[tokenizer.read]; ok {
			tokenizer.StoreToken(tokenizer.position, tokenizer.position + 1, kind)
		}

		// Numbers
		if unicode.IsDigit(tokenizer.read) {
			token_start := tokenizer.position

			for !tokenizer.finished {
				if !unicode.IsDigit(tokenizer.read) && tokenizer.read != '.' {
					break
				}

				if tokenizer.read == '.' && !unicode.IsDigit(tokenizer.lookahead) {
					panic("failed trying to tokenize number")
				}

				tokenizer.Read()
			}

			tokenizer.StoreToken(token_start, tokenizer.position, Literal_Number)
			continue
		}

		// Identifiers and keywords
		if ComposesWords(tokenizer.read) {
			token_start := tokenizer.position

			for !tokenizer.finished {
				if !ComposesWords(tokenizer.lookahead) && !unicode.IsDigit(tokenizer.lookahead) {
					token_kind := Identifier
					switch tokenizer.text[token_start:tokenizer.position + 1] {
					case "use":
						token_kind = Keyword_use
					case "f32":
						token_kind = Type_f32
					case "f64":
						token_kind = Type_f64
					case "i8":
						token_kind = Type_i8
					case "i16":
						token_kind = Type_i16
					case "i32":
						token_kind = Type_i32
					case "i64":
						token_kind = Type_i64
					case "u8":
						token_kind = Type_u8
					case "u16":
						token_kind = Type_u16
					case "u32":
						token_kind = Type_u32
					case "u64":
						token_kind = Type_u64
					case "if":
						token_kind = Keyword_if
					case "repeat":
						token_kind = Keyword_repeat
					case "break":
						token_kind = Keyword_break
					}

					tokenizer.StoreToken(token_start, tokenizer.position + 1, token_kind)
					break
				}

				tokenizer.Read()
			}
		}

		// String lexical analysis
		if tokenizer.read == '"' {
			token_start := tokenizer.position

			for !tokenizer.finished {
				tokenizer.Read()

				if tokenizer.read == '"' {
					break
				}
			}

			tokenizer.StoreToken(token_start, tokenizer.position + 1, Literal_String)
		}

		tokenizer.Read()
	}

	log.Printf("tokens read: ")

	for _, t := range tokenizer.tokens {
		fmt.Printf("type:%v piece:%v\n", t.Kind, t.Piece)
	}
}
