package commands

// Create - create a new youauto project
type Create struct{}

func (c *Create) Run(args []string) int {
	return 0
}

func (c *Create) Synopsis() string {
	return "Create a new youauto project."
}

func (c *Create) Help() string {
	return "<youauto> <command> [<args>]\n"
}
