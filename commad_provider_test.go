package command_provider_test

import (
	"errors"
	"github.com/donutloop/command-provider"
	"testing"
)

func TestCommandProviderWithOneCommand(t *testing.T) {

	buildCommand := func(test string) command_provider.CommandWrapper {
		return command_provider.CommandWrapper{
			Command: func(parameters command_provider.Parameters) (interface{}, command_provider.Warning, error) {
				return test, nil, nil
			},
		}
	}

	commandProvider := command_provider.New(5, buildCommand("Hello World"))

	data, _, _ := commandProvider.Execute()

	for _, value := range data {
		if text, ok := value.(string); !ok || text != "Hello World" {
			t.Errorf("Expected: \" Hello World\" got %v", text)
		}
	}
}

func TestCommandProviderWithMultiCommand(t *testing.T) {

	buildCommand := func(test string) command_provider.CommandWrapper {
		return command_provider.CommandWrapper{
			Command: func(parameters command_provider.Parameters) (interface{}, command_provider.Warning, error) {
				return test, nil, nil
			},
		}
	}

	commandProvider := command_provider.New(5, buildCommand("Hello World"), buildCommand("Hello World"), buildCommand("Hello World"))

	data, _, _ := commandProvider.Execute()

	for _, value := range data {
		if text, ok := value.(string); !ok || text != "Hello World" {
			t.Errorf("Expected: \" Hello World\" got %v", text)
		}
	}
}

func TestCommandProviderWithMultiCommandWarnings(t *testing.T) {

	buildCommand := func(test string) command_provider.CommandWrapper {
		return command_provider.CommandWrapper{
			Command: func(parameters command_provider.Parameters) (interface{}, command_provider.Warning, error) {
				return test, command_provider.Warning(errors.New("test")), nil
			},
		}
	}

	commandProvider := command_provider.New(10, buildCommand("Hello World"), buildCommand("Hello World"))

	_, warnings, _ := commandProvider.Execute()

	if len(warnings) != 2{
		t.Errorf("Expected: 3 times test  got %v", warnings)
	}
}
