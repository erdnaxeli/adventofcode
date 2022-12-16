import re
from dataclasses import dataclass, field, replace

VALVE_RGX = re.compile(r"Valve (..) has flow rate=(\d+); tunnels? leads? to valves? (.*)")


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
    actions: list[Action]
    opened: set[str] = field(default_factory=set)

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

    def get_next_paths(self, valves):
        match self.actions:
            case []:
                valve = self.start
            case [Open()]:
                valve = self.start
            case [*_, Move(to)]:
                valve = to
            case [*_, Move(to), Open()]:
                valve = to
            case _:
                raise ValueError("Invalid path")

        if valve.name not in self.opened and valve.flow > 0:
            yield replace(
                self,
                actions=self.actions + [Open()],
                opened=self.opened | {valve.name},
            )

        for name in valve.tunnels_to:
            valve = valves[name]
            if self._move_creates_useless_loop(name):
                # ignore this move
                continue

            yield replace(
                self,
                actions=self.actions + [Move(valve)],
            )

    def _move_creates_useless_loop(self, name):
        """A useless loop is a loop while we didn't open any valve."""
        current_valve = self.start
        watch_for_opens = False
        useless = False

        for action in self.actions:
            if current_valve.name == name:
                watch_for_opens = True
                useless = True
                if isinstance(action, Open):
                    # we skip the opening of the starting valve
                    # we want to watch for opens during the loop
                    continue

            match action:
                case Move(to):
                    current_valve = to
                case Open() if watch_for_opens:
                    useless = False
                case _:
                    pass

        return useless

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


def find_best_path(valves, paths):
    complete_paths = []
    # breakpoint()
    while True:
        # print(len(paths))
        if paths:
            current_path = paths.pop()
        else:
            break

        # print(current_path)
        if len(current_path.actions) == 30:
            # print("found")
            complete_paths.append(current_path)
            continue

        for path in current_path.get_next_paths(valves):
            paths.append(path)

    complete_paths.sort(key=lambda p: p.pressure)
    return complete_paths[-1]


def part1(valves):
    aa = valves["AA"]
    path = find_best_path(valves, [Path(start=aa, actions=[])])
    print(path)
    return path.pressure


if __name__ == "__main__":
    valves = read_valves()
    print(part1(valves))
