[![Build Status](https://travis-ci.org/donutloop/command-provider.svg?branch=master)](https://travis-ci.org/donutloop/command-provider)

# Command provider

## Usage

This is just a quick introduction

Let's start with a trivial example:

```go
    package main

    import (
        "github.com/donutloop/command-provider"
    )

	buildCommand := func(text string) command_provider.CommandWrapper {
		return command_provider.CommandWrapper{
			Command: func(parameters command_provider.Parameters) (interface{}, command_provider.Warning, error) {
				return text, nil, nil
			},
		}
	}

	commandProvider := command_provider.New(5, buildCommand("Hello World"), buildCommand("Hello World"), buildCommand("Hello World"))

	data, warnings, errors := commandProvider.Execute()
```