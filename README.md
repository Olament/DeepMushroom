# DeepMushroom
DeepMushroom is a fungal classficiation project using ResNet

## Data Sources
### iNaturalist.org
iNaturalist.org is a citizen science webstie that allows people upload the the image of unknown organism and be identified by other ecology enthusiast. We collect the image of *research grade* identification of fungus from 2017 to 2019 via their [export tool](https://www.inaturalist.org/observations/export). The images downloaded were then categorized by their species.

We use both python and golang script to download the images from iNaturalist.org. Our golang script supports mulit-thread. See here: [main.go](https://github.com/Olament/DeepMushroom/blob/master/DataCollection/main.go)

**Distribution** 

![](https://github.com/Olament/DeepMushroom/blob/master/md/distribution.png)

The data distribution is pretty skewed as you can see. For the sake the classfication, we remove the species with less than 10 images.

### MushroomExpert.com
Since the images from MushroomExpert were identified by fungal biologist, we use their images as our test set to evaluate the performance of our model.

## Model
Since we are in the very early stage of the experiment. We build the model with fast.ai library. The model with gradually switch to pure pytorch code as we fine-tuning our model.

### Metrics

| Model Architecture  | Validation Accuracy |
| ------------------- | ------------------- |
| ResNet34            | 31.07        |
