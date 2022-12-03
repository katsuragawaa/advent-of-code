use aoc_runner_derive::aoc;

#[aoc(day2, part1)]
fn part1(input: &str) -> i64 {
    return input
        .lines()
        .map(|line| match line {
            "A X" => return 4,
            "B X" => return 1,
            "C X" => return 7,
            "A Y" => return 8,
            "B Y" => return 5,
            "C Y" => return 2,
            "A Z" => return 3,
            "B Z" => return 9,
            "C Z" => return 6,
            &_ => return 0,
        })
        .sum();
}

#[aoc(day2, part2)]
fn part2(input: &str) -> i64 {
    return input
        .lines()
        .map(|line| match line {
            "A X" => return 3,
            "B X" => return 1,
            "C X" => return 2,
            "A Y" => return 4,
            "B Y" => return 5,
            "C Y" => return 6,
            "A Z" => return 8,
            "B Z" => return 9,
            "C Z" => return 7,
            &_ => return 0,
        })
        .sum();
}
