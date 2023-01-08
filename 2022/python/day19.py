from __future__ import annotations
import re
from collections.abc import Iterator
from dataclasses import dataclass
from enum import Enum, auto

from path_finder import PathFinder

ROBOT_RGX = re.compile(r"Each ([a-z]*) robot costs ((?:\d+ [a-z]*(?: and )?)+)")


class Material(Enum):
    ORE = "ore"
    CLAY = "clay"
    OBSIDIAN = "obsidian"
    GEODE = "geode"


@dataclass
class Blueprint:
    robots: dict[Action, dict[Material, int]]


def read_blueprints():
    blueprints = []
    with open("day19.txt") as file:
        for line in file:
            robots = {}
            for name, ressources_str in ROBOT_RGX.findall(line):
                ressources = {}
                for result in ressources_str.split(" and "):
                    quantity, ressource_name = result.split(" ")
                    ressources[Action(ressource_name)] = int(quantity)

                robots[Material(name)] = ressources

            blueprints.append(Blueprint(robots))

    return blueprints


class Action(Enum):
    ORE_ROBOT = "ore"
    CLAY_ROBOT = "clay"
    OBSIDIAN_ROBOT = "obsidian"
    GEODE_ROBOT = "geode"
    WAIT = auto()

    def to_material(self) -> Material:
        return Material(self.value)


class BlueprintState:
    def __init__(self, actions: list[Action], blueprint: Blueprint):
        self.actions = actions
        self.blueprint = blueprint
        self.final_state = self._compute_final_state(actions)

    def __len__(self):
        return len(self.actions)

    @property
    def geodes(self) -> int:
        return self.final_state[Material.GEODE]

    @property
    def current_state(self):
        ...

    def next_states(self) -> Iterator[BlueprintState]:
        for robot in self.blueprint.robots:
            if self._can_buy(robot):
                yield BlueprintState(self.actions + [robot], self.blueprint)

            yield BlueprintState(self.actions + [Action.WAIT], self.blueprint)

    def _can_buy(self, robot: Material) -> bool:
        for m, v in self.blueprint.robots[robot].items():
            if self.final_state[m.to_material()] < v:
                return False

        return True

    def _compute_final_state(self, actions) -> dict[Material, int]:
        state = {r: 0 for r in Material}
        robots = {r: 0 for r in Material}
        robots[Material.ORE] = 1
        for action in actions:
            match action:
                case Action.ORE_ROBOT:
                    robots[Material.ORE] += 1
                    self._buy(Material.ORE, state)
                case Action.CLAY_ROBOT:
                    robots[Material.CLAY] += 1
                    self._buy(Material.CLAY, state)
                case Action.OBSIDIAN_ROBOT:
                    robots[Material.OBSIDIAN] += 1
                    self._buy(Material.OBSIDIAN, state)
                case Action.GEODE_ROBOT:
                    robots[Material.GEODE] += 1
                    self._buy(Material.GEODE, state)
                case Action.WAIT:
                    pass

            for r, v in robots.items():
                state[r] += v

        return state

    def _buy(self, material: Material, state: dict[Material, int]) -> None:
        for m, v in self.blueprint.robots[material].items():
            state[m] -= v


class BlueprintExecutor:
    def __init__(self, blueprint: list[Blueprint]):
        self.blueprint = blueprint

    def get_start(self) -> BlueprintState:
        return BlueprintState([], self.blueprint)

    def is_done(self, path: BlueprintState):
        return len(path) == 24

    def available_paths(self, path: BlueprintState) -> Iterator[Blueprint]:
        yield from path.next_states()


def get_quality_level(bluprint_id, blueprint):
    map = BlueprintExecutor(blueprint)
    pf = PathFinder(map)
    geodes = pf.find().geodes
    return bluprint_id * geodes


def part1(blueprints):
    quality_levels = [get_quality_level(i, b) for i, b in enumerate(blueprints)]
    return sum(quality_levels)


if __name__ == "__main__":
    blueprints = read_blueprints()
    print(part1(blueprints))
