package jet

type FloatExpression interface {
	Expression
	numericExpression

	EQ(rhs FloatExpression) BoolExpression
	NOT_EQ(rhs FloatExpression) BoolExpression
	IS_DISTINCT_FROM(rhs FloatExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs FloatExpression) BoolExpression

	LT(rhs FloatExpression) BoolExpression
	LT_EQ(rhs FloatExpression) BoolExpression
	GT(rhs FloatExpression) BoolExpression
	GT_EQ(rhs FloatExpression) BoolExpression

	ADD(rhs FloatExpression) FloatExpression
	SUB(rhs FloatExpression) FloatExpression
	MUL(rhs FloatExpression) FloatExpression
	DIV(rhs FloatExpression) FloatExpression
	MOD(rhs FloatExpression) FloatExpression
	POW(rhs FloatExpression) FloatExpression
}

type floatInterfaceImpl struct {
	numericExpressionImpl
	parent FloatExpression
}

func (n *floatInterfaceImpl) EQ(rhs FloatExpression) BoolExpression {
	return eq(n.parent, rhs)
}

func (n *floatInterfaceImpl) NOT_EQ(rhs FloatExpression) BoolExpression {
	return notEq(n.parent, rhs)
}

func (n *floatInterfaceImpl) IS_DISTINCT_FROM(rhs FloatExpression) BoolExpression {
	return isDistinctFrom(n.parent, rhs)
}

func (n *floatInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs FloatExpression) BoolExpression {
	return isNotDistinctFrom(n.parent, rhs)
}

func (n *floatInterfaceImpl) GT(rhs FloatExpression) BoolExpression {
	return gt(n.parent, rhs)
}

func (n *floatInterfaceImpl) GT_EQ(rhs FloatExpression) BoolExpression {
	return gtEq(n.parent, rhs)
}

func (n *floatInterfaceImpl) LT(expression FloatExpression) BoolExpression {
	return lt(n.parent, expression)
}

func (n *floatInterfaceImpl) LT_EQ(expression FloatExpression) BoolExpression {
	return ltEq(n.parent, expression)
}

func (n *floatInterfaceImpl) ADD(expression FloatExpression) FloatExpression {
	return newBinaryFloatExpression(n.parent, expression, "+")
}

func (n *floatInterfaceImpl) SUB(expression FloatExpression) FloatExpression {
	return newBinaryFloatExpression(n.parent, expression, "-")
}

func (n *floatInterfaceImpl) MUL(expression FloatExpression) FloatExpression {
	return newBinaryFloatExpression(n.parent, expression, "*")
}

func (n *floatInterfaceImpl) DIV(expression FloatExpression) FloatExpression {
	return newBinaryFloatExpression(n.parent, expression, "/")
}

func (n *floatInterfaceImpl) MOD(expression FloatExpression) FloatExpression {
	return newBinaryFloatExpression(n.parent, expression, "%")
}

func (n *floatInterfaceImpl) POW(expression FloatExpression) FloatExpression {
	return newBinaryFloatExpression(n.parent, expression, "^")
}

//---------------------------------------------------//
type binaryFloatExpression struct {
	expressionInterfaceImpl
	floatInterfaceImpl

	binaryOpExpression
}

func newBinaryFloatExpression(lhs, rhs FloatExpression, operator string) FloatExpression {
	floatExpression := binaryFloatExpression{}

	floatExpression.binaryOpExpression = newBinaryExpression(lhs, rhs, operator)

	floatExpression.expressionInterfaceImpl.parent = &floatExpression
	floatExpression.floatInterfaceImpl.parent = &floatExpression

	return &floatExpression
}

//---------------------------------------------------//

type floatExpressionWrapper struct {
	floatInterfaceImpl
	Expression
}

func newFloatExpressionWrap(expression Expression) FloatExpression {
	floatExpressionWrap := floatExpressionWrapper{Expression: expression}
	floatExpressionWrap.floatInterfaceImpl.parent = &floatExpressionWrap
	return &floatExpressionWrap
}

func FloatExp(expression Expression) FloatExpression {
	return newFloatExpressionWrap(expression)
}
