apt install unzip
mkdir -p data
#curl -L \
#  https://osv-vulnerabilities.storage.googleapis.com/npm/all.zip \
#  -o data/npm.zip

#unzip data/npm.zip -d data

curl -L \
  https://osv-vulnerabilities.storage.googleapis.com/PyPI/all.zip \
  -o data/pypi.zip

unzip data/pypi.zip -d data
