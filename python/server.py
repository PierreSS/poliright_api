import socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('', 5000))
while 1:
    data = s.recv(1024)
    s.sendall(data + ': recu.')
s.close()
print 'Received', repr(data)