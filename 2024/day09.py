from pathlib import Path


def read_input():
    return map(int, Path("day09.txt").read_text().rstrip())


def get_disk_view(map):
    disk = []
    id_number = 0
    for i, number in enumerate(map):
        if i % 2 == 0:
            for _ in range(number):
                disk.append(id_number)

            id_number += 1
        else:
            for _ in range(number):
                disk.append(".")

    return disk


def checksum(disk):
    return sum(i * x for i, x in enumerate(disk) if x != ".")


def part1(map):
    disk = get_disk_view(map)
    i, j = 0, len(disk) - 1
    while i < j:
        if disk[i] == ".":
            if disk[j] != ".":
                disk[i], disk[j] = disk[j], disk[i]
                i += 1

            j -= 1
        else:
            i += 1

    return checksum(disk)


def part2(map):
    disk = get_disk_view(map)
    # represent the current state of the disk
    disk_map = map.copy()

    # go through the map from the end to the start
    for j in range(len(map) - 1, 1, -1):
        if j % 2 == 1:
            continue

        file_size = map[j]
        for i in range(1, j, 2):
            free_space = disk_map[i]
            if free_space == 0:
                continue

            if file_size <= free_space:
                id_number = j // 2
                file_index = sum(disk_map[:j])
                new_index = sum(disk_map[:i])

                disk[new_index : new_index + file_size] = [id_number] * file_size
                disk[file_index : file_index + file_size] = ["."] * file_size

                disk_map[i - 1] += file_size
                disk_map[i] -= file_size
                disk_map[j] = 0
                break

    return checksum(disk)


map = list(read_input())
print(part1(map))
print(part2(map))
