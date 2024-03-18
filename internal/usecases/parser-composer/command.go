package prsrcmps

import "transaction-parser/internal/usecases/parser-composer/parser"

type ComposerControl struct {
	command Command
}

func NewComposerControl(command Command) *ComposerControl {
	return &ComposerControl{command: command}
}

func (c *ComposerControl) Run() {
	c.command.execute()
}

type Command interface {
	execute()
}

type RunCommand struct {
	parsersOpts map[parser.Type]int
	cmpr        IParserComposer
}

func NewRunCommand(parsersOpts map[parser.Type]int, cmpr IParserComposer) *RunCommand {
	return &RunCommand{
		parsersOpts: parsersOpts,
		cmpr:        cmpr,
	}
}

func (c *RunCommand) execute() {
	c.cmpr.run()
}
