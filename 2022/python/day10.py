from dataclasses import dataclass, field
from abc import ABC, abstractmethod


class Instruction(ABC):
    cycles: int

    @abstractmethod
    def exec(state, cycle):
        ...


class Noop(Instruction):
    cycles = 1

    def exec(self, state, cycle):
        pass


@dataclass
class AddX(Instruction):
    value: int
    cycles = 2

    def exec(self, state, cycle):
        match cycle:
            case 0:
                pass
            case 1:
                state["X"] += self.value


def read_program():
    program = []
    with open("day10.txt") as file:
        for line in file:
            words = line.strip().split()
            match words:
                case ["noop"]:
                    program.append(Noop())
                case ["addx", value]:
                    program.append(AddX(int(value)))

    return program


def part1(program):
    state = {"X": 1}
    total_signal = 0
    cycle = 0
    for inst in program:
        for inst_cycle in range(inst.cycles):
            cycle += 1
            if cycle in (20, 60, 100, 140, 180, 220):
                total_signal += state["X"] * cycle

            inst.exec(state, inst_cycle)

    return total_signal


def part2(program):
    state = {"X": 1}
    cycle = 0
    lines = []
    for _ in range(6):
        line = []
        for _ in range(40):
            line.append(" ")

        lines.append(line)

    for inst in program:
        for inst_cycle in range(inst.cycles):
            cycle += 1
            y = (cycle - 1) // 40
            x = (cycle - 1) % 40

            if state["X"] - 1 <= x <= state["X"] + 1:
                lines[y][x] = "#"
            else:
                lines[y][x] = "."

            inst.exec(state, inst_cycle)

    return "\n".join("".join(line) for line in lines)


if __name__ == "__main__":
    program = read_program()
    print(part1(program))
    print(part2(program))
