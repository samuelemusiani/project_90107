build:
  cd front && npm run build
  rm -rf ./handler/dist
  cp -r front/dist ./handler/
  go build -o basi .
