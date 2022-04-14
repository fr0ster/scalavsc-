pub mod bar {
    use std::ops::*;

    pub fn bar<T>(a: T, b: T) -> T
    where
        T: Add<Output = T> + Copy + Clone + Div<Output = T> + Mul<Output = T>,
    {
        return a + b;
    }
}
