# Maintainer: Maxime Gauduin <alucryd@archlinux.org>
# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Jingbei Li <i@jingbei.li>
# Contributor: Jose Riha <jose1711i@gmail.com>

pkgname=python-audioread
pkgver=3.1.0
pkgrel=3
pkgdesc="Cross-library (GStreamer + Core Audio + MAD + FFmpeg) audio decoding for Python"
arch=(any)
url=https://github.com/beetbox/audioread
license=(MIT)
depends=(
  python
  python-standard-aifc
  python-standard-sunau
  python-audioop-lts
)
makedepends=(
  git
  python-build
  python-poetry-core
  python-gobject
  python-installer
  python-pymad
  python-wheel
)
checkdepends=(
  python-pytest
)
optdepends=(
  'gst-python: GStreamer backend'
  'python-gobject: GStreamer backend'
  'python-pymad: MAD backend'
)
source=("$pkgname::git+https://github.com/beetbox/audioread.git#tag=v$pkgver")
sha512sums=('48bd01a99a9acaf9cd2c0416036173b739c91ea78d96162d48cc833b73773991c3cfe13da3ac8984a496daadd438044aa588e4b467dfae3c2694830801a1a786')
b2sums=('edb2bd15df3679359b0120d4fbd6dac65fcef6fb5677bb12284e426b07de2c3ea0fbaf6d0ecc92bf0aaee9a79e352c96eef336b2dc2322b31b32b14230379c38')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  # install to temporary directory
  python -m installer --destdir="$PWD/tmp_install" dist/*.whl

  PYTHONPATH="$PWD/tmp_install$site_packages" pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
# vim:set ts=2 sw=2 et:
