
gcc \
  -I../decentdb/include \
  app.c \
  -L../decentdb/target/debug \
  -Wl,-rpath,../decentdb/target/debug \
  -ldecentdb \
  -o app
