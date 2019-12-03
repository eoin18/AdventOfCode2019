import math

def calc_fuel(input):
	return math.floor(input / 3) - 2

def main():
	result = 0
	with open('input.txt') as file:
		line = file.readline()
		while line:
			fuel = calc_fuel(int(line))
			#part2
			while fuel > 0:
				result += fuel
				fuel = calc_fuel(fuel)
			line = file.readline()
	print result

if __name__ == "__main__":
	main()
