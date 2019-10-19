import shutil
from PIL import Image

path = './data/mushroom/'

# get ride of species with less than 10 photos
for root, dirs, files in os.walk(path):
    for dir in dirs:
        length = len(os.listdir(path+dir))
        if length < 10:
            shutil.rmtree(path+dir)
        

# get ride of corrupted images
for root, dirs, files in os.walk(path):
    for dir in dirs:
        for file in os.listdir(path+dir):
            try:
                img = Image.open(path+dir+'/'+file)
            except OSError:
                os.remove(path+dir+'/'+file)