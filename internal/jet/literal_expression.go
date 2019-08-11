package jet

import (
	"fmt"
	"strconv"
	"time"
)

// Representation of an escaped literal
type LiteralExpression interface {
	Expression

	Value() interface{}
	SetConstant(constant bool)
}

type literalExpressionImpl struct {
	ExpressionInterfaceImpl
	noOpVisitorImpl

	value    interface{}
	constant bool
}

func literal(value interface{}, optionalConstant ...bool) *literalExpressionImpl {
	exp := literalExpressionImpl{value: value}

	if len(optionalConstant) > 0 {
		exp.constant = optionalConstant[0]
	}

	exp.ExpressionInterfaceImpl.Parent = &exp

	return &exp
}

func constLiteral(value interface{}) *literalExpressionImpl {
	exp := literal(value)
	exp.constant = true

	return exp
}

func (l *literalExpressionImpl) serialize(statement StatementType, out *SqlBuilder, options ...SerializeOption) error {
	if l.constant {
		out.insertConstantArgument(l.value)
	} else {
		out.insertParametrizedArgument(l.value)
	}

	return nil
}

func (l *literalExpressionImpl) Value() interface{} {
	return l.value
}

func (l *literalExpressionImpl) SetConstant(constant bool) {
	l.constant = constant
}

type integerLiteralExpression struct {
	literalExpressionImpl
	integerInterfaceImpl
}

// Int is constructor for integer expressions literals.
func Int(value int64, constant ...bool) IntegerExpression {
	numLiteral := &integerLiteralExpression{}

	numLiteral.literalExpressionImpl = *literal(value)
	if len(constant) > 0 && constant[0] == true {
		numLiteral.constant = true
	}

	numLiteral.literalExpressionImpl.Parent = numLiteral
	numLiteral.integerInterfaceImpl.parent = numLiteral

	return numLiteral
}

//---------------------------------------------------//
type boolLiteralExpression struct {
	boolInterfaceImpl
	literalExpressionImpl
}

// Bool creates new bool literal expression
func Bool(value bool) BoolExpression {
	boolLiteralExpression := boolLiteralExpression{}

	boolLiteralExpression.literalExpressionImpl = *literal(value)
	boolLiteralExpression.boolInterfaceImpl.parent = &boolLiteralExpression

	return &boolLiteralExpression
}

//---------------------------------------------------//
type floatLiteral struct {
	floatInterfaceImpl
	literalExpressionImpl
}

// Float creates new float literal expression
func Float(value float64) FloatExpression {
	floatLiteral := floatLiteral{}
	floatLiteral.literalExpressionImpl = *literal(value)

	floatLiteral.floatInterfaceImpl.parent = &floatLiteral

	return &floatLiteral
}

//---------------------------------------------------//
type stringLiteral struct {
	stringInterfaceImpl
	literalExpressionImpl
}

// String creates new string literal expression
func String(value string, constant ...bool) StringExpression {
	stringLiteral := stringLiteral{}
	stringLiteral.literalExpressionImpl = *literal(value)
	if len(constant) > 0 && constant[0] == true {
		stringLiteral.constant = true
	}

	stringLiteral.stringInterfaceImpl.parent = &stringLiteral

	return &stringLiteral
}

func formatMilliseconds(milliseconds ...int) string {
	if len(milliseconds) > 0 {
		if milliseconds[0] < 1000 {
			return fmt.Sprintf(".%03d", milliseconds[0])
		} else {
			return "." + strconv.Itoa(milliseconds[0])
		}
	}

	return ""
}

// Time creates new time literal expression
func Time(hour, minute, second int, milliseconds ...int) TimeExpression {
	timeStr := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)

	timeStr += formatMilliseconds(milliseconds...)

	return TimeExp(literal(timeStr))
}

func TimeT(t time.Time) TimeExpression {
	return TimeExp(literal(t))
}

// Timez creates new time with time zone literal expression
func Timez(hour, minute, second, milliseconds, timezone int) TimezExpression {
	timeStr := fmt.Sprintf("%02d:%02d:%02d.%03d %+03d", hour, minute, second, milliseconds, timezone)

	return TimezExp(literal(timeStr))
}

func TimezT(t time.Time) TimezExpression {
	return TimezExp(literal(t))
}

// Timestamp creates new timestamp literal expression
func Timestamp(year int, month time.Month, day, hour, minute, second int, milliseconds ...int) TimestampExpression {
	timeStr := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

	timeStr += formatMilliseconds(milliseconds...)

	return TimestampExp(literal(timeStr))
}

func TimestampT(t time.Time) TimestampExpression {
	return TimestampExp(literal(t))
}

// Timestampz creates new timestamp with time zone literal expression
func Timestampz(year, month, day, hour, minute, second, milliseconds, timezone int) TimestampzExpression {
	timeStr := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%03d %+04d",
		year, month, day, hour, minute, second, milliseconds, timezone)

	return TimestampzExp(literal(timeStr))
}

func TimestampzT(t time.Time) TimestampzExpression {
	return TimestampzExp(literal(t))
}

//Date creates new date expression
func Date(year int, month time.Month, day int) DateExpression {
	timeStr := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	return DateExp(literal(timeStr))
}

func DateT(t time.Time) DateExpression {
	return DateExp(literal(t))
}

//--------------------------------------------------//
type nullLiteral struct {
	ExpressionInterfaceImpl
	noOpVisitorImpl
}

func newNullLiteral() Expression {
	nullExpression := &nullLiteral{}

	nullExpression.ExpressionInterfaceImpl.Parent = nullExpression

	return nullExpression
}

func (n *nullLiteral) serialize(statement StatementType, out *SqlBuilder, options ...SerializeOption) error {
	out.WriteString("NULL")
	return nil
}

//--------------------------------------------------//
type starLiteral struct {
	ExpressionInterfaceImpl
	noOpVisitorImpl
}

func newStarLiteral() Expression {
	starExpression := &starLiteral{}

	starExpression.ExpressionInterfaceImpl.Parent = starExpression

	return starExpression
}

func (n *starLiteral) serialize(statement StatementType, out *SqlBuilder, options ...SerializeOption) error {
	out.WriteString("*")
	return nil
}

//---------------------------------------------------//

type wrap struct {
	ExpressionInterfaceImpl
	expressions []Expression
}

func (n *wrap) accept(visitor visitor) {
	for _, exp := range n.expressions {
		exp.accept(visitor)
	}
}

func (n *wrap) serialize(statement StatementType, out *SqlBuilder, options ...SerializeOption) error {
	out.WriteString("(")
	err := serializeExpressionList(statement, n.expressions, ", ", out)
	out.WriteString(")")
	return err
}

// WRAP wraps list of expressions with brackets '(' and ')'
func WRAP(expression ...Expression) Expression {
	wrap := &wrap{expressions: expression}
	wrap.ExpressionInterfaceImpl.Parent = wrap

	return wrap
}

//---------------------------------------------------//

type rawExpression struct {
	ExpressionInterfaceImpl
	noOpVisitorImpl

	raw string
}

func (n *rawExpression) serialize(statement StatementType, out *SqlBuilder, options ...SerializeOption) error {
	out.WriteString(n.raw)
	return nil
}

// Raw can be used for any unsupported functions, operators or expressions.
// For example: Raw("current_database()")
func Raw(raw string) Expression {
	rawExp := &rawExpression{raw: raw}
	rawExp.ExpressionInterfaceImpl.Parent = rawExp

	return rawExp
}