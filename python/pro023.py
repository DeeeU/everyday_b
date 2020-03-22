n = int(input())
s = input()
count = (n-1)//2
acs = "b"
i = 0
while i < count:
    if i % 3 == 0:
        acs = "a" + acs + "c"
    elif i % 3 == 1:
        acs = "c" + acs + "a"
    else:
        acs = "b" + acs + "b"
    i += 1

if s == acs:
    print(count)
else:
    print(-1)
