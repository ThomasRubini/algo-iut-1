package scanutils

import (
	"fmt"
	"strconv"
	"text/scanner"
)

func Text(s *scanner.Scanner) string {
	s.Scan()
	return s.TokenText()
}

func UntilEOL(s *scanner.Scanner) string {
	str := ""
	for {
		s.Scan()
		if s.TokenText() == ";" {
			return str
		} else {
			str += s.TokenText()
		}
	}
}

func Number(s *scanner.Scanner) int {
	str := Text(s)
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to int: %v", str, err))
	}
	return num
}

func Must(s *scanner.Scanner, str string) {
	s.Scan()
	if s.TokenText() != str {
		panic(fmt.Sprintf("expected %s, got %s", str, s.TokenText()))
	}
}
