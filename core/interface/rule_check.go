package _interface

type RuleCheckInterface interface {
	Check(RuleCode string) bool
}
