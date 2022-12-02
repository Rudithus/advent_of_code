use std::{
    borrow::{Borrow, BorrowMut},
    fs::{read_to_string, File},
    io::{BufRead, BufReader, Lines, Read},
    path::Path,
};

struct Elf {
    inventory: Vec<u32>,
    total: u32,
}
impl Elf {
    fn new() -> Elf {
        Elf {
            inventory: vec![],
            total: 0,
        }
    }
    fn add_to_inv(&mut self, calorie: u32) {
        self.total += calorie;
        self.inventory.push(calorie);
    }
}
fn main() {
    let file = File::open("input.txt").unwrap();
    let mut lines = BufReader::new(file).lines();

    q2(lines.borrow_mut());
}

fn q1(lines: &mut Lines<BufReader<File>>) {
    let mut max_elf = build_elf(lines.borrow_mut()).unwrap();
    while let Some(elf) = build_elf(lines.borrow_mut()) {
        if elf.total > max_elf.total {
            max_elf = elf;
        }
    }
    println!("{}", max_elf.total);
}

fn build_elf(lines: &mut Lines<BufReader<File>>) -> Option<Elf> {
    let mut elf = Elf::new();
    for line in lines {
        let line = line.unwrap();
        if line == "" {
            return Some(elf);
        }
        let calorie = line.parse::<u32>().unwrap();
        elf.add_to_inv(calorie);
    }
    None
}
fn q2(lines: &mut Lines<BufReader<File>>) {
    let mut top_elfs = vec![];
    top_elfs.push(build_elf(lines.borrow_mut()).unwrap());
    top_elfs.push(build_elf(lines.borrow_mut()).unwrap());
    top_elfs.push(build_elf(lines.borrow_mut()).unwrap());
    top_elfs.sort_by(|e1, e2| e2.total.cmp(&e1.total));

    while let Some(elf) = build_elf(lines.borrow_mut()) {
        if elf.total > top_elfs[2].total {
            top_elfs.pop();
            if top_elfs[1].total > elf.total {
                top_elfs.push(elf);
            } else {
                if top_elfs[0].total > elf.total {
                    top_elfs.insert(1, elf);
                } else {
                    top_elfs.insert(0, elf);
                }
            }
        }
    }
    let sum: u32 = top_elfs.iter().map(|e| e.total).sum();
    println!("{}", sum);
}
