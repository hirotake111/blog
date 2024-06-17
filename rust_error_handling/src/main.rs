use axum::{routing::get, Router};

fn main() {
    let a = 4;
    let b = double(a);
    let app = Router::new().route("/", get(root));
}

fn root() -> String {
    "hello world".to_string()
}

fn double(n: i32) -> i32 {
    n * 2
}
