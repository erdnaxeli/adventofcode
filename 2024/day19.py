import re


def read_input():
    with open("day19.txt") as f:
        pattern = f'^({"|".join(f.readline().rstrip().split(", "))})++$'
        f.readline()

        designs = [line.rstrip() for line in f]
        return re.compile(pattern), designs


def part1(pattern, designs):
    count = 0
    for design in designs:
        if pattern.match(design):
            count += 1
        else:
            print(design)

    return count


if __name__ == "__main__":
    pattern, designs = read_input()

    print(part1(pattern, designs))