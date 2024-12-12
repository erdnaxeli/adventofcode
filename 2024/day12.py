from collections import deque

from grid import (
    distance_m,
    get_sides_straight,
    read_grid,
)


def part1(grid):
    seen = set()
    total_cost = 0

    for point in grid:
        if point in seen:
            continue

        seen.add(point)
        area = 0
        perimeter = 0
        to_explore = set([point])

        while to_explore:
            p1 = to_explore.pop()
            area += 1
            seen.add(p1)

            for p2, _ in get_sides_straight(p1):
                if p2 not in grid:
                    perimeter += 1
                    continue

                if grid[p2] != grid[p1]:
                    perimeter += 1
                    continue

                if p2 in seen:
                    continue

                to_explore.add(p2)

        total_cost += area * perimeter

    return total_cost


def part2(grid):
    seen = set()
    total_cost = 0

    for point in grid:
        if point in seen:
            continue

        seen.add(point)
        area = 0
        perimeter = set()
        to_explore = set([point])

        while to_explore:
            p1 = to_explore.pop()
            area += 1
            seen.add(p1)

            for p2, d in get_sides_straight(p1):
                if p2 not in grid:
                    perimeter.add((p2, d))
                    continue

                if grid[p2] != grid[p1]:
                    perimeter.add((p2, d))
                    continue

                if p2 in seen:
                    continue

                to_explore.add(p2)

        sides = deque()
        while perimeter:
            p1, d1 = perimeter.pop()
            current_side = [p1]
            keep_going = True

            while keep_going:
                keep_going = False
                for p2, d2 in perimeter:
                    if d2 != d1 or p2 in current_side:
                        continue

                    if distance_m(p2, current_side[0]) == 1:
                        current_side.insert(0, p2)
                        keep_going = True
                    elif distance_m(p2, current_side[-1]) == 1:
                        current_side.append(p2)
                        keep_going = True

            for p in current_side:
                if (p, d1) in perimeter:
                    perimeter.remove((p, d1))

            sides.append(current_side)

        total_cost += area * len(sides)

    return total_cost


grid = read_grid("day12.txt", str)
print(part1(grid))
print(part2(grid))
