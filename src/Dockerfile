FROM golang AS builder

WORKDIR /src

ENV GOCACHE=/src/.cache
ENV SOURCE=favsort
ENV TARGET='++ sort Favorites ++'

COPY cmd /src/cmd
COPY artifacts /src/artifacts
COPY internal /src/internal
COPY go.mod /src/
COPY cli /cli


RUN apt-get update \
 && apt-get install -y zip \
 && go mod tidy \
 && go build ./cmd/acrc \
;

ENV GOARCH=arm
ENV GOOS=linux
ENV CGO_ENABLED=0

RUN  go build ./cmd/"$SOURCE" \
 && cd artifacts \
 && mv "Imgs/$SOURCE.png" "Imgs/$TARGET.png" \
 && mv ../"$SOURCE" "FavSort/$TARGET" \
 && ../acrc "FavSort/$TARGET" > build-id.txt \
 && mv "$SOURCE.sh" "FavSort/$TARGET.sh" \
 && ln -s "FavSort/$TARGET.sh" "$TARGET.sh" \
 && tar czf - "$TARGET.sh" "FavSort" "Imgs/$TARGET.png" | base64 -w 72 >> "Installer.sh" \
 && mv "Installer.sh" "Install FavSort.sh" \
 && zip -r ../FavSort.zip "Readme.md" "Install FavSort.sh" "build-id.txt" \
;

FROM scratch

COPY --from=builder /src/'FavSort.zip' /
