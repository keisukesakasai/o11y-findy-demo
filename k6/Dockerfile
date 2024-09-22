# Dockerfile
FROM grafana/k6:0.50.0

COPY ./k6/script.js /app/script.js
CMD ["run", "/app/script.js"]
