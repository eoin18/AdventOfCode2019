import math

def execute_intcode(noun, verb, intcode):
	intcode[1] = noun
	intcode[2] = verb
	for i in range(0, len(intcode), 4):
		opcode = int(intcode[i])
		if opcode == 99:
			break
		else:
			operand1 = int(intcode[i + 1])
			operand2 = int(intcode[i + 2])
			target = int(intcode[i + 3])
			if opcode == 1:
				intcode[target] = int(intcode[operand1]) + int(intcode[operand2])
			if opcode == 2:
				intcode[target] = int(intcode[operand1]) * int(intcode[operand2])
	return intcode[0]

def main():
	with open('input.txt') as file:
		line = file.readline()
		while line:
			for i in range(0, 99):
				for j in range(0, 99):
					intcode = line.split(',')
					result = execute_intcode(i, j, intcode)
					if result == 19690720:
						print str((100*i) + j)



if __name__ == "__main__":
	main()
