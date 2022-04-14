mod lib;
mod fib;

fn main() {
    println!("Hello, world!");
    println!("{}", lib::foo::foo(20, 22));
    println!("{}", lib::bar::bar(10,32));
    println!("{}", fib::bar::bar(10,32));
}