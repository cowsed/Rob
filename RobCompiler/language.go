package main

type RobotProgram struct {
	inputs        []Input
	outputs       []Output
	subscriptions []Subscription
	init          FunctionDefinition
	update        FunctionDefinition
}

type FunctionDefinition struct {
	name      string
	signature []Type
}

type Subscription struct{}
type Input struct{}
type Output struct{}
