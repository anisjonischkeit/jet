package jet

type BoolExpression interface {
	Expression

	// Check if this expression is equal to rhs
	EQ(rhs BoolExpression) BoolExpression
	// Check if this expression is not equal to rhs
	NOT_EQ(rhs BoolExpression) BoolExpression
	// Check if this expression is distinct to rhs
	IS_DISTINCT_FROM(rhs BoolExpression) BoolExpression
	// Check if this expression is not distinct to rhs
	IS_NOT_DISTINCT_FROM(rhs BoolExpression) BoolExpression

	// Check if this expression is true
	IS_TRUE() BoolExpression
	// Check if this expression is not true
	IS_NOT_TRUE() BoolExpression
	// Check if this expression is false
	IS_FALSE() BoolExpression
	// Check if this expression is not false
	IS_NOT_FALSE() BoolExpression
	// Check if this expression is unknown
	IS_UNKNOWN() BoolExpression
	// Check if this expression is not unknown
	IS_NOT_UNKNOWN() BoolExpression

	// expression AND operator rhs
	AND(rhs BoolExpression) BoolExpression
	// expression OR operator rhs
	OR(rhs BoolExpression) BoolExpression
}

type boolInterfaceImpl struct {
	parent BoolExpression
}

func (b *boolInterfaceImpl) EQ(expression BoolExpression) BoolExpression {
	return eq(b.parent, expression)
}

func (b *boolInterfaceImpl) NOT_EQ(expression BoolExpression) BoolExpression {
	return notEq(b.parent, expression)
}

func (b *boolInterfaceImpl) IS_DISTINCT_FROM(rhs BoolExpression) BoolExpression {
	return isDistinctFrom(b.parent, rhs)
}

func (b *boolInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs BoolExpression) BoolExpression {
	return isNotDistinctFrom(b.parent, rhs)
}

func (b *boolInterfaceImpl) AND(expression BoolExpression) BoolExpression {
	return newBinaryBoolExpression(b.parent, expression, "AND")
}

func (b *boolInterfaceImpl) OR(expression BoolExpression) BoolExpression {
	return newBinaryBoolExpression(b.parent, expression, "OR")
}

func (b *boolInterfaceImpl) IS_TRUE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS TRUE")
}

func (b *boolInterfaceImpl) IS_NOT_TRUE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS NOT TRUE")
}

func (b *boolInterfaceImpl) IS_FALSE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS FALSE")
}

func (b *boolInterfaceImpl) IS_NOT_FALSE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS NOT FALSE")
}

func (b *boolInterfaceImpl) IS_UNKNOWN() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS UNKNOWN")
}

func (b *boolInterfaceImpl) IS_NOT_UNKNOWN() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS NOT UNKNOWN")
}

//---------------------------------------------------//
type binaryBoolExpression struct {
	expressionInterfaceImpl
	boolInterfaceImpl

	binaryOpExpression
}

func newBinaryBoolExpression(lhs, rhs Expression, operator string) BoolExpression {
	boolExpression := binaryBoolExpression{}

	boolExpression.binaryOpExpression = newBinaryExpression(lhs, rhs, operator)
	boolExpression.expressionInterfaceImpl.parent = &boolExpression
	boolExpression.boolInterfaceImpl.parent = &boolExpression

	return &boolExpression
}

//---------------------------------------------------//
type prefixBoolExpression struct {
	expressionInterfaceImpl
	boolInterfaceImpl

	prefixOpExpression
}

func newPrefixBoolExpression(expression Expression, operator string) BoolExpression {
	exp := prefixBoolExpression{}
	exp.prefixOpExpression = newPrefixExpression(expression, operator)

	exp.expressionInterfaceImpl.parent = &exp
	exp.boolInterfaceImpl.parent = &exp

	return &exp
}

//---------------------------------------------------//
type postfixBoolOpExpression struct {
	expressionInterfaceImpl
	boolInterfaceImpl

	postfixOpExpression
}

func newPostifxBoolExpression(expression Expression, operator string) BoolExpression {
	exp := postfixBoolOpExpression{}
	exp.postfixOpExpression = newPostfixOpExpression(expression, operator)

	exp.expressionInterfaceImpl.parent = &exp
	exp.boolInterfaceImpl.parent = &exp

	return &exp
}

//---------------------------------------------------//

type boolExpressionWrapper struct {
	boolInterfaceImpl
	Expression
}

func newBoolExpressionWrap(expression Expression) BoolExpression {
	boolExpressionWrap := boolExpressionWrapper{Expression: expression}
	boolExpressionWrap.boolInterfaceImpl.parent = &boolExpressionWrap
	return &boolExpressionWrap
}

func BoolExp(expression Expression) BoolExpression {
	return newBoolExpressionWrap(expression)
}
