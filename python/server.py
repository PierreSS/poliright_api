import socket, json

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('', 5000))
print (s)
while 1:
    s.send("ia")
    data = s.recv(1024)
    print (data)
    """    data = {"null"} """
    data = {"relation_between": [["europe", "centrales", 0]], "important_words": ["les centralesnucl\u00e9aires"], "politician": ["Damien Abad", "Serge Bardy"], "sentiment":-1, "topic": "europe", "final_deduction": [-1, "europe"]}
    json_data = json.dumps(data)
    s.sendall(json_data)
s.close()
print ('Received', repr(data))