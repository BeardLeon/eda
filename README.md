
## USAGE

### RunServer
```
nohup go run src/eda_main.go > server.log 2>&1 &
```

### ImportFile
设置 Content-Type multipart/form-data; boundary=<calculated when request is sent>
Body中关键字`file`类型File，选择文件, 发送POST请求 ```http://0.0.0.0:9476/import```

### ExportFile
浏览器访问```http://0.0.0.0:9476/export?oid=81b0ec33```

### CreateFile
``` 
curl -H "Content-type: application/json" -X POST -d  '
    {
        "title":"post-title",
        "desc":"post-not a good desc"
    }' http://0.0.0.0:9476/file
```

### GetFile
``` 
curl -H "Content-type: application/json" -X GET http://0.0.0.0:9476/file?oid=81b0ec33
```

### InsertComponent
```
curl -H "Content-type: application/json" -X POST -d '
    {   
        "_id" : "81b0ec33",
        "id": 1,
        "name": "andGate",
        "shape": [
          { "sx": 30, "sy": 20, "ex": 100, "ey": 20 },
          { "sx": 100, "sy": 0, "ex": 100, "ey": 100 },
          { "sx": 0, "sy": 80, "ex": 100, "ey": 80 },
          { "sx": 100, "sy": 0, "ex": 186, "ey": 50 },
          { "sx": 100, "sy": 100, "ex": 186, "ey": 50 },
          { "sx": 186, "sy": 50, "ex": 250, "ey": 50 }
        ],
        "pin": [ 
          { "x": 0, "y": 20 }, 
          { "x": 0, "y": 80 }, 
          { "x": 250, "y": 50 } 
        ]
    }' http://0.0.0.0:9476/component
```

### InsertLine
```
curl -H "Content-type: application/json" -X POST -d '
    {   
        "_id" : "81b0ec33",
        "sx": 0, 
        "sy": 2, 
        "ex": 2, 
        "ey": 20 
    }' http://0.0.0.0:9476/line
```

### Updatecomponent
```
curl -H "Content-type: application/json" -X PUT -d '
    [
        {
            "_id":"81b0ec33",
             "id": 1,
            "name": "andGate",
            "shape": [
            { "sx": 30, "sy": 20, "ex": 100, "ey": 20 },
            { "sx": 100, "sy": 0, "ex": 100, "ey": 100 },
            { "sx": 0, "sy": 80, "ex": 100, "ey": 80 },
            { "sx": 100, "sy": 0, "ex": 186, "ey": 50 },
            { "sx": 100, "sy": 100, "ex": 186, "ey": 50 },
            { "sx": 186, "sy": 50, "ex": 250, "ey": 50 }
            ],
            "pin": [ 
            { "x": 0, "y": 20 }, 
            { "x": 0, "y": 80 }, 
            { "x": 250, "y": 50 } 
            ]
        },
        {  
            "_id":"81b0ec33",
             "id": 1,
            "name": "andGate",
            "shape": [
                { "sx": 30, "sy": 20, "ex": 100, "ey": 20 },
                { "sx": 100, "sy": 0, "ex": 100, "ey": 100 },
                { "sx": 0, "sy": 80, "ex": 100, "ey": 80 },
                { "sx": 100, "sy": 0, "ex": 186, "ey": 50 },
                { "sx": 100, "sy": 100, "ex": 186, "ey": 50 },
                { "sx": 186, "sy": 50, "ex": 250, "ey": 50 }
            ],
            "pin": [ 
                { "x": 0, "y": 20 }, 
                { "x": 0, "y": 80 }, 
                { "x": 110, "y": 110 } 
            ]
        }
    ]' http://0.0.0.0:9476/component
```

### UpdateLine
```
curl -H "Content-type: application/json" -X PUT -d '
    [
        {
            "_id":"81b0ec33",
            "sx":2, 
            "sy":2, 
            "ex":2, 
            "ey":20
        },
        {  
            "_id":"81b0ec33",
            "sx":1, 
            "sy":1, 
            "ex":1, 
            "ey":20
        }
    ]' http://0.0.0.0:9476/line
```

### DeleteLine
```
curl -H "Content-type: application/json" -X DELETE -d '
    {   
        "_id" : "81b0ec33",
        "sx": 1, 
        "sy": 1, 
        "ex": 1, 
        "ey": 20 
    }' http://0.0.0.0:9476/line
```

### DeleteComponent
```
curl -H "Content-type: application/json" -X DELETE -d '
    {  
        "_id":"81b0ec33",
        "id": 1,
        "name": "andGate",
        "shape": [
            { "sx": 30, "sy": 20, "ex": 100, "ey": 20 },
            { "sx": 100, "sy": 0, "ex": 100, "ey": 100 },
            { "sx": 0, "sy": 80, "ex": 100, "ey": 80 },
            { "sx": 100, "sy": 0, "ex": 186, "ey": 50 },
            { "sx": 100, "sy": 100, "ex": 186, "ey": 50 },
            { "sx": 186, "sy": 50, "ex": 250, "ey": 50 }
        ],
        "pin": [ 
            { "x": 0, "y": 20 }, 
            { "x": 0, "y": 80 }, 
            { "x": 110, "y": 110 } 
        ]
    }' http://0.0.0.0:9476/component
```