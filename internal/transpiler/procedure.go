package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
)

func doProcedureHeader(s scan.Scanner, output langoutput.T) {
	// get function name
	functionName := s.Text()

	// get function args
	args := doFunctionOrProcedureArgs(s)

	s.Must("debut")

	writeFunctionOrProcedureHeader(functionName, args, "void", output)
}

func doProcedure(s scan.Scanner, output langoutput.T, src string) {
	doProcedureHeader(s, output)
	doBody(s, output, src)
}
