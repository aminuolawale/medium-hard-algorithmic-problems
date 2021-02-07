def nonDivisibleSubset(k, s):
    """ """
    freq_array = [0 for _ in range(k)]
    for i, v in enumerate(s):
        freq_array[v % k] += 1
    freq_array[0] = min(freq_array[0], 1)
    if k % 2 == 0:
        freq_array[k // 2] = min(freq_array[k // 2], 1)
    result = freq_array[0]
    for i, v in enumerate(freq_array[1 : k // 2 + 1]):
        result += max(v, freq_array[k - i - 1])
    return result


if __name__ == "__main__":
    S = [19, 10, 12, 10, 24, 25, 22]
    k = 4
    ans = nonDivisibleSubset(k, S)
    print("the ans", ans)