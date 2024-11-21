from collections import defaultdict
from itertools import groupby


def get_input(path: str):
    with open(path, mode="r") as file:
        data = file.readline().split(",")

    return data


def get_hash_value(string: str):
    current_value = 0

    for char in string:
        ascii_code = ord(char)
        current_value += ascii_code
        current_value *= 17
        current_value %= 256

    return current_value


def get_focusing_power(data: list[str]):
    dd = defaultdict(list)

    for string in data:
        try:
            label, focal_length = string.split("=")
            dd[label] = [get_hash_value(label), int(focal_length)]
        except ValueError:
            label = string.split("-")[0]
            if label in dd:
                dd.pop(label)

    values = sorted(dd.values(), key=lambda x: x[0])
    hashmap = {
        k: [(x[0], x[1][1]) for x in enumerate(g, 1)]
        for k, g in groupby(values, key=lambda x: x[0])
    }

    focusing_power = 0
    for box, value in hashmap.items():
        for slot, fL in value:
            focusing_power += (box + 1) * slot * fL

    return focusing_power


if __name__ == "__main__":
    input_data = get_input("day-15/input.txt")
    print("PART 1 -", sum(map(get_hash_value, input_data)))
    print("PART 2 -", get_focusing_power(input_data))
