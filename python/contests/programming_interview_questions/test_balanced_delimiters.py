import pytest
from .balanced_delimiters import is_balanced


def test_is_balanced_none():
    with pytest.raises(ValueError):
        is_balanced(None) == True

def test_is_balanced():
    inputs = {
        "": True,
        "NoDelimiters": True,
        "()[]{}": True,
        "([{}])": True,
        "([]{})": True,
        "([] [])": True,
        "(([][]))": True,
        "()[mix]{}": True,
        "([]})": False,
        "([)]": False,
    }

    for text, expected in inputs.iteritems():
        assert is_balanced(text) == expected
