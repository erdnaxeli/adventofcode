from collections import defaultdict
from functools import cache


@cache
def read_input():
    grid = defaultdict(lambda: "")
    with open("day04.txt") as f:
        x = 0
        for row in f:
            y = 0
            for column in row:
                grid[(x, y)] = column
                y += 1

            x += 1

    return grid


def check_xmas(input, p1, p2, p3, p4):
    word = "".join((input[p1], input[p2], input[p3], input[p4]))
    return word == "XMAS"


def get_all_directions(x, y):
    return (
        ((x, y), (x - 1, y), (x - 2, y), (x - 3, y)),
        ((x, y), (x + 1, y), (x + 2, y), (x + 3, y)),
        ((x, y), (x, y - 1), (x, y - 2), (x, y - 3)),
        ((x, y), (x, y + 1), (x, y + 2), (x, y + 3)),
        (
            (x, y),
            (x + 1, y + 1),
            (x + 2, y + 2),
            (x + 3, y + 3),
        ),
        (
            (x, y),
            (x + 1, y - 1),
            (x + 2, y - 2),
            (x + 3, y - 3),
        ),
        (
            (x, y),
            (x - 1, y + 1),
            (x - 2, y + 2),
            (x - 3, y + 3),
        ),
        (
            (x, y),
            (x - 1, y - 1),
            (x - 2, y - 2),
            (x - 3, y - 3),
        ),
    )


def part1(input):
    xmax, ymax = max(input.keys())
    count = 0
    for x in range(xmax + 1):
        for y in range(ymax + 1):
            if input[(x, y)] != "X":
                continue

            for p1, p2, p3, p4 in get_all_directions(x, y):
                if check_xmas(input, p1, p2, p3, p4):
                    count += 1

    return count


def part2(input):
    xmax, ymax = max(input.keys())
    count = 0
    for x in range(xmax + 1):
        for y in range(ymax + 1):
            if input[(x, y)] != "A":
                continue

            w1p1 = (x - 1, y - 1)
            w1p2 = (x, y)
            w1p3 = (x + 1, y + 1)
            w2p1 = (x - 1, y + 1)
            w2p2 = (x, y)
            w2p3 = (x + 1, y - 1)

            w1 = "".join((input[w1p1], input[w1p2], input[w1p3]))
            w2 = "".join((input[w2p1], input[w2p2], input[w2p3]))

            if w1 in ("MAS", "SAM") and w2 in ("MAS", "SAM"):
                count += 1

    return count


print(part1(read_input()))
print(part2(read_input()))
