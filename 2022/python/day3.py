from itertools import islice


PRIORITIES = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWXYZ"


def read_rucksacks() -> list[tuple[str, str]]:
    rucksacks = []
    with open("day3.txt") as file:
        for line in file:
            rucksack = line.rstrip()
            half = int(len(rucksack) / 2)
            rucksacks.append((rucksack[:half], rucksack[half:]))

    return rucksacks


def get_priority(item: str) -> int:
    return PRIORITIES.index(item) + 1


def day1(rucksacks: list[tuple[str, str]]) -> int:
    priorities = 0
    for c1, c2 in rucksacks:
        common = (set(c1) & set(c2)).pop()
        priorities += get_priority(common)

    return priorities


def day2(rucksacks: list[tuple[str, str]]):
    priorities = 0
    for r1, r2, r3 in [rucksacks[i : i + 3] for i in range(0, len(rucksacks), 3)]:
        common = (set(r1[0] + r1[1]) & set(r2[0] + r2[1]) & set(r3[0] + r3[1])).pop()
        priorities += get_priority(common)

    return priorities


if __name__ == "__main__":
    rucksacks = read_rucksacks()
    print(day1(rucksacks))
    print(day2(rucksacks))
