moves = {"X": "A", "Y": "B", "Z": "C"}
score_per_move = {"A": 1, "B": 2, "C": 3}
wins = {("A", "C"), ("B", "A"), ("C", "B")}


def round_score(m1, m2):
    score = score_per_move[m2]
    if (m2, m1) in wins:
        score += 6
    elif m1 == m2:
        score += 3

    return score



def part1():
    score = 0
    with open("day2.txt") as file:
        for l in file:
            m1, m2 = l.rstrip().split(" ")
            m2 = moves[m2]
            score += round_score(m1, m2)

    return score


moves_by_end = {
    "X": {"A": "C", "B": "A", "C": "B"},
    "Y": {"A": "A", "B": "B", "C": "C"},
    "Z": {"A": "B", "B": "C", "C": "A"},
}


def part2():
    score = 0
    with open("day2.txt") as file:
        for l in file:
            m1, m2 = l.rstrip().split(" ")
            m2 = moves_by_end[m2][m1]
            score += round_score(m1, m2)

    return score


print(part1())
print(part2())
