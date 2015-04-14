
"""
Problem Statement

For this question, you will parse a string to determine if it contains only "balanced delimiters."

A balanced delimiter starts with an opening character ((, [, {), ends with a matching closing character (), ], } respectively), and has only other matching delimiters in between. A balanced delimiter may contain any number of balanced delimiters.

Examples
The following are examples of balanced delimiter strings:

()[]{} ([{}]) ([]{})

The following are examples of invalid strings:

([)] ([] []) ([})

Input is provided as a single string. Your output should be True or False according to whether the string is balanced. For example:

Input:
([{}])

Output:
True

Input Format and Restrictions
Each test case will consist of a string only containing the characters ()[]{}. The length of the string will not exceed 2KB.

Things to think about
- How will you keep track of previous delimiters?
- How will you determine if the next character is valid?
- When you reach the end of the string, how do you know if it is balanced?
"""


OPEN   = ['(','[','{']
CLOSED = [')',']','}']


def is_balanced(text):
    stack = []
    if text == None:
        raise ValueError("None value found")

    found_delimiter = False
    for c in text:
        if c in OPEN:
            stack.append(c)
        elif c in CLOSED and OPEN[CLOSED.index(c)] == stack[len(stack)-1]:
            stack.pop()
        elif c in CLOSED or c in OPEN:
            return False

    return len(stack) == 0

import fileinput
for line in fileinput.input():
    print is_balanced(line)
