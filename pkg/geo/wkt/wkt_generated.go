// Code generated by goyacc -o wkt_generated.go -p wkt wkt.y. DO NOT EDIT.

//line wkt.y:12

package wkt

import __yyfmt__ "fmt"

//line wkt.y:13

import "github.com/twpayne/go-geom"

func isValidLineString(wktlex wktLexer, flatCoords []float64, stride int) bool {
	if len(flatCoords) < 2*stride {
		wktlex.(*wktLex).Error("syntax error: non-empty linestring with only one point")
		return false
	}
	return true
}

func isValidPolygonRing(wktlex wktLexer, flatCoords []float64, stride int) bool {
	if len(flatCoords) < 4*stride {
		wktlex.(*wktLex).Error("syntax error: polygon ring doesn't have enough points")
		return false
	}
	for i := 0; i < stride; i++ {
		if flatCoords[i] != flatCoords[len(flatCoords)-stride+i] {
			wktlex.(*wktLex).Error("syntax error: polygon ring not closed")
			return false
		}
	}
	return true
}

type geomFlatCoordsRepr struct {
	flatCoords []float64
	ends       []int
}

func makeGeomFlatCoordsRepr(flatCoords []float64) geomFlatCoordsRepr {
	return geomFlatCoordsRepr{flatCoords: flatCoords, ends: []int{len(flatCoords)}}
}

func appendGeomFlatCoordsReprs(p1 geomFlatCoordsRepr, p2 geomFlatCoordsRepr) geomFlatCoordsRepr {
	if len(p1.ends) > 0 {
		p1LastEnd := p1.ends[len(p1.ends)-1]
		for i, _ := range p2.ends {
			p2.ends[i] += p1LastEnd
		}
	}
	return geomFlatCoordsRepr{flatCoords: append(p1.flatCoords, p2.flatCoords...), ends: append(p1.ends, p2.ends...)}
}

//line wkt.y:60
type wktSymType struct {
	yys       int
	str       string
	geom      geom.T
	coord     float64
	coordList []float64
	flatRepr  geomFlatCoordsRepr
}

const POINT = 57346
const POINTM = 57347
const POINTZ = 57348
const POINTZM = 57349
const LINESTRING = 57350
const LINESTRINGM = 57351
const LINESTRINGZ = 57352
const LINESTRINGZM = 57353
const POLYGON = 57354
const POLYGONM = 57355
const POLYGONZ = 57356
const POLYGONZM = 57357
const MULTIPOINT = 57358
const MULTIPOINTM = 57359
const MULTIPOINTZ = 57360
const MULTIPOINTZM = 57361
const MULTILINESTRING = 57362
const MULTILINESTRINGM = 57363
const MULTILINESTRINGZ = 57364
const MULTILINESTRINGZM = 57365
const EMPTY = 57366
const NUM = 57367

var wktToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"POINT",
	"POINTM",
	"POINTZ",
	"POINTZM",
	"LINESTRING",
	"LINESTRINGM",
	"LINESTRINGZ",
	"LINESTRINGZM",
	"POLYGON",
	"POLYGONM",
	"POLYGONZ",
	"POLYGONZM",
	"MULTIPOINT",
	"MULTIPOINTM",
	"MULTIPOINTZ",
	"MULTIPOINTZM",
	"MULTILINESTRING",
	"MULTILINESTRINGM",
	"MULTILINESTRINGZ",
	"MULTILINESTRINGZM",
	"EMPTY",
	"NUM",
	"'('",
	"')'",
	"','",
}

var wktStatenames = [...]string{}

const wktEofCode = 1
const wktErrCode = 2
const wktInitialStackSize = 16

//line yacctab:1
var wktExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const wktPrivate = 57344

const wktLast = 263

