import re


def read_input():
    with open("day03.txt") as f:
        content = f.read()

    result = re.findall(r"mul\((\d+),(\d+)\)", content)
    return [(int(m[0]), int(m[1])) for m in result]


def read_input2():
    with open("day03.txt") as f:
        content = f.read()

    result = re.findall(r"mul\((\d+),(\d+)\)|(do)\(\)|(don\'t)\(\)", content)

    do = True
    mult = []
    for match in result:
        if match[2] == "do":
            do = True
            continue
        if match[3] == "don't":
            do = False

        if not do:
            continue

        mult.append((int(match[0]), int(match[1])))

    return mult


def part1(input):
    return sum(a * b for a, b in input)


def part2(input):
    return sum(a * b for a, b in input)


print(part1(read_input()))
print(part2(read_input2()))
