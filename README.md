# backend_base_app
Backend for Base Apps, 


make sercret token
- string to base64 -> https://www.base64encode.org/
- base64 to encrypt default -?https://md5decrypt.net/en/Sha256/

application -> router ("set path api") -> controller ("set response api") -> usecase ("split handle situasion controller") -> gateway-repository ("handle database data") ->  domain.entity ("struct data")