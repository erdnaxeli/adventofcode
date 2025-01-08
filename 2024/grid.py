from collections import defaultdict

UP = (-1, 0)
RIGHT = (0, 1)
DOWN = (1, 0)
LEFT = (0, -1)


def add_points(p1, p2):
    return p1[0] + p2[0], p1[1] + p2[1]


# Manhattan distance
def distance_m(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


def get_sides_straight(p):
    for d in UP, RIGHT, DOWN, LEFT:
        yield add_points(p, d), d


def read_grid(filename, default, transform=None):
    grid = defaultdict(default)
    x, y = 0, 0
    with open(filename) as f:
        for line in f:
            for char in line.rstrip():
                if transform:
                    grid[(x, y)] = transform(char)
                else:
                    grid[(x, y)] = char

                y += 1

            x += 1
            y = 0

    return grid


def print_grid(grid):
    xmax, ymax = max(grid.keys())
    for x in range(xmax + 1):
        for y in range(ymax + 1):
            print(grid[(x, y)], end="")

        print()

    print()
