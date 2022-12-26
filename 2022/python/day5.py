from collections import namedtuple
from copy import deepcopy

Instruction = namedtuple("Instruction", ["count", "src", "dst"])


def read_stacks(repr: str) -> list[list[str]]:
    lines = repr.split("\n")
    lines.reverse()
    stacks_numbers = lines.pop(0).strip().split(" ")
    stacks_numbers.reverse()
    stacks_count = int(stacks_numbers[0])
    stacks = [[] for _ in range(stacks_count)]

    for line in lines:
        for i in range(0, stacks_count):
            idx = i * 4 + 1
            if line[idx] != " ":
                stacks[i].append(line[idx])

    return stacks


def read_instructions(repr: str) -> list[Instruction]:
    instructions = []
    for line in repr.split("\n"):
        if not line:
            continue

        parts = line.split(" ")
        instructions.append(
            Instruction(
                count=int(parts[1]), src=int(parts[3]) - 1, dst=int(parts[5]) - 1
            )
        )

    return instructions


def read_input() -> tuple[list[list[str]], list[Instruction]]:
    """Return a tuple of crates and instructions."""
    with open("day5.txt") as file:
        content = file.read()

    crates_repr, instructions_repr = content.split("\n\n")
    return (read_stacks(crates_repr), read_instructions(instructions_repr))


def move(stacks: list[list[str]], src: int, dst: int) -> None:
    stacks[dst].append(stacks[src].pop())


def part1(stacks: list[list[str]], instructions: list[Instruction]) -> str:
    for instruction in instructions:
        for _ in range(0, instruction.count):
            move(stacks, instruction.src, instruction.dst)

    return "".join(stack.pop() for stack in stacks)


def part2(stacks: list[list[str]], instructions: list[Instruction]) -> str:
    for instruction in instructions:
        to_move = [stacks[instruction.src].pop() for _ in range(0, instruction.count)]
        to_move.reverse()
        stacks[instruction.dst].extend(to_move)

    return "".join(stack.pop() for stack in stacks)


if __name__ == "__main__":
    stacks, instructions = read_input()
    print(part1(deepcopy(stacks), instructions))
    print(part2(deepcopy(stacks), instructions))
