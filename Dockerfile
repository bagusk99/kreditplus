# Pakai base image Go versi terbaru
FROM golang:1.24

# Set working directory dalam container
WORKDIR /app

# Copy file module
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code ke container
COPY . .

# Build binary aplikasi (hasilnya file 'app')
RUN go build -o app .

# Expose port aplikasi (sesuaikan dengan port yang kamu pakai)
EXPOSE 8080

# Jalankan binary aplikasi
CMD ["./app"]
