# Stage 1 (nge build golang tapi size nya gedde)
# 1. Build the application
FROM golang:1.26-alpine AS builder

# 2. set environtment variabel
# Kenapa hanya cgo_enabbled? karena untuk memanggil kode C untuk menjalankan kaya opencv, sqlite, dll.
ENV CGO_ENABLED=0

# 3 set working directory
WORKDIR /app

# 4 copy source code (. pertama ittu src, . kedua itu destination)
COPY . .

# 5. download dependecies. kaya github.com/go-sql-driver/mysql v1.7.0 dan sebagainya siap di build tanpa errors
RUN go mod download

# 6. build application dengan nama outputnbya yg bisa dirubah rubah 
RUN go build -o api-contact-form-v2

# Stage 2 (copy hasil build terus di resize dengan alpine agar lebih kecil)
#7. Create production image . biar size nya kecil. image nya di build dari project ini
FROM alpine:latest

#8.  Install tzdata for timezone support. biar bisa gunain load location asia jakarta. jadi ngambil timezone untuk linux (tzdata). kalau apk add itu soalnya si alpine ga punya kaya apt atau yum
RUN apk add --no-cache tzdata

# 9. Set the timezone environment variable (can be overridden by .env)
ENV TZ=Asia/Jakarta

# 10. Configure the timezone. ??? masih tidak paham si.
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 11. Set working directory (lagi, untuk stage 2, karna kan FROM nya beda)
WORKDIR /app

# Create user and group for application
#
# Create a group with GID 1001
RUN addgroup -g 1001 binarygroup
# Create a user with UID 1001 and assign them to the 'binarygroup' group
RUN adduser -D -u 1001 -G binarygroup userapp


# Copy the binary from the builder stage
COPY --from=builder --chown=userapp:binarygroup /app/api-contact-form-v2 .

# Switch to the userapp user
USER userapp

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./api-contact-form-v2"]