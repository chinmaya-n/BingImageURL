# Docker Image for minimal Bing Image URL service in GoLang

FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
ADD main /

ENTRYPOINT [ "/main" ]