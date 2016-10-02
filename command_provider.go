package command_provider

import "time"
import "errors"

type Warning error

type Parameters map[string]interface{}

type Command func(parameters Parameters) (interface{}, Warning, error)

type CommandWrapper struct {
	Command    Command
	Parameters Parameters
}

func New(timeForCommand time.Duration, commandWrapper ...CommandWrapper) CommandProvider {
	return CommandProvider{
		commandWrappers: append([]CommandWrapper{}, commandWrapper...),
		TimeForCommand:  timeForCommand,
	}
}

type CommandProvider struct {
	commandWrappers []CommandWrapper
	TimeForCommand  time.Duration
}

func (cp *CommandProvider) Execute() (data []interface{}, warnings []Warning, errs []error) {

	dataChan := make(chan interface{}, len(cp.commandWrappers))
	warChan := make(chan Warning, len(cp.commandWrappers))
	errChan := make(chan error)

	for _, commandWrapper := range cp.commandWrappers {
		go func(commandWrapper CommandWrapper) {
			data, war, err := commandWrapper.Command(commandWrapper.Parameters)

			if err != nil {
				errChan <- err
				return
			}

			if war != nil {
				warChan <- war
				return
			}

			dataChan <- data
		}(commandWrapper)
	}

	for i := 0; i < len(cp.commandWrappers); i++ {
		select {
		case ReceivedData := <-dataChan:
			data = append(data, ReceivedData)
		case ReceivedWar := <-warChan:
			warnings = append(warnings, ReceivedWar)
		case ReceivedErr := <-errChan:
			errs = append(errs, ReceivedErr)
		case <-time.After(time.Second * cp.TimeForCommand):
			errs = append(errs, errors.New("Command timeout"))
		}
	}

	dataChan = nil
	warChan = nil
	errChan = nil

	return
}
