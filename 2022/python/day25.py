from itertools import product


def read_fuel_requirements():
    with open("day25.txt") as file:
        return [line.strip() for line in file]


def from_snafu(snafu):
    total = 0
    for i, digit in enumerate(reversed(snafu)):
        match digit:
            case "-":
                decimal_digit = -1
            case "=":
                decimal_digit = -2
            case x:
                decimal_digit = int(x)

        total += 5**i * decimal_digit

    return total


def to_snafu(decimal):
    if decimal <= 2:
        return decimal

    d = decimal // 5
    r = decimal % 5
    if r <= 2:
        return f"{to_snafu(d)}{r}"

    return f"{to_snafu(decimal + (5 - r))[:-1]}{'=' if r == 3 else '-'}"


def part1(fuels):
    return to_snafu(sum(from_snafu(snafu) for snafu in fuels))


if __name__ == "__main__":
    fuels = read_fuel_requirements()
    print(part1(fuels))
