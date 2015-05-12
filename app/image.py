#!/usr/bin/env python
import bottle
import urllib2
import json

@bottle.post('/uop')
def image():
    data = bottle.request.json
    url = 'http://requestb.in/1kc814q1'
    req = urllib2.Request(url)
    req.add_header('Content-Type', 'application/json')
    resp = urllib2.urlopen(req, json.dumps(data))
    return resp

bottle.run(host='0.0.0.0', port=9100)
