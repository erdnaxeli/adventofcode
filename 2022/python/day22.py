import re
from copy import deepcopy
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


def fold_sphere(x, y, map, facing):
    max_y, max_x = len(map), len(map[0])
    x %= max_x
    y %= max_y
    while map[y][x] == " ":
        x, y = add_point((x, y), facing)
        x %= max_x
        y %= max_y

    return x, y, facing


def fold_cube(x, y, map, facing):
    # This is totally ad-hoc four my input.
    if x == -1:
        if 100 <= y <= 149:
            return 50, 49 - (y - 100), (1, 0)
        elif 150 <= y <= 199:
            return 50 + (y - 150), 0, (0, 1)
    elif x == 49:
        if 0 <= y <= 49:
            return 0, 100 + (49 - y), (1, 0)
        elif 50 <= y <= 99:
            return 100 - y, 100, (0, 1)
    elif x == 50 and 150 <= y <= 199:
        return 50 + (y - 150), 149, (0, -1)
    elif x == 100:
        if 50 <= y <= 99:
            return 100 + (y - 50), 49, (0, -1)
        elif 100 <= y <= 149:
            return 149, 49 - (y - 100), (-1, 0)
    elif x == 150 and 0 <= y <= 149:
        return 99,  100 + (49 - y),  (-1, 0)

    if y == -1:
        if 50 <= x <= 99:
            return 0, 150 + (x - 50), (1, 0)
        elif 100 <= x <= 149:
            return x - 100, 199, (0, -1)
    elif y == 50 and 100 <= x <= 149:
        return 99, 50 + (x - 100), (-1, 0)
    elif y == 99 and 0 <= x <= 49:
        return 50, 99 - (50 - x), (1, 0)
    elif y == 150 and 50 <= x <= 99:
        return 49, 150 + (x - 50), (-1, 0)
    elif y == 200 and 0 <= x <= 49:
        return 100 + x, 0, (0, 1)

    return x, y, facing



def apply_path(map, path, x, y, facing, fold_method):
    max_y, max_x = len(map), len(map[0])
    mapc = deepcopy(map)
    facing_values = {
        (1, 0): ">",
        (0, 1): "v",
        (-1, 0): "<",
        (0, -1): "^",
    }
    for move in path:
        match move:
            case int():
                for _ in range(move):
                    mapc[y][x] = f"\033[93m\033[1m{facing_values[facing]}\033[0m"
                    xx, yy = add_point((x, y), facing)
                    print(xx, yy)
                    # if (xx, yy) == (100, -1):
                        # breakpoint()
                    xx, yy, facing = fold_method(xx ,yy ,map, facing)
                    print(xx, yy)
                    print()
                    xx %= max_x
                    yy %= max_y
                    match map[yy][xx]:
                        case "#":
                            break
                        case ".":
                            x, y = xx, yy
                        case err:
                            raise ValueError(f"Unknown case {err}")


            case "L":
                facing = (facing[1], -facing[0])
            case "R":
                facing = (-facing[1], facing[0])

        mapc[y][x] = "\033[94mo\033[0m"
        # print_map(mapc)
        # print(move, x, y)
        # input()

    return x, y, facing


def get_password(facing, x, y):
    facing_values = {
        (1, 0): 0,
        (0, 1): 1,
        (-1, 0): 2,
        (0, -1): 3,
    }
    return 1000 * (y + 1) + 4 * (x + 1) + facing_values[facing]


def print_map(map):
    print("\n".join(" ".join(m) for m in map))


def part1(map, path):
    start_x, start_y = get_start(map)
    x, y, facing = apply_path(map, path, start_x, start_y, (1, 0), fold_sphere)
    return get_password(facing, x, y)


def part2(map, path):
    start_x, start_y = get_start(map)
    x, y, facing = apply_path(map, path, start_x, start_y, (1, 0), fold_cube)
    return get_password(facing, x, y)


if __name__ == "__main__":
    map, path = read_notes()
    map = fill_map(map)
    # print(part1(map, path))
    print(part2(map, path))
