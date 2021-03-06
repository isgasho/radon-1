package ast

type expr struct{}

func (e expr) Expr() {}

// The different AST nodes available to parse.
type (
	// An Identifier represents a name in the variable store; evaluates to the
	// value of that variable.
	Identifier struct {
		expr
		Value string
	}

	// A Number represents a number literal.
	Number struct {
		expr
		Value float64
	}

	// A Boolean represents a boolean literal.
	Boolean struct {
		expr
		Value bool
	}

	// A String represents a unicode string literal.
	String struct {
		expr
		Value string
	}

	// Nil is the nil literal; the absence of a value.
	Nil struct {
		expr
	}

	// A List is a linked list of items.
	List struct {
		expr
		Value []Expression
	}

	// A Map is a hashmap.
	Map struct {
		expr
		Value map[Expression]Expression
	}

	// A Block combines multiple statements into an expression.
	Block struct {
		expr
		Value []Statement
	}

	// A Prefix is a prefix operator expression, such as -5.
	Prefix struct {
		expr
		Operator string
		Right    Expression
	}

	// An Infix is an infix operator expression, such as 1 + 2
	Infix struct {
		expr
		Operator    string
		Left, Right Expression
	}

	// A Call calls a function with the argument. If the argument is a tuple,
	// each element of the tuple is passed as a separate argument.
	Call struct {
		expr
		Function, Argument Expression
	}

	// An If expression executes Consequence or Alternative based on Condition.
	If struct {
		expr
		Condition, Consequence, Alternative Expression
	}

	// A MatchBranch is a condition -> body branch for use in a Match expression.
	// Notice a MatchBranch isn't an Expression itself.
	MatchBranch struct {
		Condition, Body Expression
	}

	// A Match executes a different piece of code based on the input value.
	Match struct {
		expr
		Input    Expression
		Branches []MatchBranch
	}

	// A Model expression defines a new model. Parent is another model, since
	// the syntax is `model dog (name) : animal (name, "dog")`.
	Model struct {
		expr
		Parameters, Parent Expression
	}
)