var wktAct = [...]int{
	81, 144, 147, 142, 148, 143, 126, 121, 123, 108,
	127, 46, 132, 97, 45, 98, 96, 120, 141, 140,
	95, 56, 122, 191, 192, 190, 189, 188, 189, 152,
	94, 47, 186, 187, 184, 185, 182, 183, 180, 181,
	179, 178, 177, 178, 115, 175, 176, 173, 174, 90,
	113, 171, 172, 170, 167, 169, 165, 168, 165, 166,
	167, 164, 165, 162, 163, 151, 117, 160, 161, 100,
	150, 100, 99, 100, 112, 136, 110, 83, 134, 133,
	128, 149, 56, 82, 104, 83, 82, 124, 109, 101,
	102, 51, 103, 92, 201, 101, 145, 158, 159, 91,
	48, 92, 91, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 156, 157, 119, 86, 40, 119, 85,
	35, 119, 197, 205, 119, 84, 32, 140, 140, 51,
	201, 86, 40, 85, 35, 140, 80, 48, 79, 78,
	76, 77, 75, 74, 72, 73, 71, 70, 196, 69,
	68, 66, 67, 65, 64, 62, 63, 61, 60, 58,
	59, 57, 55, 117, 56, 197, 216, 100, 99, 202,
	200, 112, 204, 203, 215, 136, 208, 207, 209, 86,
	128, 211, 212, 213, 149, 214, 210, 206, 53, 101,
	51, 50, 90, 51, 85, 44, 199, 48, 195, 39,
	198, 40, 37, 34, 35, 35, 31, 138, 32, 194,
	193, 155, 154, 153, 84, 137, 114, 1, 139, 146,
	43, 129, 131, 54, 116, 29, 33, 36, 42, 49,
	52, 41, 118, 30, 130, 135, 38, 125, 105, 107,
	106, 28, 111, 93, 89, 88, 87, 7, 6, 5,
	4, 3, 2,
}

var wktPact = [...]int{
	99, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 192, 189,
	188, 185, 181, 177, 174, 148, 145, 144, 141, 140,
	137, 136, 133, 130, 129, 126, 125, 122, -1000, -1000,
	-1000, -1000, 199, -1000, -1000, 179, -1000, -1000, -1000, -1000,
	164, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 199, -1000,
	-1000, 179, -1000, -1000, -1000, -1000, 164, 74, -1000, 65,
	-1000, 65, -1000, 56, -1000, 110, -1000, 104, -1000, 104,
	-1000, 101, -1000, 121, -1000, 113, -1000, 113, -1000, -5,
	-1000, 43, 38, 2, 198, 197, 196, 96, 70, 40,
	-1000, -1000, -1000, 36, 34, 32, -1000, -1000, -1000, -1000,
	-1000, -1000, 30, 28, 26, 24, 20, 18, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	15, -1000, -1000, -1000, 13, 11, -1000, -1000, -1000, 9,
	7, 5, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 0, -1000, -1000, -1000, -2, -4, -1000, -1000, -1000,
	-1000, -1000, -1000, 195, 194, 183, -1000, 150, -1000, 179,
	-1000, 164, -1000, 68, -1000, 65, -1000, 56, -1000, -1000,
	-1000, -1000, 107, -1000, 118, -1000, 116, -1000, 104, -1000,
	-1000, 101, -1000, 114, -1000, 65, -1000, 56, -1000, 113,
	-1000, -1000, -5, 159, -1000, 159, -1000, 151, -1000, -1000,
	-1000, 150, -1000, -1000, -1000, 150, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000,
}

var wktPgo = [...]int{
	0, 262, 261, 260, 259, 258, 257, 0, 50, 44,
	242, 226, 234, 256, 255, 254, 14, 11, 31, 228,
	225, 217, 16, 13, 15, 253, 30, 20, 252, 22,
	10, 250, 249, 8, 1, 9, 7, 6, 248, 17,
	247, 245, 5, 4, 12, 3, 2, 244, 232, 231,
	18, 229, 227,
}

var wktR1 = [...]int{
	0, 52, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 25, 25, 26,
	26, 27, 27, 22, 23, 24, 47, 47, 48, 48,
	49, 49, 50, 50, 51, 51, 44, 44, 45, 45,
	46, 46, 41, 42, 43, 19, 20, 21, 16, 17,
	18, 34, 13, 13, 14, 14, 15, 15, 31, 31,
	32, 32, 38, 38, 39, 39, 40, 40, 35, 35,
	36, 36, 37, 37, 28, 28, 29, 29, 30, 30,
	33, 10, 11, 12, 7, 8, 9,
}

