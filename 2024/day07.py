from functools import cache


@cache
def read_input():
    equations = []
    with open("day07.txt") as f:
        for line in f:
            splitted = line.rstrip().split()
            equations.append([int(splitted[0][:-1]), *[int(s) for s in splitted[1:]]])

    return equations


def compute_equation(result, acc, values, debug):
    if not values:
        return result == acc

    return compute_equation(
        result, acc + values[0], values[1:], f"{debug}+{values[0]}"
    ) or compute_equation(result, acc * values[0], values[1:], f"{debug}*{values[1:]}")


def part1(equations):
    total = 0
    for equation in equations:
        result, *values = equation
        if compute_equation(result, values[0], values[1:], f"{values[0]}"):
            total += result

    return total


def compute_equation2(result, acc, values, debug):
    if acc > result:
        return False
    
    if not values:
        return result == acc

    return compute_equation2(
        result, int(f"{acc}{values[0]}"), values[1:], f"{debug} || {values[0]}"
    ) or compute_equation2(
        result, acc + values[0], values[1:], f"{debug} + {values[0]}"
    ) or compute_equation2(result, acc * values[0], values[1:], f"{debug} * {values[0]}")


def part2(equations):
    total = 0
    for equation in equations:
        result, *values = equation
        if compute_equation2(result, values[0], values[1:], f"{values[0]}"):
            total += result

    return total


print(part1(read_input()))
print(part2(read_input()))
