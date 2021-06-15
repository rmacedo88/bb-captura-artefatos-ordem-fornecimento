echo "
"

echo "  TESTANDO Main/Main.go"
echo "  test case: configurarTamanhoHashCommit"
go tool test2json -t /tmp/___Main_test_go.test -test.v -test.paniconexit0 -test.run Test_configurarTamanhoHashCommit
echo "  test case: main"
go tool test2json -t /tmp/___Main_test_go.test -test.v -test.paniconexit0 -test.run Test_main
echo "
"

echo "  TESTANDO GitLog/LinhaLog.go"
echo "  test case: buildLinhaLog"
go tool test2json -t /tmp/___Test_buildLinhaLog_in_bb_captura_artefatos_ordem_fornecimento_GitLog.test -test.v -test.paniconexit0 -test.run Test_buildLinhaLog
echo "
"
