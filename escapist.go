package escapist

const (
	charSingleQuote = '\''
	charDblQuote    = '"'
	charAmpersand   = '&'
	charLessThan    = '<'
	charGreaterThan = '>'
	charTick        = '`'
	charExclamation = '!'
	charAt          = '@'
	charDollar      = '$'
	charPercent     = '%'
	charLParen      = '('
	charRParen      = ')'
	charEquals      = '='
	charAddition    = '+'
	charLCurly      = '{'
	charRCurly      = '}'
	charLBracket    = '['
	charRBracket    = ']'

	codeSingleQuote = "&#39;"
	codeDblQuote    = "&#34;"
	codeAmpersand   = "&#38;"
	codeLessThan    = "&#60;"
	codeGreaterThan = "&#62;"
	codeTick        = "&#96;"
	codeExclamation = "&#33;"
	codeAt          = "&#64;"
	codeDollar      = "&#36;"
	codePercent     = "&#37;"
	codeLParen      = "&#40;"
	codeRParen      = "&#41;"
	codeEquals      = "&#61;"
	codeAddition    = "&#43;"
	codeLCurly      = "&#123;"
	codeRCurly      = "&#125;"
	codeLBracket    = "&#91;"
	codeRBracket    = "&#93;"
)

type getNextFunc func([]byte) int
type appendFunc func([]byte, byte) []byte

// Escape will perform a basic escape an inbound byteslice
// If a change occurs the returned slice will be new rather than mutating the original
// Else, byteslices which do not contain escapable characters will return directly with no allocations
// Note: A basic escape looks for ', ", <, >, and &
func Escape(in []byte) (out []byte) {
	return escape(in, true)
}

// EscapeAdv will perform an advanced escape an inbound byteslice
// If a change occurs the returned slice will be new rather than mutating the original
// Else, byteslices which do not contain escapable characters will return directly with no allocations
// Note: An advanced escape looks for ', ", <, >, &, `, !, @, $, %, (, ), =, +, \{, \}, \[, and \]
func EscapeAdv(in []byte) (out []byte) {
	return escape(in, false)
}

// escape is the primary func for this library.
// Takes an inbound byteslice and a basic boolean, basic determines if basic or advanced escaping will be used
func escape(in []byte, basic bool) (out []byte) {
	var (
		changed bool      // Changed boolean, if remains false - no replacements were made
		lm      int       // Last marker
		nc      int       // Next character
		cap     = len(in) // Capacity

		gfn getNextFunc
		afn appendFunc
	)

	if basic {
		// Escaping type is basic, use basic set of funcs
		gfn = getNextEscChar
		afn = appendEscapedBytes
	} else {
		// Escaping type is advanced, use advanced set of funcs
		gfn = getNextAdvEscChar
		afn = appendAdvEscapedBytes
	}

	// Seek for the next character to escape, iterate until no character is found
	for nc = gfn(in); nc > -1; nc = gfn(in[lm:]) {
		if !changed {
			// Think of a good way to guess final size to avoid unnecessary allocations
			// Right now the outbound slice is created to be 30% larger than the inbound slice
			out = make([]byte, 0, len(in)*130/100)
			changed = true
		}

		// Catch up our 'out' slice to have all the data from the last marker position UP TO (not including) our new escape character
		out = append(out, in[lm:lm+nc]...)

		// Append the bytes for the escaped version of our character
		out = afn(out, in[nc+lm])

		// Set last marker to be the sum of last marker, new char, and one (to reference next index)
		if lm = lm + nc + 1; lm >= cap {
			// Our new last marker is either at the end of our inbound slice, break
			break
		}
	}

	if !changed {
		// Nothing has changed, set out to in
		out = in
	}
	return
}

// getNextEscChar will return the index of the next escape character within an inbound slice
// Note: This matches for basic escaping
func getNextEscChar(in []byte) (i int) {
	var b byte
	for i, b = range in {
		switch b {
		case charSingleQuote, charDblQuote, charAmpersand, charLessThan, charGreaterThan:
			return i
		}
	}

	return -1
}

// appendEscapedBytes will append bytes representing the escaped character
// Note: This appends for basic escaping
func appendEscapedBytes(in []byte, b byte) []byte {
	switch b {
	case charSingleQuote:
		return append(in, codeSingleQuote...)
	case charDblQuote:
		return append(in, codeDblQuote...)
	case charAmpersand:
		return append(in, codeAmpersand...)
	case charLessThan:
		return append(in, codeLessThan...)
	case charGreaterThan:
		return append(in, codeGreaterThan...)

	default:
		panic("We received an unexpected character: '" + string(b) + "'")
	}
}

// getNextAdvEscChar will return the index of the next escape character within an inbound slice
// Note: This matches for advanced escaping
func getNextAdvEscChar(in []byte) (i int) {
	var b byte
	for i, b = range in {
		switch b {
		case charSingleQuote, charDblQuote, charAmpersand, charLessThan,
			charGreaterThan, charTick, charExclamation, charAt, charDollar,
			charPercent, charLParen, charRParen, charEquals, charAddition,
			charLCurly, charRCurly, charLBracket, charRBracket:
			return i
		}
	}

	return -1
}

// appendAdvEscapedBytes will append bytes representing the escaped character
// Note: This appends for advanced escaping
func appendAdvEscapedBytes(in []byte, b byte) []byte {
	switch b {
	case charSingleQuote:
		return append(in, codeSingleQuote...)
	case charDblQuote:
		return append(in, codeDblQuote...)
	case charAmpersand:
		return append(in, codeAmpersand...)
	case charLessThan:
		return append(in, codeLessThan...)
	case charGreaterThan:
		return append(in, codeGreaterThan...)
	case charTick:
		return append(in, codeTick...)
	case charExclamation:
		return append(in, codeExclamation...)
	case charAt:
		return append(in, codeAt...)
	case charDollar:
		return append(in, codeDollar...)
	case charPercent:
		return append(in, codePercent...)
	case charLParen:
		return append(in, codeLParen...)
	case charRParen:
		return append(in, codeRParen...)
	case charEquals:
		return append(in, codeEquals...)
	case charAddition:
		return append(in, codeAddition...)
	case charLCurly:
		return append(in, codeLCurly...)
	case charRCurly:
		return append(in, codeRCurly...)
	case charLBracket:
		return append(in, codeLBracket...)
	case charRBracket:
		return append(in, codeRBracket...)
	default:
		panic("We received an unexpected character: '" + string(b) + "'")
	}
}
