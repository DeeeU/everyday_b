a,b = map(int,input().split())
t= 1
count = 0
while t < b:
  t += a - 1
  count += 1
print(count)
