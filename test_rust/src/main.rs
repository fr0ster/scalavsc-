// use std::ops::Add;
use std::ops::*;

fn main() {
    println!("Hello, world!");
    println!("{}", foo(20,22));
}

fn foo<T>(a: T, b: T) -> T 
where T: Add<Output=T> + Copy + Clone + Div<Output=T> + Mul<Output=T>
{
    return a + b;
}

// fn foo<T:Add>(x : T) -> T{
//     return x + x;
// }
