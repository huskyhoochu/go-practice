for i in range(1,30+1):
	if i % 15 == 0:
		print("FizzBuzz")
	elif i % 3 == 0:
		print("Fuzz")
	elif i % 5 == 0:
		print("Buzz")
	else:
		print(i)