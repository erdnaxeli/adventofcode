from functools import cache


@cache
def read_input():
    equations = []
    with open("day07.txt") as f:
        for line in f:
            splitted = line.rstrip().split()
            equations.append([int(splitted[0][:-1]), *[int(s) for s in splitted[1:]]])

    return equations


def is_correct(operators, result, values, acc):
    if acc > result:
        return False

    if not values:
        return result == acc

    for op in operators:
        if is_correct(operators, result, values[1:], op(acc, values[0])):
            return True

    return False


def find_correct(operators, equations):
    return sum(
        equation[0]
        for equation in equations
        if is_correct(operators, equation[0], equation[1:], 0)
    )


print(
    find_correct(
        [lambda a, b: a + b, lambda a, b: a * b],
        read_input(),
    )
)
print(
    find_correct(
        [lambda a, b: a + b, lambda a, b: a * b, lambda a, b: int(f"{a}{b}")],
        read_input(),
    )
)
