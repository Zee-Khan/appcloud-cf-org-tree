package main

import "github.com/fatih/color"

// Text formatting and colors for print text
var (
	bold      = color.New(color.Bold).SprintFunc()
	greenBold = color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()
	cyanBold  = color.New(color.FgHiCyan).Add(color.Bold).SprintFunc()
	redBold   = color.New(color.FgHiRed).Add(color.Bold).SprintFunc()
	blueBold   = color.New(color.FgHiBlue).Add(color.Bold).SprintFunc()
	yellowBold   = color.New(color.FgHiYellow).Add(color.Bold).SprintFunc()
)
