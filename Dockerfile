FROM alpine
RUN apk add --no-cache ca-certificates
ADD motivationapp /motivationapp
EXPOSE 443
ENTRYPOINT ["/motivationapp"]