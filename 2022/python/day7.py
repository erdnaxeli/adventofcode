from __future__ import annotations

import sys
from collections import defaultdict, deque
from dataclasses import dataclass, field
from functools import cached_property
from itertools import chain


@dataclass
class File:
    size: int


@dataclass
class Directory:
    files: dict[str, File] = field(default_factory=dict)
    subdirectories: dict[str, Directory] = field(
        default_factory=lambda: defaultdict(Directory)
    )

    @cached_property
    def size(self) -> int:
        return sum(
            file.size
            for file in chain(self.files.values(), self.subdirectories.values())
        )


def read_commands(root: Directory, commands: list[str]) -> None:
    directory = root
    path = []
    for command in commands:
        if command.startswith("cd "):
            filename = command[3:]
            if filename == "..":
                directory = path.pop()
            else:
                path.append(directory)
                directory = directory.subdirectories[filename]
        elif command.startswith("ls"):
            result = command.split("\n")
            for line in result[1:]:
                if line.startswith("dir "):
                    directory.subdirectories[line[4:]] = Directory()
                else:
                    size, name = line.split(" ")
                    directory.files[name] = File(int(size))


def read_input() -> Directory:
    with open("day7.txt") as file:
        commands = file.read().strip().split("\n$ ")

    if commands[0] != "$ cd /":
        raise ValueError()
    else:
        root = Directory()
        read_commands(root, commands[1:])
        return root


def part1(root: Directory) -> int:
    directories = []
    to_scan = deque()
    to_scan.append(root)

    while to_scan:
        directory = to_scan.pop()
        if directory.size <= 100000:
            directories.append(directory)

        to_scan.extend(directory.subdirectories.values())

    return sum(directory.size for directory in directories)


def part2(root: Directory) -> int:
    directories = []
    to_scan = deque()
    to_scan.append(root)
    to_free = 30000000 - (70000000 - root.size)

    while to_scan:
        directory = to_scan.pop()
        if directory.size >= to_free:
            directories.append(directory.size)

        to_scan.extend(directory.subdirectories.values())

    return min(directories)


if __name__ == "__main__":
    sys.setrecursionlimit(2000)
    root = read_input()
    print(part1(root))
    print(part2(root))
