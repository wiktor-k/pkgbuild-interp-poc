# Maintainer: Filipe Laíns (FFY00) <lains@archlinux.org>
# Maintainer: Maxime Gauduin <alucryd@archlinux.org>

pkgname=python-httptools
pkgver=0.8.0
pkgrel=1
pkgdesc='Fast HTTP parser'
arch=(x86_64)
url=https://github.com/MagicStack/httptools
license=(MIT)
depends=(
  python
  llhttp
)
makedepends=(
  cython
  git
  python-setuptools
  python-wheel
)
source=(git+https://github.com/MagicStack/httptools#tag=v${pkgver}
        git+https://github.com/nodejs/http-parser.git)
b2sums=('7d90b76b7482c48ccd0e29e7ea11eb59956bb0b7dc4e1b9185f21ac32954ad1d5563e175ab897c5a2dc13b901426a0cd3bf41a20b52101b0cf491b704c3aeeea'
        'SKIP')

prepare() {
  cd httptools
  sed 's/CYTHON_DEPENDENCY =.*/CYTHON_DEPENDENCY = "Cython"/' -i setup.py
  git rm vendor/llhttp
  git submodule init
  git config submodule.vendor/http-parser.url "$srcdir"/http-parser
  git -c protocol.file.allow=always submodule update
}

build() {
  cd httptools
  python setup.py build build_ext --use-system-llhttp
}

package() {
  cd httptools
  python setup.py install --root="$pkgdir" --optimize=1 --skip-build
  install -Dm 644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}

# vim: ts=2 sw=2 et:
