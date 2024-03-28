import hashlib


def stringHash(source):
    """计算字符串的SHA-256哈希，并转换为64位二进制字符串"""
    sha256 = hashlib.sha256(source.encode('utf-8')).hexdigest()
    return bin(int(sha256, 16))[2:].zfill(256)[:64]


def simHash(token):
    """计算文本的SimHash，并转换为十六进制"""
    tokens = token.split()
    v = [0] * 64
    for t in tokens:
        t_hash = stringHash(t)
        for i in range(len(t_hash)):
            if t_hash[i] == '1':
                v[i] += 1
            else:
                v[i] -= 1
    fingerprint = ''
    for i in range(len(v)):
        if v[i] >= 0:
            fingerprint += '1'
        else:
            fingerprint += '0'
    return binary_to_hex(fingerprint)


def binary_to_hex(binary_str):
    """二进制字符串转换为十六进制"""
    return hex(int(binary_str, 2))[2:].zfill(16)


def hex_to_binary(hex_str):
    """十六进制字符串转换为二进制"""
    return bin(int(hex_str, 16))[2:].zfill(len(hex_str) * 4)


def hamming_distance(s1, s2):
    """计算两个字符串的汉明距离"""
    if len(s1) != len(s2):
        raise ValueError("Strings must be of the same length")
    return sum(el1 != el2 for el1, el2 in zip(s1, s2))


def similarity(token1, token2):
    """根据汉明距离计算相似度（十六进制输入）"""
    hex1 = simHash(token1)
    hex2 = simHash(token2)
    print(hex1)
    print(hex2)
    bin1 = hex_to_binary(hex1)
    bin2 = hex_to_binary(hex2)
    distance = hamming_distance(bin1, bin2)
    return 1 - distance / len(bin1)


# 示例
tokens1 = "this is a test phrase"
tokens2 = "this is a test phrasi"

print(f"Similarity: {similarity(tokens1, tokens2)}")
