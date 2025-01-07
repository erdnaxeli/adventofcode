from dataclasses import dataclass
from typing import Optional


@dataclass
class Machine:
    A: tuple[int, int]
    B: tuple[int, int]
    Prize: tuple[int, int]


def read_input() -> list[Machine]:
    machines: list[Machine] = []

    with open("day13.txt") as f:
        while True:
            line = f.readline()
            if not line:
                break

            a_elts = line.rstrip().split(": ")[1].split(", ")
            a = (int(a_elts[0].split("+")[1]), int(a_elts[1].split("+")[1]))
            b_elts = f.readline().rstrip().split(": ")[1].split(", ")
            b = (int(b_elts[0].split("+")[1]), int(b_elts[1].split("+")[1]))
            prize_elts = f.readline().rstrip().split(": ")[1].split(", ")
            prize = (int(prize_elts[0].split("=")[1]), int(prize_elts[1].split("=")[1]))

            machines.append(Machine(a, b, prize))
            f.readline()

    return machines


def solve(machine: Machine) -> Optional[int]:
    px, py = machine.Prize
    ax, ay = machine.A
    bx, by = machine.B

    # P = xA + yB  (x is a_count, y is b_count)
    # We got a system with two equations:
    # px = x*ax + y*bx
    # py = x*ay + y*by
    #
    # x = (px - y*bx) / ax
    # Replace x in the second equation and solve. We got:
    b_count = (ax * py - ay * px) / (ax * by - ay * bx)
    a_count = (px - b_count * bx) / ax

    if a_count.is_integer() and b_count.is_integer():
        return int(3 * a_count + b_count)

    return None


def part1(machines: list[Machine]) -> int:
    total_cost = 0
    for machine in machines:
        cost = solve(machine)
        if cost is not None:
            total_cost += cost

    return total_cost


def part2(machines: list[Machine]) -> int:
    total_cost = 0
    for machine in machines:
        machine.Prize = (
            machine.Prize[0] + 10000000000000,
            machine.Prize[1] + 10000000000000,
        )
        cost = solve(machine)
        if cost is not None:
            total_cost += cost

    return total_cost


if __name__ == "__main__":
    machines = read_input()
    print(part1(machines))
    print(part2(machines))
