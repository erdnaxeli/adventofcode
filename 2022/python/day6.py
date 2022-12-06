from collections import Counter
from typing import Optional

def read_buffer() -> str:
    with open("day6.txt") as file:
        return file.read().strip()


def part1(buffer: str) -> Optional[int]:
    for i in range(4, len(buffer)):
        if len(Counter(buffer[i-4:i])) == 4:
            return i


def part2(buffer: str) -> Optional[int]:
    for i in range(14, len(buffer)):
        if len(Counter(buffer[i-14:i])) == 14:
            return i


if __name__ == '__main__':
    buffer = read_buffer()
    print(part1(buffer))
    print(part2(buffer))