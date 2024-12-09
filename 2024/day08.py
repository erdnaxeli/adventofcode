from collections import defaultdict
from itertools import combinations
from math import floor, sqrt

from day06 import read_grid


def get_antennas(grid):
    antennas = defaultdict(list)

    for p, value in grid.items():
        if value != ".":
            antennas[value].append(p)

    return antennas


def part1(grid):
    antennas = get_antennas(grid)
    antinodes = set()

    for freq in antennas:
        for antenna_a, antenna_b in combinations(antennas[freq], 2):
            distance = (antenna_a[0] - antenna_b[0], antenna_a[1] - antenna_b[1])
            p1 = (antenna_a[0] + distance[0], antenna_a[1] + distance[1])
            p2 = (antenna_b[0] - distance[0], antenna_b[1] - distance[1])

            if p1 in grid:
                antinodes.add(p1)
            if p2 in grid:
                antinodes.add(p2)

    return len(antinodes)


def reduce_vector(x, y):
    if x % 2 == 0 and y % 2 == 0:
        x //= 2
        y //= 2

    for d in range(3, floor(sqrt(abs(min(x, y)))) + 1, 2):
        if x % d == 0 and y % d == 0:
            x //= d
            y //= d

    return x, y


def part2(grid):
    antennas = get_antennas(grid)
    antinodes = set()

    for freq in antennas:
        for antenna_a, antenna_b in combinations(antennas[freq], 2):
            # We need to _reduce_ the vector to the smallest one that goes on a grid position,
            # e.g. (9, 15) becomes (3,5).
            # This is actually not needed for the given input, but it could have happened.
            distance = reduce_vector(
                antenna_a[0] - antenna_b[0], antenna_a[1] - antenna_b[1]
            )

            p = antenna_a
            while p in grid:
                antinodes.add(p)
                p = (p[0] + distance[0], p[1] + distance[1])

            p = antenna_a
            p = (p[0] - distance[0], p[1] - distance[1])
            while p in grid:
                antinodes.add(p)
                p = (p[0] - distance[0], p[1] - distance[1])

    return len(antinodes)


if __name__ == "__main__":
    input = read_grid("day08.txt", str)
    print(part1(input))
    print(part2(input))
