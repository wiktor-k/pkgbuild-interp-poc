# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Maxime Gauduin <alucryd@archlinux.org>
# Contributor: Frederik “Freso” S. Olesen <freso.dk@gmail.com>

pkgname=python-mediafile
pkgver=0.17.0
pkgrel=1
pkgdesc='Elegant audio file tagging'
arch=(any)
url='https://github.com/beetbox/mediafile'
license=(MIT)
depends=(
  python
  python-mutagen
  python-filetype
)
makedepends=(
  git
  python-build
  python-poetry-core
  python-installer
)
checkdepends=(python-pytest)
source=("$pkgname::git+https://github.com/beetbox/mediafile#tag=v$pkgver")
sha512sums=('f3d182e065343697f384c0eb8685d4f9f7eaff12c6bd32be80f194104837672acfb2d2055cc6948c7065cd9c0a4f21df8a6c1782652eb18bf684c282410cd746')
b2sums=('bed66b9d4852735d55b506877e9bc013d2564f2401836920db8ac955bf1c74106c8fa7f745751a3695b41c47df709ee01b100ea17d4fe041f565f5c711c1e89b')

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}

# vim: ts=2 sw=2 et:
