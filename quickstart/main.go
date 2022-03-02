package main

import (
	demo "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func main() {
	d := demo.New()
	d.Add(updatecliDemo(), "Updatecli QuickStart", "Quick Start demo")
	d.Setup(setup)
	d.Run()
}

// To execute before Run()
func setup(ctx *cli.Context) error {
	return demo.Ensure(
		"clear",
	)
}

// updatecliDemo is our quick start demo
func updatecliDemo() *demo.Run {
	// A new run contains a title and an optional description
	r := demo.NewRun(
		"Updatecli QuickStart",
		"This video executes quickstart instruction",
		"as described from",
		"https://www.updatecli.io/docs/prologue/quick-start/",
	)

	// A single step can consist of a description and a command to be executed
	r.Step(demo.S(
		"Updatecli's purpose is to let you decide",
		"how files should be automatically updated,",
		"and update them when needed,",
		"each time updatecli is executed"), nil)

	r.Step(demo.S(
		"Let's see the data file that we need to update",
	), demo.S(
		"cat data.yaml",
	))

	r.Step(demo.S(
		"Within this file, we notice a yaml key `container.tag`",
		"set to the latest Jenkins weekly version.",
		"Now let's see what our updatecli manifest looks like",
		"and how it will update our data file",
	), demo.S(
		"cat manifest.yaml",
	))

	r.Step(demo.S(
		"Here we have three major categories of information.",
		"`sources` that define where information is coming from.",
		"`conditions` that define requirements before updating our targets.",
		"`targets`: which define what will be updated.",
		"",
		"Each resource has a `kind` that defines the behavior of the resource",
		"And a `spec` that defines parameters for a specific resource `kind`",
	), nil)

	r.Step(demo.S(
		"In our example, we defined:",
		"One source of `kind` 'Jenkins' with the parameter `release` set 'weekly.'",
		"Differently said, we are looking for the latest Jenkins weekly version.",
		"",
		"One condition of kind `dockerImage` where we want to test",
		"that the docker image `jenkins/jenkins:<latest weekly version>` exists",
		"",
		"One target of kind `yaml` where want to test",
		"that the yaml named 'data.yaml', has a key 'container.tag'",
		"set to the latest Jenkins weekly version",
		"",
	), nil)

	r.Step(demo.S(
		"Let's now see if something needs to be changed",
	), demo.S(
		"updatecli diff --config manifest.yaml",
	))

	r.Step(demo.S(
		"Updatecli tells us that an update is availabe",
		"All we have to do is to execute updatecli",
		"once again but in apply mode",
		"updatecli apply --config manifest.yaml",
		"",
	), nil)

	r.Step(demo.S(
		"You can now go further by:",
		"Looking at available plugins: https://www.updatecli.io/plugins/",
		"Digging into core concepts: https://www.updatecli.io/docs/core/",
		"Helping us growing updatecli: https://github.com/updatecli/updatecli",
		"",
		"Thank you for watching this ❤️ ",
		"",
	), nil)

	return r
}
