use aoc_runner_derive::aoc;

#[aoc(day1, part1)]
fn part1(input: &str) -> i64 {
    let mut max = 0;
    let mut current = 0;
    for i in input.lines() {
        if i.is_empty() {
            if max < current {
                max = current
            }
            current = 0;
            continue;
        }
        current += i.parse::<i64>().unwrap();
    }

    return max;
}

#[aoc(day1, part2)]
fn part2(input: &str) -> i64 {
    let mut acc = vec![];
    let mut curr = 0;

    for i in input.lines() {
        if i.is_empty() {
            acc.push(curr);
            curr = 0;
            continue;
        }
        curr += i.parse::<i64>().unwrap();
    }

    acc.sort_by(|a, b| b.cmp(a));

    return acc[0] + acc[1] + acc[2];
}
