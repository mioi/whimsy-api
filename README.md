# whimsy-api

![Tests](https://github.com/mioi/whimsy-api/workflows/Test/badge.svg)
![Go Version](https://img.shields.io/badge/go-1.20%2B-blue.svg)

A simple Go API server that provides random names of plants, animals, and colors. Perfect for generating whimsical names and data for your projects.

## Endpoints

- `GET /plants` - Get all plants
- `GET /animals` - Get all animals
- `GET /colors` - Get all colors
- `GET /names` - Get all categories combined
- `GET /plants/random?count=5` - Get random plants
- `GET /animals/random?count=5` - Get random animals
- `GET /colors/random?count=5` - Get random colors
- `GET /names/random?count=10&parts=2` - Generate random names
- `GET /health` - Health check

## Local Development

```bash
go run main.go
```

## Testing

```bash
go test ./...
```

## Deploy to Google Cloud Run

1. Build and push the Docker image:
```bash
gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/whimsy-api
```

2. Deploy to Cloud Run:
```bash
gcloud run deploy whimsy-api \
  --image gcr.io/YOUR_PROJECT_ID/whimsy-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated
```

Replace `YOUR_PROJECT_ID` with your actual Google Cloud project ID.
