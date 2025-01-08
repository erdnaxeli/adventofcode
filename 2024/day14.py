import re
import time
from collections import Counter
from dataclasses import dataclass
from pathlib import Path


@dataclass
class Robot:
    p: tuple[int, int]
    v: tuple[int, int]


def read_input() -> list[Robot]:
    regex = re.compile(r"^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$", re.MULTILINE)
    content = Path("day14.txt").read_text()

    return [
        Robot((int(px), int(py)), (int(vx), int(vy)))
        for px, py, vx, vy in regex.findall(content)
    ]


def robots_iterate(
    robots: list[Robot], len_x: int, len_y: int, times: int
) -> list[Robot]:
    return [
        Robot(
            (
                (robot.p[0] + robot.v[0] * times) % len_x,
                (robot.p[1] + robot.v[1] * times) % len_y,
            ),
            robot.v,
        )
        for robot in robots
    ]


def part1(robots: list[Robot]) -> int:
    robots = robots_iterate(robots, 101, 103, 100)

    count_nw, count_ne, count_sw, count_se = 0, 0, 0, 0
    for robot in robots:
        if robot.p[0] == 50 or robot.p[1] == 51:
            continue

        if robot.p[0] < 50:
            if robot.p[1] < 51:
                count_nw += 1
            else:
                count_sw += 1
        else:
            if robot.p[1] < 51:
                count_ne += 1
            else:
                count_se += 1

    return count_nw * count_ne * count_sw * count_se


def part2(robots: list[Robot]) -> int:
    i = 0
    while True:
        i += 1
        robots = robots_iterate(robots, 101, 103, 1)

        # We suppose that the drawing we are looking for does not contain
        # two or more robot on the same position.
        if len(set(robot.p for robot in robots))  == len(robots):
            robots.sort(key=lambda r: (r.p[1], r.p[0]))

            x, y = 0, 0
            for robot in robots:
                if robot.p[1] == y and robot.p[0] < x:
                    # it can happen if two robots are at the same location
                    continue

                while y != robot.p[1]:
                    print()
                    y += 1
                    x = 0

                while x != robot.p[0]:
                    print(" ", end="")
                    x += 1

                print("#", end="")
                x += 1

            # print()
            # time.sleep(0.2)
            try:
                input(f"\n{i=} continue?")
            except EOFError:
                print()
                return i


if __name__ == "__main__":
    robots = read_input()
    print(part1(robots))
    print(part2(robots))
