from collections import namedtuple

Point = namedtuple("Point", ("x", "y"))


def read_input():
    with open("day9.txt") as file:
        return [l.strip().split(" ") for l in file]


def move_head(head, direction):
    match direction:
        case "R":
            head = Point(head.x + 1, head.y)
        case "L":
            head = Point(head.x - 1, head.y)
        case "D":
            head = Point(head.x, head.y - 1)
        case "U":
            head = Point(head.x, head.y + 1)

    return head


def move_tail(head, tail):
    if tail.x == head.x and abs(head.y - tail.y) > 1:
        tail = Point(
            tail.x,
            tail.y + (head.y - tail.y) / abs(head.y - tail.y),
        )
    elif tail.y == head.y and abs(head.x - tail.x) > 1:
        tail = Point(
            tail.x + (head.x - tail.x) / abs(head.x - tail.x),
            tail.y,
        )
    elif abs(head.x - tail.x) > 1 or abs(head.y - tail.y) > 1:
        tail = Point(
            tail.x + (head.x - tail.x) / abs(head.x - tail.x),
            tail.y + (head.y - tail.y) / abs(head.y - tail.y),
        )

    return tail


def part1(motions):
    head = Point(0, 0)
    tail = Point(0, 0)
    tail_positions = {tail}

    for direction, repetition in motions:
        for _ in range(int(repetition)):
            head = move_head(head, direction)
            tail = move_tail(head, tail)
            tail_positions.add(tail)

    return len(tail_positions)


def part2(motions):
    head = Point(0, 0)
    tails = [Point(0, 0) for _ in range(9)]
    tail_positions = {tails[-1]}

    for direction, repetition in motions:
        for _ in range(int(repetition)):
            head = move_head(head, direction)
            tails[0] = move_tail(head, tails[0])

            for i in range(1, len(tails)):
                tails[i] = move_tail(tails[i - 1], tails[i])

            tail_positions.add(tails[-1])

    return len(tail_positions)


if __name__ == "__main__":
    motions = read_input()
    print(part1(motions))
    print(part2(motions))
