package loops

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler/translate"
)

func DoInfiniteLoop(s scan.Scanner, output langoutput.T) {
	output.Write("while(true) {")
}

func DoBreak(s scan.Scanner, output langoutput.T) {
	s.Must(";")
	output.Write("break;")
}

func DoWhile(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	cond := translate.Expr(s.Expr())
	s.Must(")")
	s.Must("faire")
	output.Writef("while(%v) {", cond)
}

func DoUntil(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	cond := translate.Expr(s.Expr())
	s.Must(")")

	output.Writef("while(!(%v))", cond)
	if s.Match("faire") { // start of a loop, or end of a loop ?
		output.Write("{")
	}

}

func DoRepeatUntil(s scan.Scanner, output langoutput.T) {
	output.Write("do {")
}