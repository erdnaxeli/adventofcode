from itertools import combinations


def read_cubes():
    return set(read_cubes_yield())


def read_cubes_yield():
    with open("day18.txt") as file:
        for line in file:
            yield tuple(int(x) for x in line.strip().split(","))


def get_neighbors(cube):
    x, y, z = cube
    return {
        (x - 1, y, z),
        (x + 1, y, z),
        (x, y - 1, z),
        (x, y + 1, z),
        (x, y, z - 1),
        (x, y, z + 1),
    }


def part1(cubes: set[tuple[int, int, int]]):
    area = 0
    for cube in cubes:
        for neighbor in get_neighbors(cube):
            if neighbor not in cubes:
                area += 1

    return area


def find_blocks(cubes):
    """
    Return a list on blocks.

    A block is a set of continuous cubes.
    """
    cubes = cubes.copy()
    blocks = []
    block = set()
    cubes_to_analyze = [cubes.pop()]
    while True:
        while cubes_to_analyze:
            x, y, z = cubes_to_analyze.pop()
            block.add((x, y, z))
            candidates = get_neighbors((x, y, z))
            cubes_to_analyze.extend(candidates & cubes)
            cubes.difference_update(candidates)

        blocks.append(block)
        block = set()
        if not cubes:
            break

        cubes_to_analyze.append(cubes.pop())

    return blocks


def find_air_cubes(block):
    air_cubes = set()
    min_x = min(x for x, _, _ in block)
    max_x = max(x for x, _, _ in block)
    for x in range(min_x, max_x):
        min_y = min(y for xx, y, _ in block if xx == x)
        max_y = max(y for xx, y, _ in block if xx == x)
        for y in range(min_y, max_y):
            min_z = min(z for xx, yy, z in block if (xx, yy) == (x, y))
            max_z = max(z for xx, yy, z in block if (xx, yy) == (x, y))
            for z in range(min_z, max_z):
                if (x, y, z) not in block:
                    air_cubes.add((x, y, z))

    return air_cubes


def part2(cubes: set[tuple[int]]):
    air_cubes = find_air_cubes(cubes)
    air_blocks = find_blocks(air_cubes)
    air_area = 0

    for air_block in air_blocks:
        if any(
            not get_neighbors(air_cube).issubset(cubes | air_block)
            for air_cube in air_block
        ):
            continue

        air_area += part1(air_block)

    return part1(cubes) - air_area


if __name__ == "__main__":
    cubes = read_cubes()
    print(part1(cubes))
    print(part2(cubes))
