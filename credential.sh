#!/bin/sh
content="{\"installed\":{\"client_id\":\"893357034813-2ovpauugcnc334l7as1lf4033srj5eba.apps.googleusercontent.com\",\"project_id\":\"nihongo-hit\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_secret\":\"${SECRET}\",\"redirect_uris\":[\"http://localhost\"]}}"
echo "$content" > credentials.json
