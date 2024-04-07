use std::{
    fs::File,
    io::Read,
    net::{SocketAddr, UdpSocket},
    process::exit,
};

fn main() -> Result<(), String> {
    let args: Vec<String> = std::env::args().collect();
    if args.len() < 2 {
        show_help();
        exit(0);
    }
    let domain_name = &args[1];
    let domains = vec![domain_name.clone()];
    // TODO: server's IP address should be passed via cli argument
    let dns_server_address = [8, 8, 8, 8];
    let server = SocketAddr::from((dns_server_address, 53));
    let client = SocketAddr::from(([0, 0, 0, 0], 0));
    let sock = UdpSocket::bind(client).map_err(|err| err.to_string())?;
    let query = dns_client::Query::new(get_random_u16(), &domains);

    // // send UDP packet to the server
    let buf: Vec<u8> = query.into();
    sock.send_to(&buf, server).map_err(|err| err.to_string())?;
    let mut buf = [0; 512];
    sock.recv_from(&mut buf).map_err(|err| err.to_string())?;
    let response = dns_client::Response::try_from(&buf)?;

    println!(
        "====\nIP address for {} is {}",
        domain_name,
        response.answers[0].rdata.to_string(),
    );

    Ok(())
}

fn show_help() {
    println!(
        "
USAGE:
$ dns_client <domain_name_you_want_to_resolve>
    "
    )
}

fn get_random_u16() -> u16 {
    // TODO: this should not work on Windows
    let mut file = File::open("/dev/urandom").unwrap();
    let mut buffer = [0u8; 2];
    file.read_exact(&mut buffer).unwrap();
    ((buffer[0] as u16) << 8) + (buffer[1] as u16)
}
