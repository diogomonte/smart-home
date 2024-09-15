package rule

import "testing"

func TestSimpleExpression_Evaluate(t *testing.T) {
	tests := []struct {
		name     string
		expr     SimpleExpression
		context  map[string]interface{}
		expected bool
	}{
		{
			name: "Test Equal (==) - Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorEquals, Value: 25},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: true,
		},
		{
			name: "Test Equal (==) - No Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorEquals, Value: 30},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: false,
		},
		{
			name: "Test Not Equal (!=) - Match",
			expr: SimpleExpression{Field: "status", Operator: OperatorNotEquals, Value: "married"},
			context: map[string]interface{}{
				"status": "single",
			},
			expected: true,
		},
		{
			name: "Test Greater Than (>) - Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorGreaterThan, Value: 18},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: true,
		},
		{
			name: "Test Greater Than (>) - No Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorGreaterThan, Value: 30},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: false,
		},
		{
			name: "Test Less Than (<) - Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorLessThan, Value: 30},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: true,
		},
		{
			name: "Test Less Than (<) - No Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorLessThan, Value: 18},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: false,
		},
		{
			name: "Test Greater Or Equal (>=) - Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorGreaterOrEqual, Value: 25},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: true,
		},
		{
			name: "Test Greater Or Equal (>=) - No Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorGreaterOrEqual, Value: 30},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: false,
		},
		{
			name: "Test Less Or Equal (<=) - Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorLessOrEqual, Value: 25},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: true,
		},
		{
			name: "Test Less Or Equal (<=) - No Match",
			expr: SimpleExpression{Field: "age", Operator: OperatorLessOrEqual, Value: 20},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: false,
		},
		{
			name: "Test Missing Field - No Match",
			expr: SimpleExpression{Field: "nonexistent", Operator: OperatorEquals, Value: "value"},
			context: map[string]interface{}{
				"age": 25,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.expr.Evaluate(tt.context)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
