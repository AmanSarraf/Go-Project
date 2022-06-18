# Azure SDK for compute vision service

Set Azure computer vision service url
```cmd
 export AZURE_VISION_URL="URL"
 ```

set Azure compute vision key
```cmd
export AZURE_VISION_KEY=unique keyvalue
```
Set Port 
```cmd
export PORT=3000
```
Run main

```cmd
go run .
```

switch to api menue and pass URL of image as json body with POST request
```cmd
curl -X POST localhost:3000/compvis \
   -d '{
    "url": "https://raw.githubusercontent.com/MicrosoftDocs/azure-docs/master/articles/cognitive-services/Computer-vision/Images/readsample.jpg"
}'
```
### Enjoy the end result!!! ^w^
