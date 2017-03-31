import vici
import socket

s = socket.socket(socket.AF_UNIX)
s.connect("/var/run/charon.vici")
v = vici.Session(s)

ver = v.version()

print "{daemon} {version} ({sysname}, {release}, {machine})".format(**ver)
