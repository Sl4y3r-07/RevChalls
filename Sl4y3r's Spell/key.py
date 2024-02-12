key ='sl4Y3r07'
enc='8cb5cdbc37678078'
result=''
dec=''
for i in range(len(key)):
    result=int(enc[i*2:(i+1)*2],16)-i;
    dec += chr((result)^ord(key[i]))

print( ''.join(format(ord(char), '02x')+' ' for char in dec))