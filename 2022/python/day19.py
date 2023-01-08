import re
from collections import defaultdict
from dataclasses import dataclass
from enum import Enum
from math import factorial

ROBOT_RGX = re.compile(r"Each ([a-z]*) robot costs ((?:\d+ [a-z]*(?: and )?)+)")


class Material(Enum):
    ORE = "ore"
    CLAY = "clay"
    OBSIDIAN = "obsidian"
    GEODE = "geode"


@dataclass
class Blueprint:
    robots: dict[Material, dict[Material, int]]


class BlueprintState:
    def __init__(
        self,
        blueprint,
        materials=None,
        robots=None,
        tick=0,
        target=None,
        building_robot=None,
    ):
        self.blueprint = blueprint
        self.materials = materials or {material: 0 for material in Material}
        self.robots = robots or defaultdict(lambda: 0, {Material.ORE: 1})
        self.tick = tick
        self.target = target
        self.building_robot = building_robot

    def next_states(self):
        if self.target is None:
            # If the current state does not have a target, we return states at
            # the same tick but with a target.
            cant = []
            can = []
            for material, robot in self.blueprint.robots.items():
                robots_to_consider = self.robots.copy()
                if self.building_robot:
                    robots_to_consider[self.building_robot] += 1

                if any(robots_to_consider[needed] == 0 for needed in robot):
                    # we don't have the robots to produce the needed ressources
                    cant.append(material.name)
                    continue

                if self._can_buy(robot):
                    # If we can buy a robot when we just bought another, it means
                    # we could have bought this one first and still bought the
                    # other.
                    # The paths are symmetric, so we skip this one.
                    can.append(material.name)
                    continue

                yield BlueprintState(
                    blueprint=self.blueprint,
                    materials=self.materials,
                    robots=self.robots,
                    tick=self.tick,
                    target=(material, robot),
                    building_robot=self.building_robot,
                )

            # if cant or can:
            #     print(" " * self.tick, cant, can)
        else:
            # The state has a target.
            material_produced, robot = self.target
            if not self.building_robot and self._can_buy(robot):
                # If a robot is not building already, try to build the target.
                # print(
                #     "." * self.tick,
                #     self.tick + 1,
                #     self.target[0].name,
                #     {k.name: v for k, v in self.materials.items()},
                # )
                yield BlueprintState(
                    blueprint=self.blueprint,
                    materials={
                        material: self.materials[material] - robot.get(material, 0)
                        for material in self.materials
                    },
                    robots=self.robots,
                    tick=self.tick,
                    target=None,
                    building_robot=material_produced,
                )
            else:
                next_materials = {
                    material: self.materials[material] + self.robots[material]
                    for material in Material
                }
                yield BlueprintState(
                    blueprint=self.blueprint,
                    materials=next_materials,
                    robots={
                        robot: self.robots[robot] + (robot == self.building_robot)
                        for robot in self.robots
                    },
                    tick=self.tick + 1,
                    target=self.target,
                    building_robot=None,
                )

    def _can_buy(self, robot):
        return all(
            self.materials[material] >= quantity for material, quantity in robot.items()
        )

    def __repr__(self):
        return (
            "BlueprintState(materials="
            f"{ {k.name: v for k, v in self.materials.items() if v != 0} }, "
            f"robots={ {k.name: v for k, v in self.robots.items() if v != 0} }, "
            f"tick={self.tick}, "
            f"target={self.target[0].name if self.target else None}, "
            f"building={self.building_robot.name if self.building_robot else None}"
            ")"
        )


def read_blueprints():
    blueprints = []
    with open("day19.txt") as file:
        for line in file:
            robots = {}
            for name, ressources_str in ROBOT_RGX.findall(line):
                ressources = {}
                for result in ressources_str.split(" and "):
                    quantity, ressource_name = result.split(" ")
                    ressources[Material(ressource_name)] = int(quantity)

                robots[Material(name)] = ressources

            blueprints.append(Blueprint(robots))

    return blueprints


def worth_exploring(quality_level, state):
    best_quality_level_expected = state.materials[Material.GEODE] + sum(
        range(
            state.robots[Material.GEODE],
            (24 - state.tick) + state.robots[Material.GEODE],
        )
    )
    return best_quality_level_expected > quality_level


def part1(blueprints):
    quality_levels = []
    for blueprint in blueprints:
        print(blueprint)
        quality_level = 0
        states = [BlueprintState(blueprint)]
        while states:
            # print(len(states))
            state = states.pop()
            if state.tick == 24:
                if state.materials[Material.GEODE] > quality_level:
                    quality_level = state.materials[Material.GEODE]
                    print(quality_level)

                continue

            if not worth_exploring(quality_level, state):
                # print("drop", state.tick)
                continue

            n = list(state.next_states())
            states.extend(n)

        quality_levels.append(quality_level)

    print(list(enumerate(quality_levels)))
    return sum(
        (i + 1) * quality_level for i, quality_level in enumerate(quality_levels)
    )


if __name__ == "__main__":
    blueprints = read_blueprints()
    print(part1(blueprints))
