#!/usr/bin/env python
import bottle
import urllib
import json
from PIL import Image
import io

@bottle.post('/uop')
def image():
    data = bottle.request.json
    try:
        src_url = data['src']['url']
    except:
        abort(400, 'no src url found')

    fin = urllib.urlopen(src_url)
    img = Image.open(fin).convert('L')
    fin.close()

    fout = io.BytesIO()
    img.save(fout, "PNG")
    return fout.getvalue()

bottle.run(host='0.0.0.0', port=9100)

