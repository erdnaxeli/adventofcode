package main

import (
	"log"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type Pulse int

const (
	HighPulse = iota
	LowPulse
)

type PulseMessage struct {
	Source string
	Pulse  Pulse
	Dest   string
}

type Module interface {
	Receive(pm PulseMessage) []PulseMessage
	AddInput(input string)
}

type FlipFlopModule struct {
	outputs []string
	isOn    bool
}

func NewFlipFlopModule(outputs []string) *FlipFlopModule {
	return &FlipFlopModule{isOn: false, outputs: outputs}
}

func (m *FlipFlopModule) Receive(pm PulseMessage) []PulseMessage {
	var result []PulseMessage
	if pm.Pulse == HighPulse {
		return result
	}

	var resultPulse Pulse
	if !m.isOn {
		resultPulse = HighPulse
	} else {
		resultPulse = LowPulse
	}

	m.isOn = !m.isOn
	for _, output := range m.outputs {
		result = append(result, PulseMessage{Pulse: resultPulse, Dest: output})
	}

	return result
}

func (*FlipFlopModule) AddInput(input string) {}

type ConjunctionModule struct {
	inputs  map[string]Pulse
	outputs []string
}

func NewConjunctionModule(outputs []string) *ConjunctionModule {
	return &ConjunctionModule{
		inputs:  make(map[string]Pulse),
		outputs: outputs,
	}
}

func (m *ConjunctionModule) Receive(pm PulseMessage) []PulseMessage {
	m.inputs[pm.Source] = pm.Pulse

	var pulse Pulse = LowPulse
	for _, p := range m.inputs {
		if p == LowPulse {
			pulse = HighPulse
			break
		}
	}

	var result []PulseMessage
	for _, output := range m.outputs {
		result = append(result, PulseMessage{Pulse: pulse, Dest: output})
	}

	return result
}

func (m *ConjunctionModule) AddInput(input string) {
	m.inputs[input] = LowPulse
}

type BroadcastModule struct {
	outputs []string
}

func NewBroadcastModule(outputs []string) BroadcastModule {
	return BroadcastModule{outputs: outputs}
}

func (m BroadcastModule) Receive(pm PulseMessage) []PulseMessage {
	var result []PulseMessage
	for _, output := range m.outputs {
		result = append(result, PulseMessage{Pulse: pm.Pulse, Dest: output})
	}

	return result
}

func (m BroadcastModule) AddInput(input string) {}

type ButtonModule struct{}

func NewButtonModule() ButtonModule {
	return ButtonModule{}
}

func (m ButtonModule) Receive(pm PulseMessage) []PulseMessage {
	return []PulseMessage{{Pulse: LowPulse, Dest: "broadcaster"}}
}

func (ButtonModule) AddInput(input string) {}

func (s solver) Day20p1(input aoc.Input) string {
	modules := make(map[string]Module)
	modules["button"] = NewButtonModule()
	// list of inputs for module not yet created
	modules = ParseModules(modules, input)

	lowPulses := 0
	highPulses := 0
	buttonPushes := 0
	// this is a very memory inefficient FIFO, never do that at home
	var messages []PulseMessage
	for {
		if len(messages) == 0 {
			if buttonPushes == 1000 {
				break
			}

			messages = []PulseMessage{{Dest: "button"}}
			buttonPushes++
		}

		message := messages[0]
		messages = messages[1:]
		// we ignore the pulse sent to the button, as we don't actually sent it a pulse
		if message.Dest != "button" {
			if message.Pulse == LowPulse {
				lowPulses++
			} else {
				highPulses++
			}
		}
		module, ok := modules[message.Dest]
		if !ok {
			// unknown module,
			continue
		}
		newMessages := module.Receive(message)

		for i := range newMessages {
			newMessages[i].Source = message.Dest
		}

		messages = append(messages, newMessages...)
	}

	return aoc.ResultI(lowPulses * highPulses)
}

func (s solver) Day20p2(input aoc.Input) string {
	modules := make(map[string]Module)
	modules["button"] = NewButtonModule()
	// list of inputs for module not yet created
	modules = ParseModules(modules, input)

	buttonPushes := 0
	// this is a very memory inefficient FIFO, never do that at home
	var messages []PulseMessage
	for {
		if len(messages) == 0 {
			messages = []PulseMessage{{Dest: "button"}}
			buttonPushes++
			// log.Print(buttonPushes)
		}

		message := messages[0]
		messages = messages[1:]
		log.Printf("%+v", message)

		if message.Dest == "rx" && message.Pulse == LowPulse {
			break
		}

		module, ok := modules[message.Dest]
		if !ok {
			// unknown module,
			continue
		}
		newMessages := module.Receive(message)

		for i := range newMessages {
			newMessages[i].Source = message.Dest
			log.Printf("\t%+v", newMessages[i])
		}

		messages = append(messages, newMessages...)
	}

	return aoc.ResultI(buttonPushes)
}

func ParseModules(modules map[string]Module, input aoc.Input) map[string]Module {
	inputs := make(map[string][]string)

	for _, line := range input.ToStringSlice() {
		parts := line.SplitOn(" -> ")
		outputs := parts[1].SplitOnS(", ")
		var moduleName string

		if parts[0] == "broadcaster" {
			modules["broadcaster"] = NewBroadcastModule(outputs)
		} else if parts[0][0] == '%' {
			moduleName = string(parts[0][1:])
			modules[moduleName] = NewFlipFlopModule(outputs)
		} else if parts[0][0] == '&' {
			moduleName = string(parts[0][1:])
			modules[moduleName] = NewConjunctionModule(outputs)
		}

		for _, output := range outputs {
			if module, ok := modules[output]; ok {
				module.AddInput(moduleName)
			} else {
				// module was not created yet
				inputs[output] = append(inputs[output], moduleName)
			}
		}
	}

	for moduleName, moduleInputs := range inputs {
		for _, input := range moduleInputs {
			module, ok := modules[moduleName]
			if !ok {
				// unknown module, it means it is a module without any output
				continue
			}
			module.AddInput(input)
		}
	}

	return modules
}
