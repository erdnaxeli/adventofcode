from functools import cache, cmp_to_key


@cache
def read_input():
    rules = []
    updates = []
    with open("day05.txt") as f:
        for line in f:
            line = line.rstrip()
            if line == "":
                continue

            result = line.split("|")
            if len(result) == 2:
                rules.append((int(result[0]), int(result[1])))
            else:
                result = line.split(",")
                updates.append([int(x) for x in result])

    return rules, updates


def is_valid(update, rules):
    for rule in rules:
        if rule[0] not in update or rule[1] not in update:
            continue

        i1 = update.index(rule[0])
        i2 = update.index(rule[1])

        if not i1 < i2:
            return False

    return True


def part1(input):
    rules, updates = input
    valid_updates = []
    for update in updates:
        if is_valid(update, rules):
            valid_updates.append(update)

    return sum(x[(len(x) - 1) // 2] for x in valid_updates)


def part2(input):
    rules, updates = input
    invalid_updates = []
    for update in updates:
        if not is_valid(update, rules):
            invalid_updates.append(update)

    rules_by_page = {}
    for rule in rules:
        r1, r2 = rule
        rules_by_page[r1] = rules_by_page.get(r1, []) + [r2]

    def cmp(a, b):
        if b in rules_by_page.get(a, []):
            return -1
        else:
            return 1

    corrected_updates = [sorted(update, key=cmp_to_key(cmp)) for update in invalid_updates]
    return sum(x[(len(x) - 1) // 2] for x in corrected_updates)


print(part1(read_input()))
print(part2(read_input()))
