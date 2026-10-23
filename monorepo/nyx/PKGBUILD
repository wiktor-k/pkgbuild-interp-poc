# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Contributor: Alex Xu <alex_y_xu@yahoo.ca>

pkgname=nyx
pkgver=2.1.0
pkgrel=10
pkgdesc='Command-line status monitor for tor'
url='https://nyx.torproject.org/'
arch=('any')
license=('GPL-3.0-or-later')
depends=('python' 'python-stem' 'net-tools' 'procps-ng' 'lsof')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel')
optdepends=('tor: tor-server to talk to')
source=(https://files.pythonhosted.org/packages/source/${pkgname::1}/${pkgname}/${pkgname}-${pkgver}.tar.gz
        "python3.11.patch::https://gitweb.torproject.org/nyx.git/patch/?id=dcaddf2ab7f9d2ef8649f98bb6870995ebe0b893")
sha512sums=('69a645cc811a1b1fbd9cf0501d8f23f1a52e0e92a0598097bb0c9efb040e1895e1dd431bad93be36db0e85ca7ac32ee5c866d50d2812069a39e5160b44575a37'
            '00db02e9d7aa8787a2b167a5b4364168c5fdc60f7baa1edd7416466b4f981ca69ba4488011cf87420c0b49b7232c2efca8915873b32f68127819f4e073c9a9bb')

prepare() {
  cd ${pkgname}-${pkgver}
  # remove tests requiring full terminal
  rm -r test/{panel,{popups,subwindow,installation}.py}
  # make test exit on error
  sed -E 's|(test_runner.run\(tests\))|__import__("sys").exit(0 if \1.wasSuccessful() else 1)|' -i run_tests.py

  # fix tests with python 3.12
  find -type f -exec sed -e 's|assertRaisesRegexp|assertRaisesRegex|' -i {} \;


  patch -Np1 < ../python3.11.patch
}

build() {
  cd ${pkgname}-${pkgver}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname}-${pkgver}
  ./run_tests.py
}

package() {
  cd ${pkgname}-${pkgver}
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 nyx.1 -t "${pkgdir}/usr/share/man/man1"
}

# vim: ts=2 sw=2 et:
