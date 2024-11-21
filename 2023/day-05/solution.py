def get_input(path: str) -> list[str]:
    with open(path, mode="r") as file:
        data = file.read().split("\n\n")

    return data


def getMaps(data: list[str]) -> list[list[str]]:
    return [data[i].split(":")[1].lstrip().split("\n") for i in range(1, len(data))]


def solve(data: list[str]):
    seeds = list(map(int, data[0].split(":")[1].split()))

    all_maps = getMaps(data)

    lowest = seeds[0]
    for seed in seeds:
        for mapp in all_maps:
            newSeed = get(seed, mapp)
            seed = newSeed

        lowest = min(lowest, seed)

    return lowest


def get(seed: int, mapping: list[str]):
    for destination, source, length in map(str.split, mapping):
        if int(source) <= seed < int(source) + int(length):
            return seed - int(source) + int(destination)

    return seed


if __name__ == "__main__":
    input_data = get_input("day-05/part-1.txt")
    print(solve(input_data))
