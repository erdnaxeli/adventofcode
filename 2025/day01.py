from pathlib import Path

moves = [
    (line[0], int(line[1:]))
    for line in Path("./day01.txt").read_text().split("\n")
    if line
]

result = 50
zeros = 0

for direction, count in moves:
    if direction == "L":
        result -= count
    else:
        result += count

    if result % 100 == 0:
        zeros += 1

print(zeros)

result = 50
zeros = 0
for direction, count in moves:
    if direction == "L":
        new_result = result - count
        if new_result % 100 == 0 or (
            result % 100 > 0 and new_result % 100 > result % 100
        ):
            zeros += 1
    else:
        new_result = result + count
        if new_result % 100 < result % 100:
            zeros += 1

    zeros += count // 100
    # print(f"{result} + {direction}{count} -> {new_result}, {zeros}")
    result = new_result

print(zeros)
