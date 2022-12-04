use std::{
    collections::HashSet,
    fs::File,
    io::{BufRead, BufReader, Lines},
};

fn main() {
    let file = File::open("input.txt").unwrap();
    let lines = BufReader::new(file).lines();

    q2(lines);
}

fn calculate_priority(byte: u8) -> u8 {
    if byte > 96 {
        byte - 96
    } else {
        byte - 38
    }
}

fn q1(lines: Lines<BufReader<File>>) {
    let mut total_priority: u32 = 0;
    for line in lines {
        let mut set: HashSet<u8> = HashSet::new();

        let line = line.unwrap();
        let chars = line.as_bytes();
        let (first_compartment, second_compartment) = chars.split_at(chars.len() / 2);

        for item in first_compartment {
            set.insert(*item);
        }

        for item in second_compartment {
            if set.insert(*item) {
                set.remove(item);
            } else {
                set.remove(item);
                total_priority += calculate_priority(*item) as u32;
                continue;
            }
        }
    }
    println!("{}", total_priority);
}
fn q2(lines: Lines<BufReader<File>>) {
    let mut total_priority: u32 = 0;

    let mut curr_group = vec![];
    for line in lines {
        curr_group.push(line.unwrap());
        if curr_group.len() == 3 {
            let s1: HashSet<u8> = curr_group[0].as_bytes().iter().cloned().collect();
            let mut s2: HashSet<u8> = HashSet::new();

            // ?????
            // let s2: HashSet<u8> = curr_group[1]
            //     .as_bytes()
            //     .iter()
            //     .cloned()
            //     .take_while(|p| s1.contains(p))
            //     .collect();

            for c in curr_group[1].as_bytes() {
                if s1.contains(c) {
                    s2.insert(*c);
                }
            }

            for c in curr_group[2].as_bytes() {
                if s2.contains(c) {
                    total_priority += calculate_priority(*c) as u32;

                    break;
                }
            }

            curr_group.clear();
        }
    }

    println!("{}", total_priority);
}
