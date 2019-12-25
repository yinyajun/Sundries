import re


def find_div(code):
    # .匹配除\n\r外的任意字符
    # 要匹配\n使用[\s\S]
    pattern = r'(<div>[\s\S]*?</div>)'
    pattern = re.compile(pattern)
    return re.sub(pattern, lambda p: div_formatter(p.group(1)), code)


def div_formatter(div_code):
    div_code = div_code.strip()
    pattern = r'(<style)([\s\S]*?)(</style>)'
    return re.sub(pattern, '', div_code).replace('\n', '')


if __name__ == '__main__':
    code = """..."""
    print(find_div(code))
