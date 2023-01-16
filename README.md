# Email-sample



### Check health of service

```bash
http :5000/health
```
### Create token for athentication

```bash
http post :5000/generate-token requestorEmail=singhshishank2012@gmail.com
```

### Send email

```bash
export TOKEN=Get from Generate endpoint
curl -s -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    localhost:5000/v1/sendmail \
    --data-raw '{
        "from":"singshishank2012@gmail.com",
        "from_name":"admin",
        "to":"singshishank2012@gmail.com",
        "to_name":"mahdi",
        "subject":"salam",
        "content":"khobi?",
        ]
    }' | jq
```
 ### Docker
Build docker image with Command
docker build --rm -t email-sample-image .

Run the command
docker run -p 5000:5000 email-sample-image


 
