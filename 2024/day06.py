from collections import defaultdict
from copy import deepcopy


def read_grid(filename, default):
    grid = defaultdict(default)
    x, y = 0, 0
    with open(filename) as f:
        for line in f:
            for char in line.rstrip():
                grid[(x, y)] = char
                y += 1

            x += 1
            y = 0

    return grid


def find_in_grid(grid, char):
    for x, y in grid:
        if grid[(x, y)] == "^":
            break
    else:
        raise "not found"

    return x, y


def part1(grid):
    x, y = find_in_grid(grid, "^")
    direction = (-1, 0)
    points = {(x, y)}

    while True:
        nx, ny = x + direction[0], y + direction[1]
        if grid[(nx, ny)] == "":
            return len(points)
        elif grid[(nx, ny)] == "#":
            direction = rotate_right(direction)
        else:
            x, y = nx, ny
            points.add((x, y))


def rotate_right(direction):
    return direction[1], direction[0] * -1


def will_it_loop(x, y, direction, grid, points):
    while True:
        points[(x, y)].append(direction)
        # print_grid(grid, points)
        nx, ny = x + direction[0], y + direction[1]
        if grid[(nx, ny)] == "":
            # print("NON")
            return False
        elif grid[(nx, ny)] == "#":
            direction = rotate_right(direction)
        else:
            if direction in points[(nx, ny)]:
                # print("OUI")
                return True

            x, y = nx, ny
            grid[(x, y)] = "+"


def print_grid(grid, points):
    xmax, ymax = max(grid.keys())
    for x in range(xmax + 1):
        for y in range(xmax + 1):
            if points[(x, y)]:
                print("*", end="")
            else:
                print(grid[(x, y)], end="")

        print()

    input()
    print()


def part2(grid):
    x, y = find_in_grid(grid, "^")
    direction = (-1, 0)
    points = defaultdict(list)
    blocks = 0

    while True:
        points[(x, y)].append(direction)
        nx, ny = x + direction[0], y + direction[1]
        if grid[(nx, ny)] == "":
            return blocks
        elif grid[(nx, ny)] == "#":
            direction = rotate_right(direction)
        else:
            # print_grid(grid, points)
            if not points[(nx, ny)]:
                test_grid = grid.copy()
                test_grid[(nx, ny)] = "#"
                if will_it_loop(
                    x, y, rotate_right(direction), test_grid, deepcopy(points)
                ):
                    blocks += 1

            x, y = nx, ny


grid = read_grid("day06.txt", lambda: "")
print(part1(grid))
print(part2(grid))
