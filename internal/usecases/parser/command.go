package parser

type Command interface {
	Execute()
}

type StartCommand struct {
	prsr IHistory
}

func NewStartCommand(prsr IHistory) *StartCommand {
	return &StartCommand{prsr: prsr}
}

func (c *StartCommand) Execute() {
	c.prsr.Start()
}
