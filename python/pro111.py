n = int(input())
while n <= 999 :
  if list(str(n)[0]) == list(str(n)[1]) == list(str(n)[2]):
    print(n)
    exit()
  else:
    n += 1
