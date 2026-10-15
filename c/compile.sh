
gcc \
  -I../vendor/decentdb/include \
  app.c \
  -L../vendor/decentdb/target/debug \
  -Wl,-rpath,../vendor/decentdb/target/debug \
  -ldecentdb \
  -o app
