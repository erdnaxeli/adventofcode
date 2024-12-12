from functools import cache

input = [4, 4841539, 66, 5279, 49207, 134, 609568, 0]


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
    return stone // factor, stone % factor


def part1(stones):
    for _ in range(25):
        stones = [new_stone for stone in stones for new_stone in blink(stone)]

    return len(stones)


def part2(stones):
    for i in range(75):
        print(i, len(stones))
        stones = [new_stone for stone in stones for new_stone in blink(stone)]

    return len(stones)


print(part1(input.copy()))
print(part2(input.copy()))
