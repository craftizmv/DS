def cut_into_segments_recursive(n, a, b, c):
    def solve(n, a, b, c):
        # base case
        if n == 0:
            return 0
        if n < 0:
            return -1

        answer1 = solve(n - a, a, b, c)
        answer2 = solve(n - b, a, b, c)
        answer3 = solve(n - c, a, b, c)

        max_answer = max(answer1, answer2, answer3)
        if max_answer == -1:
            return -1
        return max_answer + 1

    max_segment_possible = solve(n, a, b, c)
    if max_segment_possible < 0:
        return