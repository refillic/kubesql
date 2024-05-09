%{
	package main

	import "fmt"
	import "flag"
%}

FOR_SYSTEM_TIME  (?i)FOR SYSTEM_TIME
FOR_VERSION      (?i)FOR VERSION
KEYWORD      (?i)(ALL|AS|ASC|BY|CALL|CONTAINED|CROSS|DATA|DEFAULT|DELETE|DESC|DISTINCT|DIV|DUPLICATE|DUMPFILE|ENCLOSED|ESCAPED|EXCEPT|EXISTS|FOR|FORCE|FROM|GROUP|HAVING|INNER|INTERSECT|INSERT|JOIN|KEYS?|LAST_INSERT_ID|LEFT|LIMIT|LINES|LOAD|LOCK|MODE|NATURAL|NEXT|OF|OFFSET|ON|OPTIONALLY|ORDER|OUTER|OUTFILE|RIGHT|SELECT|SET|SHARE|SQL_CACHE|SQL_CALC_FOUND_ROWS|SQL_NO_CACHE|STARTING|STRAIGHT_JOIN|STREAM|SYSTEM_TIME|TERMINATED|UNION|UNIQUE|UNLOCK|UPDATE|USE|USING|VALUES?|VERSIONS?|WHERE)
DIGIT    [0-9]
ID       [a-z][a-z0-9]*

%%

DIV|    {

}

{DIGIT}+	{
			var i int
			fmt.Sscanf(yytext, "%d", &i)
			fmt.Printf("An integer: %s (%d)\n", yytext, i)
		}

{DIGIT}+"."{DIGIT}*	{
			var i float32
			fmt.Sscanf(yytext, "%g", &i)
			fmt.Printf("A float: %s (%g)\n", yytext, i)
		}



if|then|begin|end|procedure|function	{
			fmt.Printf("A keyword: %s\n", yytext)
		}

{ID}			fmt.Printf("An identifier: %s\n", yytext)

"+"|"-"|"*"|"/"		fmt.Printf("An operator: %s\n", yytext)

"{"[^{}\n]*"}"		/* eat up one-line comments */

[ \t\n]+		/* eat up whitespace */

.			fmt.Printf("Unrecognized character: %s\n", yytext)

%%

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		yyin, _ = os.Open(flag.Arg(0))
	} else {
		yyin = os.Stdin
	}

	yylex()
}