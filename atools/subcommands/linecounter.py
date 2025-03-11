import os


def count_lines(directory, extension):
    """Count the amount of lines for a certain extension in a given directory."""
    total_lines = 0
    for root, _, files in os.walk(directory):
        for file in files:
            if file.endswith(extension):
                with open(os.path.join(root, file), "r", encoding="utf-8") as f:
                    total_lines += sum(1 for _ in f)
    return total_lines


def register_subcommand(subparsers):
    """Register the 'linecounter' subcommand."""
    parser = subparsers.add_parser(
        "linecounter", help="Count lines of code in files with a given extension"
    )
    parser.add_argument("extension", help="File extension (e.g., .py, .txt)")
    parser.add_argument(
        "--dir", default=".", help="Directory to scan (default: current)"
    )
    parser.set_defaults(func=handle_linecounter)


def handle_linecounter(args):
    """Handle the 'linecounter' command."""
    lines = count_lines(args.dir, args.extension)
    print(f"Total lines in {args.extension} files: {lines}")
