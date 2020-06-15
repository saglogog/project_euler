# The idea is to solve in reverse: start with the largest possible palindrome (999.999) 
# and go to the immediately lower: 998.899 etc. by checking each time if has two 3-digit divisors.

i = 999

while (i>=100):
	i = i-1
	y = int(str(i) + str(i)[::-1])
#	print (y)
	i2 = 100
	while (y%i2 != 0):
		i2 = i2+1
	y2 = y/i2
	if (len (str(i2)) == len (str(y2))-2 and len(str(i2)) == 3):
		print (i2, y2, y)
	#	print (y)

# I probably have made a mistake with the whiles on the above code as 
# the largest palindrome I get is the one made up from 863 * 869 = 749947
# Thus run the below code and get largest possible palindrome which is 906609 = 993 * 913 and the correct answer


i1 = 863
i2 = 869

for i1 in range(863, 1000):
    for i2 in range(869, 1000):
        x = i2 * i1
        x = str(x)
        if(x[:3] == x[5]+x[4]+x[4]):
            print (x, i1, i2)
