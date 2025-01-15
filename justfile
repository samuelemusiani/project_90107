build:
  cd front && npm install && npm run build
  rm -rf ./handler/dist
  cp -r front/dist ./handler/
  go build -o basi .
