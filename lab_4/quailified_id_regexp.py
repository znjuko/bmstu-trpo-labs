import re

def get_qualified_id(target):
    return re.search(r'^([a-zA-Z_]\w*\.)*([a-zA-Z_]\w*)$', target)