import pytest
from .mth_to_last import (Node, DoubleLinkList)


def test_empty_list():
    actual = DoubleLinkList().size()
    expected = 0
    assert actual == expected


def test_add_none():
    with pytest.raises(ValueError):
        ls = DoubleLinkList()
        ls.add(None)


def test_add():
    ls = DoubleLinkList()
    map(ls.add, [10, "11", 12.0])

    actual = ls.size()
    expected = 3
    assert actual == expected


def test_mtn_to_last():
    pass


def test_add_list():
    expected = [10, 200, 3, 40000, 5]
    ls = DoubleLinkList(expected)
    index = 0
    for actual in ls.iter():
        assert actual == expected[index]
        index += 1


def test_mth_to_last():
    data = [10, 200, 3, 40000, 5]
    ls = DoubleLinkList(data)
    expected = 200
    actual = ls.mth_to_last(4)
    assert actual == expected


def test_mth_to_last_last():
    data = [10, 200, 3, 40000, 5]
    ls = DoubleLinkList(data)
    expected = 10
    actual = ls.mth_to_last(5)
    assert actual == expected


def test_mth_to_last_first():
    data = [10, 200, 3, 40000, 5]
    ls = DoubleLinkList(data)
    expected = 5
    actual = ls.mth_to_last(1)
    assert actual == expected



def test_mth_to_last_nil():
    data = [42]
    ls = DoubleLinkList(data)
    expected = 'NIL'
    actual = ls.mth_to_last(2)
    assert actual == expected


def test_mth_to_last_l_eq_0():
    data = [42]
    ls = DoubleLinkList(data)
    expected = 'NIL'
    actual = ls.mth_to_last(0)
    assert actual == expected


def test_mth_to_last_empty():
    data = []
    ls = DoubleLinkList(data)
    expected = 'NIL'
    actual = ls.mth_to_last(1)
    assert actual == expected