var wktR2 = [...]int{
	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 4, 4, 4,
	4, 4, 4, 2, 2, 2, 2, 4, 4, 4,
	4, 4, 4, 2, 2, 2, 2, 4, 4, 4,
	4, 4, 4, 2, 2, 2, 2, 3, 1, 3,
	1, 3, 1, 1, 1, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 3,
	3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 3, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 3, 2, 3, 4,
}

var wktChk = [...]int{
	-1000, -52, -1, -2, -3, -4, -5, -6, 4, 5,
	6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, -10, -11,
	-12, 24, 26, -11, 24, 26, -11, 24, -12, 24,
	26, -19, -20, -21, 24, -16, -17, -18, 26, -20,
	24, 26, -20, 24, -21, 24, 26, 26, 24, 26,
	24, 26, 24, 26, 24, 26, 24, 26, 24, 26,
	24, 26, 24, 26, 24, 26, 24, 26, 24, 26,
	24, -7, -8, -9, 25, 25, 25, -13, -14, -15,
	-7, -8, -9, -25, -26, -27, -22, -23, -24, -16,
	-17, -18, -26, -26, -27, -38, -31, -32, -35, -29,
	-30, -28, -33, -8, -11, -9, -12, -7, -10, 24,
	-39, -36, -29, -33, -39, -40, -37, -30, -33, -49,
	-47, -48, -44, -42, -43, -41, -34, -20, -21, -19,
	24, -50, -45, -42, -34, -50, -51, -46, -43, -34,
	27, 27, 27, 25, 25, 25, 27, 28, 27, 28,
	27, 28, 27, 28, 27, 28, 27, 28, 27, 27,
	27, 27, 28, 27, 28, 27, 28, 27, 28, 27,
	27, 28, 27, 28, 27, 28, 27, 28, 27, 28,
	27, 27, 28, 25, 25, 25, -7, 25, -8, -9,
	-22, 26, -23, -24, -35, 26, -29, -30, -36, -37,
	-44, -42, -43, -45, -46, 25, 25,
}

var wktDef = [...]int{
	0, -2, 1, 2, 3, 4, 5, 6, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 7, 8,
	9, 13, 0, 10, 14, 0, 11, 15, 12, 16,
	0, 17, 18, 19, 23, 85, 86, 87, 0, 20,
	24, 0, 21, 25, 22, 26, 0, 0, 33, 0,
	34, 0, 35, 0, 36, 0, 43, 0, 44, 0,
	45, 0, 46, 0, 53, 0, 54, 0, 55, 0,
	56, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	93, 95, 97, 0, 0, 0, 58, 60, 62, 63,
	64, 65, 0, 0, 0, 0, 0, 0, 103, 99,
	101, 108, 109, 116, 117, 118, 119, 114, 115, 120,
	0, 105, 110, 111, 0, 0, 107, 112, 113, 0,
	0, 0, 71, 67, 69, 76, 77, 83, 84, 82,
	91, 0, 73, 78, 79, 0, 0, 75, 80, 81,
	121, 122, 123, 124, 0, 0, 88, 0, 89, 0,
	90, 0, 27, 0, 28, 0, 29, 0, 30, 31,
	32, 37, 0, 38, 0, 39, 0, 40, 0, 41,
	42, 0, 47, 0, 48, 0, 49, 0, 50, 0,
	51, 52, 0, 125, 125, 0, 92, 0, 94, 96,
	57, 0, 59, 61, 102, 0, 98, 100, 104, 106,
	70, 66, 68, 72, 74, 126, 124,
}

var wktTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	26, 27, 3, 3, 28,
}

var wktTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25,
}

var wktTok3 = [...]int{
	0,
}

var wktErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	wktDebug        = 0
	wktErrorVerbose = true
)

type wktLexer interface {
	Lex(lval *wktSymType) int
	Error(s string)
}

