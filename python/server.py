import socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('', 5000))
s.sendall('OK'+'\n')
data = s.recv(1024)
s.sendall(data + 'aadz')
s.close()
print 'Received', repr(data)