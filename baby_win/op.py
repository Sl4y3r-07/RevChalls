a="pa55wd"
b="1Ntr0_T0_W1NR3V"
c=''   
for i in range(len(b)):
     
     c += chr(ord(b[i])^ord(a[i%len(a)]))
    
print(c)
#1Ntr0_T0_W1NR3V