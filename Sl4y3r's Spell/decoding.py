# str='fa 3d 7c 21 42 7d 30 44'.replace(" ",'')
# str2= '89 50 4e 47 0d 0a 1a 0a'.replace(" ",'')
# res='' 
# xor =''
# for i in range(8):
#     res= int(str[i*2:(i+1)*2],16)-i
#     xor+= chr(res^int(str2[i*2:(i+1)*2],16))
# print(str)
# print(xor)

#fa 3d 7c 21 42 7d 30 44 Enc Header
#89 50 4e 47 0d 0a 1a 0a PNG Header

with open('./n00b/cool.encrypted.txt','rb') as file:
     str= file.read()
key='sl4Y3r07'
xor=''
for i in range(len(str)):
    xor += chr((str[i])^ord(key[i%len(key)]))

print(xor)