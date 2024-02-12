with open('./n00b/flag.txt', 'rb') as file:
      encrypted_text = file.read()

add_op = lambda b, i: (b + i) % 256
ror_op = lambda b, i: ((b >> i) | (b << (8 - (i % 8)))) & 255


def decrypt_block(encrypted_text):
    decrypted = ""
    for i in range(0, len(encrypted_text), 8):
        block_encrypted = encrypted_text[i:i+8]
        block_decrypted = ""
        for j in range(8):
            block_decrypted += chr(ror_op(add_op(block_encrypted[j], j), j))
        decrypted += block_decrypted.rstrip('\x00')  
    return decrypted

print(decrypt_block(encrypted_text))
