# json
```json
{
    "request": [
    {
        "advertiserId": 101,
        "advertiserName": "name1",
        "site": "http://www.test.com",
        "qualifications":[
        {
            "name": "营业执照", 
            "url": "http://material.client.com/123.jpg"
        },{
            "name": "网络文化经营许可证", 
            "url": "http://material.client.com/456.jpg"
        }]
    }],

    "authHeader": {
        "dspId": 2516,
        "token": "a878e28e7b7a6019908ae66a6e6943df"
    }
}
```


{"request":[{"advertiserId": 101,"advertiserName": "name1","site": "http://www.test.com","qualifications":[{"name": "营业执照","url": "http://material.client.com/123.jpg"},{"name": "网络文化经营许可证","url": "http://material.client.com/456.jpg"}]}],"authHeader": {"dspId": 2516,"token": "a878e28e7b7a6019908ae66a6e6943df"}}

# curl 
curl -i -X POST -H 'Content-Type:application/json' -d '{"request":[{"advertiserId": 101,"advertiserName": "name1","site": "http://www.test.com","qualifications":[{"name": "营业执照","url": "http://material.client.com/123.jpg"},{"name": "网络文化经营许可证","url": "http://material.client.com/456.jpg"}]}],"authHeader": {"dspId": 2516,"token": "a878e28e7b7a6019908ae66a6e6943df"}}' 'http://api.accuenmedia.com.cn/v2/advertiser/update'
