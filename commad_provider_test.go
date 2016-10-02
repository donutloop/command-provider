package command_provider_test

import (
	"testing"
	"github.com/donutloop/command-provider"
)

func TestCommandProviderWithOneCommand(t *testing.T) {

	buildCommand := func(test string) command_provider.CommandWrapper{
		return command_provider.CommandWrapper{
			Command:func(parameters command_provider.Parameters) (interface{}, command_provider.Warning, error) {
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

	buildCommand := func(test string) command_provider.CommandWrapper{
		return command_provider.CommandWrapper{
			Command:func(parameters command_provider.Parameters) (interface{}, command_provider.Warning, error) {
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
