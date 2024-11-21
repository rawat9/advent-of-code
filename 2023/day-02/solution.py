from collections import namedtuple, defaultdict

Red = namedtuple("Red", ["count", "color"])
Green = namedtuple("Green", ["count", "color"])
Blue = namedtuple("Blue", ["count", "color"])


def get_input(path: str) -> list[str]:
    with open(path, mode="r") as file:
        data = file.read().split("\n")

    return data


info = defaultdict(list[Red | Green | Blue])

MAX_LIMITS = {"red": 12, "green": 13, "blue": 14}


def buildDict(games: list[str]):
    for game in games:
        g, sets = game.split(":")
        _, id = g.split()

        for set in sets.split(";"):
            for r in set.strip().split(", "):
                count, type = r.split()

                match type:
                    case "blue":
                        info[id].append(Blue(int(count), "blue"))
                    case "red":
                        info[id].append(Red(int(count), "red"))
                    case "green":
                        info[id].append(Green(int(count), "green"))

    return info


def calc_possible_games() -> int:
    possible = 0
    for g, cubes in info.items():
        valid = True
        for cube in cubes:
            for color, c in MAX_LIMITS.items():
                if cube.color == color and int(cube.count) > c:
                    valid = False

        if valid:
            possible += int(g)

    return possible


def calc_fewest() -> int:
    total = 0
    for cones in info.values():
        red, blue, green = 0, 0, 0
        for cone in cones:
            match cone.color:
                case "red":
                    red = max(red, cone.count)
                case "blue":
                    blue = max(blue, cone.count)
                case "green":
                    green = max(green, cone.count)

        total += red * blue * green

    return total


if __name__ == "__main__":
    input_data = get_input("day-02/input.txt")
    buildDict(input_data)
    print("PART 1 -", calc_possible_games())
    print("PART 2 -", calc_fewest())
