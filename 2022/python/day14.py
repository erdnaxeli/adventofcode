from collections import namedtuple, defaultdict
from functools import cached_property
from itertools import pairwise
from math import copysign

Point = namedtuple("Point", ("x", "y"))


class MapIterator:
    """An iterator over the positions of the sand falling."""

    def __init__(self, map, sand_generator, floor=None):
        self.map = map
        self.floor_y = floor
        self.sand = None
        self.sand_generator = sand_generator

    def __iter__(self):
        return self

    def __next__(self):
        if self.sand is None:
            self.sand = self.sand_generator
            new = True
        else:
            sand = self._fall()
            new = False
            if sand is None:
                if self.sand == self.sand_generator:
                    # we can't go anywhere
                    raise StopIteration

                self.sand = self.sand_generator
                new = True
            else:
                if self.sand != self.sand_generator:
                    self.map[self.sand] = "."

                self.sand = sand
                self.map[self.sand] = "o"

        return self.sand, new

    def is_in_abyss(self, point):
        return point.y >= self.max_y

    @cached_property
    def max_y(self):
        return max(p.y for p in map)

    def _fall(self):
        sand = self._fall_down()
        if not self._is_blocked(sand):
            return sand

        sand = self._fall_left()
        if not self._is_blocked(sand):
            return sand

        sand = self._fall_right()
        if not self._is_blocked(sand):
            return sand

        return None

    def _fall_down(self):
        return Point(self.sand.x, self.sand.y + 1)

    def _fall_left(self):
        return Point(self.sand.x - 1, self.sand.y + 1)

    def _fall_right(self):
        return Point(self.sand.x + 1, self.sand.y + 1)

    def _is_blocked(self, point):
        if self.floor_y is not None and point.y == self.floor_y:
            return True
        else:
            return self.map[point] in ("#", "o")


def read_paths():
    with open("day14.txt") as file:
        return [read_path(line.strip()) for line in file]


def read_path(descr):
    path = []
    parts = descr.split(" -> ")
    for part in parts:
        path.append(read_point(part))

    return expand_path(path)


def read_point(descr):
    x, y = descr.split(",")
    return Point(x=int(x), y=int(y))


def expand_path(path):
    full_path = []
    for p1, p2 in pairwise(path):
        if p1.x == p2.x:
            for y in range(p1.y, p2.y, int(copysign(1, p2.y - p1.y))):
                full_path.append(Point(x=p1.x, y=y))
        elif p1.y == p2.y:
            for x in range(p1.x, p2.x, int(copysign(1, p2.x - p1.x))):
                full_path.append(Point(x=x, y=p1.y))
        else:
            raise ValueError

        full_path.append(p2)

    return full_path


def build_map(paths):
    map = defaultdict(lambda: ".")
    for path in paths:
        for point in path:
            map[point] = "#"

    return map


def print_map(map):
    min_x = min(p.x for p in map)
    min_y = min(p.y for p in map)
    max_x = max(p.x for p in map)
    max_y = max(p.y for p in map)
    for y in range(min_y, max_y + 1):
        for x in range(min_x, max_x + 1):
            print(map[(x, y)], end="")

        print()


def part1(map):
    map_it = MapIterator(map, Point(500, 0))
    units = 0
    for sand, new in map_it:
        if new:
            units += 1

        if map_it.is_in_abyss(sand):
            break

    # we don't count the last unit that will never rest
    return units - 1


def part2(map):
    max_y = max(p.y for p in map)
    map_it = MapIterator(map, Point(500, 0), max_y + 2)
    units = 0
    for sand, new in map_it:
        if new:
            units += 1

    print_map(map_it.map.copy())
    return units


if __name__ == "__main__":
    map = build_map(read_paths())
    map[Point(500, 0)] = "+"
    print_map(map.copy())
    print(part1(map.copy()))
    print(part2(map.copy()))
