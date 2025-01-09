from collections.abc import Iterator


class Computer:
    program: list[int]
    out: list[int]

    A: int
    B: int
    C: int

    def __init__(self, a, b, c, program: list[int]):
        self.program = program

        self.A = a
        self.B = b
        self.C = c

    def execute(self) -> Iterator[int]:
        ptr = 0
        while True:
            if ptr >= len(self.program):
                return

            instruction = self.program[ptr]
            match instruction:
                case 0:
                    # adv
                    print("A = A // 2**", end="")
                    self.A = self.A // (2 ** self._get_combo_operand(ptr + 1))
                case 1:
                    # bxl
                    print(f"B = B^{self.program[ptr + 1]}")
                    self.B = self.B ^ self.program[ptr + 1]
                case 2:
                    # bst
                    print("B = ", end="")
                    self.B = self._get_combo_operand(ptr + 1) % 8
                case 3:
                    # jnz
                    if self.A != 0:
                        print(f"jump to {self.program[ptr + 1]}")
                        ptr = self.program[ptr + 1]
                        continue
                case 4:
                    # bxc
                    print("B = B ^ C")
                    self.B = self.B ^ self.C
                case 5:
                    # out
                    print("out = ", end="")
                    yield self._get_combo_operand(ptr + 1) % 8
                case 6:
                    # bdv
                    print("B = A // 2**", end="")
                    self.B = self.A // 2 ** self._get_combo_operand(ptr + 1)
                case 7:
                    # cdv
                    print("C = A // 2**", end="")
                    self.C = self.A // 2 ** self._get_combo_operand(ptr + 1)
                case _:
                    raise ValueError("Invalid program")

            ptr += 2

    def _get_combo_operand(self, ptr: int) -> int:
        op = self.program[ptr]
        if 0 <= op <= 3:
            print(op)
            return op
        elif op == 4:
            print("A")
            return self.A
        elif op == 5:
            print("B")
            return self.B
        elif op == 6:
            print("C")
            return self.C
        else:
            raise ValueError("Invalid program")


def part1() -> str:
    computer = Computer(
        38_610_541, 0, 0, [2, 4, 1, 1, 7, 5, 1, 5, 4, 3, 5, 5, 0, 3, 3, 0]
    )
    return ",".join(str(i) for i in computer.execute())


def part2() -> int:
    program = [2, 4, 1, 1, 7, 5, 1, 5, 4, 3, 5, 5, 0, 3, 3, 0]

    a = 35_184_372_088_832
    while True:
        computer = Computer(a, 0, 0, program)

        i = -1
        for i, v in enumerate(computer.execute()):
            if i > 12:
                print(a, i, v, program[i])
            if v != program[i]:
                break
        else:
            if i == len(program)-1:
                return a

        a += 1


if __name__ == "__main__":
    print(part1())
    # print(part2())
