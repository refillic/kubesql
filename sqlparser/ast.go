package sqlparser

type LineNumberRange struct {
	Start uint
	End uint
}

type Node struct {
	Lines LineNumberRange
	Children []Node
}

type SelectClause struct {
	Node
	Items []SelectClauseItem
}

type SelectClauseItem struct {
	Node
	Column Expression
	Alias string
}

type Expression struct {
	Node
}

type FunctionCall struct {
	Expression
	Function string
	Arguments []Expression
}

type IDPath struct {
	Expression
	Path string
}

type StringLiteral struct {
	Expression
	Value string
}

type IntegerLiteral struct {
	Expression
	Value int64
}

type DecimalLiteral struct {
	Expression
	Value int64
}

type BooleanLiteral struct {
	Expression
	Value bool
}

type AbstractSyntaxTree struct {
	Root Node
}