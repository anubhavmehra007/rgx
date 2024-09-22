//Types

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

//Types End


/*
Function parse
Input string
Ouput *parseContext
/*

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


/*
Function process
Input string, *parseContext
Ouput void
/*

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


