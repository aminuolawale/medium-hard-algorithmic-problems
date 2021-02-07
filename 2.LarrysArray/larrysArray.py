def larrysArray(A):
    """ """
    a = A
    total = 0
    for i, p in enumerate(a):
        count = 0
        for j, q in enumerate(a[i + 1 :]):
            if p > q:
                count += 1
        total += count
    return "YES" if total % 2 == 0 else "NO"


if __name__ == "__main__":
    A = [1, 2, 3, 5, 4]
    result = larrysArray(A)
    print("this is the result", result)