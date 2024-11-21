from math import prod


def get_input(path: str) -> list[str]:
    with open(path, mode="r") as file:
        data = file.read().split("\n")

    return data


def calculate_ways(time: int, record: int):
    distance = 0
    ways = 0
    for t in range(time):
        remaining = time - t
        distance = t * remaining

        if distance > record:
            ways += 1

    return ways


if __name__ == "__main__":
    input_data = get_input("day-06/input.txt")

    time, distance = (
        input_data[0].split(":")[1].split(),
        input_data[1].split(":")[1].split(),
    )
    print("PART 1 -", prod(map(calculate_ways, map(int, time), map(int, distance))))

    time = int("".join(time))
    distance = int("".join(distance))
    print("PART 2 -", prod(map(calculate_ways, [time], [distance])))
