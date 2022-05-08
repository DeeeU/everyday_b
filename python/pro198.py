n = input()
for i in range(len(n)):
    a = i * "0" + n
    if n == n[::-1]:
        print("yes")
        exit()

print("no")
