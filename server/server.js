var net = require('net')
var mjml = require('mjml')

var server = net.createServer((conn) => {
  conn.on('data', function (data) {
    var transmit = data.slice(0, 4).readUInt32BE(0)
    var buffer = data.slice(4).toString()

    if (buffer.length === transmit) {
      var output = mjml.mjml2html(buffer)
      console.log('write', output.length)

      const buf = new Buffer.allocUnsafe(4)
      buf.writeUInt32BE(output.length.toString(10))
      this.write(buf)
      this.write(output)

      this.end()
    } else {
      console.log('error')
    }
  })

  conn.on('close', function () {
    console.log('socket close')
  })
}).on('error', (err) => {
  console.log(err)
  throw err
}).listen({
  host: 'localhost',
  port: 8686,
  exclusive: true
}, () => {
  console.log('server listening')
})
