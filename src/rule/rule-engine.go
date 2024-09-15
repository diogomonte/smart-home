package rule

type Expression interface {
	Evaluate(context interface{}) bool
}

type Operator string

const (
	OperatorEquals         Operator = "=="
	OperatorNotEquals      Operator = "!="
	OperatorGreaterThan    Operator = ">"
	OperatorLessThan       Operator = "<"
	OperatorGreaterOrEqual Operator = ">="
	OperatorLessOrEqual    Operator = "<="
)

type SimpleExpression struct {
	Field    string
	Operator Operator
	Value    interface{}
}

func (se SimpleExpression) Evaluate(context interface{}) bool {
	ctx, ok := context.(map[string]interface{})
	if !ok {
		return false
	}

	fieldValue, ok := ctx[se.Field]
	if !ok {
		return false
	}

	switch se.Operator {
	case OperatorEquals:
		return fieldValue == se.Value
	case OperatorNotEquals:
		return fieldValue != se.Value
	case OperatorGreaterThan:
		return fieldValue.(float64) > float64(se.Value.(int))
	case OperatorLessThan:
		return fieldValue.(float64) < se.Value.(float64)
	case OperatorGreaterOrEqual:
		return fieldValue.(float64) >= se.Value.(float64)
	case OperatorLessOrEqual:
		return fieldValue.(float64) <= se.Value.(float64)
	default:
		return false
	}
}

type Rule struct {
	Expression Expression
	Action     func()
}

func (r Rule) EvaluateAndExecute(context interface{}) {
	if r.Expression.Evaluate(context) {
		r.Action()
	}
}

type RuleEngine struct {
	Rules []Rule
}

func (re *RuleEngine) AddRule(rule Rule) {
	re.Rules = append(re.Rules, rule)
}

// Run executes all rules in the engine
func (re RuleEngine) Run(context interface{}) {
	for _, rule := range re.Rules {
		rule.EvaluateAndExecute(context)
	}
}
