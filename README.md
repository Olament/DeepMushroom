# DeepMushroom
DeepMushroom is a fungal classficiation project using ResNet

## Data Sets
### iNaturalist.org
iNaturalist.org is a citizen science webstie that allows people upload the the image of unknown organism and be identified by other ecology enthusiast. We collect the image of *research grade* identification of fungus from 2017 to 2019 via their [export tool](https://www.inaturalist.org/observations/export). The images downloaded were then categorized by their species.

We use both python and golang script to download the images from iNaturalist.org. Our golang script supports mulit-thread. See here: [main.go](https://github.com/Olament/DeepMushroom/blob/master/DataCollection/main.go)

### MushroomExpert.com