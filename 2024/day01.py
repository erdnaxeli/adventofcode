from collections import Counter
from functools import cache


@cache
def read_input():
    l1, l2 = [], []
    with open("day01.txt") as f:
        for line in f:
            id1, *_, id2 = line.rstrip().split(" ")
            l1.append(int(id1))
            l2.append(int(id2))

    return l1, l2


def part1():
    l1, l2 = read_input()
    return sum(abs(id1 - id2) for id1, id2 in zip(sorted(l1), sorted(l2)))


def part2():
    l1, l2 = read_input()
    counts = Counter(l2)
    return sum(id * counts.get(id, 0) for id in l1)


print(part1())
print(part2())
