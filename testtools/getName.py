import random

# 常见的姓氏列表
surnames = ['赵', '钱', '孙', '李', '周', '吴', '郑', '王', '冯', '陈', '褚', '卫', '蒋', '沈', '韩', '杨', '朱', '秦', '尤', '许']

# 常见的名字列表（两个字）
two_char_names = ['伟', '芳', '娜', '秀英', '敏', '静', '丽', '强', '磊', '军', '洋', '勇', '艳', '杰', '娟', '涛', '明', '超', '秀兰', '霞']

# 常见的名字列表（三个字）
three_char_names = ['桂英', '凤姐', '玉英', '英子', '美华', '冬梅', '春梅', '秀华', '国强', '国华', '玉兰', '海燕', '海伦', '伟强', '伟华', '秀兰', '美丽', '红梅', '红霞', '小明']

# 生成随机汉字
def generate_random_chinese_char():
    head = random.randint(0xB0, 0xD7)
    body = random.randint(0xA1, 0xFE)
    return bytes.fromhex(f'{head:x}{body:x}').decode('gbk')

# 生成指定长度的汉字名字
def generate_chinese_name(length):
    name = ''
    for _ in range(length):
        name += generate_random_chinese_char()
    return name

# 生成 100 个随机名字（包括两个字和三个字）
output_file = '/Users/hao/Desktop/go/names.txt'
with open(output_file, 'w') as f:
    for i in range(100):
        surname = random.choice(surnames)  # 随机选择一个常见的姓氏
        if random.choice([True, False]):  # 随机选择是两个字还是三个字的名字
            given_name = random.choice(two_char_names)  # 两个字的名字
        else:
            given_name = random.choice(three_char_names)  # 三个字的名字
        name = f"{surname}{given_name}"
        f.write(name + '\n')
        print(f"{i+1}. {name}")

print(f"Names written to {output_file}")
