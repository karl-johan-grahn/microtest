FROM scratch

ENV PORT 8090
EXPOSE $PORT

COPY microtest /
CMD ["/microtest"]
