docker build -t avatarsolution/ipos-mblb-backend .
docker stop ipos-mblb-backend
docker run --rm -dp 7100:8080 --name ipos-mblb-backend  avatarsolution/ipos-mblb-backend
