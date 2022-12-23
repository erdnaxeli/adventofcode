from collections import defaultdict
from enum import Enum, auto
from typing import Optional

from day9 import Point


def read_scan() -> set[Point]:
    return set(read_scan_yield())


def read_scan_yield():
    with open("day23.txt") as file:
        for y, line in enumerate(file):
            for x, c in enumerate(line.strip()):
                if c == "#":
                    yield Point(x, y)


class Direction(Enum):
    N = auto()
    E = auto()
    S = auto()
    W = auto()


class MapIterator:
    def __init__(self, map: set[Point]):
        self.map = map
        self._directions = [Direction.N, Direction.S, Direction.W, Direction.E]

    def run(self):
        moves = self._propose_moves()
        resolved_moves = self._resolve_moves(moves)
        if not resolved_moves:
            raise StopIteration()

        self._do_moves(resolved_moves)
        self._cycle_directions()

    def count_empty(self) -> int:
        min_x = min(elf.x for elf in self.map)
        max_x = max(elf.x for elf in self.map)
        min_y = min(elf.y for elf in self.map)
        max_y = max(elf.y for elf in self.map)

        all_points = {
            Point(x, y)
            for x in range(min_x, max_x + 1)
            for y in range(min_y, max_y + 1)
        }
        return len(all_points.difference(self.map))

    def _propose_moves(self) -> dict[Point, list[Point]]:
        moves: dict[Point, list[Point]] = defaultdict(list)
        for elf in self.map:
            if move := self._move_elf(elf):
                moves[move].append(elf)

        return moves

    def _move_elf(self, elf: Point) -> Optional[Point]:
        if not self._elves_around(elf):
            return None

        for direction in self._directions:
            if not self._elves_in_direction(elf, direction):
                return self._move_elf_direction(elf, direction)

        return None

    def _elves_around(self, elf: Point) -> bool:
        points_around = {
            Point(elf.x + x, elf.y + y)
            for x in (-1, 0, 1)
            for y in (-1, 0, 1)
            if (x, y) != (0, 0)
        }
        return self._elves_in(points_around)

    def _elves_in(self, points: set[Point]) -> bool:
        return bool(self.map & points)

    def _elves_in_direction(self, elf: Point, direction: Direction):
        match direction:
            case Direction.N:
                return self._elves_in(
                    {Point(x, elf.y - 1) for x in (elf.x - 1, elf.x, elf.x + 1)}
                )
            case Direction.E:
                return self._elves_in(
                    {Point(elf.x + 1, y) for y in (elf.y - 1, elf.y, elf.y + 1)}
                )
            case Direction.S:
                return self._elves_in(
                    {Point(x, elf.y + 1) for x in (elf.x - 1, elf.x, elf.x + 1)}
                )
            case Direction.W:
                return self._elves_in(
                    {Point(elf.x - 1, y) for y in (elf.y - 1, elf.y, elf.y + 1)}
                )

    def _move_elf_direction(self, elf: Point, direction: Direction):
        match direction:
            case Direction.N:
                return Point(elf.x, elf.y - 1)
            case Direction.E:
                return Point(elf.x + 1, elf.y)
            case Direction.S:
                return Point(elf.x, elf.y + 1)
            case Direction.W:
                return Point(elf.x - 1, elf.y)

    def _resolve_moves(
        self, moves: dict[Point, list[Point]]
    ) -> list[tuple[Point, Point]]:
        return [(elves[0], point) for point, elves in moves.items() if len(elves) == 1]

    def _do_moves(self, moves: list[tuple[Point, Point]]) -> None:
        for elf, point in moves:
            self.map.remove(elf)
            self.map.add(point)

    def _cycle_directions(self):
        self._directions.append(self._directions.pop(0))


def part1(map):
    it = MapIterator(map)
    for _ in range(10):
        it.run()

    return it.count_empty()


def part2(map):
    it = MapIterator(map)
    rounds = 1
    while True:
        try:
            it.run()
        except StopIteration:
            return rounds

        rounds += 1


if __name__ == "__main__":
    map = read_scan()
    print(part1(map.copy()))
    print(part2(map.copy()))
