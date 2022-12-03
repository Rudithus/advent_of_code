use std::{
    char,
    fs::File,
    io::{BufRead, BufReader, Lines, Read},
};

enum Move {
    Rock = 1,
    Paper = 2,
    Scissor = 3,
}

enum Outcome {
    Lose = 0,
    Draw = 3,
    Win = 6,
}

fn main() {
    let file = File::open("input.txt").unwrap();
    let mut lines = BufReader::new(file).lines();

    q2(lines);
}

fn map_opponent(c: char) -> Result<Move, String> {
    match c {
        'A' => Ok(Move::Rock),
        'B' => Ok(Move::Paper),
        'C' => Ok(Move::Scissor),
        _ => Err(format!("wtf is {}", c)),
    }
}
fn map_own(c: char) -> Result<Move, String> {
    match c {
        'X' => Ok(Move::Rock),
        'Y' => Ok(Move::Paper),
        'Z' => Ok(Move::Scissor),
        _ => Err(format!("wtf is {}", c)),
    }
}

fn map_to_outcome(c: char) -> Result<Outcome, String> {
    match c {
        'X' => Ok(Outcome::Lose),
        'Y' => Ok(Outcome::Draw),
        'Z' => Ok(Outcome::Win),
        _ => Err(format!("wtf is {}", c)),
    }
}
fn play_q1(m1: &Move, m2: &Move) -> Outcome {
    match (m1, m2) {
        (Move::Rock, Move::Rock) => Outcome::Draw,
        (Move::Paper, Move::Paper) => Outcome::Draw,
        (Move::Scissor, Move::Scissor) => Outcome::Draw,
        (Move::Rock, Move::Scissor) => Outcome::Win,
        (Move::Paper, Move::Rock) => Outcome::Win,
        (Move::Scissor, Move::Paper) => Outcome::Win,
        (Move::Rock, Move::Paper) => Outcome::Lose,
        (Move::Paper, Move::Scissor) => Outcome::Lose,
        (Move::Scissor, Move::Rock) => Outcome::Lose,
    }
}

fn play_q2(m: &Move, outcome: &Outcome) -> Move {
    match (m, outcome) {
        (Move::Rock, Outcome::Draw) => Move::Rock,
        (Move::Paper, Outcome::Draw) => Move::Paper,
        (Move::Scissor, Outcome::Draw) => Move::Scissor,
        (Move::Rock, Outcome::Lose) => Move::Scissor,
        (Move::Paper, Outcome::Lose) => Move::Rock,
        (Move::Scissor, Outcome::Lose) => Move::Paper,
        (Move::Rock, Outcome::Win) => Move::Paper,
        (Move::Paper, Outcome::Win) => Move::Scissor,
        (Move::Scissor, Outcome::Win) => Move::Rock,
    }
}

fn q1(lines: Lines<BufReader<File>>) {
    let mut total_score = 0;
    for line in lines {
        let line = line.unwrap();
        let chars = line.as_bytes();

        let opponent_move = map_opponent(chars[0] as char).unwrap();
        let own_move = map_own(chars[2] as char).unwrap();
        let outcome = play_q1(&own_move, &opponent_move);

        let score = outcome as u32 + own_move as u32;
        total_score += score;
    }
    println!("{}", total_score);
}

fn q2(lines: Lines<BufReader<File>>) {
    let mut total_score = 0;
    for line in lines {
        let line = line.unwrap();
        let chars = line.as_bytes();

        let opponent_move = map_opponent(chars[0] as char).unwrap();
        let desired_outcome = map_to_outcome(chars[2] as char).unwrap();
        let own_move = play_q2(&opponent_move, &desired_outcome);

        let score = desired_outcome as u32 + own_move as u32;
        total_score += score;
    }
    println!("{}", total_score);
}
