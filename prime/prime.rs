//use std::int::*;
//use std::os::*;
use std::vec::*;
//use std::str::*;

fn generate(num: int) -> ~[int] {
    let mut primes = ~[];
    let mut n = num;
    let mut candidate = 2;

    while n > 1 {
        while n%candidate == 0 {
            primes.push(candidate);
            n /= candidate;
        }
        candidate += 1;
    }

    return primes
}

#[test]
fn testOne() {
    let expected = ~[];
    let actual = generate(1);
    assert!(expected == actual,
            "Expected %s, got %s",
            vecToString(expected),
            vecToString(actual));
}

#[test]
fn testTwo() {
    let expected = ~[2];
    let actual = generate(2);
    assert!(expected == actual,
            "Expected %s, got %s",
            vecToString(expected),
            vecToString(actual));
}

#[test]
fn testThree() {
    let expected = ~[3];
    let actual = generate(3);
    assert!(expected == actual,
            "Expected %s, got %s",
            vecToString(expected),
            vecToString(actual));
}

#[test]
fn testFour() {
    let expected = ~[2,2];
    let actual = generate(4);
    assert!(expected == actual,
            "Expected %s, got %s",
            vecToString(expected),
            vecToString(actual));
}

#[test]
fn testSix() {
    let expected = ~[2,3];
    let actual = generate(6);
    assert!(expected == actual,
            "Expected %s, got %s",
            vecToString(expected),
            vecToString(actual));
}

#[test]
fn testEight() {
    let expected = ~[2,2,2];
    let actual = generate(8);
    assert!(expected == actual,
            "Expected %s, got %s",
            vecToString(expected),
            vecToString(actual));
}

#[test]
fn testNine() {
    let expected = ~[3,3];
    let actual = generate(9);
    assert!(expected == actual,
            "Expected %s, got %s",
            vecToString(expected),
            vecToString(actual));
}

fn vecToString(ints: ~[int]) -> ~str {
    let mut retStr = std::str::with_capacity(ints.len()*2+1 as uint);
    retStr.push_str("[");

    for i in ints.iter() {
        retStr.push_str(fmt!("%d",*i));
        retStr.push_str(",");
    }

    retStr.pop_char();
    retStr.push_str("]");
    retStr
} 

/*
fn list(ints: ~[int]) -> ~[int] {
    let mut expect = with_capacity(ints as uint);
    let mut i int = 0;
    while 

    return ~[1];
}
*/