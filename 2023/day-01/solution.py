import re


def get_input(path: str) -> list[str]:
    with open(path, mode="r") as file:
        data = file.read().split("\n")

    return data


def calibrationSum(values: list[str]) -> int:
    return sum(map(getNumber, values))


def convertToDigits(value: str) -> str:
    for k, v in digits.items():
        value = value.replace(k, v)

    return value


def getNumber(value: str) -> str:
    nums = re.findall("\d", value)
    return int(nums[0] + nums[-1])


digits = {
    "one": "o1e",
    "two": "t2o",
    "three": "t3e",
    "four": "f4r",
    "five": "f5e",
    "six": "s6x",
    "seven": "s7n",
    "eight": "e8t",
    "nine": "n9e",
}

if __name__ == "__main__":
    input_data = get_input("day-01/input.txt")
    print("PART 1 -", calibrationSum(input_data))
    print("PART 2 -", calibrationSum(map(convertToDigits, input_data)))
