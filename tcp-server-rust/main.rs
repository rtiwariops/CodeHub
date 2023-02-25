use serde::{Deserialize, Serialize};
use std::io::{BufRead, BufReader};
use std::net::{TcpListener, TcpStream};
use std::thread;

#[derive(Serialize, Deserialize)]
struct Employee {
    first_name: String,
    last_name: String,
    employee_id: u32,
}

fn main() -> std::io::Result<()> {
    println!("**************** TCP Server ********************");
    println!("------------------------------------------------");

    let listener = TcpListener::bind("127.0.0.1:5566")?;
    let mut thread_vec = Vec::new();

    for stream in listener.incoming() {
        let stream = stream?;
        let thread = thread::spawn(move || {
            match process_object(stream) {
                Ok(_) => (),
                Err(e) => println!("{:?}", e),
            }
        });
        thread_vec.push(thread);
    }

    for thread in thread_vec {
        thread.join().unwrap();
    }

    Ok(())
}

fn process_object(stream: TcpStream) -> std::io::Result<()> {
    let mut input = BufReader::new(stream.try_clone()?);

    let mut emp_string = String::new();
    input.read_line(&mut emp_string)?;
    println!("Received JSON string: {}", emp_string);

    let emp: Employee = serde_json::from_str(&emp_string)?;
    println!("Employee Name: {}", emp.first_name);

    Ok(())
}
