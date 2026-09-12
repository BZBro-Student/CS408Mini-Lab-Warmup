# Project #: Lab Warmup

* Author: Broden Benson
* Class: CS408
* Semester: Fall 2026

## Overview

On an easy to view low strain background, displays Grades in an easy to
view format and displays module assignments to allow for quick assesment
of completion. This is useful as Canvas navigation tends to be obtuse and,
on the list of bad things instructure has done, not including a dark mode is 
in my opinion the worst one.

## Initial Setup

Install Go: https://go.dev/

Once Go is installed install the needed tools using:

go install github.com/a-h/templ/cmd/templ@latest
go install github.com/air-verse/air@latest

Install dependencies using:
go mod download

Install TailwindCSS application and add it to your system path:
https://tailwindcss.com/blog/standalone-cli

Setup a .env file in the root directory using the provided example as a guide



## Compiling and Using

For a dev run use the command:

make dev

For a standard run use the command:

make run