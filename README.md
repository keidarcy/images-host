# zzw-food-gallery

- https://keidarcy.github.io/zzw-food-gallery/

## steps

- local
  - airdrop heic from iphone to mac
  - `go run cmd/main.go`
    - convert heic to jpg
    - upload to s3
    - clean up local files
- ci
  - push code or trigger scheduled job
  - `go run cmd/main.go ci`
    - list s3 bucket
    - render html
  - deploy to github page
