from pathlib import Path

from grid import DOWN, LEFT, RIGHT, UP, add_points, read_grid


def read_input(part: int) -> tuple[dict, list[str]]:
    grid = read_grid(f"day15_p{part}_grid.txt", lambda: "#")
    moves = list("".join(Path("day15_moves.txt").read_text().rstrip().split()))

    return grid, moves


class WareHouseRobot:
    def __init__(self, grid: dict):
        self.grid = grid

    def _move_object(self, p: tuple[int, int], d: tuple[int, int]) -> tuple[int, int]:
        np = add_points(p, d)
        if self.grid[np] == ".":
            self.grid[np] = self.grid[p]
            self.grid[p] = "."
            return np
        elif self.grid[np] == "#":
            return p
        elif self.grid[np] in ["O", "[", "]"]:
            nnp = self.move_object(np, d)

            if nnp == np:
                # the object can't move
                return p
            else:
                self.grid[np] = self.grid[p]
                self.grid[p] = "."
                return np
        else:
            raise ValueError(f"unknown map item {self.grid[np]}")

    def move_object(self, p: tuple[int, int], d: tuple[int, int]) -> tuple[int, int]:
        if self.grid[p] in [".", "#"]:
            # there is nothing to move
            return p
        elif self.grid[p] in ["@", "O"]:
            return self._move_object(p, d)
        elif self.grid[p] == "[":
            p_r = add_points(p, RIGHT)

            if d in [UP, DOWN]:
                prev_grid = self.grid.copy()
                np = self._move_object(p, d)

                if np == p:
                    # the first part of the box did not move
                    return p

                np_r = self._move_object(p_r, d)
                if np_r == p_r:
                    # the second part did not move
                    # we undo the move of the first part
                    self.grid = prev_grid
                    return p

                return np
            elif d == LEFT:
                np = self._move_object(p, d)
                if np == p:
                    return p

                # we know this move can't fail
                self._move_object(p_r, d)

                return np
            elif d == RIGHT:
                np_r = self._move_object(p_r, d)
                if np_r == p_r:
                    return p

                # we know this move can't fail
                return self._move_object(p, d)
            else:
                raise ValueError(f"unknown direction {d=}")
        elif self.grid[p] == "]":
            p_l = add_points(p, LEFT)
            np_l = self.move_object(p_l, d)

            return add_points(np_l, RIGHT)
        else:
            raise ValueError(f"unknown map item: {p=} {self.grid[p]=}")

    def apply_move(self, p: tuple[int, int], move: str) -> tuple[int, int]:
        if self.grid[p] != "@":
            raise ValueError(f"incorrect robot position given: {p=} {self.grid[p]=}")

        match move:
            case "^":
                d = UP
            case ">":
                d = RIGHT
            case "v":
                d = DOWN
            case "<":
                d = LEFT
            case _:
                raise ValueError(f"unknown move {move}")

        return self.move_object(p, d)

    def apply_moves(self, moves: list[str]):
        for robot_position, v in self.grid.items():
            if v == "@":
                break
        else:
            raise ValueError("the robot position could not be found")

        # print_grid(self.grid)
        for move in moves:
            robot_position = self.apply_move(robot_position, move)
            # print(move)
            # print_grid(self.grid)


def part1(grid: dict, moves: list[str]) -> int:
    robot = WareHouseRobot(grid)
    robot.apply_moves(moves)

    return sum(100 * p[0] + p[1] for p in grid if grid[p] == "O")


def part2(grid: dict, moves: list[str]) -> int:
    robot = WareHouseRobot(grid)
    robot.apply_moves(moves)

    return sum(100 * p[0] + p[1] for p in robot.grid if robot.grid[p] == "[")


if __name__ == "__main__":
    print(part1(*read_input(1)))
    print(part2(*read_input(2)))
