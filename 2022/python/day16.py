import re
from dataclasses import dataclass, field, replace
from itertools import groupby

VALVE_RGX = re.compile(
    r"Valve (..) has flow rate=(\d+); tunnels? leads? to valves? (.*)"
)


@dataclass
class Valve:
    name: str
    flow: int
    tunnels_to: list[str]


class Action:
    pass


@dataclass
class Move(Action):
    to: Valve

    def __str__(self):
        return f"Move({self.to.name})"


@dataclass
class Open(Action):
    pass


@dataclass
class Path:
    start: Valve
    actions: list[Action] = field(default_factory=list)
    opened: set[str] = field(default_factory=set)

    @property
    def position(self):
        for action in reversed(self.actions):
            match action:
                case Move(to):
                    return to.name
                case _:
                    continue

        return self.start.name

    @property
    def pressure(self):
        current_valve = self.start
        pressure = 0
        for i, action in zip(range(30, 0, -1), self.actions):
            match action:
                case Move(valve):
                    current_valve = valve
                case Open():
                    # The valve will be open from the next minute to the end.
                    pressure += (i - 1) * current_valve.flow
                case _:
                    raise ValueError("Invalid action")

        return pressure

    def __str__(self):
        return f"Path({', '.join(str(action) for action in self.actions)})"


def read_valves():
    return dict(yield_valves())


def yield_valves():
    with open("day16.txt") as file:
        for line in file:
            yield read_valve(line.strip())


def read_valve(line):
    match = VALVE_RGX.match(line)
    valve = Valve(
        name=match.group(1),
        flow=int(match.group(2)),
        tunnels_to=match.group(3).split(", "),
    )
    return valve.name, valve


def compute_paths_from(valve, valves, path):
    paths = []
    for next_valve in valves[valve].tunnels_to:
        if next_valve in path:
            continue

        paths.append(path + [next_valve])
        for p in compute_paths_from(next_valve, valves, path + [next_valve]):
            paths.append(p)

    shortest_paths = []
    for valve, paths_to_valve in groupby(
        sorted(paths, key=lambda p: p[-1]), key=lambda p: p[-1]
    ):
        shortest_path = sorted(paths_to_valve, key=len)[0]
        shortest_paths.append(shortest_path)

    return shortest_paths


def compute_paths(valves):
    valves_paths = {}
    for valve in valves.keys():
        paths = compute_paths_from(valve, valves, [valve])
        valves_paths[valve] = paths

    return valves_paths


def part1(valves):
    valves_to_valves = compute_paths(valves)
    aa = valves["AA"]
    paths = [Path(start=aa)]
    best_pressure = 0
    flows = sorted([v.flow for v in valves.values()], reverse=True)

    i = 0
    while paths:
        i += 1
        # print(best_pressure, len(paths))
        path = paths.pop()
        if len(path.actions) >= 30:
            path.actions = path.actions[:30]
            if path.pressure > best_pressure:
                best_pressure = path.pressure
                print("done", best_pressure)

        actions_to_do = 30 - len(path.actions)
        # this is an ideal case that will not append
        # we should remove the flows already used to have a better filter
        best_pressure_to_expect = path.pressure + sum(
            i * f for i, f in zip(range(actions_to_do, 0, -1), flows)
        )
        if best_pressure_to_expect <= best_pressure:
            continue

        new_path_found = False
        for next_path in valves_to_valves[path.position]:
            actions = []
            if next_path[-1] in path.opened:
                continue
            if valves[next_path[-1]].flow == 0:
                continue

            for valve in next_path[1:]:
                actions.append(Move(valves[valve]))

            new_path_found = True
            actions.append(Open())
            paths.append(
                replace(
                    path, actions=path.actions + actions, opened=path.opened | {valve}
                )
            )

        if not new_path_found and path.pressure > best_pressure:
            best_pressure = path.pressure
            print("end", best_pressure)

    print(i)
    return best_pressure


if __name__ == "__main__":
    valves = read_valves()
    print(part1(valves))
