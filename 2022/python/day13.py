from functools import cmp_to_key
from itertools import chain, zip_longest


def read_pairs():
    with open("day13.txt") as file:
        return [read_pair(part.strip()) for part in file.read().split("\n\n")]


def read_pair(descr):
    a, b = descr.split("\n")
    return eval(a), eval(b)


def compare(a, b):
    match a, b:
        case int(), int():
            return (a > b) - (a < b)
        case list(), list():
            for aa, bb in zip_longest(a, b):
                order = compare(aa, bb)
                if order == 0:
                    continue

                return order
            return 0
        case int(), list():
            return compare([a], b)
        case list(), int():
            return compare(a, [b])
        case None, int() | list():
            return -1
        case int() | list(), None:
            return 1


def part1(pairs):
    sum = 0
    for idx, (a, b) in enumerate(pairs):
        if compare(a, b) == -1:
            sum += idx + 1

    return sum


def part2(pairs):
    packets = sorted(
        (packet for packets in chain(pairs, [[[[2]], [[6]]]]) for packet in packets),
        key=cmp_to_key(compare),
    )
    return (packets.index([[2]]) + 1) * (packets.index([[6]]) + 1)


if __name__ == "__main__":
    pairs = read_pairs()
    print(part1(pairs))
    print(part2(pairs))
