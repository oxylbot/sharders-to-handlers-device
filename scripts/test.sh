#!/bin/sh

# Get the linter
go get -u golang.org/x/lint/golint

# Execute the linter on the project (by default lints the project)
golint