type wktParser interface {
	Parse(wktLexer) int
	Lookahead() int
}

type wktParserImpl struct {
	lval  wktSymType
	stack [wktInitialStackSize]wktSymType
	char  int
}

func (p *wktParserImpl) Lookahead() int {
	return p.char
}

func wktNewParser() wktParser {
	return &wktParserImpl{}
}

const wktFlag = -1000

func wktTokname(c int) string {
	if c >= 1 && c-1 < len(wktToknames) {
		if wktToknames[c-1] != "" {
			return wktToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func wktStatname(s int) string {
	if s >= 0 && s < len(wktStatenames) {
		if wktStatenames[s] != "" {
			return wktStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func wktErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !wktErrorVerbose {
		return "syntax error"
	}

	for _, e := range wktErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + wktTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := wktPact[state]
	for tok := TOKSTART; tok-1 < len(wktToknames); tok++ {
		if n := base + tok; n >= 0 && n < wktLast && wktChk[wktAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if wktDef[state] == -2 {
		i := 0
		for wktExca[i] != -1 || wktExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; wktExca[i] >= 0; i += 2 {
			tok := wktExca[i]
			if tok < TOKSTART || wktExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if wktExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += wktTokname(tok)
	}
	return res
}

func wktlex1(lex wktLexer, lval *wktSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = wktTok1[0]
		goto out
	}
	if char < len(wktTok1) {
		token = wktTok1[char]
		goto out
	}
	if char >= wktPrivate {
		if char < wktPrivate+len(wktTok2) {
			token = wktTok2[char-wktPrivate]
			goto out
		}
	}
	for i := 0; i < len(wktTok3); i += 2 {
		token = wktTok3[i+0]
		if token == char {
			token = wktTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = wktTok2[1] /* unknown char */
	}
	if wktDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", wktTokname(token), uint(char))
	}
	return char, token
}

func wktParse(wktlex wktLexer) int {
	return wktNewParser().Parse(wktlex)
}

func (wktrcvr *wktParserImpl) Parse(wktlex wktLexer) int {
	var wktn int
	var wktVAL wktSymType
	var wktDollar []wktSymType
	_ = wktDollar // silence set and not used
	wktS := wktrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	wktstate := 0
	wktrcvr.char = -1
	wkttoken := -1 // wktrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		wktstate = -1
		wktrcvr.char = -1
		wkttoken = -1
	}()
	wktp := -1
	goto wktstack

ret0:
	return 0

ret1:
	return 1

wktstack:
	/* put a state and value onto the stack */
	if wktDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", wktTokname(wkttoken), wktStatname(wktstate))
	}

	wktp++
	if wktp >= len(wktS) {
		nyys := make([]wktSymType, len(wktS)*2)
		copy(nyys, wktS)
		wktS = nyys
	}
	wktS[wktp] = wktVAL
	wktS[wktp].yys = wktstate

wktnewstate:
	wktn = wktPact[wktstate]
	if wktn <= wktFlag {
		goto wktdefault /* simple state */
	}
	if wktrcvr.char < 0 {
		wktrcvr.char, wkttoken = wktlex1(wktlex, &wktrcvr.lval)
	}
	wktn += wkttoken
	if wktn < 0 || wktn >= wktLast {
		goto wktdefault
	}
	wktn = wktAct[wktn]
	if wktChk[wktn] == wkttoken { /* valid shift */
		wktrcvr.char = -1
		wkttoken = -1
		wktVAL = wktrcvr.lval
		wktstate = wktn
		if Errflag > 0 {
			Errflag--
		}
		goto wktstack
	}

wktdefault:
	/* default state action */
	wktn = wktDef[wktstate]
	if wktn == -2 {
		if wktrcvr.char < 0 {
			wktrcvr.char, wkttoken = wktlex1(wktlex, &wktrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if wktExca[xi+0] == -1 && wktExca[xi+1] == wktstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			wktn = wktExca[xi+0]
			if wktn < 0 || wktn == wkttoken {
				break
			}
		}
		wktn = wktExca[xi+1]
		if wktn < 0 {
			goto ret0
		}
	}
	if wktn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			wktlex.Error(wktErrorMessage(wktstate, wkttoken))
			Nerrs++
			if wktDebug >= 1 {
				__yyfmt__.Printf("%s", wktStatname(wktstate))
				__yyfmt__.Printf(" saw %s\n", wktTokname(wkttoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for wktp >= 0 {
				wktn = wktPact[wktS[wktp].yys] + wktErrCode
				if wktn >= 0 && wktn < wktLast {
					wktstate = wktAct[wktn] /* simulate a shift of "error" */
					if wktChk[wktstate] == wktErrCode {
						goto wktstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if wktDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", wktS[wktp].yys)
				}
				wktp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if wktDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", wktTokname(wkttoken))
			}
			if wkttoken == wktEofCode {
				goto ret1
			}
			wktrcvr.char = -1
			wkttoken = -1
			goto wktnewstate /* try again in the same state */
		}
	}

	/* reduction by production wktn */
	if wktDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", wktn, wktStatname(wktstate))
	}

	wktnt := wktn
	wktpt := wktp
	_ = wktpt // guard against "declared and not used"

	wktp -= wktR2[wktn]
	// wktp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if wktp+1 >= len(wktS) {
		nyys := make([]wktSymType, len(wktS)*2)
		copy(nyys, wktS)
		wktS = nyys
	}
	wktVAL = wktS[wktp+1]

	/* consult goto table to find next state */
	wktn = wktR1[wktn]
	wktg := wktPgo[wktn]
	wktj := wktg + wktS[wktp].yys + 1

	if wktj >= wktLast {
		wktstate = wktAct[wktg]
	} else {
		wktstate = wktAct[wktj]
		if wktChk[wktstate] != -wktn {
			wktstate = wktAct[wktg]
		}
	}
	// dummy call; replaced with literal code
	switch wktnt {

	case 1:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:106
		{
			wktlex.(*wktLex).ret = wktDollar[1].geom
		}
	case 7:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:119
		{
			wktVAL.geom = geom.NewPointFlat(geom.XY, wktDollar[2].coordList)
		}
	case 8:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:123
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 9:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:127
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 10:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:131
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYM, wktDollar[2].coordList)
		}
	case 11:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:135
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 12:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:139
		{
			wktVAL.geom = geom.NewPointFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 13:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:143
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XY)
		}
	case 14:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:147
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XYM)
		}
	case 15:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:151
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XYZ)
		}
	case 16:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:155
		{
			wktVAL.geom = geom.NewPointEmpty(geom.XYZM)
		}
	case 17:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:161
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XY, wktDollar[2].coordList)
		}
	case 18:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:165
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 19:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:169
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 20:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:173
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYM, wktDollar[2].coordList)
		}
	case 21:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:177
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZ, wktDollar[2].coordList)
		}
	case 22:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:181
		{
			wktVAL.geom = geom.NewLineStringFlat(geom.XYZM, wktDollar[2].coordList)
		}
	case 23:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:185
		{
			wktVAL.geom = geom.NewLineString(geom.XY)
		}
	case 24:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:189
		{
			wktVAL.geom = geom.NewLineString(geom.XYM)
		}
	case 25:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:193
		{
			wktVAL.geom = geom.NewLineString(geom.XYZ)
		}
	case 26:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:197
		{
			wktVAL.geom = geom.NewLineString(geom.XYZM)
		}
	case 27:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:203
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XY, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 28:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:207
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 29:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:211
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 30:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:215
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 31:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:219
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 32:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:223
		{
			wktVAL.geom = geom.NewPolygonFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 33:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:227
		{
			wktVAL.geom = geom.NewPolygon(geom.XY)
		}
	case 34:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:231
		{
			wktVAL.geom = geom.NewPolygon(geom.XYM)
		}
	case 35:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:235
		{
			wktVAL.geom = geom.NewPolygon(geom.XYZ)
		}
	case 36:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:239
		{
			wktVAL.geom = geom.NewPolygon(geom.XYZM)
		}
	case 37:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:245
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XY, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 38:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:249
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZ, wktDollar[3].coordList)
		}
	case 39:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:253
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZM, wktDollar[3].coordList)
		}
	case 40:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:257
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYM, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 41:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:261
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 42:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:265
		{
			wktVAL.geom = geom.NewMultiPointFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, geom.NewMultiPointFlatOptionWithEnds(wktDollar[3].flatRepr.ends))
		}
	case 43:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:269
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XY)
		}
	case 44:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:273
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XYM)
		}
	case 45:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:277
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XYZ)
		}
	case 46:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:281
		{
			wktVAL.geom = geom.NewMultiPoint(geom.XYZM)
		}
	case 47:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:287
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XY, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 48:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:291
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 49:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:295
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 50:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:299
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 51:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:303
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZ, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 52:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:307
		{
			wktVAL.geom = geom.NewMultiLineStringFlat(geom.XYZM, wktDollar[3].flatRepr.flatCoords, wktDollar[3].flatRepr.ends)
		}
	case 53:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:311
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XY)
		}
	case 54:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:315
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XYM)
		}
	case 55:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:319
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XYZ)
		}
	case 56:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:323
		{
			wktVAL.geom = geom.NewMultiLineString(geom.XYZM)
		}
	case 57:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:329
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 59:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:336
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 61:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:343
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 63:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:350
		{
			if !isValidPolygonRing(wktlex, wktDollar[1].coordList, 2) {
				return 1
			}
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 64:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:359
		{
			if !isValidPolygonRing(wktlex, wktDollar[1].coordList, 3) {
				return 1
			}
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 65:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:368
		{
			if !isValidPolygonRing(wktlex, wktDollar[1].coordList, 4) {
				return 1
			}
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 66:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:379
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 68:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:386
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 70:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:393
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 72:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:400
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 74:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:407
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 82:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:426
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 83:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:432
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 84:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:438
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 85:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:444
		{
			if !isValidLineString(wktlex, wktDollar[1].coordList, 2) {
				return 1
			}
		}
	case 86:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:452
		{
			if !isValidLineString(wktlex, wktDollar[1].coordList, 3) {
				return 1
			}
		}
	case 87:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:460
		{
			if !isValidLineString(wktlex, wktDollar[1].coordList, 4) {
				return 1
			}
		}
	case 88:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:468
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 89:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:474
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 90:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:480
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 91:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:486
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(nil)
		}
	case 92:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:492
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 94:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:499
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 96:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:506
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 98:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:515
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 100:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:522
		{
			wktVAL.coordList = append(wktDollar[1].coordList, wktDollar[3].coordList...)
		}
	case 102:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:529
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 104:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:536
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 106:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:543
		{
			wktVAL.flatRepr = appendGeomFlatCoordsReprs(wktDollar[1].flatRepr, wktDollar[3].flatRepr)
		}
	case 108:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:550
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 110:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:557
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 112:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:564
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(wktDollar[1].coordList)
		}
	case 120:
		wktDollar = wktS[wktpt-1 : wktpt+1]
//line wkt.y:583
		{
			wktVAL.flatRepr = makeGeomFlatCoordsRepr(nil)
		}
	case 121:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:589
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 122:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:595
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 123:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:601
		{
			wktVAL.coordList = wktDollar[2].coordList
		}
	case 124:
		wktDollar = wktS[wktpt-2 : wktpt+1]
//line wkt.y:607
		{
			wktVAL.coordList = []float64{wktDollar[1].coord, wktDollar[2].coord}
		}
	case 125:
		wktDollar = wktS[wktpt-3 : wktpt+1]
//line wkt.y:613
		{
			wktVAL.coordList = []float64{wktDollar[1].coord, wktDollar[2].coord, wktDollar[3].coord}
		}
	case 126:
		wktDollar = wktS[wktpt-4 : wktpt+1]
//line wkt.y:619
		{
			wktVAL.coordList = []float64{wktDollar[1].coord, wktDollar[2].coord, wktDollar[3].coord, wktDollar[4].coord}
		}
	}
	goto wktstack /* stack new state and value */
}
