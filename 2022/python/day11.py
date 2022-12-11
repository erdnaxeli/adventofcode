from dataclasses import dataclass
from collections import defaultdict
from math import floor, prod
from typing import Callable


@dataclass
class Monkey:
    items: list[int]
    operation: Callable[[int], int]
    test_value: int
    test_true: int
    test_false: int

    def run(self, divide=True):
        for item in self.items:
            item = self.operation(item)

            if divide:
                item //= 3

            if item % self.test_value == 0:
                yield item, self.test_true
            else:
                yield item, self.test_false

        self.items = []

    def add(self, item):
        self.items.append(item)


def read_monkeys():
    monkeys = []
    for monkey_str in read_monkeys_file():
        monkeys.append(read_monkey(monkey_str))

    return monkeys


def read_monkeys_file():
    with open("day11.txt") as file:
        return file.read().strip().split("\n\n")


def read_monkey(description):
    lines = description.split("\n")
    items = [int(i) for i in lines[1][len("  Starting items: ") :].split(", ")]
    operation_body = lines[2][len("  Operation: new = ") :]
    operation = eval(f"lambda old: {operation_body}")
    test_value = int(lines[3][len("  Test: divisible by ") :])
    test_true = int(lines[4][len("    If true: throw to monkey ") :])
    test_false = int(lines[5][len("    If false: throw to monkey ") :])

    monkey = Monkey(items, operation, test_value, test_true, test_false)
    return monkey


def part1(monkeys):
    inspections = defaultdict(lambda: 0)
    for _ in range(20):
        for idx, monkey in enumerate(monkeys):
            for item, dst in monkey.run():
                inspections[idx] += 1
                monkeys[dst].add(item)

    return prod(sorted(inspections.values())[-2:])


def part2(monkeys):
    inspections = defaultdict(lambda: 0)
    divider = prod(m.test_value for m in monkeys)
    print(divider)

    for _ in range(1000):
        for idx, monkey in enumerate(monkeys):
            for item, dst in monkey.run(divide=False):
                inspections[idx] += 1

                monkeys[dst].add(item)

    return prod(sorted(inspections.values())[-2:])


if __name__ == "__main__":
    monkeys = read_monkeys()
    print(part1(monkeys))
    print(part2(monkeys))
