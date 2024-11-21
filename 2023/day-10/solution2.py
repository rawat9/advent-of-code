import numpy as np
from collections import deque as queue
from pprint import pprint


def get_input(path: str):
    with open(path, mode="r") as file:
        data = np.matrix(data=[list(line) for line in file.read().split("\n")])

    return data


# TODO: implement this correctly
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

def is_valid(x: int, y: int):
    if visited[x][y]:
        return False

    if matrix[x, y] == ".":
        return False

    return True


def bfs(x: int, y: int):
    q = queue()

    q.append((x, y))
    visited[x][y] = True

    while len(q) > 0:
        cell = q.popleft()
        x = cell[0]
        y = cell[1]

        if matrix[x, y] in ("F", "L", "7", "J", "-", "|"):
            for dx, dy in ALLOWED[matrix[x, y]]:
                adjX, adjY = x + dx, y + dy
                if is_valid(adjX, adjY):
                    q.append((adjX, adjY))
                    visited[x][y] = True


if __name__ == "__main__":
    matrix = get_input("day-10/part-1.txt")
    N, M = matrix.shape
    visited = np.array([[False for _ in range(N)] for _ in range(M)])
    # start = predict_starting_point(matrix)
    start = [2, 0]
    bfs(*start)
    print()
    pprint(visited)
    print((visited.sum() // 2) + 1)
