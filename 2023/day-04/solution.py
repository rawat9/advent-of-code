def get_input(path: str) -> list[str]:
    with open(path, mode="r") as file:
        data = file.read().strip().split("\n")

    return data


def calc_points(data: list[str]):
    total_points = 0
    cards = [1] * len(data)

    for index, d in enumerate(data):
        _, lists = d.split(": ")
        winning_numbers, numbers = lists.split(" | ")
        common = set(winning_numbers.split()).intersection(numbers.split())

        if common:
            total_points += 2 ** (len(common) - 1)

        for i in range(len(common)):
            cards[index + i + 1] += cards[index]

    return total_points, sum(cards)


if __name__ == "__main__":
    input_data = get_input("day-04/input.txt")
    part1, part2 = calc_points(input_data)
    print("PART 1 -", part1)
    print("PART 2 -", part2)
