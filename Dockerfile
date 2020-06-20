FROM bash:latest
RUN mkdir /app
COPY main /app/
EXPOSE 9000
CMD ["/app/main"]