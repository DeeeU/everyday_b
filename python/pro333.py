def get_distance_type(v1, v2):
    if v1 == v2:
        return 0  # 同じ点

    diff = abs(vertices[v1] - vertices[v2])
    # 円周上なので、逆回りも考慮
    min_diff = min(diff, 5 - diff)

    if min_diff == 1:
        return 1  # 隣接（辺の長さ）
    elif min_diff == 2:
        return 2  # 対角線の長さ
    else:
        return 0  # 同じ点（念のため）

s1,s2 = map(str, input())
t1,t2 = map(str, input())

vertices = {'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4}

ans_s, ans_t = abs(hash_dict[s1] - hash_dict[s2]) % 2, abs(hash_dict[t1] - hash_dict[t2]) % 2
if ans_s == ans_t:
  print('Yes')
else:
  print('No')

# 距離タイプを比較
distance_s = get_distance_type(s1, s2)
distance_t = get_distance_type(t1, t2)

if distance_s == distance_t:
    print('Yes')
else:
    print('No')
