# DSA Problems practice


# Sliding Window :

#  1. Maximum Sum Subarray of Size/Sum K

arr = [9,4,20,3,10,5]
k = 33

left = 0
count = 0
currsum = 0
for i in range(len(arr)):
    currsum += arr[i]

    #  when currsum increases than k then increase left pointer and decrease currsum by arr[left]

    while currsum > k and left <= i:
        currsum -= arr[left]
        left += 1

#  if currsum is equal to k then increase count by 1 
    if currsum == k:
        count += 1
print(count)


# Recursion

# Tail Recursion

def printrev(n):
    if n <= 0: # base condition
        return
    print(n) # execution statement
    printrev(n-1) # recursive call

printrev(20)

# Head Recursion

def printrev(n):
    if n <= 0: # base condition
        return
    printrev(n-1) # recursive call
    print(n) 

printrev(20) 

# Factorial of a number using recursion

def factorial(n):
    if n == 1:
        return 1
    return n * factorial(n-1)
print(factorial(5))

# Fibonacci Series using recursion

def fibonacci(n):
    if n <= 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fibonacci(n-1) + fibonacci(n-2)  
print(fibonacci(10))



