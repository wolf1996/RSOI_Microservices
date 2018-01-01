import http.server
import socketserver
import urllib.parse as parse
import requests


PORT = 9191
hostname = "localhost"

sess = None
class MyHandler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        self.address_string
        self.send_response(200)
        self.send_header('Content-type','text/html')
        self.end_headers()
        print(self.path)
        prsd = parse.urlparse(self.path)
        q = prsd[4]
        code = parse.parse_qs(q)['code'][0]
        print(q)
        print("let's start")
        print(code)
        sess = requests.session()
        resp = sess.post("http://127.0.0.1:8080/shiftcode", json={'code_flow':str(code),'domain':'1'})
        print(resp)
        resp = sess.get("http://127.0.0.1:8080/hello")
        print(resp)
        bdy = resp.text
        self.wfile.write(bytes("<html><head><title>Title goes here.</title></head>", "utf-8"))
        self.wfile.write(bytes("<body><p>{}.</p>".format(bdy), "utf-8"))
        self.wfile.write(bytes("</body></html>", "utf-8"))
        return
        pass

httpd = socketserver.TCPServer((hostname, PORT), MyHandler)

print("serving at port", PORT)
httpd.serve_forever()

