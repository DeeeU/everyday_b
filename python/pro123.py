t=0
s=10
for i in range(5):
    k=int(input())
    if k%10==0:
        t+=k
    else:
        t+=((k//10)+1)*10
        if k%10<s:
            s=k%10

if s!=10:
    print(t-10+s)
else:
    print(t)
