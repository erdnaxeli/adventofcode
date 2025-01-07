from functools import cache

STONES = [4, 4841539, 66, 5279, 49207, 134, 609568, 0]


@cache
def blink(stone):
    if stone == 0:
        return [1]
    if stone == 1:
        return [2024]

    stone_s = str(stone)
    if len(stone_s) % 2 == 1:
        return [stone * 2024]

    factor = 10 ** (len(stone_s) // 2)
    return [stone // factor, stone % factor]


def part1(stones):
    for _ in range(25):
        stones = [new_stone for stone in stones for new_stone in blink(stone)]

    return len(stones)


def part2(stones):
    stones_count = {stone: 1 for stone in stones}

    for _ in range(75):
        tmp = {}
        for stone in stones_count:
            for new_stone in blink(stone):
                tmp[new_stone] = tmp.get(new_stone, 0) + stones_count[stone]

        stones_count = tmp

    return sum(stones_count.values())


if __name__ == "__main__":
    print(part1(STONES.copy()))
    print(part2(STONES.copy()))
