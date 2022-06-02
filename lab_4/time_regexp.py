import re

def get_time(target):
    return re.search(r'\b((1[0-2]|0?[1-9]):([0-5][0-9]) ([AaPp][Mm]))', target)
