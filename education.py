sum_positive = 0

while True:
    num = int(input())
    if num == 0:
        break
    if num > 0:
        sum_positive += num

if sum_positive > 0:
    print(sum_positive)
else:
    print(0)