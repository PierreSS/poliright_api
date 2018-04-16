import socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('', 80))
s.sendall('OK'+'\n')
data = s.recv(1024)
s.sendall(data + 'aadz')
s.close()
print 'Received', repr(data)