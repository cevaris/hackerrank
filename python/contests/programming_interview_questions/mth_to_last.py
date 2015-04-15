"""
Problem Statement

For this question, you will write a program that, given a positive integer M and a list of integers L, outputs the list element M links away from the end of the list. For this program, we will use 1-indexing. That means mth_to_last(1) is the "1st-to-last" element, or simply the last element in the list.


If you are given an invalid index, output NIL instead.

Examples
Input:
4 10 200 3 40000 5

Output:
200

Input:
2 42

Output:
NIL

General Approach
Construct a linked list from the inputs. While it's certainly possible to solve this using array indices, the point is to practice linked list traversals.
Use that linked list to find the Mth to last element.
Things to think about
Is your list singly- or doubly-linked? How does this affect your algorithm? Does this change the complexity of your algorithm?
What if your list was circular? Would this change how you check for edge cases?
Input Format and Restrictions
Each test case will consist of two lines. The first line contains the value of M. The second line contains the values of L, each separated by a space character.

The inputs will always satisfy the following restrictions:

0 < M < 2^32 - 1,
Each element of the list satisfies 0 <= L[i] <= 2^32 - 1,
The number of elements in the list satisfies 0 < \|L\| < 1024. The bonus test cases may be much larger!
"""


class Node(object):
    def __init__(self, prev, value, next):
        self.prev = prev
        self.next = next
        self.value = value

class DoubleLinkList(object):
    def __init__(self, ls=None):
        self.first = self.last = None
        if ls and len(ls):
            self.add(ls)

    def _add(self, elem):
        if elem == None:
            raise ValueError("None value found")

        curr = Node(None, elem, None)

        if self.last == None:
            self.first = curr
            self.last = curr
        else:
            self.last.next = curr
            curr.prev = self.last
            self.last = curr

    def add(self, elem ):
        if type(elem) == list:
            for e in elem:
                self._add(e)
        else:
            self._add(elem)
        return True

    def iter(self):
        curr = self.first
        while curr:
            yield curr.value
            curr = curr.next

    def size(self):
        size = 0
        for a in self.iter():
            size += 1
        return size

    def mth_to_last(self, i):
        curr = self.last
        result = 'NIL'

        if i <= 0 or not curr:
            return result

        while curr and curr.prev and i > 1:
            curr = curr.prev
            i -= 1

        if i == 1:
            result = curr.value
        return result


import fileinput

stdin = fileinput.input()
index = int(stdin.next())
data = stdin.next().split(' ')
ls = DoubleLinkList(data)
print ls.mth_to_last(index)
