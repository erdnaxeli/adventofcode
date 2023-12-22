package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
	"golang.org/x/exp/maps"
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

	// helpers to print dot graph
	GetPrefix() string
	GetColor() string
	GetOutputs() []string
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

func (*FlipFlopModule) GetPrefix() string {
	return "%"
}

func (m *FlipFlopModule) GetColor() string {
	if m.isOn {
		return "blue"
	} else {
		return "white"
	}
}

func (m *FlipFlopModule) GetOutputs() []string {
	return m.outputs
}

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

func (*ConjunctionModule) GetPrefix() string {
	return "&"
}

func (*ConjunctionModule) GetColor() string {
	return "white"
}

func (m *ConjunctionModule) GetOutputs() []string {
	return m.outputs
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

func (BroadcastModule) GetPrefix() string {
	return ""
}

func (BroadcastModule) GetColor() string {
	return "white"
}

func (m BroadcastModule) GetOutputs() []string {
	return m.outputs
}

type ButtonModule struct{}

func NewButtonModule() ButtonModule {
	return ButtonModule{}
}

func (m ButtonModule) Receive(pm PulseMessage) []PulseMessage {
	return []PulseMessage{{Pulse: LowPulse, Dest: "broadcaster"}}
}

func (ButtonModule) AddInput(input string) {}

func (ButtonModule) GetPrefix() string {
	return ""
}

func (ButtonModule) GetColor() string {
	return "white"
}

func (m ButtonModule) GetOutputs() []string {
	return []string{"broadcaster"}
}

func (s solver) Day20p1(input aoc.Input) string {
	modules := make(map[string]Module)
	modules["button"] = NewButtonModule()
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
	// The graph of connected module is divided in for parts. Each part is a succession
	// of flip flop connected to a conjunction, then to another conjunction (with act
	// like an inverter). The for part are then connected to a conjunction, which is
	// connected to rx.
	//
	// broadcaster ---> some flip flops -> conjunction -> inverter ---> conjunction -> rx
	//              \-> some flip flops -> conjunction -> invertor -/
	//              \-> some flip flops -> conjunction -> invertor -/
	//              \-> some flip flops -> conjunction -> invertor -/
	//
	// For each part we need all the flip flop connected to the first conjunction to
	// switch all in sync from off to on, so they send a high signal. Then the
	// conjunction send a low signal which reset all the flip flop and goes to the
	// inverter, which send a high signal to the last conjunction.
	//
	// In order to send a low signal to rx, we need all parts to send in sync a high
	// signal to the last conjunction.
	//
	// Each of the four parts is arranged in a way so that they all have a different
	// cycle length. Once the cycle length determined, we just have to wait until they
	// are all synchronized.
	r := getLeastCommonMultiple([]int{3739, 3797, 3889, 3761})
	return aoc.ResultF64(r)

	// The following code was only use to print the result of the graph for each of the
	// first 4096 button pushes.
	//
	// To generate a video:
	// * run this code
	// * generate a png file for each dot file:
	//	   parallel -j12 dot -Gsize=9,15\! -Gdpi=200 -Tpng -O {} ::: *.dot
	// * generate a video from those png files:
	//     ffmpeg -framerate 24 -i day20_%04d.dot.png -c:v libx264 -r 30
	//       -pix_fmt yuv420p -vf "drawtext=fontfile=Arial.ttf: text='%{frame_num}':
	//         start_number=1: x=(w-tw)/2: y=h-(2*lh): fontcolor=black: fontsize=20:
	//         box=1: boxcolor=white: boxborderw=5"
	//       day20.mp4

	modules := make(map[string]Module)
	modules["button"] = NewButtonModule()
	modules = ParseModules(modules, input)

	buttonPushes := 0
	// this is a very memory inefficient FIFO, never do that at home
	var messages []PulseMessage
	for {
		if len(messages) == 0 {
			file, _ := os.Create(fmt.Sprintf("day20_%04d.dot", buttonPushes))
			defer file.Close()
			fmt.Fprintln(file, "digraph G {")
			keys := maps.Keys(modules)
			slices.Sort(keys)
			for _, name := range keys {
				module := modules[name]
				fmt.Fprintf(
					file,
					"\t%s [label=\"%s%s\"; style=filled; fillcolor=%s]\n",
					name,
					module.GetPrefix(),
					name,
					module.GetColor(),
				)
				for _, output := range module.GetOutputs() {
					fmt.Fprintf(file, "\t%s -> %s\n", name, output)
				}
			}
			fmt.Fprintln(file, "}")

			messages = []PulseMessage{{Dest: "button"}}
			if buttonPushes == 4096 {
				return ""
			}

			buttonPushes++
			if buttonPushes%100_000 == 0 {
				log.Print(buttonPushes)
			}
		}

		message := messages[0]
		messages = messages[1:]

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
		}

		messages = append(messages, newMessages...)
	}

	return aoc.ResultI(buttonPushes)
}

func ParseModules(modules map[string]Module, input aoc.Input) map[string]Module {
	// list of inputs for module not yet created
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
