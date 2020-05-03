x = int(int(input()))
base = 100
count = 0
while(x > base):
    base += base // 100
    count += 1
print(count)
