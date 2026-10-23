# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-eventlet
pkgver=0.41.2
pkgrel=1
pkgdesc='Highly concurrent networking library'
url='https://eventlet.net'
arch=('any')
license=('MIT')
depends=(
  'python'
  'python-dnspython'
  'python-greenlet'
)
makedepends=(
  'python-build'
  'python-hatch-vcs'
  'python-hatchling'
  'python-installer'
  'python-sphinx'
  'python-sphinxcontrib-apidoc'
)
checkdepends=(
  'python-httplib2'
  'python-psycopg2'
  'python-pyopenssl'
  'python-pytest'
  'python-pyzmq'
)
optdepends=(
  'python-httplib2: non-blocking HTTP support'
  'python-psycopg2: non-blocking PostgreSQL support'
  'python-pyopenssl: non-blocking SSL support'
  'python-pyzmq: non-blocking ZeroMQ support'
)
source=("https://github.com/eventlet/eventlet/archive/$pkgver/${pkgname#python-}-$pkgver.tar.gz")
sha512sums=('5c36bee94ff5f6da4a863754ad7baad226d5a0b60c5c91cfc639264f931e3458c668333e1749673cb0ed21af2bec7e74aba51cfe0e5ea572e5d793767a3d3daa')
b2sums=('e46a771829f1dab8f4259a4d31274b33f86c97ef6d32fee0c64aca356f788a80d0a015f52036983d8d514f46453a4c816f448b4f43a4aa7456e0d0c9468c172a')

build() {
  cd ${pkgname#python-}-$pkgver
  export SETUPTOOLS_SCM_PRETEND_VERSION=$pkgver
  python -m build --wheel --no-isolation

  python -m venv --system-site-packages doc-env
  doc-env/bin/python -m installer dist/*.whl
  local site_packages=$(doc-env/bin/python -c "import site; print(site.getsitepackages()[0])")
  PYTHONPATH="$site_packages" sphinx-build -d /tmp -b text doc/source doc/build/html
}

check() {
  cd ${pkgname#python-}-$pkgver
  # test_ssl_close already failed under Python3.12 "known regression"
  pytest tests --deselect=tests/ssl_test.py::SSLTest::test_ssl_close 
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE

  install -vdm 755 "$pkgdir/usr/share/doc/$pkgname"
  cp -vr doc/build/html "$pkgdir/usr/share/doc/$pkgname"
  cp -vr examples "$pkgdir/usr/share/doc/$pkgname"
}
