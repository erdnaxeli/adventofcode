from functools import cache


def read_input():
    with open("day19.txt") as f:
        towels = f.readline().rstrip().split(", ")
        f.readline()

        designs = [line.rstrip() for line in f]
        return towels, designs


class Part1:
    def __init__(self, towels, designs):
        self.towels = towels
        self.designs = designs
        self.cache = {}

    @cache
    def can_do(self, design):
        if not design:
            return True

        for towel in self.towels:
            if design.startswith(towel) and self.can_do(design[len(towel) :]):
                return True

        return False

    def solve(self):
        count = 0
        for design in self.designs:
            if self.can_do(design):
                count += 1

        return count


class Part2:
    def __init__(self, towels, designs):
        self.towels = towels
        self.designs = designs
        self.cache = {}

    @cache
    def can_do(self, design):
        if not design:
            return 1

        count = 0
        for towel in self.towels:
            if design.startswith(towel):
                count += self.can_do(design[len(towel):])

        return count

    def solve(self):
        count = 0
        for design in self.designs:
            count += self.can_do(design)

        return count


if __name__ == "__main__":
    towels, designs = read_input()

    print(Part1(towels, designs).solve())
    print(Part2(towels, designs).solve())
