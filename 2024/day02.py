import itertools
from functools import cache


@cache
def read_input():
    def it():
        with open("day02.txt") as f:
            for line in f:
                yield [int(x) for x in line.split(" ")]

    return list(it())


def is_safe(report):
    sign = 0
    for l1, l2 in itertools.pairwise(report):
        if sign == 0:
            sign = l1 - l2
        elif (l1 - l2) * sign < 0:
            return False

        if l1 == l2:
            return False
        if abs(l1 - l2) > 3:
            return False

    return True


def is_safe_tolerate(report):
    if is_safe(report):
        return True

    for idx in range(0, len(report)):
        if is_safe([lvl for i, lvl in enumerate(report) if i != idx]):
            return True

    return False


def part1(input):
    safe_reports = [report for report in input if is_safe(report)]
    return len(safe_reports)


def part2(input):
    safe_reports = [report for report in input if is_safe_tolerate(report)]
    return len(safe_reports)


print(part1(read_input()))
print(part2(read_input()))
