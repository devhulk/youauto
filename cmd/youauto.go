package main

import (
	"fmt"
	"log"
	"os"

	"github.com/devhulk/youauto/commands"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("youauto", "0.1.0")
	c.Args = os.Args[1:]
	// Write to $GDRIVE_YOUTUBE_DIR
	c.Commands = map[string]cli.CommandFactory{
		"create": func() (cli.Command, error) {
			home := os.Getenv("HOME")
			fmt.Println(home)
			project := fmt.Sprintf("%v/videos/%v", home, c.Args[1])
			fmt.Println(project)
			err := os.MkdirAll(project, 0755)
			if err != nil {
				log.Fatal(err)
			}
			os.Mkdir(fmt.Sprintf("%v/a", project), 0755)
			os.Mkdir(fmt.Sprintf("%v/b", project), 0755)
			os.Mkdir(fmt.Sprintf("%v/graphics", project), 0755)
			return &commands.Create{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

}
