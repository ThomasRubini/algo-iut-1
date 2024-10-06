package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/tabanalyser"
	"algo-iut/internal/transpiler/loops"
	"strings"
)

func doReturn(s scan.Scanner, output langoutput.T) {
	value := s.UntilEOL()
	output.Writef("return %s;", value)
}

func doAfficher(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	args := getFunctionArgs(s)
	s.Must(")")
	s.Must(";")
	output.Writef("std::cout << %s << std::endl;", strings.Join(args, " << "))
}

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s scan.Scanner, output langoutput.T, src string) {
	tabsPrefix := tabanalyser.Do(src)
	for {
		if !doLine(s, output, tabsPrefix) {
			break
		}
	}
}

func doLigneSuivante(s scan.Scanner, output langoutput.T) {
	s.Must(";")
	output.Write("std::cout << std::endl;")
}

func doLine(s scan.Scanner, output langoutput.T, tabsPrefix []string) bool {
	// write tabs/space prefix
	prefix := tabsPrefix[s.Pos().Line-1]
	output.Write(prefix)

	tok := s.Peek()
	switch tok {
	// conditions
	case "si":
		s.Advance()
		doCondition(s, output)
	case "sinon":
		s.Advance()
		output.Write("} else {")
	case "sinon_si":
		s.Advance()
		doElseIf(s, output)
	case "fsi":
		s.Advance()
		output.Write("}")
	// loops
	case "pour":
		s.Advance()
		loops.DoPourLoop(s, output)
	case "tant_que":
		s.Advance()
		loops.DoWhile(s, output)
	case "repeter":
		s.Advance()
		loops.DoRepeat(s, output)
	case "jusqua":
		s.Advance()
		loops.DoUntil(s, output)
	case "repeat":
		s.Advance()
		loops.DoRepeatUntil(s, output)
	case "boucle":
		s.Advance()
		loops.DoInfiniteLoop(s, output)
	case "sortie":
		s.Advance()
		loops.DoBreak(s, output)
	case "ffaire":
		s.Advance()
		output.Write("}")
	case "fboucle":
		s.Advance()
		output.Write("}")
	// afficher special case
	case "afficher":
		s.Advance()
		doAfficher(s, output)
	// others
	case "ligne_suivante":
		s.Advance()
		doLigneSuivante(s, output)
	case "declarer":
		s.Advance()
		doDeclare(s, output)
	case "renvoie":
		s.Advance()
		doReturn(s, output)
	case "fin":
		s.Advance()
		output.Write("}\n")
		return false
	default:
		doLValueStart(s, output)
	}
	output.Write("\n")

	return true
}
