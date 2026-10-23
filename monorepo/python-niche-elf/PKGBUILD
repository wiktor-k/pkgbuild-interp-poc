# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

pkgname=python-niche-elf
pkgver=0.3.7
pkgrel=1
pkgdesc='Library that optimizes some niche operations commonly used by debugger extensions'
url='https://github.com/pwndbg/niche-elf'
arch=(any)
license=(GPL-3.0-or-later)
depends=(
  python
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=("${url}/archive/${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha512sums=('ab53e82745144a62652bbfc3f0b3c469def0b798c0c2ba4fe948f67ff55dfba80d5881d80bdb55b8ed63affb4e8cc445f30e1ec9435a7089d8d0839a7713425a')
b2sums=('71ae8ea33b6b0167ca52ab15589702dcfe756b6a91881c56b3a0f1dfb0800c925184259fc42d0fa5db1e78908f6395062a6fed1a0adf5f81c5dabe1441188f9a')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
  install -Dm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
}

# vim:set ts=2 sw=2 et:
