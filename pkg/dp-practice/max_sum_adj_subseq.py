
def max_sum_adj_subseq(input):
    if not input:
        return 0

    return solve_adj(input, 0)

def solve_adj(input, idx):
    if idx >= len(input):
        return 0

    # Include current element and move to idx + 2
    incl_case = input[idx] + solve_adj(input, idx + 2)

    # Exclude current element and move to idx + 1
    excl_case = solve_adj(input, idx + 1)

    return max(incl_case, excl_case)


def solve_adj(input, idx, memo):
    if idx >= len(input):
        return 0

    if memo[idx] != -1:
        return memo[idx]

    incl_case = input[idx] + solve_adj(input, idx + 2, memo)
    excl_case = solve_adj(input, idx + 1, memo)

    memo[idx] = max(incl_case, excl_case)
    return memo[idx]