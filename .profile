if [ ! `type -P go` ]; then
    echo "go doesn't exist"
    exit 1
fi

if [[ -n "$DB2HOME" ]]; then
    echo "DB2 driver exists"
    exit 1
fi

go get -d github.com/ibmdb/go_ibm_db

cd $HOME/go/src/github.com/ibmdb/go_ibm_db/installer

go run setup.go

export DB2HOME=$HOME/go/src/github.com/ibmdb/go_ibm_db/installer/clidriver
export CGO_CFLAGS=-I$DB2HOME/include
export CGO_LDFLAGS=-L$DB2HOME/lib
export LD_LIBRARY_PATH=$HOME/go/src/github.com/ibmdb/go_ibm_db/installer/clidriver/lib

