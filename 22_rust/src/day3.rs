use std::vec;

use aoc_runner_derive::aoc;

#[aoc(day3, part1)]
fn part1(input: &str) -> u32 {
    return input
        .lines()
        .map(|line| {
            let (left, right) = line.split_at(line.len() / 2);
            let mut same: Vec<char> = left
                .chars()
                .filter_map(|c| right.contains(c).then(|| c))
                .collect();

            same.dedup();

            return same.iter().fold(0, |acc, c| {
                let c_num = *c as u32;
                if c_num <= 90 {
                    return acc + c_num - 64 + 26;
                }

                return acc + c_num - 96;
            });
        })
        .sum();
}

#[aoc(day3, part2)]
fn part2(input: &str) -> u32 {
    let mut acc = vec![];
    for (idx, _) in input.lines().enumerate() {
        if idx % 3 != 0 {
            continue;
        }

        let x = input.lines().nth(idx).unwrap();
        let y = input.lines().nth(idx + 1).unwrap();
        let z = input.lines().nth(idx + 2).unwrap();

        let mut xy: Vec<char> = x.chars().filter_map(|c| y.contains(c).then(|| c)).collect();
        xy.dedup();

        let mut same: Vec<char> = xy
            .iter()
            .filter_map(|c| z.contains(*c).then(|| *c))
            .collect();
        same.dedup();

        acc.append(&mut same);
    }

    return acc.iter().fold(0, |acc, c| {
        let c_num = *c as u32;
        if c_num <= 90 {
            return acc + c_num - 64 + 26;
        }

        return acc + c_num - 96;
    });
}
