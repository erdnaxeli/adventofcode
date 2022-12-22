import re
from pathlib import Path


def read_path(path):
    for value in re.findall(r"\d+|[RL]", path):
        yield int(value) if value.isnumeric() else value


def read_notes():
    match Path("day22.txt").read_text().split("\n"):
        case [*map, "", path, ""]:
            return [list(m) for m in map], list(read_path(path))
        case err:
            raise ValueError(err)


def fill_map(map):
    max_x = max(len(row) for row in map)
    for row in map:
        row.extend([" "] * (max_x - len(row)))

    return map


def get_start(map):
    for y in range(0, len(map)):

        for x in range(0, len(map[0])):
            if map[y][x] == ".":
                return (x, y)

    raise ValueError("No start found")


def add_point(a, b):
    return (a[0] + b[0], a[1] + b[1])


def apply_path(map, path, x, y, facing):
    max_y, max_x = len(map), len(map[0])
    for move in path:
        match move:
            case int():
                for _ in range(move):
                    xx, yy = add_point((x, y), facing)
                    xx %= max_x
                    yy %= max_y
                    match map[yy][xx]:
                        case "#":
                            break
                        case ".":
                            x, y = xx, yy
                        case " ":
                            while map[yy][xx] == " ":
                                xx, yy = add_point((xx, yy), facing)
                                xx %= max_x
                                yy %= max_y

                            if map[yy][xx] != "#":
                                x, y = xx, yy
                            else:
                                break
            case "L":
                facing = (facing[1], -facing[0])
            case "R":
                facing = (-facing[1], facing[0])

    return x, y, facing


def part1(map, path):
    start_x, start_y = get_start(map)
    x, y, facing = apply_path(map, path, start_x, start_y, (1, 0))
    facing_values = {
        (1, 0): 0,
        (0, 1): 1,
        (-1, 0): 2,
        (0, -1): 3,
    }

    return 1000 * (y+1) + 4 * (x+1) + facing_values[facing]


def print_map(map):
    print("\n".join("".join(m) for m in map))


if __name__ == "__main__":
    map, path = read_notes()
    map = fill_map(map)
    print(part1(map, path))
