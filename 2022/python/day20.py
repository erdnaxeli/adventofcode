def read_encrypted():
    return list(read_encrypted_yield())


def read_encrypted_yield():
    with open("day20.txt") as file:
        for line in file:
            yield int(line.strip())


def part1(encrypted):
    result = [(e, False) for e in encrypted]
    position = 0
    size = len(encrypted)
    for _ in range(size):
        while True:
            if result[position][1]:
                position = (position + 1) % size
            else:
                e, _ = result.pop(position)
                break

        new_position = (position + e) % (size - 1)
        result.insert(new_position, (e, True))

        if new_position <= position:
            position += 1

    zero = result.index((0, True))
    return (
        result[(zero + 1000) % size][0]
        + result[(zero + 2000) % size][0]
        + result[(zero + 3000) % size][0]
    )


if __name__ == "__main__":
    encrypted = read_encrypted()
    print(part1(encrypted))
