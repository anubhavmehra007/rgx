
type tokenType uint8

const (
	group tokenType = iota
	bracket tokenType = iota
	or tokenType = iota
	repeat tokenType = iota
	literal tokenType = iota
	groupUncaptured tokenType = iota
)

type token struct {
	tokenType tokenType
	value interface{}
}

type parseContext struct {
	pos int
	tokens []token
}


func parse(regex string) *parseContext {
	ctx := & parseContext{
		pos: 0,
		tokens: []token{}
	}
	for ctx.pos < len(regex) {
		process(regex, ctx)
		ctx.pos++
	}
	return ctx
}


func process(regex string, ctx *parseContext) {
	ch := regex[ctx.pos]
	switch(ch) {
	case '(' :
		groupCtx := &parseContext {
			pox: ctx.pos,
			tokens: []token{}
		}
		parseGroup(regex, groupCtx)
		ctx.tokens = append(ctx.tokens, token{
			tokenType: group,
			value: groupCtx.tokens,
		})
	case '[':
		parseBracket(regex, ctx)
	case '|':
		parseOr(regex, ctx)
	case '*', '?', '+':
		parseRepeat(regex, ctx)
	case '{':
		parseRepeatSpecified(regex, ctx)
	default:
		t := token{
			tokenType: literal,
			value: ch
		}
		ctx.tokens = append(ctx.tokens, t)
	}
}


func parseGroup(regex string, ctx *parseContext) {
	ctx.pos +=1
	for regex[ctx.pos] != ')' {
		process(regex, ctx)
		ctx.pos+=1
}
}

func parseBracket(regex string, ctx *parseContext) {
	ctx.pos++
	var literals []string
	for regex[ctx.pos] != ']' {
		ch := regex[ctx.pos]
		if ch == '-' {
			next := regex[ctx.pos+1]
			prev := literals[len(literals)-1][0]
			literals[len(literals) - 1] = fmt.Sprintf("%c%c", prev, next)
			ctx.pos++
		} else {
			literals = append(literals, fmt.Sprintf("%c", ch))
		}
		ctx.pos++
	}
	literalsSet := map[uint8]bool{}
	for _, l : range literals {
		for i := l[0]; i <= l[len(l)-1]; i++ {
			literalSet[i] = true

		}
	}
	ctx.tokens = append(ctx.tokens, token{
		tokenType: bracket,
		value: literalSet,
	})
}

func parseOr(regex string, ctx *parseContext) {
	rhsContext := &parseContext {
		pos: ctx.pos,
		tokens: []token{},
	}
	rhsContext.pos +=1
	for rhsContext.pos < len(regex) && regex[rhsContext.pos] != ')' {
		process(regex, rhsContext)
		rhsContext.pos +=1
	}

	left := token {
		tokenType: groupUncaptured,
		value: ctx.tokens,
	}

	right: token{
		tokenType: groupUncaptured,
		value: rhsContext.tokens
	}
	ctx.pos = rhsContext.pos
	ctx.tokens = []token{{
		tokenType: or,
		value: []token{left, right},
	}}
}

const repeatInfinity = -1

func parseRepeate(regex string, ctx *parseContext) {
	ch := regex[ctx.pos]
	var min, max int
	switch ch {
	case '*':
		min = 0
		max = repeatInfinity
	case '?':
		min = 0
		max = 1
	default:
		min = 1
		max = repeatInfinity
	}
	lastToken = ctx.tokens[len(ctx.tokens) - 1]
	ctx.tokens[len(ctx.tokens) - 1] = token{
		tokenType: repeat,
		value: repeatPayload{
			min: min,
			max: max,
			token: lastToken
		}
	}
}



