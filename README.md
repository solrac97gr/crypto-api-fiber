# START GUIDE
- Create a new firebase project from [firebase console](https://console.firebase.google.com/u/0/)
- Initialize a new Free Database, Must be in the **Realtime database service**
- Claim your Firebase credentials <code>(Settings>Services Accounts> Go > Create New Credentials)</code> and replace it in the <code>serviceAccountKey.example.go</code>
- Rename the file as <code>serviceAccountKey.go</code>
-  In the same file replace your DB Url
- Execute Docker <code>docker-compose -f docker-compose.yml up</code>
- Access to the API from http://localhost:8080

The nex you will see it's a docker instance with 5 replicas of the API this will help for improve the performance in the Encrypt process

**Your database will looks like this:**

```
{
  "log" : {
    "-MaCG3sW9Q4nXOdHHjjP" : {
      "Method" : "POST",
      "URL" : "/v1/encrypt"
    },
    "-MaCGFCaJriPjOVWCOpw" : {
      "Method" : "GET",
      "URL" : "/v1/decrypt/-MaCG4-7Ngf9tfEDqcKf"
    },
   
  },
  "text" : {
    "-MaCG4-7Ngf9tfEDqcKf" : {
      "EncryptedText" : "2c1b47f08f09d55eb0f167783f4f885ddff746eb44fd8d67e245ee4377e1510da8ddd4ed5ada4e9528d16164a310e2da997fbcc24ae2ce48542de5758f763b7ecec1e9ab3bc771a800f4f56542cfa1fa5856df9b4b6643ce0917",
      "EncryptionKey" : "25c0aca4-0a0d-44c4-a573-946f2cbdab75"
    },
  }
}
```