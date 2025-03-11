import argparse
from atools.subcommands import linecounter, portcheck, ipinfo


def main():
    parser = argparse.ArgumentParser(prog="atools", description="Arne's CLI Tools")

    subparsers = parser.add_subparsers(
        title="Available Commands", dest="command", metavar="<command>", required=True
    )

    # Register subcommands
    linecounter.register_subcommand(subparsers)
    portcheck.register_subcommand(subparsers)
    ipinfo.register_subcommand(subparsers)

    args = parser.parse_args()
    args.func(args)


if __name__ == "__main__":
    main()
