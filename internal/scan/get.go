package scan

import (
	"algo-iut/internal/transpiler/translate"
	"fmt"
	"slices"
	"strconv"
)

func (s *impl) Text() string {
	str := s.Peek()
	s.Advance()
	return str
}

func (s *impl) Number() int {
	str := s.Text()
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to int: %v", str, err))
	}
	return num
}

var keywords = []string{
	// loops
	"tant_que", "pour", "jusqua", "boucle", "fboucle",
	"faire", "ffaire", "sortie",
	// function stuff
	"fonction", "procedure", "renvoie",
	// others
	"debut", "fin",
}

func (s *impl) LValue() string {
	tok := s.Text()
	if slices.Contains(keywords, tok) {
		s.InvalidToken("Expected lvalue, found reserved keyword")
	}

	if s.Match("[") {
		inside := s.Expr()
		s.Must("]")
		return fmt.Sprintf("%v[%v]", tok, translate.Expr(inside))
	} else {
		return tok
	}
}
