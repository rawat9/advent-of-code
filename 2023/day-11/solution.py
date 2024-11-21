from itertools import combinations
import numpy as np


def get_input(path: str):
    with open(path, mode="r") as file:
        data = np.matrix(data=[list(line) for line in file.read().split("\n")])

    return data


calc_distance = lambda p1, p2: abs(p2[0] - p1[0]) + abs(p2[1] - p1[1])


def find_empty():
    rows = np.where(np.all(matrix == ".", axis=1))[0].tolist()  # rows
    cols = np.where(np.all(matrix == ".", axis=0))[1].tolist()  # columns
    return rows, cols


def find_galaxies():
    indices = np.where(matrix == "#")
    return {galaxy: v for galaxy, v in enumerate(zip(indices[0], indices[1]), start=1)}


def update_points(coefficient=1):
    rows, cols = find_empty()
    original = list(map(list, find_galaxies().values()))

    for index in original:
        r = index[0]
        index[0] += len(list(filter(lambda x: x < r, rows))) * coefficient

        c = index[1]
        index[1] += len(list(filter(lambda x: x < c, cols))) * coefficient

    indices.update(
        {galary: newValue for galary, newValue in enumerate(original, start=1)}  # type: ignore
    )
    return indices


def calculate_shortest_path():
    return sum(
        map(
            lambda p: calc_distance(indices[p[0]], indices[p[1]]),
            combinations(range(1, len(indices.keys()) + 1), 2),
        )
    )


if __name__ == "__main__":
    matrix = get_input("day-11/input.txt")
    N, M = matrix.shape
    indices = find_galaxies()
    update_points()
    print("PART 1 -", calculate_shortest_path())
    update_points(1000000 - 1)
    print("PART 2 -", calculate_shortest_path())
