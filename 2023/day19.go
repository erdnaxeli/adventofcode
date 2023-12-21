package main

import (
	"regexp"
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type MachinePart struct {
	x, m, a, s int
}

func (m MachinePart) getCategory(c string) int {
	switch c {
	case "x":
		return m.x
	case "m":
		return m.m
	case "a":
		return m.a
	case "s":
		return m.s
	default:
		panic("unknown category")
	}
}

type WorkflowIteration struct {
	machinePartRange MachinePartRange
	nextWorkflow     string
}

type MachinePartRange struct {
	xMin, mMin, aMin, sMin int
	xMax, mMax, aMax, sMax int
}

type Operator int

const (
	Default Operator = iota
	Gt
	Lt
)

type Rule struct {
	category string
	operator Operator
	value    int
	dest     string
}

var (
	MACHINE_PART_RGX = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
	WORKFLOW_RGX     = regexp.MustCompile(`(.+){(.*)}`)
	RULE_RGX         = regexp.MustCompile(`(.+)(<|>)(\d+):(.+)`)
)

func (s solver) Day19p1(input aoc.Input) string {
	parts := input.Delimiter("\n\n").ToStringSlice()
	var machineParts []MachinePart
	workflows := make(map[string][]Rule)

	for _, line := range parts[0].SplitOn("\n") {
		match := WORKFLOW_RGX.FindStringSubmatch(string(line))
		name := match[1]
		rules := strings.Split(match[2], ",")

		for _, rule := range rules[:len(rules)-1] {
			match := RULE_RGX.FindStringSubmatch(rule)
			var operator Operator
			if match[2] == ">" {
				operator = Gt
			} else {
				operator = Lt
			}

			workflows[name] = append(workflows[name], Rule{
				category: match[1],
				operator: operator,
				value:    aoc.Atoi(match[3]),
				dest:     match[4],
			})
		}

		workflows[name] = append(workflows[name], Rule{dest: rules[len(rules)-1]})
	}

	for _, line := range parts[1].SplitOn("\n") {
		match := MACHINE_PART_RGX.FindStringSubmatch(string(line))
		machineParts = append(machineParts, MachinePart{
			x: aoc.Atoi(match[1]),
			m: aoc.Atoi(match[2]),
			a: aoc.Atoi(match[3]),
			s: aoc.Atoi(match[4]),
		})
	}

	var accepted []MachinePart
	for _, part := range machineParts {
		workflow := "in"

		for {
			for _, rule := range workflows[workflow] {
				var success bool
				if rule.category == "" {
					// default case
					success = true
				} else {
					success = ApplyWorkflowOperator(part.getCategory(rule.category), rule.operator, rule.value)
				}

				if success {
					switch rule.dest {
					case "A":
						accepted = append(accepted, part)
						workflow = ""
					case "R":
						workflow = ""
					default:
						workflow = rule.dest
					}

					break
				}
			}

			if workflow == "" {
				break
			}
		}
	}

	sum := 0
	for _, p := range accepted {
		sum += p.x + p.m + p.a + p.s
	}

	return aoc.ResultI(sum)
}

func (s solver) Day19p2(input aoc.Input) string {
	parts := input.Delimiter("\n\n").ToStringSlice()
	workflows := make(map[string][]Rule)

	for _, line := range parts[0].SplitOn("\n") {
		match := WORKFLOW_RGX.FindStringSubmatch(string(line))
		name := match[1]
		rules := strings.Split(match[2], ",")

		for _, rule := range rules[:len(rules)-1] {
			match := RULE_RGX.FindStringSubmatch(rule)
			var operator Operator
			if match[2] == ">" {
				operator = Gt
			} else {
				operator = Lt
			}

			workflows[name] = append(workflows[name], Rule{
				category: match[1],
				operator: operator,
				value:    aoc.Atoi(match[3]),
				dest:     match[4],
			})
		}

		workflows[name] = append(workflows[name], Rule{dest: rules[len(rules)-1]})
	}

	queue := []WorkflowIteration{
		{
			machinePartRange: MachinePartRange{
				xMin: 1,
				mMin: 1,
				aMin: 1,
				sMin: 1,
				xMax: 4000,
				mMax: 4000,
				aMax: 4000,
				sMax: 4000,
			},
			nextWorkflow: "in",
		},
	}
	var result []MachinePartRange

	for len(queue) > 0 {
		wi := queue[len(queue)-1]
		if r := wi.machinePartRange; r.xMin == 0 || r.mMax == 0 || r.aMin == 0 || r.sMin == 0 {
			// invalid range
			continue
		}

		queue = queue[:len(queue)-1]
		rules := workflows[wi.nextWorkflow]

		for _, rule := range rules {
			// the new WorkflowIteration if the rule succeeds
			successWi := wi
			// the new WorkflowIteration if the rule fails
			failWi := wi

			switch rule.category {
			case "x":
				sMin, sMax, fMin, fMax := getNewCategoryRange(wi.machinePartRange.xMin, wi.machinePartRange.xMax, rule)
				successWi.machinePartRange.xMin = sMin
				successWi.machinePartRange.xMax = sMax
				failWi.machinePartRange.xMin = fMin
				failWi.machinePartRange.xMax = fMax

			case "m":
				sMin, sMax, fMin, fMax := getNewCategoryRange(wi.machinePartRange.mMin, wi.machinePartRange.mMax, rule)
				successWi.machinePartRange.mMin = sMin
				successWi.machinePartRange.mMax = sMax
				failWi.machinePartRange.mMin = fMin
				failWi.machinePartRange.mMax = fMax
			case "a":
				sMin, sMax, fMin, fMax := getNewCategoryRange(wi.machinePartRange.aMin, wi.machinePartRange.aMax, rule)
				successWi.machinePartRange.aMin = sMin
				successWi.machinePartRange.aMax = sMax
				failWi.machinePartRange.aMin = fMin
				failWi.machinePartRange.aMax = fMax
			case "s":
				sMin, sMax, fMin, fMax := getNewCategoryRange(wi.machinePartRange.sMin, wi.machinePartRange.sMax, rule)
				successWi.machinePartRange.sMin = sMin
				successWi.machinePartRange.sMax = sMax
				failWi.machinePartRange.sMin = fMin
				failWi.machinePartRange.sMax = fMax
			default:
				// default rule, nothing to do
			}

			switch rule.dest {
			case "A":
				result = append(result, successWi.machinePartRange)
			case "R":
				// successWi is rejected
				// failWi continues to the next rule
			default:
				successWi.nextWorkflow = rule.dest
				failWi.nextWorkflow = rule.dest
				queue = append(queue, successWi)
			}

			wi = failWi
		}
	}

	sum := 0
	for _, mpRange := range result {
		sum += (1 + mpRange.xMax - mpRange.xMin) *
			(1 + mpRange.mMax - mpRange.mMin) *
			(1 + mpRange.aMax - mpRange.aMin) *
			(1 + mpRange.sMax - mpRange.sMin)
	}

	return aoc.ResultI(sum)
}

func ApplyWorkflowOperator(l int, op Operator, r int) bool {
	switch op {
	case Gt:
		return l > r
	case Lt:
		return l < r
	default:
		panic("called with default operator")
	}
}

func getNewCategoryRange(rMin int, rMax int, rule Rule) (int, int, int, int) {
	switch rule.operator {
	case Gt:
		if rMax < rule.value {
			return 0, 0, rMin, rMax
		} else if rMin > rule.value {
			return rMin, rMax, 0, 0
		} else {
			return rule.value + 1, rMax, rMin, rule.value
		}
	case Lt:
		if rMin > rule.value {
			return 0, 0, rMin, rMax
		} else if rMax < rule.value {
			return rMin, rMax, 0, 0
		} else {
			return rMin, rule.value - 1, rule.value, rMax
		}
	case Default:
		return rMin, rMax, 0, 0
	default:
		panic("???")
	}
}
