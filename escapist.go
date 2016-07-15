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

// Escape likes to escape
func Escape(in []byte) (out []byte) {
	return escape(in, true)
}

// EscapeAdv likes to escape very thoroguhly
func EscapeAdv(in []byte) (out []byte) {
	return escape(in, false)
}

// Escape likes to escape
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
		gfn = getNextEscChar
		afn = appendEscapedBytes
	} else {
		gfn = getNextAdvEscChar
		afn = appendAdvEscapedBytes
	}

	// Seek for the next character to escape, iterate until no character is found
	for nc = gfn(in); nc > -1; nc = gfn(in[nc+1:]) {
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
		panic("Run Forrest, run! Something bad happened, we received '" + string(b) + "'")
	}
}

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
		panic("Run Forrest, run! Something bad happened, we received '" + string(b) + "'")
	}
}
