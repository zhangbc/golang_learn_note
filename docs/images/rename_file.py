#!/usr/bin/env python
#coding:utf-8

import os


def rename_files(pfix='go_pro_', path=None):
    """
    批量更新文件夹名称
    :param pfix:
    :param path:
    :return:
    """

    if not path:
        path = os.getcwd()

    for fi in os.listdir(path):
        if fi.startswith('Screen'):
            # Screen Shot 2022-10-07 at 00.44.16
            os.renames(fi, fi.replace('Screen Shot', pfix).replace('-', '')
            .replace(' ', '').replace('.', '', 2).replace('at', '').strip())

if __name__ == '__main__':
    print('starting...')
    rename_files()
    print('done')