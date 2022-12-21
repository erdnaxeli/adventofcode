def read_monkeys():
    return dict(read_monkeys_yield())


def read_monkeys_yield():
    with open("day21.txt") as file:
        for line in file:
            monkey, job = line.strip().split(": ")
            if " " in job:
                yield monkey, job.split(" ")
            else:
                yield monkey, int(job)


def solve(monkeys, monkey):
    stack = [(monkey, monkeys[monkey])]
    while True:
        match stack.pop():
            case (monkey, int(x)):
                if not stack:
                    return x

                match stack.pop():
                    case (monkey_waiting, (str(m1), op, str(m2))) if m1 == monkey:
                        stack.append((monkey_waiting, (x, op, m2)))
                        stack.append((m2, monkeys[m2]))
                    case (monkey_waiting, (int(m1), op, str(m2))) if m2 == monkey:
                        stack.append((monkey_waiting, (m1, op, x)))
                    case err:
                        return ValueError(
                            f"Invalid op state {err} (current monkey {monkey})"
                        )
            case (monkey, (int(x), op, int(y))):
                match op:
                    case "+":
                        stack.append((monkey, x + y))
                    case "-":
                        stack.append((monkey, x - y))
                    case "*":
                        stack.append((monkey, x * y))
                    case "/":
                        stack.append((monkey, x // y))
                    case _:
                        raise ValueError("Unknown operator")
            case (monkey, (str(m1), op, str(m2))):
                stack.append((monkey, (m1, op, m2)))
                stack.append((m1, monkeys[m1]))
            case err:
                raise ValueError(f"Invalid stack value {err}")


def part1(monkeys):
    result = solve(monkeys, "root")
    return result


if __name__ == "__main__":
    monkeys = read_monkeys()
    print(part1(monkeys))
