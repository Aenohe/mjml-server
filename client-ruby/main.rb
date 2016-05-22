require "socket"

conn = TCPSocket.open("localhost", 8686)

file = File.open("template.mjml")
template = file.read
file.close

sizeSend = [template.length].pack('l>')

conn.write sizeSend
conn.puts template

data = ''
while line = conn.gets
  data += line
end

sizeReceive = data[0..4].unpack('l>')
output = data[4..-1]

if sizeReceive.join.to_i == output.length
  file = File.open("output.html", "w")
  file.write output
  file.close
end

conn.close
