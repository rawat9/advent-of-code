from itertools import pairwise


def get_input(path: str) -> list[list[int]]:
    with open(path, mode="r") as file:
        data = [list(map(int, line.split())) for line in file.read().split("\n")]

    return data


def get_next_value(history: list[int], backwards=False):
    def inner(history, n=0):
        if all(0 == h for h in history):
            return n

        if backwards:
            sequence = [(i - j) for i, j in pairwise(history)]
            n += sequence[0]
        else:
            sequence = [(j - i) for i, j in pairwise(history)]
            n += sequence[-1]

        return inner(sequence, n)

    return (history[0] if backwards else history[-1]) + inner(history)


if __name__ == "__main__":
    input_data = get_input("day-09/input.txt")
    print("PART 1 -", sum(map(get_next_value, input_data)))
    print("PART 2 -", sum(map(lambda x: get_next_value(x, True), input_data)))
