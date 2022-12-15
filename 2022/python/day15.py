from collections import defaultdict, namedtuple
from dataclasses import dataclass
import re
from functools import cached_property


Point = namedtuple("Point", ("x", "y"))


@dataclass
class Report:
    sensor: Point
    beacon: Point

    @cached_property
    def distance(self):
        return abs(self.sensor.x - self.beacon.x) + abs(self.sensor.y - self.beacon.y)


def read_reports():
    regex = re.compile(
        r"Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)"
    )
    with open("day15.txt") as file:
        for line in file:
            match = regex.match(line)
            yield Report(
                sensor=Point(
                    x=int(match.group(1)),
                    y=int(match.group(2)),
                ),
                beacon=Point(x=int(match.group(3)), y=int(match.group(4))),
            )


def find_impossible_xs(reports, target_y, min_x=None):
    line = set()
    for report in reports:
        sensor = report.sensor
        distance = report.distance
        dy = target_y - sensor.y
        if abs(dy) > distance:
            continue

        dx = distance - abs(dy)
        for x in range(sensor.x - dx, sensor.x + dx):
            line.add(x)

    return line


def part1(reports):
    target_y = 2000000
    line = find_impossible_xs(reports, target_y)
    return len(line)


def part2(reports):
    for target_y in range(4000001):
        if target_y % 1 == 0:
            print(target_y)
        line = find_impossible_xs(reports, target_y)
        print("?")
        line = [x for x in line if 0 <= x <= 4000000]
        print("!")
        if len(line) == 4000001:
            continue

        print(len(line))
        print("scream")
        for x in range(4000001):
            if x not in line:
                return x * 4000000 + target_y


if __name__ == "__main__":
    reports = list(read_reports())
    print(part1(reports))
    print(part2(reports))
