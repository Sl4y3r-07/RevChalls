# # with open('./n00b/flag.txt','rb') as file:
# #     str = file.read()
# #byte=''
# # for i in range (len(str)):
# #     byte += format(str[i],'02x')
# key ='sl4Y3r07'
# xor=''
# enc=''
# # for i in range(len(key)):
# #      xor =chr(ord(key[i])^int(byte[i * 2: (i + 1) * 2], 16))
# #      enc += chr(ord(xor) + i)
# #xor_hex = ''.join(format(ord(char), '02x')+' ' for char in enc)



str='n00bCTF{d4mn!!PNG_h34d3rs_4r3_c00l!!}'
print(len(str))
key ='sl4Y3r07'
xor=''
enc=''
for i in range(len(str)):
    xor += chr(ord(str[i])^ord(key[i%len(key)]))
   
    
print(xor)
with open("./n00b/cool.encrypted.txt",'wb') as file:
    file.write(bytearray(xor,'utf-8'))

