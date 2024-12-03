import re

with open("input.txt") as f:
    input = f.read().strip()

def nums(s):
    return list(map(int, re.findall(r"\d+", s)))

part1 = 0
ops = re.findall(r"mul\(\d+,\d+\)", input)
for op in ops:
    l, r = nums(op)
    part1 += l * r
print(f"part1: {part1}")

part2 = 0
ops = re.findall(r"mul\(\d+,\d+\)|do\(\)|don't\(\)", input)
include = True
for op in ops:
    if op == "do()":
        include = True
    elif op == "don't()":
        include = False
    else:
        if include:
            l, r = nums(op)
            part2 += l * r
print(f"part2: {part2}")