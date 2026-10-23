# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Massimiliano Torromeo <massimiliano.torromeo@gmail.com>
# Contributor: William J Bowman <bluephoenix47@gmail.com>

pkgname=python-certifi
pkgver=2026.10.19
pkgrel=1
pkgdesc="Python package for providing Mozilla's CA Bundle (using system CA store)"
arch=(any)
url="https://github.com/certifi/python-certifi"
license=(MPL-2.0)
depends=(
  ca-certificates
  python
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-pytest)
source=("git+$url.git#tag=$pkgver")
sha512sums=('de2c802877bcde3803ceb242d43501de620175d2586d1fa40ac8cfb94fbea1c21bd3b96bbaea2e769768518e5411ec726ec5fe67e01525fd63b282ea20b72a9a')
b2sums=('75c47f01393893fed0c34ec837d474fadd906b030b7c735c311e9fdc26541047a9f3c8ce02822ef2310361535cb8e5436c3747f7b6c3f2969a315fa56699c235')

prepare() {
  cd $pkgname
  # Use system CA store. Replacing the copy in the source tree so the test
  # suite is actually run against it.
  ln -sf /etc/ssl/certs/ca-certificates.crt certifi/cacert.pem
  # Our CA store has non-ASCII comments, but we are not packaging for JVM
  # https://github.com/certifi/python-certifi/issues/50
  sed -i 's/encoding="ascii"/encoding="utf-8"/' certifi/core.py
}

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname
  pytest
}

package() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl

  # Replace CA store here again because the symlink was installed as a file
  ln -sf /etc/ssl/certs/ca-certificates.crt "$pkgdir/$site_packages/certifi/cacert.pem"

  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
