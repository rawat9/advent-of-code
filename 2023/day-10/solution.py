import numpy as np
from collections import deque as queue
from pprint import pprint


def get_input(path: str):
    with open(path, mode="r") as file:
        data = np.matrix(data=[list(line) for line in file.read().split("\n")])

    return data


def predict_starting_point(matrix: np.matrix):
    indices = np.where(matrix == "F")
    return [indices[0][0], indices[1][0]]


ALLOWED = {
    "F": [(1, 0), (0, 1)],
    "L": [(-1, 0), (0, 1)],
    "7": [(0, -1), (1, 0)],
    "J": [(0, -1), (-1, 0)],
    "-": [(0, -1), (0, 1)],
    "|": [(1, 0), (-1, 0)],
}


def dfs(x: int, y: int, count=1):
    visited[x][y] = count
    pprint(visited)

    if matrix[x, y] == "F":
        for dx, dy in ALLOWED["F"]:
            if is_valid(x + dx, y + dy):
                dfs(x + dx, y + dy, count + 1)

    if matrix[x, y] == "L":
        for dx, dy in ALLOWED["L"]:
            if is_valid(x + dx, y + dy):
                dfs(x + dx, y + dy, count + 1)

    if matrix[x, y] == "7":
        for dx, dy in ALLOWED["7"]:
            if is_valid(x + dx, y + dy):
                dfs(x + dx, y + dy, count + 1)

    if matrix[x, y] == "J":
        for dx, dy in ALLOWED["J"]:
            if is_valid(x + dx, y + dy):
                dfs(x + dx, y + dy, count + 1)

    if matrix[x, y] == "-":
        for dx, dy in ALLOWED["-"]:
            if is_valid(x + dx, y + dy):
                dfs(x + dx, y + dy, count + 1)

    if matrix[x, y] == "|":
        for dx, dy in ALLOWED["|"]:
            if is_valid(x + dx, y + dy):
                dfs(x + dx, y + dy, count + 1)


def is_valid(x: int, y: int):
    if x < 1 or x > N - 1 or y < 1 or y > M - 1 or visited[x][y] >= 1:
        return False

    if matrix[x, y] == ".":
        return False

    return True


def bfs(x: int, y: int):
    q = queue()
    count = 0

    q.append((x, y, count))
    visited[x][y] = count

    while len(q) > 0:
        cell = q.popleft()
        print(cell)
        x = cell[0]
        y = cell[1]
        count = cell[2]

        if matrix[x, y] in ("F", "L", "7", "J", "-", "|"):
            for dx, dy in ALLOWED[matrix[x, y]]:
                if is_valid(x + dx, y + dy):
                    q.append((x + dx, y + dy, count + 1))
                    visited[x + dx][y + dy] = count + 1


if __name__ == "__main__":
    matrix = get_input("day-10/part-1.txt")
    N, M = matrix.shape
    visited = [[0 for _ in range(N)] for _ in range(M)]
    start = predict_starting_point(matrix)
    bfs(*start)
    print()
    pprint(visited)
