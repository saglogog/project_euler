# import sys

# divident_arg = int(sys.argv[1]) 
# divisor_arg = int(sys.argv[2])

# print (divident)
# print (divisor)

# https://en.wikipedia.org/wiki/Divisibility_rule

# it is assumed that all possible divisors will be checked for 
# the exercise, thus divisors that depend on others are not checked to save time
def check_no_divisibility(divident, divisor):
	# divisibility_ind = True

	if divisor == 2:
		# to check last digit get the quotient of the division by ten
		if (divident%10==0) or (divident%10==2) or (divident%10==4) or  (divident%10==6) or (divident%10==8): pass
		else: return False
	if divisor == 3:
		i = 0
		for digit in str(divident):
			i += int(digit)
		if i%3 == 0:
			pass
		else: return False
	if divisor == 4:
		if (divident%100)%4==0: pass
		else: return False
	if divisor == 5:
		if (divident%10==0) or divident%10 == 5: pass
		else: return False
	
	# if it is divided by 2 and 3 then it is divided by 6
	
	if divisor == 7:
		if divident%7==0: pass
		else: return False	
	if divisor == 8:
		if (divident%1000)%8==0: pass
		else: return False
	if divisor == 9:
		i = 0
		for digit in str(divident):
			i += int(digit)
		if i%9 == 0:
			pass
		else: return False
	if divisor == 10:
		if divident%10 == 0: pass
		else: return False
	if divisor == 11:
		if divident%11==0: pass
		else: return False

	# if it is divisible by 3 and 4 it is divisible by 12
	
	if divisor == 13:
		if divident%13==0: pass
		else: return False

	# div by 14 if div by 2, 7
	# div by 15 if div by 3, 5

	if divisor == 16:
		if (divident%10000)%16: pass
		else: return False
	if divisor == 17:
		if (2*(divident//100) - (divident%100))%17==0: pass
		else: return False
	
	# div by 18 if div by 2, 9

	if divisor == 19:
		if (divident//100 + 4*(divident%100))%19==0: pass
		else: return False
	
	if divisor == 20:
		if int(str(divident)[-2])%2==0: pass
		else: return False	

	return True


# print (str(check_no_divisibility (divident_arg, divisor_arg)))

def check_no_divisibility_2(divident):
	

def find_small_div():
	# start from 2520 which is definetely minimum
	cur_no = 232700000
	found = False
	while found == False:
		cur_no += 1
		print (cur_no)
		for i in range (1 ,22):
			if check_no_divisibility(cur_no, i) == False:
				break
			if i == 21:
				found=True
				return cur_no
	return "no_dice"

print (str(find_small_div()))
		 
		
