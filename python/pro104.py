s = input()
count_c = 0

if s[0] == 'A':
  if s[2:-1].count('C') == 1:
    for i in range(1, len(s)):
      if 97 <= ord(s[i]) <= 122:
        continue
      else:
        count_c += 1
      if count_c > 1:
        print("WA")
  else:
    print("WA")

else:
  print("WA")

if count_c == 1:
  print("AC")
