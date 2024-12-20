from day16 import PathFinder
from grid import DOWN, LEFT, RIGHT, UP, add_points


def read_input():
    bytes = []
    with open("day18.txt") as f:
        for line in f:
            x, y = line.rstrip().split(",")
            bytes.append((int(x), int(y)))

    grid = {}
    for x in range(71):
        for y in range(71):
            if (x, y) in bytes[:1024]:
                grid[(x, y)] = "#"
            else:
                grid[(x, y)] = "."

    print(grid[(69, 69)])
    return grid


class PathFinderMemory(PathFinder):
    def can_walk(self, p):
        return p in self.grid and self.grid[p] != "#"

    def is_end(self, p):
        return p == (69, 69)

    def get_possibilities(self, p, d):
        for direction in LEFT, UP, RIGHT, DOWN:
            yield add_points(p, direction), direction

    def update_score(self, previous_score, p1, d1, p2, d2):
        return previous_score + 1

    def cmp_score(self, s1, s2):
        return s1 - s2


def part1(grid):
    paths, _ = PathFinderMemory(grid).solve((0, 0), DOWN, (70, 70), (0, 0, 70, 70))
    return len(paths[0])


grid = read_input()
print(part1(grid))
