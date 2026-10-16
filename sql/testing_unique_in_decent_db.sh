DECENTDB=../vendor/decentdb/target/debug/decentdb
rm -rf prueba.ddb*
$DECENTDB exec --db="prueba.ddb" --sql="$(cat testing_unique.sql)"
