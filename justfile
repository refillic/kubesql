generate-parser:
    goyacc -o sqlparser/parser.go sqlparser/sql-grammar.y

test: generate-parser
    go test