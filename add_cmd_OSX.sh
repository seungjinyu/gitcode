echo "building gitcode binary file"
go build . 

echo "adding cmd to /usr/local/bin"
cp -av gitcode /usr/local/bin

echo "removing gitcode binary on repo"
rm gitcode