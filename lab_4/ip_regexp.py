import re

def get_ip(target):
    return re.search(r'\b(([0-1]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])[.]){3}'
                     r'([0-1]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])\b', target)