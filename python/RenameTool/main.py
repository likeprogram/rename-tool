import argparse
import os
import re


def parse_args():
    arg_parser = argparse.ArgumentParser(description='rename tool')
    arg_parser.add_argument('--dir', type=str, default=os.path.curdir, help='Source directory path')
    return arg_parser.parse_args()


def rename_tool():
    args = parse_args()
    source = args.dir
    # if source is None:
    #     source = os.path.curdir
    print(f'source: {source}')
    dest = os.path.join(source, "dest")
    pat = re.compile(r'.+_(bin.+\.dat)')
    for entry in os.listdir(source):
        if os.path.isfile(os.path.join(source, entry)):
            print(f'Found {entry}')
            match = pat.match(entry)
            if match is not None and len(match.groups()) == 1:
                new_file_name = match.group(1)
                if not os.path.exists(dest):
                    os.makedirs(dest)
                os.renames(os.path.join(source, entry), os.path.join(dest, new_file_name))


if __name__ == '__main__':
    rename_tool()

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
