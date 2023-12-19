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
	return ""
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
