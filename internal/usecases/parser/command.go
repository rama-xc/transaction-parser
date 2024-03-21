package parser

type Command interface {
	Execute()
}

type StartCommand struct {
	prsr IHistory
	resp chan Ping
}

func NewStartCommand(prsr IHistory, resp chan Ping) *StartCommand {
	return &StartCommand{prsr: prsr, resp: resp}
}

func (c *StartCommand) Execute() {

	c.prsr.start(
		c.resp,
	)

}