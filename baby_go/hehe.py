str='flag{G0_rev3Rs1ng_b4byy!}'
res=''
st=''
for i in range(len(str)):
    res += chr(ord(str[i])^(i+1))
for i in range(len(res)): 
    if i%2==0:
       st += chr(ord(res[i])-1)
    else:
        st+= res[i]
print(res)
print(st)