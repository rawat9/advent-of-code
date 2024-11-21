from collections import defaultdict
from math import prod
import re


def get_input(path: str) -> list[str]:
    with open(path, mode="r") as file:
        data = file.read().strip().split("\n")

    return data


directions = [
    (0, -1),  # left
    (0, 1),  # right
    (1, 0),  # down
    (-1, 0),  # up
    (1, -1),  # down-left
    (1, 1),  # down-right
    (-1, -1),  # up-left
    (-1, 1),  # up-right
]


dd = defaultdict(list)


def solve(data: list[str]):
    part_num = 0

    symbols = {
        (i, j)
        for i, row in enumerate(data)
        for j, v in enumerate(row)
        if v != "." and not v.isnumeric()
    }

    for r in range(len(data)):
        for m in re.finditer("\d+", data[r]):
            num, start, end = int(m.group()), m.start(), m.end()

            boundary = {
                (r + dx, j + dy) for dx, dy in directions for j in range(start, end)
            }

            if symbols & boundary:  # interesection (a symbol + a number)
                part_num += num
            for symbol in symbols & boundary:
                dd[symbol].append(num)

    return part_num, sum(prod(v) for v in dd.values() if len(v) == 2)


if __name__ == "__main__":
    input_data = get_input("day-03/input.txt")
    part1, part2 = solve(input_data)
    print("PART 1 -", part1)
    print("PART 2 -", part2)
