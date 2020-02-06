# DeepMushroom
DeepMushroom is a fungal classficiation project using [ResNet](https://en.wikipedia.org/wiki/Residual_neural_network)

## Data Sources
### iNaturalist.org
iNaturalist.org is a citizen science website that allows people to upload images of unknown organisms for identification by other ecology enthusiasts. We collected images of identified fungi uploaded between 2017 and 2019 via the iNaturalist [export tool](https://www.inaturalist.org/observations/export). The images downloaded were then categorized by their species.

We use both python and golang scripts to download the images from iNaturalist.org. See here: [main.go](https://github.com/Olament/DeepMushroom/blob/master/datacollection/main.go)

**Distribution** 

![](https://github.com/Olament/DeepMushroom/blob/master/md/distribution.png)

The data distribution is heavily skewed towards the few most common species. We remove the fungal species with less than 10 images for two reasons:
- If one species has less than 10 identification on iNaturalist.org, it indicates that it is not frequently occuring. Therefore, there is less value in the identification of such species.
- There is not enough data to effectively train the identification model. A species with less than 10 images will hurt the overall accuracy of our model

### MushroomExpert.com
Since the images from MushroomExpert were identified by mycologists, we can use their images as a reliable validator to test the performance of our model.

## Model
Since we are in the very early stage of the experiment we built the model with the [fast.ai](https://www.fast.ai/) library. The model will gradually switch to our own models utilizing [pytorch](https://pytorch.org/) as we progress.

### Metrics

|     Architecture    | Validation Accuracy | Validation Top-5 Accuracy | Test Accurarcy | Test Top-5 Accuracy |
|:-------------------:|:-------------------:|:-------------------------:|:--------------:|:-------------------:|
|       ResNet34      |        70.68        |           86.36           |      31.94     |        48.11        |
|       ResNet50      |        79.67        |           91.76           |      38.77     |        59.14        |
| ResNet50+Focal Loss |        80.24        |           92.32           |      39.48     |        60.45        |

#### Top 10 Most Confused Fungal Species

|        Prediction        |       Ground Truth       |
|:------------------------:|:------------------------:|
|    Fomitopsis mounceae   |    Fomitopsis pinicola   |
|   Pleurotus pulmonarius  |    Pleurotus ostreatus   |
| Dacrymyces chrysospermus |   Tremella mesenterica   |
|   Tremella mesenterica   | Dacrymyces chrysospermus |
|  Laetiporus gilbertsonii |   Laetiporus sulphureus  |
|     Stereum hirsutum     |    Stereum complicatum   |
|     Tremella aurantia    |   Tremella mesenterica   |
|    Ganoderma megaloma    |   Ganoderma applanatum   |
|  Laetiporus cincinnatus  |   Laetiporus sulphureus  |
|   Ganoderma applanatum   |     Ganoderma brownii    |
