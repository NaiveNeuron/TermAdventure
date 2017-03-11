from __future__ import print_function
from jinja2 import Environment, FileSystemLoader
import argparse
import json
import sys
import os


def gen_lvls(name, count, format="{}{:02d}"):
    levels = []
    for i in range(1, count + 1):
        levels.append(format.format(name, i))
    return ", ".join(levels)


def get_names(file_names, json_paths):
    dirs = {}
    if json_paths is not None and len(file_names) > len(json_paths):
        print("Every template has to have json file assigned!",
              file=sys.stderr)
        exit(1)

    for i, filename in enumerate(file_names):
        name = os.path.basename(filename)
        if json_paths is None:
            json_name = os.path.splitext(filename)[0] + '.json'
        else:
            json_name = json_paths[i]
        if os.path.isfile(json_name) is False:
            print("Every template has to have json file assigned!",
                  file=sys.stderr)
            exit(1)

        with open(json_name) as json_file:
            variables = json.loads(json_file.read())

        if not os.path.dirname(filename):
            dir_name = "."
        else:
            dir_name = os.path.dirname(filename)

        if dir_name in dirs:
            dirs[dir_name].append({'template_name': name,
                                   'variables': variables})
        else:
            dirs[dir_name] = [{'template_name': name,
                               'variables': variables}]
    return dirs


def main(filenames, var_jsons):
    dirs = get_names(filenames, var_jsons)
    for directory, templates in dirs.iteritems():
        env = Environment(loader=FileSystemLoader(directory), trim_blocks=True)
        env.filters['gen_lvls'] = gen_lvls
        for files in templates:
            out_filename = format(os.path.splitext(files['template_name'])[0])
            template = env.get_template(files['template_name'])
            with open("{}.gta".format(out_filename), 'w+') as gta_file:
                gta_file.write(template.render(files['variables']))


if __name__ == '__main__':
    desc = '''
           GTA templater - program which converts jinja2 templates to .gta
           files.
           '''
    parser = argparse.ArgumentParser(description=desc)

    parser.add_argument('filename', nargs='+', metavar='filename',
                        help='''
                             Path to template files.
                             ''')

    parser.add_argument('-j', '--json', metavar='filename', nargs='+',
                        dest='var_json',
                        help='''
                             Path to a json file containing definitions of all
                             variables used in templates. (Default name is the
                             same as template name)
                             ''')
    args = parser.parse_args()
    main(args.filename, args.var_json)
