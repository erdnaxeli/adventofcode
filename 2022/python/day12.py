from collections import defaultdict
from copy import deepcopy


def read_map():
    with open("day12.txt") as file:
        return {
            (x, y): point
            for (y, line) in enumerate(file)
            for x, point in enumerate(line.strip())
        }


def distance(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


def print_map_path(map, path):
    for p in path:
        map[p] = f"\x1b[6;30;42m{map[p]}\x1b[0m"

    max_x, max_y = sorted(map.keys())[-1]
    for y in range(max_y + 1):
        for x in range(max_x + 1):
            print(map[(x, y)], end="")

        print()


def compute_shortest_path(starts, end, map):
    paths = [[start] for start in starts]
    shortest_path = None
    map_len = {}

    while True:
        if not paths:
            break

        paths.sort(key=lambda path: distance(path[-1], end), reverse=True)
        path = paths.pop()
        if shortest_path is not None and len(path) >= len(shortest_path):
            continue

        position = path[-1]
        for next_step in (
            (position[0] + 1, position[1]),
            (position[0], position[1] + 1),
            (position[0] - 1, position[1]),
            (position[0], position[1] - 1),
        ):
            if next_step in path:
                # loop
                continue
            elif next_step not in map:
                # not in map
                continue
            elif next_step in map_len and len(path) + 1 >= map_len[next_step]:
                # a shorter path goes there
                continue
            elif ord(map[next_step]) - ord(map[position]) > 1:
                continue
            elif next_step == end:
                shortest_path = path + [next_step]
                continue

            paths.append(path + [next_step])
            map_len[next_step] = len(path) + 1

    return shortest_path


def part1(map):
    start = next(p for p in map if map[p] == "S")
    end = next(p for p in map if map[p] == "E")
    map[start] = "a"
    map[end] = "z"

    path = compute_shortest_path([start], end, map)
    print_map_path(map, path)
    return len(path) - 1


def part2(map):
    starts = (p for p in map if map[p] in ("a", "S"))
    startS = next(p for p in map if map[p] == "S")
    end = next(p for p in map if map[p] == "E")
    map[startS] = "a"
    map[end] = "z"

    path = compute_shortest_path(starts, end, map)
    print_map_path(map, path)
    return len(path) - 1


if __name__ == "__main__":
    map = read_map()
    print(part1(deepcopy(map)))
    print(part2(deepcopy(map)))
