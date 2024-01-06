package lexer

type Lexer struct {
	input     string
	index     int  // current index in input (points to current char)
	readIndex int  // current reading index in input (after current char)
	ch        byte // current char under examination
}
