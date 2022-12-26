from abc import ABC, abstractmethod
from collections import namedtuple
from itertools import cycle

from tqdm import tqdm

Point = namedtuple("Point", ("x", "y"))


def read_jets():
    with open("day17.txt") as file:
        return file.read().strip()


class Shape(ABC):
    @abstractmethod
    def spawn(self, point: Point) -> set[Point]:
        """
        Return the set of the coordinates of the shape blocks.

        The given point is the initial position of the shape:
        * x is the position from the left wall
        * y is the position from the floor
        """
        ...


class HBar(Shape):
    def spawn(self, point: Point) -> set[Point]:
        return {
            point,
            Point(point.x + 1, point.y),
            Point(point.x + 2, point.y),
            Point(point.x + 3, point.y),
        }


class Cross(Shape):
    def spawn(self, point: Point) -> set[Point]:
        return {
            Point(point.x + 1, point.y + 2),
            Point(point.x, point.y + 1),
            Point(point.x + 1, point.y + 1),
            Point(point.x + 2, point.y + 1),
            Point(point.x + 1, point.y),
        }


class Corner(Shape):
    def spawn(self, point: Point) -> set[Point]:
        return {
            Point(point.x + 2, point.y + 2),
            Point(point.x + 2, point.y + 1),
            Point(point.x, point.y),
            Point(point.x + 1, point.y),
            Point(point.x + 2, point.y),
        }


class VBar(Shape):
    def spawn(self, point: Point) -> set[Point]:
        return {
            Point(point.x, point.y + 3),
            Point(point.x, point.y + 2),
            Point(point.x, point.y + 1),
            Point(point.x, point.y),
        }


class Square(Shape):
    def spawn(self, point: Point) -> set[Point]:
        return {
            Point(point.x, point.y + 1),
            Point(point.x + 1, point.y + 1),
            Point(point.x, point.y),
            Point(point.x + 1, point.y),
        }


class Tetris:
    def __init__(self, moves: list[str], shapes: list[Shape]):
        self.stopped_shapes = 0
        self.highest_y = -1

        self._moves = cycle(moves)
        self._shapes = cycle(shapes)
        self._stopped_blocks: set[Point] = set()
        self._moving_blocks: set[Point] = set()

    def run(self):
        if not self._moving_blocks:
            self._spawn_shape()

        new_positions, stopped = self._move()
        if stopped:
            self._stopped_blocks |= set(new_positions)
            for block in self._moving_blocks:
                self.highest_y = max(block.y, self.highest_y)

            self._moving_blocks = set()
            self.stopped_shapes += 1
        else:
            self._moving_blocks = new_positions

    def print(self):
        if self._moving_blocks:
            max_y = max(self.highest_y, max(p.y for p in self._moving_blocks))
        else:
            max_y = self.highest_y
        for y in range(max_y, -1, -1):
            for x in range(0, 7):
                if Point(x, y) in self._moving_blocks:
                    print("@", end="")
                elif Point(x, y) in self._stopped_blocks:
                    print("#", end="")
                else:
                    print(".", end="")

            print()

    def _move(self):
        stopped = False
        after_move = self._move_move(self._moving_blocks)
        if not self._is_valid(after_move):
            after_move = self._moving_blocks

        after_down = self._move_down(after_move)
        if not self._is_valid(after_down):
            after_down = after_move
            stopped = True

        return after_down, stopped

    def _move_move(self, blocks):
        match next(self._moves):
            case "<":
                diff = -1
            case ">":
                diff = 1

        return set(Point(block.x + diff, block.y) for block in blocks)

    def _move_down(self, blocks):
        return set(Point(block.x, block.y - 1) for block in blocks)

    def _is_valid(self, blocks):
        return all(self._is_valid_block(block) for block in blocks)

    def _is_valid_block(self, block):
        return block not in self._stopped_blocks and block.y >= 0 and 0 <= block.x <= 6

    def _spawn_shape(self):
        self._moving_blocks = next(self._shapes).spawn(Point(2, self.highest_y + 4))


def part1(jets):
    tetris = Tetris(
        jets,
        [HBar(), Cross(), Corner(), VBar(), Square()],
    )

    while tetris.stopped_shapes < 2022:
        tetris.run()

    return tetris.highest_y + 1


def part2(jets):
    tetris = Tetris(
        jets,
        [HBar(), Cross(), Corner(), VBar(), Square()],
    )

    with tqdm(total=1_000_000_000_000, smoothing=0, unit_scale=True) as pbar:
        while tetris.stopped_shapes < 1_000_000_000_000:
            pbar.update(1)
            tetris.run()

    return tetris.highest_y + 1


if __name__ == "__main__":
    jets = read_jets()
    print(part1(jets))
    print(part2(jets))
