package tokenizer

type TokenKind int

const (
	Identifier TokenKind = iota
	Literal_Number
	Literal_String
	Delimiter_Open_Parenthesis
	Delimiter_Close_Parenthesis
	Delimiter_Open_Chevrons
	Delimiter_Close_Chevrons
	Delimiter_Open_Brackets
	Delimiter_Close_Brackets
	Delimiter_Open_Braces
	Delimiter_Close_Braces
	Delimiter_Colon
	Delimiter_Comma
	Keyword_use
	Keyword_if
	Keyword_repeat
	Keyword_break
	Operator_Assignment
	Operator_Equals
	Operator_Bigger_Than
	Operator_Smaller_Than
	Operator_Access
	Operator_Add
	Operator_Subtract
	Type_u8
	Type_u16
	Type_u32
	Type_u64
	Type_i8
	Type_i16
	Type_i32
	Type_i64
	Type_f32
	Type_f64
	Type_void
)

var token_runes = map[rune]TokenKind {
	'.': Operator_Access,
	',': Delimiter_Comma,
	'(': Delimiter_Open_Parenthesis,
	')': Delimiter_Close_Parenthesis,
	'[': Delimiter_Open_Brackets,
	']': Delimiter_Close_Brackets,
	'<': Delimiter_Open_Chevrons,
	'>': Delimiter_Close_Chevrons,
	'{': Delimiter_Open_Braces,
	'}': Delimiter_Close_Braces,
	':': Delimiter_Colon,
	'=': Operator_Assignment,
	'+': Operator_Add,
	'-': Operator_Subtract,
}

type Token struct {
	Kind TokenKind
	Start int
	End int
	Piece string
}
