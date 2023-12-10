# Project 2: Shell Builtins

## Description

For this project I added commands to a simple shell. 

The shell is already written, but I chose five (5) shell builtins (or shell-adjacent) commands to rewrite into Go, and integrate into the Go shell.

Two shell builtins were already written:

- `cd`
- `env`

## Builtins

I wrote the alias, cat, echo, help, and pwd functions.

The test coverage hits exactly 76% as required by the assignment. this is shown in the coverage.html report. I got this report by running 
- go test -coverprofile=coverage.out
- go tool cover -html=coverage.out
