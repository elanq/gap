#!/usr/bin/python
# coding: UTF-8

# display current screen resolution

from AppKit import NSScreen
import sys


def main():
    rect = NSScreen.mainScreen().frame()
    H = int(rect.size.height)
    W = int(rect.size.width)
    print("{};{}").format(W, H)


if __name__ == '__main__':
    sys.exit(main())
