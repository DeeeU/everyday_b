n = int(input())
s = int(n ** 0.5)
ans = set()
for i in range(2, s+1):
    a = i * i
    while a <= n:
        ans.add(a)
        a *= i
print(n - len(ans))
