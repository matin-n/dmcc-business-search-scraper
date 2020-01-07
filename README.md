# dmcc-business-search-scraper
Extract contact details from [dmcc.ae](https://www.dmcc.ae/business-search)

## How was the data obtained?
The business directory has a frame which links to the actual directory page:
<p align="center"> 
<img src="https://i.imgur.com/X1U6TdS.png">
</p>

The directory page contains `vid` and `csrf` tokens which are used to authenticate the API to return the contents:
<p align="center"> 
<img src="https://i.imgur.com/qZIr802.png">
<img src="https://i.imgur.com/xJo09k9.png">
</p>


A post request is made to `/apexremote` along with the `vid` and `csrf` token. Once the request is completed the requested business information will be given:

```
{
  "action": "Business_directory_Controller",
  "method": "fetchAllCompanies",
  "data": [
    {
      "pagesize": 10,
      "startIndex": 0,
      "endIndex": 499
    }
  ],
  "type": "rpc",
  "tid": 2,
  "ctx": {
    "csrf": "CSRF_TOKEN_HERE",
    "vid": "VID_HERE",
    "ns": "",
    "ver": 36
  }
}
```

The scraped data shows that there is a total of 11,052 results within the database. However, we are only getting 500 at a time with the default post parameters:
<p align="center"> 
<img src="https://i.imgur.com/X0q2Cqq.png">
</p>


One thing to note is that the post request allows for modification of the `startIndex` `endIndex`. We can modify the index values to scrape more data within a single query. However, there is a 15MB limitation within the server:
<p align="center"> 
<img src="https://i.imgur.com/Q8RcJ2U.png">
</p>

So, we can split the request into two requests by changing the `startIndex` and `endIndex` from an index value of `0-6000` and `6000-13000`. This allows us to scrape the whole with just two queries while maintaining the 15MB limitation:

<p align="center"> 
<img src="https://i.imgur.com/BkT4VsM.png">
<img src="https://i.imgur.com/g21nVvv.png">
</p>



----
## Thanks
* [GJSON](https://github.com/tidwall/gjson)
* [Pretty](https://github.com/tidwall/pretty)
* [Fiddler4](https://www.telerik.com/fiddler)
