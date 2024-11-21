from collections import defaultdict
from math import lcm
from itertools import cycle


def get_input(path: str) -> list[str]:
    with open(path, mode="r") as file:
        data = file.read().split("\n")

    return data


def build_network(data: list[str]):
    network = defaultdict(list)

    for i in range(2, len(data)):
        node, l_r = data[i].split(" = ")
        left, right = l_r.removeprefix("(").removesuffix(")").split(", ")
        network[node] = [left, right]

    return network


def stepsUntilZZZ(start, instructions):
    destination = "ZZZ"

    curr_node = start
    count = 0

    for dir in cycle(instructions):
        if curr_node == destination:
            return count

        if dir == "R":
            curr_node = network[curr_node][1]

        if dir == "L":
            curr_node = network[curr_node][0]

        count += 1

    return network


def stepsUntilEndswithZ(node, instructions):
    count = 0
    curr_node = node

    for dir in cycle(instructions):
        if curr_node.endswith("Z"):
            return count

        if dir == "R":
            curr_node = network[curr_node][1]

        if dir == "L":
            curr_node = network[curr_node][0]

        count += 1

    return 0


if __name__ == "__main__":
    input_data = get_input("day-08/input.txt")
    network = build_network(input_data)

    print("PART 1 -", stepsUntilZZZ("AAA", input_data[0]))

    nodes_endswith_a = list(filter(lambda node: node.endswith("A"), network.keys()))
    print(
        "PART 2 -",
        lcm(*[stepsUntilEndswithZ(node, input_data[0]) for node in nodes_endswith_a]),
    )
