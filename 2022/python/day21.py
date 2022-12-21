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


def solve_op(x, op, y):
    match op:
        case "+":
            return x + y
        case "-":
            return x - y
        case "*":
            return x * y
        case "/":
            return x // y
        case "=":
            match x, y:
                case int(), str():
                    return x
                case str(), int():
                    return y
                case _:
                    raise ValueError("Invalid equality comparaison")
        case _:
            raise ValueError("Unknown operator")


def solve(monkeys, monkey, ignore=None):
    if ignore is None:
        ignore = []

    stack = [(monkey, monkeys[monkey])]
    while True:
        match stack.pop():
            case (monkey, (str(m1), op, str(m2))):
                stack.append((monkey, (m1, op, m2)))
                if m1 not in ignore:
                    stack.append((m1, monkeys[m1]))
                else:
                    stack.append((m2, monkeys[m2]))
            case (monkey, (int(x), op, int(y))):
                stack.append((monkey, solve_op(x, op, y)))
            case (monkey, (x, op, y)):
                stack.append((monkey, (op, (x, y))))
            case (monkey, x):
                if not stack:
                    return x

                match stack.pop():
                    case (monkey_waiting, (str(m1), op, str(m2))) if m1 == monkey:
                        stack.append((monkey_waiting, (x, op, m2)))
                        if m2 not in ignore:
                            stack.append((m2, monkeys[m2]))
                    case (monkey_waiting, (m1, op, str(m2))) if m2 == monkey:
                        stack.append((monkey_waiting, (m1, op, x)))
                    case err:
                        return ValueError(
                            f"Invalid op state {err} (current monkey {monkey})"
                        )
            case err:
                raise ValueError(f"Invalid stack value {err}")


def find_var(rpn, var):
    match rpn:
        case int() | str():
            return []
        case _, (_, x) if x == var:
            return [1]
        case _, (x, _) if x == var:
            return [0]
        case _, (left, right):
            if path := find_var(left, var):
                return [0] + path
            elif path := find_var(right, var):
                return [1] + path
            else:
                return []
        case err:
            raise ValueError(err)


def rpn_compute(rpn, value, path):
    if not path:
        return value

    var_branch = rpn[1][path[0]]
    # 1 != 1 => False => 0
    # 0 != 1 => True => 1
    other_branch = rpn[1][int(path[0] != 1)]
    match rpn[0]:
        case "+":
            return rpn_compute(var_branch, value - other_branch, path[1:])
        case "-":
            if path[0]:
                return rpn_compute(var_branch, other_branch - value, path[1:])
            else:
                return rpn_compute(var_branch, other_branch + value, path[1:])
        case "*":
            return rpn_compute(var_branch, value / other_branch, path[1:])
        case "/":
            if path[0]:
                return rpn_compute(var_branch, other_branch / value, path[1:])
            else:
                return rpn_compute(var_branch, value * other_branch, path[1:])

        case err:
            raise ValueError(err)


def rpn_solve(equation, var):
    if equation[0] != "=":
        raise ValueError("Invalid equation")

    if len(equation[1][0]) == 1:
        value = equation[1][0]
        rpn = equation[1][1]
    else:
        value = equation[1][1]
        rpn = equation[1][0]

    path = find_var(rpn, var)
    print(path)
    return rpn_compute(rpn, value, path)


def part1(monkeys):
    result = solve(monkeys, "root")
    return result


def part2(monkeys):
    monkeys = monkeys.copy()
    del monkeys["humn"]
    x, _, y = monkeys["root"]
    monkeys["root"] = (x, "=", y)
    result = solve(monkeys, "root", ignore=["humn"])
    print(result)
    return rpn_solve(result, "humn")


if __name__ == "__main__":
    monkeys = read_monkeys()
    print(part1(monkeys))
    print(part2(monkeys))
