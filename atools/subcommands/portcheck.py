import socket


def check_port(port):
    """Check if a port is free (available) or occupied."""
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        return s.connect_ex(("127.0.0.1", port)) != 0  # Returns True if free


def register_subcommand(subparsers):
    """Register the 'portcheck' subcommand."""
    parser = subparsers.add_parser(
        "portcheck", help="Check if a port is occupied or free"
    )
    parser.add_argument("port", type=int, help="Port number to check")
    parser.set_defaults(func=handle_portcheck)


def handle_portcheck(args):
    """Handle the 'portcheck' command."""
    port_status = check_port(args.port)
    status_text = "free" if port_status else "occupied"
    print(f"Port {args.port} is {status_text}.")
