FROM alpine:edge AS build
RUN apk update
RUN apk upgrade
RUN apk add go=1.18.2-r0 gcc=11.2.1_git20220219-r2 g++=11.2.1_git20220219-r2 make rust cargo
WORKDIR /app
# ENV GOPATH /app
RUN go mod init app
ADD ./ /app
RUN export CC=gcc && cd lib/regression && cargo build --release && cp lib/regression/target/release/libregression.a lib/
RUN CGO_ENABLED=1 GOOS=linux go install -a line_integrals_fuel_efficency
RUN go build line_integrals_fuel_efficiency
CMD ["app/line_integrals_fuel_efficency"]

# FROM alpine:edge
# WORKDIR /app
# RUN cd /app
# COPY --from=build /app/bin/line_integrals_fuel_efficency /app/bin/line_integrals_fuel_efficency
# CMD ["bin/line_integrals_fuel_efficency"]