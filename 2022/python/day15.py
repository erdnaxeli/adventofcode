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


def find_impossible_xs(reports, target_y):
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
    target_y = 2_000_000
    line = find_impossible_xs(reports, target_y)
    return len(line)


def part2(reports):
    lines = defaultdict(list)
    max_coordinate = 4_000_000

    for report in reports:
        sensor = report.sensor
        distance = report.distance
        for y in range(
            max(0, sensor.y - distance),
            min(max_coordinate + 1, sensor.y + distance + 1),
        ):
            dy = abs(sensor.y - y)
            dx = distance - dy

            lines[y].append(
                (max(0, sensor.x - dx), min(max_coordinate + 1, sensor.x + dx))
            )

    # return len(lines)
    for y in lines:
        free = [(0, max_coordinate)]
        for start, end in lines[y]:
            free_new = []
            for sf, ef in free:
                if end < sf:
                    free_new.append((sf, ef))
                elif start <= sf:
                    if end < ef:
                        free_new.append((end + 1, ef))
                elif ef < start:
                    free_new.append((sf, ef))
                elif sf < start:
                    free_new.append((sf, start - 1))
                    if end < ef:
                        free_new.append((end + 1, ef))

            free = free_new

        if free:
            return free[0][0] * 4_000_000 + y


if __name__ == "__main__":
    reports = list(read_reports())
    print(part1(reports))
    print(part2(reports))
