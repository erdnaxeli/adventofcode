def get_assignments() -> list[tuple[int, int, int, int]]:
    assignments = []
    with open("day4.txt") as file:
        for line in file:
            a1, a2 = line.rstrip().split(",")
            a1_start, a1_end = a1.split("-")
            a2_start, a2_end = a2.split("-")
            assignments.append((int(a1_start), int(a1_end), int(a2_start), int(a2_end)))

    return assignments


def contains(r1: range, r2: range) -> bool:
    return r2.start in r1 and r2.stop - 1 in r1


def overlaps(r1: range, r2: range) -> bool:
    return r1.start in r2 or r2.start in r1


def part1(assignments: list[tuple[int, int, int, int]]) -> int:
    count = 0
    for a1_start, a1_end, a2_start, a2_end in assignments:
        a1 = range(a1_start, a1_end + 1)
        a2 = range(a2_start, a2_end + 1)
        if contains(a1, a2) or contains(a2, a1):
            count += 1

    return count


def part2(assignments: list[tuple[int, int, int, int]]) -> int:
    count = 0
    for a1_start, a1_end, a2_start, a2_end in assignments:
        a1 = range(a1_start, a1_end + 1)
        a2 = range(a2_start, a2_end + 1)
        if overlaps(a1, a2):
            count += 1

    return count


assignments = get_assignments()
print(part1(assignments))
print(part2(assignments))
