import socket 

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM) 
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
s.bind(('127.0.0.1',12345)) 
s.listen(1) 

while 1: 
    client, address = s.accept() 
    print('accespted')
    data = client.recv(1024) 
    client.send(data) 
    client.close()
