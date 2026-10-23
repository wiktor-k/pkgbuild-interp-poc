# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-django-compressor
_name="${pkgname#python-}"
pkgver=4.6
pkgrel=5
pkgdesc="Compresses linked and inline javascript or CSS into a single cached file"
arch=(any)
url="https://github.com/django-compressor/django-compressor"
license=(MIT)
depends=(
  python
  python-django
  python-django-appconf
  python-rcssmin
  python-rjsmin
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-beautifulsoup4
  python-brotli
  python-calmjs.parse
  python-csscompressor
  python-django-sekizai
  python-html5lib
  python-jinja
  python-lxml
)
optdepends=(
  'python-beautifulsoup4: for compressor.parser.BeautifulSoupParser and compressor.parser.LxmlParser'
  'python-brotli: for compressor.storage.BrotliCompressorFileStorage'
  'python-calmjs.parse: for compressor.filters.jsmin.CalmjsFilter'
  'python-csscompressor: for compressor.filters.cssmin.CSSCompressorFilter'
  'python-django-sekizai: for including template code into main template'
  'python-html5lib: for compressor.parser.Html5LibParser'
  'python-jinja: for jinja2 templating support'
  'python-lxml: for compressor.parser.LxmlParser'
)
source=(
  $pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz
)
sha512sums=('326e2a6c2b28e1a93b9897ea19ae959034cef488d40d23eebd8d956f350e8f7c908649ea03c7fc4786bd8af27f825b81774d6a038db347335c7d2e94ead5c1b7')
b2sums=('e414899110df922d1dde66ebfe16c3e70dcba23fe78054282e82fa277d76465af0ccbc1c0a692e5d1c542a9d6fed31c2ac216d3b0b9ebd608333a1998eba3a3f')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  PYTHONPATH="$PWD" django-admin test --settings=compressor.test_settings compressor
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {AUTHORS,README.rst} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
