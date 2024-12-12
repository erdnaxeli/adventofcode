from grid import read_grid

DIRECTIONS = [(-1, 0), (0, 1), (1, 0), (0, -1)]


def get_trailhead_score(grid, sx, sy):
    points = [(sx, sy)]
    tops = set()

    while points:
        px, py = points.pop()

        for direction in DIRECTIONS:
            x, y = px + direction[0], py + direction[1]

            if (x, y) not in grid or grid[(x, y)] - grid[(px, py)] != 1:
                continue

            if grid[(x, y)] == 9:
                tops.add((x, y))
                continue

            points.append((x, y))

    return len(tops)


def get_trailhead_rating(grid, sx, sy):
    points = [(sx, sy)]
    rating = 0

    while points:
        px, py = points.pop()

        for direction in DIRECTIONS:
            x, y = px + direction[0], py + direction[1]

            if (x, y) not in grid or grid[(x, y)] - grid[(px, py)] != 1:
                continue

            if grid[(x, y)] == 9:
                rating += 1
                continue

            points.append((x, y))

    return rating


def part1(grid):
    count = 0
    for p, value in grid.items():
        if value == 0:
            score = get_trailhead_score(grid, p[0], p[1])
            count += score

    return count


def part2(grid):
    return sum(
        get_trailhead_rating(grid, x, y) for (x, y), value in grid.items() if value == 0
    )


grid = read_grid("day10.txt", str, int)
print(part1(grid))
print(part2(grid))
