s = list(input())
for i in "abcdefghijklmnopqrstuvwxyz":
  if i not in s:
      print(i)
      break
else:
  print("None")
