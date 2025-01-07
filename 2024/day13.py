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


def _solve(
    machine: Machine,
    a_count: int,
    b_count: int,
    position: tuple[int, int],
    cost: int,
    cache: dict[tuple[int, int], Optional[int]],
) -> Optional[int]:
    if (a_count, b_count) in cache:
        return cache[(a_count, b_count)]

    if position[0] > machine.Prize[0] or position[1] > machine.Prize[1]:
        return None

    if position[0] == machine.Prize[0] and position[1] == machine.Prize[1]:
        return cost

    cost_a = _solve(
        machine,
        a_count + 1,
        b_count,
        (position[0] + machine.A[0], position[1] + machine.A[1]),
        cost + 3,
        cache,
    )
    cost_b = _solve(
        machine,
        a_count,
        b_count + 1,
        (position[0] + machine.B[0], position[1] + machine.B[1]),
        cost + 1,
        cache,
    )

    new_cost = min((c for c in (cost_a, cost_b) if c is not None), default=None)
    cache[(a_count, b_count)] = new_cost
    return new_cost


def solve(machine: Machine) -> Optional[int]:
    return _solve(machine, 0, 0, (0, 0), 0, {})


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
        machine.Prize[0] += 10000000000000
        machine.Prize[1] += 10000000000000
        print(machine)
        cost = solve(machine)
        print(cost)
        if cost is not None:
            total_cost += cost

    return total_cost


if __name__ == "__main__":
    machines = read_input()
    print(part1(machines))
    # print(part2(machines))
