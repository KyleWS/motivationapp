FROM alpine
RUN apk add --no-cache ca-certificates bash git curl
COPY motivationapp /motivationapp
EXPOSE 443
ENTRYPOINT ["/motivationapp"]