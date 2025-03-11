import socket
import requests


def get_private_ip():
    """Get the private IP address of the machine."""
    return socket.gethostbyname(socket.gethostname())


def get_public_ip():
    """Get the public IP address of the machine."""
    return requests.get("https://api64.ipify.org").text


def register_subcommand(subparsers):
    """Register the 'ipinfo' subcommand."""
    parser = subparsers.add_parser(
        "ipinfo", help="Display the IP info of the current machine"
    )
    parser.set_defaults(func=handle_ipinfo)


def handle_ipinfo(args):
    """Handle the 'linecounter' command."""
    private_ip = get_private_ip()
    public_ip = get_public_ip()
    hostname = socket.gethostname()
    print(f"{'Hostname:':<15} {hostname}")
    print(f"{'Private IP:':<15} {private_ip}")
    print(f"{'Public IP:':<15} {public_ip}")
