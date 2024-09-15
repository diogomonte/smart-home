package device

import "github.com/montediogo/home/src/rule"

func HandleDeviceMessage(deviceType string, messageBody interface{}) {
	switch deviceType {
	case DeviceTypePlant:
		handlePlantEvent(messageBody)
	}
}

func handlePlantEvent(messageBody interface{}) {
	rules := rule.Rule{
		Expression: rule.SimpleExpression{
			Field:    "moisture",
			Operator: rule.OperatorGreaterThan,
			Value:    100,
		},
		Action: func() {
			println("action for message")
		},
	}

	ruleEngine := rule.RuleEngine{}
	ruleEngine.AddRule(rules)
	ruleEngine.Run(messageBody)
}
