

list1  = [1,3,3,6]
list2 = [3,4,7,1]

firstList = set(list1)
secList = set(list2)

dup = secList - firstList

result = firstList.union(list(dup))
print(result)


# get value from terminal
text = input("digite seu nome\n")
 #check the name is a palindrome
if text == text[::-1]:
    print("é um palindrome")
else:
    print("não é um palindrome")




