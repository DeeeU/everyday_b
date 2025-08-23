# n, *rest = map(int, input().split())
# hash_dict = dict(zip([6,8,12], rest))

# if (min(hash_dict, key=hash_dict.get)) // n == 0:
#   print(min(hash_dict, key=hash_dict.get))
# else:
#   print(min(hash_dict, key=hash_dict.get) // n * min(hash_dict, key=hash_dict.get))


n, s, m, l = map(int, input().split())
ans = float('inf')
for i in range(101):
  for j in range(101):
    for k in range(101):
      if i*6 + j*8 + k*12 >= n:
        ans = min(ans, i*s + j*m + k*l)
print(ans)
