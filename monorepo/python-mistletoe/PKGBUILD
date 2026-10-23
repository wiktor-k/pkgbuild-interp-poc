# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-mistletoe
pkgver=1.6.0
pkgrel=1
pkgdesc='A fast, extensible Markdown parser in pure Python'
arch=(any)
url='https://github.com/miyuchina/mistletoe'
license=(MIT)
depends=(python)
makedepends=(
  git
  python-build
  python-installer
  python-wheel
  python-setuptools
)
checkdepends=(
  python-parameterized
  python-pygments
)
source=("$pkgname::git+$url.git#tag=v$pkgver")
sha512sums=('6f897dbffd65c2106d7322a15f85342686d70c5314b4f928651dbfdc5395d16539a0040daf22a8960f2f00195afc303fc49a1151cff99c21276e7696392f1cb9')
b2sums=('00f81a7dc4fba20612c1aa5f1d30b6c84bd62d6696ef364e582fc2cdf26c0a8a6b468a54c85ff99230d040f2486a51f265430afe83e5d95117a9b07e4c5a50d6')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  python -m unittest discover test
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
