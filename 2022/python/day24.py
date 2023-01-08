from __future__ import annotations

from dataclasses import dataclass
from enum import Enum

from path_finder import PathFinder


class PointType(Enum):
    FREE = "."
    BLIZZARD_UP = "^"
    BLIZZARD_RIGHT = ">"
    BLIZZARD_DOWN = "v"
    BLIZZARD_LEFT = "<"
    WALL = "#"


@dataclass(frozen=True)
class State:
    x: int
    y: int
    blizzards_up: set[tuple[int, int]]
    blizzards_right: set[tuple[int, int]]
    blizzards_down: set[tuple[int, int]]
    blizzards_left: set[tuple[int, int]]


class Path(list):
    @property
    def current_state(self):
        return self[-1]


@dataclass
class Map:
    blizzards_up: set[tuple[int, int]]
    blizzards_right: set[tuple[int, int]]
    blizzards_down: set[tuple[int, int]]
    blizzards_left: set[tuple[int, int]]
    max_x: int
    max_y: int
    start: tuple[int, int, int]
    end: tuple[int, int]

    def __post_init__(self):
        state = self._generate_state(self.start[0], self.start[1])
        self._cycles = [
            (
                state.blizzards_up,
                state.blizzards_right,
                state.blizzards_down,
                state.blizzards_left,
            )
        ]
        next_state = self._move_blizzards(state)
        cycles = 1
        while next_state != state:
            self._cycles.append(
                (
                    next_state.blizzards_up,
                    next_state.blizzards_right,
                    next_state.blizzards_down,
                    next_state.blizzards_left,
                )
            )
            next_state = self._move_blizzards(next_state)
            cycles += 1

        self._blizzards_cycle_steps = cycles

    def available_paths(self, path):
        x, y, cycle = path[-1]
        state = State(
            x,
            y,
            self._cycles[(cycle + 1) % len(self._cycles)][0],
            self._cycles[(cycle + 1) % len(self._cycles)][1],
            self._cycles[(cycle + 1) % len(self._cycles)][2],
            self._cycles[(cycle + 1) % len(self._cycles)][3],
        )
        for x, y in (
            (state.x, state.y),
            (state.x, state.y - 1),
            (state.x + 1, state.y),
            (state.x, state.y + 1),
            (state.x - 1, state.y),
        ):
            if not self._is_free(x, y, state):
                continue

            yield Path(path + [(x, y, (cycle + 1) % len(self._cycles))])

    def get_start(self):
        return Path([self.start])

    def is_done(self, path):
        x, y, _ = path.current_state
        return (x, y) == self.end

    def evaluate_best_cost(self, path):
        x, y, _ = path[-1]
        distance_to_end = abs(x - self.end[0]) + abs(y - self.end[1])
        return len(path) + distance_to_end

    def print_path(self, path):
        for i, point in enumerate(path):
            self._print_point(i, point)
            print(i)

    def _print_point(self, i, point):
        for y in range(0, self.max_y + 1):
            print("  " * i, end="")
            for x in range(0, self.max_x + 1):
                if x == point[0] and y == point[1]:
                    print("@", end="")
                elif (y == 0 or y == self.max_y) and (x, y) not in (
                    self.start,
                    self.end,
                ):
                    print("#", end="")
                elif x == 0 or x == self.max_x:
                    print("#", end="")
                elif (x, y) in self._cycles[point[2] % len(self._cycles)][0]:
                    print("^", end="")
                elif (x, y) in self._cycles[point[2] % len(self._cycles)][1]:
                    print(">", end="")
                elif (x, y) in self._cycles[point[2] % len(self._cycles)][2]:
                    print("v", end="")
                elif (x, y) in self._cycles[point[2] % len(self._cycles)][3]:
                    print("<", end="")
                else:
                    print(".", end="")

            print()

    def _is_free(self, x, y, point):
        return not self._is_wall(x, y) and not self._is_blizzard(x, y, point)

    def _is_wall(self, x, y):
        return (x, y) not in ((self.start[0], self.start[1]), self.end) and (
            x <= 0 or y <= 0 or x >= self.max_x or y >= self.max_y
        )

    def _is_blizzard(self, x, y, point):
        return (x, y) in self._get_blizzards(point)

    def _get_blizzards(self, point):
        return (
            point.blizzards_up
            | point.blizzards_right
            | point.blizzards_down
            | point.blizzards_left
        )

    def _generate_state(self, x, y):
        return State(
            x,
            y,
            self.blizzards_up,
            self.blizzards_right,
            self.blizzards_down,
            self.blizzards_left,
        )

    def _move_blizzards(self, point):
        bu = self._move_blizzards_up(point)
        br = self._move_blizzards_right(point)
        bd = self._move_blizzards_down(point)
        bl = self._move_blizzards_left(point)

        return State(point.x, point.y, bu, br, bd, bl)

    def _move_blizzards_up(self, point):
        return frozenset(
            self._move_blizzards_in_direction(
                point.blizzards_up, dy=-1, ry=self.max_y - 1
            )
        )

    def _move_blizzards_right(self, point):
        return frozenset(
            self._move_blizzards_in_direction(point.blizzards_right, dx=1, rx=1)
        )

    def _move_blizzards_down(self, point):
        return frozenset(
            self._move_blizzards_in_direction(point.blizzards_down, dy=1, ry=1)
        )

    def _move_blizzards_left(self, point):
        return frozenset(
            self._move_blizzards_in_direction(
                point.blizzards_left, dx=-1, rx=self.max_x - 1
            )
        )

    def _move_blizzards_in_direction(self, blizzards, dx=0, dy=0, rx=None, ry=None):
        for x, y in blizzards:
            xx, yy = x + dx, y + dy

            if self._is_wall(xx, yy):
                if rx:
                    xx = rx
                else:
                    yy = ry

            yield xx, yy


def read_map():
    bu, br, bd, bl = set(), set(), set(), set()
    with open("day24.txt") as file:
        for y, line in enumerate(file):
            for x, value in enumerate(line.strip()):
                match value:
                    case "^":
                        bu.add((x, y))
                    case ">":
                        br.add((x, y))
                    case "v":
                        bd.add((x, y))
                    case "<":
                        bl.add((x, y))

    return Map(bu, br, bd, bl, x, y, (1, 0, 0), (x - 1, y))


def part1(map):
    pf = PathFinder(map)
    return len(pf.find()) - 1


def part2(map):
    pf = PathFinder(map)
    path1 = pf.find()

    map.start = path1[-1]
    map.end = (path1[0][0], path1[0][1])
    path2 = pf.find()

    map.start = path2[-1]
    map.end = (path2[0][0], path2[0][1])
    path3 = pf.find()
    return len(path1) - 1 + len(path2) - 1 + len(path3) - 1


if __name__ == "__main__":
    map = read_map()
    print(part1(map))
    print(part2(map))
