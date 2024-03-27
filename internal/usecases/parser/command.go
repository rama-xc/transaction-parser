package parser

type Command interface {
	Execute()
}

type RunCommand struct {
	prsr IHistory
}

func NewRunCommand(prsr IHistory) *RunCommand {
	return &RunCommand{prsr: prsr}
}

func (c *RunCommand) Execute() {
	c.prsr.run()
}

//type OptionsCommand struct {
//	prsr IHistory
//
//	wrks int
//}
//
//func NewOptionsCommand(prsr IHistory, wrks int) *OptionsCommand {
//	return &OptionsCommand{prsr: prsr, wrks: wrks}
//}
//
//func (c *OptionsCommand) Execute() {
//	c.prsr.options(
//		c.wrks,
//	)
//}
