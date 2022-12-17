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


@dataclass
class ElephantHelpState:
    start: Valve
    actions_elephant: list[Action] = field(default_factory=list)
    actions_myself: list[Action] = field(default_factory=list)
    opened: set[str] = field(default_factory=set)
    position_myself: str = None
    position_elephant: str = None
    pressure: int = 0

    def __post_init__(self):
        if self.position_elephant is None:
            self.position_elephant = self.start.name

        if self.position_myself is None:
            self.position_myself = self.start.name

    def next_state_elephant(self, jump_elephant):
        to_open = jump_elephant[-1]
        if to_open.name in self.opened:
            return None

        next_actions = (
            self.actions_elephant + [Move(v) for v in jump_elephant] + [Open()]
        )
        if len(next_actions) > 26:
            return None

        return replace(
            self,
            actions_elephant=next_actions,
            opened=self.opened | {to_open.name},
            position_elephant=to_open.name,
            pressure=self.pressure + (26 - len(next_actions)) * to_open.flow,
        )

    def next_state_myself(self, jump_myself):
        to_open = jump_myself[-1]
        if to_open.name in self.opened:
            return None

        next_actions = self.actions_myself + [Move(v) for v in jump_myself] + [Open()]
        if len(next_actions) > 26:
            return None

        return replace(
            self,
            actions_myself=next_actions,
            opened=self.opened | {to_open.name},
            position_myself=to_open.name,
            pressure=self.pressure + (26 - len(next_actions)) * to_open.flow,
        )


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

        if valves[next_valve].flow > 0:
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
        if valve != "AA" and valves[valve].flow == 0:
            continue

        paths = compute_paths_from(valve, valves, [valve])
        valves_paths[valve] = paths

    return valves_paths


def part1(valves):
    valves_to_valves = compute_paths(valves)
    aa = valves["AA"]
    paths = [Path(start=aa)]
    best_pressure = 0
    flows = sorted([v.flow for v in valves.values()], reverse=True)

    while paths:
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
            i * f for i, f in zip(range(actions_to_do, 0, -2), flows)
        )
        if best_pressure_to_expect <= best_pressure:
            continue

        new_path_found = False
        for next_path in valves_to_valves[path.position]:
            actions = []
            if next_path[-1] in path.opened:
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

    return best_pressure


def part2(valves):
    valves_to_valves = compute_paths(valves)
    aa = valves["AA"]
    paths = [ElephantHelpState(start=aa)]
    best_pressure = 0

    while paths:
        path = paths.pop()
        if len(path.actions_elephant) >= 26 and len(path.actions_myself) >= 26:
            path.actions_elephant = path.actions_elephant[:26]
            path.actions_myself = path.actions_myself[:26]
            if path.pressure > best_pressure:
                best_pressure = path.pressure
                print("done", best_pressure)

        flows = sorted(
            [
                v.flow
                for v in valves.values()
                if v.name not in path.opened and v.flow != 0
            ],
            reverse=True,
        )
        best_pressure_to_expect = path.pressure
        for i in range(
            26 - min(len(path.actions_elephant), len(path.actions_myself)),
            26 - max(len(path.actions_elephant), len(path.actions_myself)),
            -2,
        ):
            try:
                best_pressure_to_expect += i * flows.pop(0)
            except IndexError:
                break

        for i in range(
            26 - max(len(path.actions_elephant), len(path.actions_myself)), 0, -2
        ):
            try:
                best_pressure_to_expect += i * flows.pop(0)
                best_pressure_to_expect += i * flows.pop(0)
            except IndexError:
                break

        if best_pressure_to_expect <= best_pressure:
            continue

        run_elephant = True
        run_myself = True
        if path.position_elephant == path.position_myself:
            run_elephant = len(path.actions_elephant) < len(path.actions_myself)
            run_myself = not run_elephant

        new_elephant_path_found = False
        if run_elephant:
            for jump_elephant in valves_to_valves[path.position_elephant]:
                next_path = path.next_state_elephant(
                    [valves[v] for v in jump_elephant[1:]]
                )
                if next_path is None:
                    continue

                new_elephant_path_found = True
                paths.append(next_path)

        new_myself_path_found = False
        if run_myself:
            for jump_myself in valves_to_valves[path.position_myself]:
                next_path = path.next_state_myself([valves[v] for v in jump_myself[1:]])
                if next_path is None:
                    continue

                new_myself_path_found = True
                paths.append(next_path)

        if (
            not new_elephant_path_found
            and not new_myself_path_found
            and path.pressure > best_pressure
        ):
            best_pressure = path.pressure
            print("end", best_pressure)

    return best_pressure


if __name__ == "__main__":
    valves = read_valves()
    # print(part1(valves))
    print(part2(valves))
