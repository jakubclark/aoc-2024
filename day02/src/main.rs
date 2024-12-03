fn main() {
    let input = get_input();

    let mut part1 = 0;
    for record in input.iter() {
        if is_safe(record) {
            part1 += 1;
        }
    }
    println!("part1: {part1}");

    let mut part2 = 0;
    for record in input.iter() {
        if is_safe(record) {
            part2 += 1;
            continue;
        }
        for i in 0..record.len() {
            let filtered = [&record[..i], &record[i + 1..]].concat();
            if is_safe(&filtered) {
                part2 += 1;
                break;
            }
        }
    }
    println!("part2: {part2}");
}

fn is_safe(record: &[i64]) -> bool {
    let all_decreasing = record
        .iter()
        .zip(record.iter().skip(1))
        .all(|(prev, curr)| prev > curr);

    let all_increasing = record
        .iter()
        .zip(record.iter().skip(1))
        .all(|(prev, curr)| prev < curr);

    if all_increasing || all_decreasing {
        record
            .iter()
            .zip(record.iter().skip(1))
            .all(|(prev, curr)| {
                let delta = (prev - curr).abs();
                delta >= 1 && delta <= 3
            })
    } else {
        false
    }
}

fn get_input() -> Vec<Vec<i64>> {
    include_str!("input.txt")
        .lines()
        .map(|line| {
            line.split_whitespace()
                .map(|n| n.parse().unwrap())
                .collect()
        })
        .collect()
}
