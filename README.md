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

## API Endpoints Used

[university domain]/api/v1/users/self/favorites/courses

[university domain]/api/v1/courses/[course ID]/modules

## Reflection

Full Stack is hard. Especially using tools that I have never used in the past, the amount of docs reading
I did could probably encompass all of the docs reading I've done in the first 2 years of my CS degree all 
condensed over the long labor day weekend. Go is new to me with a very basic introduction coming from my
programming languages class, but it is super cool (C without the headache of C). Templ, the templating system
I used is also incredibly cool since instead of straight up serving the HTML page to the client it compiles to native Go 
code can be used to selectively render parts of the site at any given time. HTMX in combination with this
allows for super snappy and lightweight interactivity without needing to get JavaScript involved. All of this
comes with a steep learning curve, but I feel as though it gave me a pretty good understanding of the structure 
of the stack. 

On top of learning more about my stack I also got to get some hands on API usage experience which is something I 
have not had to use before. The code is a little dense so I used some AI generated boilerplate but the structure 
between the requests look to be pretty similar so I think I could adapt the code I have to work with other APIs I 
may use in the future. One thing that was hard was pulling my return data into my templ templates. Templ code variables
are immutable anything passed in stays that way so any data that needs parsed or modified has to be done before being handed
to a temple method. Learning this will be useful undoubtedly.

If I had more time my visual presentation would be something I work more on. UI design is not my strong-suit especially 
with a new system, maybe some more graphical flare in the progress section, but I worry that may be too much? I also would
want to work on overall file structure. I made a basic goth template from projects I had seen so a lot of files are empty
vestigial structure.
