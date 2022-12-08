def read_map() -> list[str]:
    with open("day8.txt") as file:
        return file.read().strip().split("\n")


def is_visible(map, x, y) -> bool:
    tree = map[x][y]
    max_x = len(map)
    max_y = len(map[0])

    return (
        all(t < tree for t in (map[x][yy] for yy in range(y)))
        or all(t < tree for t in (map[xx][y] for xx in range(x)))
        or all(t < tree for t in (map[x][yy] for yy in range(y + 1, max_y)))
        or all(t < tree for t in (map[xx][y] for xx in range(x + 1, max_x)))
    )


def get_scenic_score(map, x, y) -> bool:
    tree = map[x][y]
    max_x = len(map)
    max_y = len(map[0])

    v_up = 0
    for yy in range(y - 1, -1, -1):
        v_up += 1
        if map[x][yy] >= tree:
            break

    v_left = 0
    for xx in range(x - 1, -1, -1):
        v_left += 1
        if map[xx][y] >= tree:
            break

    v_down = 0
    for yy in range(y + 1, max_y):
        v_down += 1
        if map[x][yy] >= tree:
            break

    v_right = 0
    for xx in range(x + 1, max_x):
        v_right += 1
        if map[xx][y] >= tree:
            break

    return v_up * v_left * v_down * v_right


def part1(map: list[str]) -> int:
    return len(
        [
            (x, y)
            for x in range(len(map))
            for y in range(len(map[0]))
            if is_visible(map, x, y)
        ]
    )


def part2(map: list[str]) -> int:
    return max(
        get_scenic_score(map, x, y) for x in range(len(map)) for y in range(len(map[0]))
    )


if __name__ == "__main__":
    map = read_map()
    print(part1(map))
    print(part2(map))
