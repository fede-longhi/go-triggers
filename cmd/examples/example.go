package main

import "github.com/fede-longhi/go-triggers/triggers"

func main() {
	t := triggers.NewTrigger("test")

	printAction := triggers.NewPrintAction()
	printAction.Message = "test"
	t.AddAction(printAction)

	c := triggers.NewCompareCondition(">", 10.0)
	t.SetCondition(c)

	e := triggers.NewFloatEvent(12.0)

	t.Update(*e)
}
