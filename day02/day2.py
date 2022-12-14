file = "input"
fhandle = open(file)

score_table = {"X": 1, "Y": 2, "Z": 3}
decode = {"A": "ROCK", "B": "PAPER", "C": "SCISSORS", "X": "ROCK", "Y": "PAPER", "Z": "SCISSORS"}
opponent_to_yours = {"ROCK": "X", "PAPER": "Y", "SCISSORS": "Z"}

part1_score = 0
part2_score = 0

def lose(opponent_raw):
    opponent = decode[opponent_raw]

    if opponent == "ROCK":
        yours = 3
    elif opponent == "PAPER":
        yours = 1 
    elif opponent == "SCISSORS":
        yours = 2 

    return 0 + yours 

def draw(opponent_raw):
    opponent = decode[opponent_raw]
    yours = opponent_to_yours[opponent]
    return 3 + score_table[yours]

def win(opponent_raw):
    opponent = decode[opponent_raw]

    if opponent == "ROCK":
        yours = 2
    elif opponent == "PAPER":
        yours = 3
    elif opponent == "SCISSORS":
        yours = 1

    return 6 + yours

def determineOutcome(opponent_raw, yours_raw):
    opponent = decode[opponent_raw]
    yours = decode[yours_raw]

    if opponent == yours:
        return 3

    if opponent == "ROCK":
        if yours == "SCISSORS":
            return 0
        elif yours == "PAPER":
            return 6
    elif opponent == "PAPER":
        if yours == "ROCK":
            return 0
        elif yours == "SCISSORS":
            return 6
    elif opponent == "SCISSORS":
        if yours == "PAPER":
            return 0
        elif yours == "ROCK":
            return 6

for line in fhandle:
    inp = line.split()
    opponents_choice = inp[0]
    your_choice = inp[1]
    outcome = inp[1]
    if outcome == "X":
        part2_score += lose(opponents_choice)
    if outcome == "Y":
        part2_score += draw(opponents_choice)
    if outcome == "Z":
        part2_score += win(opponents_choice)

    part1_score += determineOutcome(opponents_choice, your_choice) + score_table[your_choice]

print(f"Part one: {part1_score}\nPart 2: {part2_score}")