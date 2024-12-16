from abc import ABC, abstractmethod
from os import getenv

from grid import DOWN, LEFT, RIGHT, UP, add_points, read_grid


class PathFinder(ABC):
    def __init__(self, grid):
        self.grid = grid

    def solve(self, p, d):
        paths = [([(p, d)], 0)]
        best_paths = []
        best_score = None
        best_score_per_point = {}

        i = 0
        while paths:
            path, score = paths.pop()

            if getenv("DEBUG"):
                if i == 0:
                    self.print(path)
                    i = int(input().rstrip() or 100)
                else:
                    i -= 1

            p1, d1 = path[-1]

            for p2, d2 in self.get_possibilities(p1, d1):
                if any(True for p, d in path if p == p2):
                    continue

                if not self.can_walk(p2):
                    continue

                new_path = path.copy()
                new_path.append((p2, d2))
                new_score = self.update_score(score, p1, d1, p2, d2)

                if best_score is not None and self.cmp_score(best_score, new_score) < 0:
                    continue

                if self.is_end(p2):
                    if new_score == best_score:
                        best_paths.append(new_path)
                    else:
                        best_paths = [new_path]
                        best_score = new_score

                    continue

                if (p2, d2) in best_score_per_point and self.cmp_score(
                    best_score_per_point[(p2, d2)], new_score
                ) < 0:
                    continue

                best_score_per_point[(p2, d2)] = new_score

                paths.append((new_path, new_score))

        return best_paths, best_score

    def print(self, path):
        xmax, ymax = max(grid.keys())
        for x in range(xmax + 1):
            for y in range(ymax + 1):
                if any(True for p, _ in path if p == (x, y)):
                    print("\033[1m\033[42m*\033[0m", end="")
                else:
                    print(grid[(x, y)], end="")

            print()

        print()

    @abstractmethod
    def can_walk(self, p2):
        pass

    @abstractmethod
    def cmp_score(self, param, new_score):
        pass

    @abstractmethod
    def get_possibilities(self, p1, d1):
        pass

    @abstractmethod
    def is_end(self, p2):
        pass

    @abstractmethod
    def update_score(self, score, p1, d1, p2, d2):
        pass


class PathFinderMaze(PathFinder):
    def can_walk(self, p):
        return self.grid[p] != "#"

    def is_end(self, p):
        return self.grid[p] == "E"

    def get_possibilities(self, p, d):
        for direction in LEFT, DOWN, RIGHT, UP:
            if direction == (-d[0], -d[1]):
                # we can't rotate 180°
                continue

            yield add_points(p, direction), direction

    def update_score(self, previous_score, p1, d1, p2, d2):
        if d1 == d2:
            # straight line
            return previous_score + 1
        else:
            # rotation
            return previous_score + 1001

    def cmp_score(self, s1, s2):
        return s1 - s2


def part1(grid):
    for p in grid:
        if grid[p] == "S":
            break

    _, score = PathFinderMaze(grid).solve(p, RIGHT)
    return score


def part2(grid):
    for p in grid:
        if grid[p] == "S":
            break

    paths, _ = PathFinderMaze(grid).solve(p, RIGHT)
    return len({p[0] for path in paths for p in path})


grid = read_grid("day16.txt", lambda: "#")
# print(part1(grid))
print(part2(grid))
