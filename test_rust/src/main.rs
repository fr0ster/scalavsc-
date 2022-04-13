fn main() {
    println!("Hello, world!");
    println!("{}", foo::foo(20, 22));
}

mod foo {
    use std::ops::*;

    pub fn foo<T>(a: T, b: T) -> T
    where
        T: Add<Output = T> + Copy + Clone + Div<Output = T> + Mul<Output = T>,
    {
        return a + b;
    }
}
