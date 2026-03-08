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

def printrev(n):
    if n <= 0:
        return
    print(n)
    printrev(n-1)

printrev(20)
