package command

type Command struct {
	command  string
	params   []Param
	argument string
}
type Param struct {
	name  string
	value string
	full  bool
}

func New(command string) *Command {
	return &Command{
		command: command,
	}
}

func (c *Command) AddShortParam(name, value string) *Command {
	c.AddParam(name, value, false)
	return c
}

func (c *Command) AddFullParam(name, value string) *Command {
	c.AddParam(name, value, true)
	return c
}

func (c *Command) AddParam(name, value string, full bool) *Command {
	c.params = append(c.params, Param{
		name:  name,
		value: value,
		full:  full,
	})
	return c
}

func (c *Command) Argument(argument string) *Command {
	c.argument = argument
	return c
}

func (c *Command) Build() string {
	command := c.command + " "
	for _, param := range c.params {
		if param.full {
			command += "--" + param.name + "=" + param.value + " "
		} else {
			command += "-" + param.name + "=" + param.value + " "
		}
	}
	command += c.argument
	return command
}
