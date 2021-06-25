echo "
"

echo " Testando o pacote Main
"
go test -v --short ./Main
echo "
"

echo " Testando o pacote CLI
"
go test -v --short ./CLI
echo "
"

echo " Testando o pacote Gitlog
"
go test -v --short ./GitLog
echo "
"

echo " Testando o pacote OrdemFornecimento
"
go test -v --short ./OrdemFornecimento
echo "
"

echo " Testando o pacote Report
"
go test -v --short ./Report
echo "
"

echo " Testando o pacote Sys
"
go test -v --short ./Sys
echo "
"
