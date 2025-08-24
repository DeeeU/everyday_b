# G < M

K,G,M = map(int, input().split())
ans_G, ans_M = 0, 0
for i in range(K):
  # グラスが水で満たされているとき、すなわちグラスにちょうど G ml 入っているとき、グラスの水をすべて捨てる。
  # そうでなく、マグカップが空であるとき、マグカップを水で満たす。
  # 上のいずれでもないとき、マグカップが空になるかグラスが水で満たされるまで、マグカップからグラスに水を移す。
  if ans_G == G:
    ans_G = 0
  elif ans_M == 0:
    ans_M = M
  else:
    if G < ans_G + ans_M:
      ans_G, ans_M = G, ans_G + ans_M - G
    else:
      ans_G, ans_M = ans_M, 0

print(ans_G, ans_M)
