FROM mongo:latest

# Optional: Copy initialization scripts
COPY ./mongo-init/ /docker-entrypoint-initdb.d/

# Expose port
EXPOSE 27017
