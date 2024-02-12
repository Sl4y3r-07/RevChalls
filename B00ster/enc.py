# flag = "n00bCTF{H3y_r3v3rs3r!_N0t_a_n00b_anym0r3:)_C0ng00..hehe}"
# res = b''

# for i in range(len(flag)):
#     rotated_val = ((ord(flag[i]) << i) | (ord(flag[i]) >> ((i % 32)))) & 0xFF
#     char_val = rotated_val + i
#     res += bytes([char_val])
# print(res)
# with open('./n00b/flag.txt', 'wb') as file:
#     file.write(res)


sub_op = lambda b, i: (b - i) % 256
rol_op = lambda b, i: ((b << i) & 255) | (b >> (8 - (i % 8)))


def encrypt_block(text):
    encrypted = bytearray()
    for i in range(0, len(text), 8):
        block = text[i:i+8].ljust(8, '\x00')  # Padding for blocks shorter than 8 characters
        block_encrypted = bytearray()
        for j in range(8):
            block_encrypted.append(sub_op(rol_op(ord(block[j]), j), j))
        encrypted.extend(block_encrypted)
    return encrypted

text = "n00bCTF{H3y_r3v3rs3r!_N0t_a_n00b_anym0r3:)Cong0}"
encrypted_text = encrypt_block(text)
print(encrypted_text)
with open('./n00b/flag.txt', 'wb') as file:
        file.write(encrypted_text)


