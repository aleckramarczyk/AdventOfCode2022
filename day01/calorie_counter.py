file = "input"
fhandle = open(file)

elves = {}

elf_number = 0
current_calories = 0

for line in fhandle:
    if line == '\n':
        elves[elf_number] = current_calories
        elf_number += 1
        current_calories = 0
    else:
        current_calories += int(line.strip())

current_max = 0
for elf in elves:
    if elves[elf] > current_max:
        current_max = elves[elf]

print(f"The elf with the most calories has {current_max}")

top_three_calories = 0

for i in range(3):
    current_max = 0
    for elf in elves:
        if elves[elf] > current_max:
            current_max = elves[elf] 
    top_three_calories += current_max
    for key, value in elves.copy().items():
        if value == current_max:
            elves.pop(key)

print(f"The top three elves have {top_three_calories} calories total")